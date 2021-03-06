package sds_package_manager

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"

	"github.com/golang/glog"
	api "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/client/clientset/versioned"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/ccutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/helmutil"
	"io/ioutil"
	"os"
	"github.com/juju/loggo"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cma"
)

var (
	logger loggo.Logger
)


type SDSPackageManagerController struct {
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
	informer cache.Controller

	client *versioned.Clientset
}

func NewSDSPackageManagerController(config *rest.Config) (output *SDSPackageManagerController) {
	if config == nil {
		config = k8sutil.DefaultConfig
	}
	client := versioned.NewForConfigOrDie(config)

	// create sdscluster list/watcher
	sdsPackageManagerListWatcher := cache.NewListWatchFromClient(
		client.CmaV1alpha1().RESTClient(),
		api.SDSPackageManagerResourcePlural,
		"default",
		fields.Everything())

	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the SDSCluster key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the SDSPackageManager than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(sdsPackageManagerListWatcher, &api.SDSPackageManager{}, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})

	output = &SDSPackageManagerController{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
		client:   client,
	}
	output.SetLogger()
	return
}

func (c *SDSPackageManagerController) processNextItem() bool {
	// Wait until there is a new item in the working queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two SDSClusters with the same key are never processed in
	// parallel.
	defer c.queue.Done(key)

	// Invoke the method containing the business logic
	err := c.processItem(key.(string))
	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, key)
	return true
}

// processItem is the business logic of the controller.
func (c *SDSPackageManagerController) processItem(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a SDSPackageManager, so that we will see a delete for one SDSPackageManager
		fmt.Printf("SDSPackageManager -->%s<-- does not exist anymore\n", key)
	} else {
		SDSPackageManager := obj.(*api.SDSPackageManager)
		clusterName := SDSPackageManager.GetName()
		fmt.Printf("SDSPackageManager -->%s<-- does exist (name=%s)!\n", key, clusterName)

		switch SDSPackageManager.Status.Phase {
		case api.PackageManagerPhaseNone, api.PackageManagerPhasePending:
			c.deployTiller(SDSPackageManager)
		case api.PackageManagerPhaseInstalling:
			c.waitForTiller(SDSPackageManager)
		}
	}
	return nil
}

// handleErr checks if an error happened and makes sure we will retry later.
func (c *SDSPackageManagerController) handleErr(err error, key interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the key on every successful synchronization.
		// This ensures that future processing of updates for this key is not delayed because of
		// an outdated error history.
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < 5 {
		glog.Infof("Error syncing krakenCluster %v: %v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	runtime.HandleError(err)
	glog.Infof("Dropping krakenCluster %q out of the queue: %v", key, err)
}

func (c *SDSPackageManagerController) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	glog.Info("Starting SDSCluster controller")

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	glog.Info("Stopping SDSCluster controller")
}

func (c *SDSPackageManagerController) runWorker() {
	for c.processNextItem() {
	}
}

func (c *SDSPackageManagerController) deployTiller(packageManager *api.SDSPackageManager) (bool, error){
	config, err := c.getRestConfigForRemoteCluster(packageManager.Spec.Name, packageManager.Namespace, nil)
	if err != nil {
		return false, err
	}

	k8sutil.CreateNamespace(k8sutil.GenerateNamespace(packageManager.Spec.Namespace), config)
	k8sutil.CreateServiceAccount(k8sutil.GenerateServiceAccount("tiller-sa"), packageManager.Spec.Namespace, config)
	if packageManager.Spec.Permissions.ClusterWide {
		k8sutil.CreateClusterRole(helmutil.GenerateClusterAdminRole("tiller-"+packageManager.Spec.Namespace), config)
		k8sutil.CreateClusterRoleBinding(k8sutil.GenerateSingleClusterRolebinding("tiller-"+packageManager.Spec.Namespace, "tiller-sa", packageManager.Spec.Namespace,"tiller-"+packageManager.Spec.Namespace ), config)
	} else {
		logger.Infof("Not cluster wide")
		namespaces := append(packageManager.Spec.Permissions.Namespaces, packageManager.Spec.Namespace)
		for _, namespace := range namespaces {
			logger.Infof("Creating namespace %s", namespace)
			k8sutil.CreateNamespace(k8sutil.GenerateNamespace(namespace), config)
			k8sutil.CreateRole(helmutil.GenerateAdminRole(packageManager.Spec.Namespace+"-tiller"), namespace, config)
			k8sutil.CreateRoleBinding(k8sutil.GenerateSingleRolebinding(packageManager.Spec.Namespace+"-tiller", "tiller-sa", packageManager.Spec.Namespace, packageManager.Spec.Namespace+"-tiller"), namespace, config)
		}
	}
	k8sutil.CreateJob(helmutil.GenerateTillerInitJob(
		helmutil.TillerInitOptions{
			BackoffLimit:   4,
			Name:           "tiller-install-job",
			Namespace:      packageManager.Spec.Namespace,
			ServiceAccount: "tiller-sa",
			Version:        packageManager.Spec.Version}), packageManager.Spec.Namespace, config)

	packageManager.Status.Phase = api.PackageManagerPhaseInstalling
	_, err = c.client.CmaV1alpha1().SDSPackageManagers(packageManager.Namespace).Update(packageManager)
	if err == nil {
		logger.Infof("Deployed tiller on -->%s<--", packageManager.Spec.Name)
	} else {
		logger.Infof("Could not update the status error was %s", err)
	}

	return true, nil
}

func retrieveClusterRestConfig(name string, namespace string, config *rest.Config) (*rest.Config, error) {
	cluster, err := ccutil.GetKrakenCluster(name, namespace, config)
	if err != nil {
		return nil, err
	}
	// Let's create a tempfile and line it up for removal
	file, err := ioutil.TempFile(os.TempDir(), "kraken-kubeconfig")
	defer os.Remove(file.Name())
	file.WriteString(cluster.Status.Kubeconfig)

	clusterConfig, err := clientcmd.BuildConfigFromFlags("", file.Name())
	if os.Getenv("CLUSTERMANAGERAPI_INSECURE_TLS") == "true" {
		clusterConfig.TLSClientConfig = rest.TLSClientConfig{Insecure:true}
	}

	if err != nil {
		logger.Errorf("Could not load kubeconfig for cluster -->%s<-- in namespace -->%s<--", name, namespace)
		return nil, err
	}
	return clusterConfig, nil
}

func (c *SDSPackageManagerController) getRestConfigForRemoteCluster(clusterName string, namespace string, config *rest.Config) (*rest.Config, error) {
	cluster, err := ccutil.GetKrakenCluster( clusterName, namespace, config)
	if err != nil {
		glog.Errorf("Failed to retrieve KrakenCluster CR -->%s<-- in namespace -->%s<--, error was: ", clusterName, namespace, err)
		return nil, err
	}
	if cluster.Status.Kubeconfig == "" {
		glog.Errorf("Could not install tiller yet for cluster -->%s<-- cluster is not ready, status is -->%s<--", cluster.Name, cluster.Status.State)
		return nil, err
	}

	remoteConfig, err := retrieveClusterRestConfig(clusterName, namespace, config)
	if err != nil {
		glog.Errorf("Could not install tiller yet for cluster -->%s<-- cluster is not ready, error is: %v", clusterName, err)
		return nil, err
	}

	return remoteConfig, nil
}

func (c *SDSPackageManagerController)  SetLogger() {
	logger = util.GetModuleLogger("pkg.controllers.sds_package_manager", loggo.INFO)
}

func (c *SDSPackageManagerController) waitForTiller(packageManager *api.SDSPackageManager) (result bool, err error) {
	config, err := c.getRestConfigForRemoteCluster(packageManager.Spec.Name, packageManager.Namespace, nil)
	if err != nil {
		return false, err
	}

	clientset, _ := kubernetes.NewForConfig(config)
	timeout := 0
	for timeout < 2000 {
		job, err := clientset.BatchV1().Jobs(packageManager.Spec.Namespace).Get("tiller-install-job", v1.GetOptions{})
		if err == nil {
			if job.Status.Succeeded > 0 {
				packageManager.Status.Phase = api.PackageManagerPhaseImplemented
				packageManager.Status.Ready = true
				_, err = c.client.CmaV1alpha1().SDSPackageManagers(packageManager.Namespace).Update(packageManager)
				if err == nil {
					logger.Infof("Tiller running on -->%s<--", packageManager.Spec.Name)
					c.updateSDSCluster(packageManager)
				} else {
					logger.Infof("Could not update the status error was %s", err)
				}
				return true, nil
			}
		}
		time.Sleep(5 * time.Second)
		timeout++
	}
	return false, nil
}

func (c *SDSPackageManagerController) updateSDSCluster(packageManager *api.SDSPackageManager) (result bool, err error) {
	// TODO This is dubious, but for the PoC, good enough
	sdsCluster, err := cma.GetSDSCluster(packageManager.Spec.Name, "default", nil)
	if err != nil {
		logger.Infof("Failed to get SDSCluster for SDSPackageManager %s, error was: ", packageManager.Spec.Name, err)
	}

	changes := false
	if sdsCluster.Status.TillerInstalled == false {
		changes = true
		sdsCluster.Status.TillerInstalled = true
	}
	switch sdsCluster.Status.Phase {
	case api.ClusterPhaseWaitingForCluster, api.ClusterPhaseNone, api.ClusterPhasePending, api.ClusterPhaseHaveCluster, api.ClusterPhaseDeployingPackageManager:
		changes = true
		sdsCluster.Status.Phase = api.ClusterPhaseHavePackageManager
	}

	if changes {
		_, err = cma.UpdateSDSCluster(sdsCluster, sdsCluster.Namespace, nil)
		if err != nil {
			logger.Infof("Could not update SDSCluster for KrakenCluster %s, error was: ", sdsCluster.Name, err)
		}
	}

	return true, nil
}
