// Code generated by go-bindata.
// sources:
// api/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\x38\x70\x03\xb6\x01\x49\x9c\x66\x6f\x79\xda\x90\x02\x5d\x80\x65\x08\x9a\x61\x79\xd8\x82\x82\xa6\xce\x12\x1b\x8a\xc7\x92\x27\x7b\xde\xe0\xff\x3e\x50\xb2\x2c\x4a\xb6\x53\x37\xce\xda\x04\xed\x4b\x63\xf1\x78\xba\xef\xee\x13\xef\x3b\xfe\x3b\x02\x10\x61\x2e\xf3\x1c\xbd\x38\x07\x71\x76\x72\x2a\x8e\xe2\x33\x6d\xa7\x24\xce\x21\xae\x03\x08\xd6\x6c\x30\xae\x5f\x98\x2a\x30\x7a\xb8\x92\x56\xe6\xe8\xe1\xe7\xeb\xcb\xda\x1e\x40\xcc\xd0\x07\x4d\x36\x5a\xcd\x4e\x4f\x5e\xad\x1c\x01\x08\x45\x96\xa5\xe2\xb5\x37\x00\x61\x65\x59\xbb\xbb\xd2\xaa\x90\x68\xe0\x0f\xb4\xf8\x8f\x96\xab\x1d\x00\xa2\xf2\x26\xae\x17\xcc\x2e\x9c\x8f\xc7\xb9\xe6\xa2\x9a\x9c\x28\x2a\xc7\xb3\xa1\x29\x96\x52\xd7\xc6\xe5\x6a\xe9\xa7\x3c\x3e\x89\xc6\xa2\xb6\x59\x8e\x00\x96\x35\xa8\xa0\x0a\x2c\x31\x88\x73\xf8\xb3\x09\x2d\xfa\x6f\xc3\xac\xdf\x15\x77\xdc\xd5\xb6\x8a\x6c\xa8\x7a\xc6\xd2\x39\xa3\x95\x64\x4d\x76\xfc\x3e\x90\xed\x6c\x9d\xa7\xac\x52\x7b\xda\x4a\x2e\x42\x97\xd9\xb1\x74\x7a\x3c\x7b\x35\x56\x4d\x62\xd3\x24\xe5\x98\xe6\x2c\x86\x5f\x95\xa5\xf4\x8b\x88\xf5\x56\x1b\x03\x1e\xd9\x6b\x9c\x21\x70\x81\x10\x58\x72\x15\x80\xa6\x20\x61\xe5\x0c\xa4\xcd\x40\x73\x80\xfb\x6a\x82\x8a\xec\x54\xe7\x30\x25\x0f\x8a\xac\x45\xc5\x7a\xa6\x79\xb1\xce\x23\x80\x20\x87\xbe\x0e\xf9\x32\x8b\xef\x78\x83\xbc\x2a\x77\x6a\xe4\x31\x38\xb2\x01\x43\x2f\x36\x00\x71\x76\x7a\x3a\x78\x04\x20\x32\x0c\xca\x6b\xc7\x2b\x62\x24\x8e\x1a\x44\xb1\x20\x72\x63\x1b\x80\xf8\xd6\xe3\x34\xee\xf8\x66\x9c\xe1\x54\x5b\x1d\x3d\x84\x36\x4b\xef\xca\x86\x7e\xef\xa4\xd3\x5d\x94\x6f\xd1\x99\x85\xe8\x39\x5a\x8e\xb6\xfd\xbd\x4c\xe0\x38\xe9\x65\x89\x8c\xbe\x2b\x5e\xf3\x6f\x00\xa4\x65\x6c\xfd\xff\xd1\x83\x20\x7f\x93\x25\xc6\x3a\xc4\xaa\xb4\x95\x60\x82\x09\x82\x21\xba\xc7\x0c\x2a\x77\x32\x74\xa1\xeb\x9d\x1f\x2a\xf4\x8b\xe1\x92\xc7\x0f\x95\xf6\x18\x4b\x32\x95\x26\xe0\x60\x99\x17\xae\x0e\x2c\xb0\xd7\x36\x17\x5b\x01\xdf\x25\x80\x59\xe6\x43\xa8\xed\x57\xdd\x6d\xbe\x1b\x0d\x32\x25\x32\x34\xc8\xf8\x30\x1f\x1b\x9b\x8e\x7f\x0f\x70\xeb\x75\x6d\xfa\x02\xe8\xd5\x0b\xf4\xb9\x30\xec\xb6\x90\x0c\x3a\xa4\x0c\xfb\x2e\x40\xdc\x18\x89\x96\x61\x60\x4f\x8b\x97\xc7\x31\x57\x7d\xe4\xc0\x73\x9e\x66\x3a\x36\x99\xbd\x38\x76\xe1\x51\xbe\x08\x8e\xf5\x02\xfd\x2c\x1c\x9b\x50\xb6\xc1\x81\x86\x1e\xdb\x56\x12\x76\xb0\xaf\x86\xe4\x78\xea\x04\x5c\x85\x7c\x1f\xf8\x8f\xe7\xdb\x28\xc9\xde\xba\xff\xe6\xc8\x8e\x32\x45\x95\xed\x09\x15\x47\x61\x37\x27\x6f\x74\xe9\x0c\x82\xa3\x0c\xea\x8d\xd0\x12\xeb\x04\x60\xd5\xa0\xe3\xef\x0c\xe6\x9a\x8b\xfa\x63\xb5\x55\x39\x41\x1f\x9b\x83\xa3\x2c\x80\xb6\xcd\x53\x59\x62\x70\x52\x61\x43\xf0\x0c\xb3\x87\xbb\xf2\x35\x65\x17\x75\xa0\xcf\x98\xd3\x49\x98\x5f\x23\xa3\x13\xf8\x5f\x86\xcf\x05\x1a\x43\x73\xf2\x26\xfb\x54\x3a\xff\x12\x77\xc2\x6d\xdc\x0a\x37\xe8\x67\x5a\x21\x1c\xc3\xbc\xe1\xb3\x43\xc9\x20\x21\xf7\x88\xac\x6d\x1e\xdb\x4d\xcb\xe0\x7d\xc8\x5b\xfb\xae\x5d\x3f\x67\xee\x76\x51\x7e\x8d\xd4\xed\xd0\x7f\x26\xe6\xae\xc7\xb3\x24\xb0\x6e\x40\xda\xa7\x5d\x24\x0c\x6f\xa5\x0a\x4d\xde\xa3\xea\x4e\xc8\x38\xa2\x39\xf4\xac\x07\x24\x6b\x8b\xd0\xa3\xdd\x40\xee\x1c\xf5\xd6\xda\x49\x78\xb7\xca\x5f\xab\x14\xcc\x3a\xcc\xcb\xad\x5f\xeb\x7e\x62\xe0\x00\x78\x74\xbf\x0b\xdc\x84\xc8\xa0\xb4\x7d\x74\x53\xf2\xa5\xe4\x9d\xcb\x6b\xf0\xb7\x05\x72\x11\x1b\x99\x07\x4b\xdc\xcb\xc1\x5c\x86\x34\x03\x30\x59\x00\x17\x3a\x40\xe4\x2d\x06\x16\x5b\x3f\x99\x66\x78\x7d\x54\x1d\x7e\x5f\xcd\xbe\x1b\xc5\xc8\xd1\xae\xce\x9e\x47\x54\x61\x8b\xec\x7f\x7e\x55\xb8\xa0\xca\x64\x3d\xcc\x13\x6c\xd5\x7f\x8f\x7b\x4f\x93\xe9\x9b\xf5\x0d\x43\x7c\xe5\x66\x3d\xf7\xcd\xed\x70\x62\x7f\x7e\x89\xbd\xec\x8d\x56\xad\x4c\x0b\x8b\xc0\x58\x3e\x79\x5a\xd3\x61\x2e\xf4\x52\xac\x86\x07\x67\xfa\xc6\xee\x4e\xe7\xe0\xb7\x26\xd7\x43\x4c\xed\xed\x50\xdb\xda\x37\x83\xf8\x84\x3a\xa7\x1a\xe8\xc0\x43\xba\x16\xc8\x8f\x3e\xa9\x1b\x79\xcd\x04\x53\x64\xd5\x68\xf1\xb5\x6c\xff\xcb\xfe\x8a\x72\x86\x80\xa5\xe3\x45\xb4\xa9\x47\x64\x90\xc6\x74\xca\x3c\x1c\x86\xff\x60\xa2\xc7\x71\x61\x17\x78\x6d\x19\xf3\xde\x84\xdb\xa3\xba\xb6\xfc\xe3\xd9\xae\xd4\x7c\x64\x1e\xf9\x9e\xfc\x20\x0f\x3f\x3c\x22\x11\x7d\x45\xf1\xff\x36\xeb\xcd\xe8\xd6\x5e\x7b\x2d\x23\x55\xb9\xbd\xa3\x6c\x2f\x1c\x07\xd7\xb3\xc4\x10\x64\xfe\x54\x60\x5a\x01\x1d\xc9\x9b\x00\x1b\x5e\x7f\xe3\xdf\x8c\xde\x4a\xf3\x9a\x54\x22\xb0\x06\x7a\xfa\x8a\x3c\x82\x9c\x50\xc5\x90\xbf\xbd\xbe\x38\x7e\x23\x19\xe7\xb2\x95\xa1\x0f\x5c\xcc\xe7\xde\xa9\x63\x54\xd4\x1c\x93\xcd\xcf\x7c\xb5\x39\x86\x30\x5a\x8e\xfe\x0b\x00\x00\xff\xff\xc3\x50\xc9\x97\x6e\x18\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 6254, mode: os.FileMode(420), modTime: time.Unix(1525746595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"api.swagger.json": apiSwaggerJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
