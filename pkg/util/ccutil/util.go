package ccutil

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util"
	"k8s.io/client-go/rest"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"github.com/samsung-cnct/cluster-controller/pkg/client/clientset/versioned"
)

var (
	logger loggo.Logger
)

func prepareRestClient(config *rest.Config) *versioned.Clientset {
	if config == nil {
		config = k8sutil.DefaultConfig
	}

	return versioned.NewForConfigOrDie(config)
}

func SetLogger() {
	logger = util.GetModuleLogger("pkg.util.ccutil", loggo.INFO)
}

