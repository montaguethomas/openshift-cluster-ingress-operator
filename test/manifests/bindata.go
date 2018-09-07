// Code generated by go-bindata.
// sources:
// test/assets/app-ingress/deployment.yaml
// test/assets/app-ingress/namespace.yaml
// test/assets/app-ingress/route-default.yaml
// test/assets/app-ingress/route-internal.yaml
// test/assets/app-ingress/service.yaml
// test/assets/cluster-ingress-default.yaml
// test/assets/cluster-ingress-internal.yaml
// DO NOT EDIT!

package manifests

import (
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

var _testAssetsAppIngressDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\x83\x31\x0c\x84\xf7\x3c\xc5\xbd\xc0\xaf\xaa\x6b\x66\x46\x66\x76\x93\x1e\xfd\x23\x9c\xc4\x8a\x0d\x12\x6f\x8f\x42\x4b\xc5\x2f\x71\x93\xed\xd3\x77\x3a\xbf\xd7\x7e\xc9\x78\xa2\xe9\xf8\x6a\xec\x91\xc4\xea\x0b\xa7\xd7\xd1\x33\xc4\xcc\x4f\x9f\xe7\xd4\x18\x72\x91\x90\x9c\x80\x2e\x8d\x3f\xce\x7d\x76\x93\xc2\x8c\xa2\x1f\x1e\x9c\x5b\xed\xd7\x49\xf7\x2d\xe8\x91\xdc\x58\x16\x33\x69\x5a\x8b\x78\xc6\x39\x01\x4e\x65\x89\x31\x97\x03\x34\x89\xb2\x3f\xcb\x2b\xd5\x6f\x07\xac\xf0\x8c\x43\x10\x10\x6c\xa6\x12\xbc\x43\x7f\x0a\x2d\xe9\x81\xff\x3f\x01\xf8\xad\xb3\x54\x46\x0f\xa9\x9d\xf3\x41\x6d\x87\xd7\x6e\xaa\x4d\xae\xcc\x18\xc6\xee\x7b\x7d\x8b\xd3\x4e\xd5\xb1\x3d\xf6\xf4\x1d\x00\x00\xff\xff\x27\xc6\x8c\x66\x3f\x01\x00\x00")

func testAssetsAppIngressDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppIngressDeploymentYaml,
		"test/assets/app-ingress/deployment.yaml",
	)
}

func testAssetsAppIngressDeploymentYaml() (*asset, error) {
	bytes, err := testAssetsAppIngressDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-ingress/deployment.yaml", size: 319, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppIngressNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xb1\x0d\xc4\x30\x08\x05\xd0\x9e\x29\x58\xc0\xc5\xb5\x0c\x71\xe5\xf5\x5f\xf6\xd7\x09\x25\x26\x96\x21\x99\x3f\xef\xf0\x18\xa6\x5f\x4c\xe6\x42\xa7\x60\xf9\x8f\x3b\xfd\x0a\xd3\xe7\x23\x93\x85\x81\x82\x89\x6a\x60\xd2\xb4\x9f\x77\x16\x77\xf3\xf8\x6f\x66\xb6\x62\x96\xbc\x01\x00\x00\xff\xff\x3e\x03\xf6\x6f\x46\x00\x00\x00")

func testAssetsAppIngressNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppIngressNamespaceYaml,
		"test/assets/app-ingress/namespace.yaml",
	)
}

func testAssetsAppIngressNamespaceYaml() (*asset, error) {
	bytes, err := testAssetsAppIngressNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-ingress/namespace.yaml", size: 70, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppIngressRouteDefaultYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x4e\xc5\x30\x0c\x45\xf7\x7c\x85\x7f\xa0\x41\xac\xd9\x10\xea\x06\x0c\x2d\xb0\x22\x2b\x71\x5b\x8b\x36\xb6\x62\xb7\xdf\x8f\x4a\xdf\xf0\xde\x78\xef\xd1\xbd\x3a\xa8\xfc\x4d\xcd\x58\x6a\x82\x26\xbb\x53\x14\xa5\x6a\x0b\x4f\x1e\x59\x9e\x8e\xe7\xf0\xcb\xb5\x24\x18\x4e\x16\x36\x72\x2c\xe8\x98\x02\x40\xc5\x8d\x12\x14\x9a\x70\x5f\xfd\x96\x4d\x31\x53\x82\xbc\xee\xe6\xd4\x3a\xae\x73\x23\xb3\xce\xc9\x3c\x98\x52\x3e\x77\x8b\x98\x27\x40\xd5\x07\x1c\x51\xd5\xe2\xeb\xdb\xd7\xf8\xd9\x0f\x3f\x1f\x2f\xef\x7d\x94\xc6\x33\xd7\x6e\xce\x14\x0b\x1d\x77\x5e\x59\xb6\x00\xe0\x72\xbe\x01\x5c\x7e\x23\xb5\x83\x33\xfd\x37\x97\x19\xaa\x86\xbf\x00\x00\x00\xff\xff\xd4\xf1\xc2\x18\xde\x00\x00\x00")

func testAssetsAppIngressRouteDefaultYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppIngressRouteDefaultYaml,
		"test/assets/app-ingress/route-default.yaml",
	)
}

func testAssetsAppIngressRouteDefaultYaml() (*asset, error) {
	bytes, err := testAssetsAppIngressRouteDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-ingress/route-default.yaml", size: 222, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppIngressRouteInternalYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\xb1\x4e\xc3\x40\x10\x44\xfb\xfb\x8a\xfd\x01\x1f\xa2\xbd\x0e\xa1\x74\x40\x91\x00\x2d\x5a\xce\x83\xb3\xc2\xde\x5b\xdd\x6e\xfc\xfd\xc8\x18\xa4\xa4\x9d\xd1\xbc\xd1\x63\x93\x77\x74\x97\xa6\x85\x7a\xbb\x04\x72\x33\xa8\x9f\xe5\x2b\xb2\xb4\xbb\xf5\x3e\x7d\x8b\x8e\x85\x8e\x5b\x97\x16\x04\x8f\x1c\x5c\x12\x91\xf2\x82\x42\xa2\x81\xae\x3c\xff\x05\x6e\x5c\x51\xa8\xce\x17\x0f\xf4\x41\x74\xea\x70\x1f\x02\x1e\x89\x68\xe6\x4f\xcc\xbe\x8d\x89\xa0\xab\xf4\xa6\x0b\x34\xae\x28\x6e\xa8\x5b\x7f\x6e\x1e\x85\xd8\xec\x06\x91\xd9\xcc\xf3\xe3\xd3\xdb\xe9\xf5\x70\xfc\x78\x79\x78\x3e\x0c\xff\xcb\xdc\xba\x4c\xa2\xc3\x54\x91\x47\xac\x57\x16\xb5\x2d\x89\x28\xda\x7e\xbb\xdb\x9c\xd0\x57\xa9\xf8\x4d\x76\x0f\x36\x4b\x3f\x01\x00\x00\xff\xff\xa4\xbd\x50\xf5\x0c\x01\x00\x00")

func testAssetsAppIngressRouteInternalYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppIngressRouteInternalYaml,
		"test/assets/app-ingress/route-internal.yaml",
	)
}

func testAssetsAppIngressRouteInternalYaml() (*asset, error) {
	bytes, err := testAssetsAppIngressRouteInternalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-ingress/route-internal.yaml", size: 268, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsAppIngressServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8b\x41\xaa\xc3\x30\x0c\x05\xf7\x3e\x85\x2e\x10\xc8\xdf\x05\x6d\xff\x05\x02\x2d\xdd\x0b\xe7\x11\x4c\x1d\x5b\x48\x6a\xce\x5f\x1c\xb2\xe9\x6e\xa4\x79\xf3\x2e\x6d\x63\x7a\xc0\xce\x92\x91\x44\xcb\x0b\xe6\xa5\x37\xa6\xf3\x2f\x1d\x08\xd9\x24\x84\x13\x51\x93\x03\x4c\xa2\x7a\xb3\xab\x64\x30\xe5\xfa\xf1\x80\x4d\xa5\xed\x06\xf7\x29\xe0\x91\x5c\x91\x47\xe3\xa8\xc8\xd1\x6d\x30\x8d\x96\xe9\x67\x47\xa4\xdd\xc2\x87\x9e\x48\xad\x47\xcf\xbd\x32\x3d\xff\xd7\x2b\x18\x92\x69\x99\xaf\x23\xc4\x76\xc4\x7a\xbf\x96\x39\x7d\x03\x00\x00\xff\xff\xd9\x18\x76\x42\xba\x00\x00\x00")

func testAssetsAppIngressServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsAppIngressServiceYaml,
		"test/assets/app-ingress/service.yaml",
	)
}

func testAssetsAppIngressServiceYaml() (*asset, error) {
	bytes, err := testAssetsAppIngressServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/app-ingress/service.yaml", size: 186, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsClusterIngressDefaultYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcf\xb1\x4e\x03\x31\x0c\x80\xe1\x3d\x4f\xe1\x07\x68\x0f\x75\xcd\x56\x95\x0e\x95\xa0\x03\x05\x16\x84\x90\xb9\x73\xef\xac\x26\x76\x94\xf8\x22\x4e\x88\x77\x47\xd1\x31\xb0\xfe\xb2\x3f\xcb\x98\xf8\x95\x72\x61\x15\x0f\x2c\x63\xa6\x52\x3a\x4d\x24\x65\xe2\xab\x75\xac\x77\x75\x87\x21\x4d\xb8\x73\x37\x96\xc1\xc3\x21\xcc\xc5\x28\x9f\xd6\x51\x17\xc9\x70\x40\x43\xef\x00\x04\x23\x79\x18\xe8\x8a\x73\x30\x57\x12\xf5\xad\xfe\xa1\xf7\x1a\x91\xc5\x03\xa6\x54\xba\xc3\xc3\xcb\xe5\xf9\xf8\xf4\x71\xde\x3f\x1e\x3b\xcd\x3c\xb2\x6c\xc7\x9e\xba\x81\xea\xbf\xe3\xbd\x46\x07\x90\x75\x36\xba\x50\xa0\xde\x34\x37\x10\x20\xa2\xf5\xd3\xf1\x2b\x35\x97\x55\xca\x5a\xb7\xf0\x7d\xa3\xc5\x03\x49\xe5\xac\x12\x49\x6c\x03\x9a\x28\x63\x5b\x84\xb3\xda\x49\x36\x50\x31\xcc\x54\x3c\xbc\xb1\x18\x65\xc1\xf0\xfe\xe3\x00\x26\x1e\xa7\x7d\x45\x0e\xf8\xc9\x81\x6d\x59\x45\x5b\x12\xb5\x8f\x75\x1e\xdc\x6f\x00\x00\x00\xff\xff\xa3\x18\xf3\xc2\x28\x01\x00\x00")

func testAssetsClusterIngressDefaultYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsClusterIngressDefaultYaml,
		"test/assets/cluster-ingress-default.yaml",
	)
}

func testAssetsClusterIngressDefaultYaml() (*asset, error) {
	bytes, err := testAssetsClusterIngressDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/cluster-ingress-default.yaml", size: 296, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testAssetsClusterIngressInternalYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xb1\x6a\x03\x31\x0c\x86\x77\x3f\x85\x5e\x20\x2e\x59\x6f\x0b\x69\x86\x42\xda\xa1\x69\xbb\x16\xc5\xa7\x9e\x45\x6d\xc9\x58\xba\x83\xbc\x7d\x39\xae\x81\xac\x3f\x9f\x3e\x3e\x61\xe3\x2f\xea\xc6\x2a\x03\xb0\x4c\x9d\xcc\xa2\x36\x12\xcb\xfc\xe3\x91\xf5\x69\xd9\x63\x69\x19\xf7\xe1\x97\x65\x1c\xe0\x58\x66\x73\xea\x2f\x1b\x1a\x2a\x39\x8e\xe8\x38\x04\x00\xc1\x4a\xab\xc4\xa9\x0b\x96\x60\x8d\xd2\x3a\xff\x5b\x9f\xb5\x22\xcb\x00\xd8\x9a\xc5\xe3\xf9\xf3\xf2\x71\x7a\xff\x7e\x3b\xbc\x9e\x76\xf7\x8b\xa8\x9d\x27\x96\xdd\x94\x28\x8e\xb4\x3c\x64\x24\xad\x01\xa0\xeb\xec\x74\xa1\x42\xc9\xb5\xaf\x66\x80\x8a\x9e\xf2\x19\xaf\x54\x6c\x1b\x00\x48\x16\xee\x2a\x95\xc4\x1f\x62\x00\x32\x4f\xf9\xb0\x20\x17\xbc\x72\x61\xbf\x6d\xbc\xdf\x1a\xad\x4f\xe9\x3c\x86\xbf\x00\x00\x00\xff\xff\xe4\xb0\xa3\x98\x0b\x01\x00\x00")

func testAssetsClusterIngressInternalYamlBytes() ([]byte, error) {
	return bindataRead(
		_testAssetsClusterIngressInternalYaml,
		"test/assets/cluster-ingress-internal.yaml",
	)
}

func testAssetsClusterIngressInternalYaml() (*asset, error) {
	bytes, err := testAssetsClusterIngressInternalYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/assets/cluster-ingress-internal.yaml", size: 267, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"test/assets/app-ingress/deployment.yaml": testAssetsAppIngressDeploymentYaml,
	"test/assets/app-ingress/namespace.yaml": testAssetsAppIngressNamespaceYaml,
	"test/assets/app-ingress/route-default.yaml": testAssetsAppIngressRouteDefaultYaml,
	"test/assets/app-ingress/route-internal.yaml": testAssetsAppIngressRouteInternalYaml,
	"test/assets/app-ingress/service.yaml": testAssetsAppIngressServiceYaml,
	"test/assets/cluster-ingress-default.yaml": testAssetsClusterIngressDefaultYaml,
	"test/assets/cluster-ingress-internal.yaml": testAssetsClusterIngressInternalYaml,
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
	"test": &bintree{nil, map[string]*bintree{
		"assets": &bintree{nil, map[string]*bintree{
			"app-ingress": &bintree{nil, map[string]*bintree{
				"deployment.yaml": &bintree{testAssetsAppIngressDeploymentYaml, map[string]*bintree{}},
				"namespace.yaml": &bintree{testAssetsAppIngressNamespaceYaml, map[string]*bintree{}},
				"route-default.yaml": &bintree{testAssetsAppIngressRouteDefaultYaml, map[string]*bintree{}},
				"route-internal.yaml": &bintree{testAssetsAppIngressRouteInternalYaml, map[string]*bintree{}},
				"service.yaml": &bintree{testAssetsAppIngressServiceYaml, map[string]*bintree{}},
			}},
			"cluster-ingress-default.yaml": &bintree{testAssetsClusterIngressDefaultYaml, map[string]*bintree{}},
			"cluster-ingress-internal.yaml": &bintree{testAssetsClusterIngressInternalYaml, map[string]*bintree{}},
		}},
	}},
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

