// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package utils generated by go-bindata.// sources:
// tpl/api.tpl
// tpl/api_common.tpl
// tpl/api_service.tpl
// tpl/proto.tpl
package utils

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tplApiTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x4d\x8b\xd3\x40\x18\xc7\xef\x81\x7c\x87\x87\xf4\xd2\x82\x34\x08\xe2\xa1\x20\xac\x2f\x20\x0b\x22\x52\xf5\x24\xe2\xce\x36\x4f\xbb\x23\xc9\x4c\x9c\x99\x2e\xbb\x1b\x72\x10\xd1\xad\xa7\x1e\x7c\x43\x90\x75\xbb\xa8\x78\xb2\x7b\xd9\xc3\xe2\xc1\x2f\x63\x63\xf3\x2d\x64\xd2\xa9\x4d\xd3\xd4\x68\x4f\xd3\xe7\xe5\xff\x7b\x9e\x7f\x26\x91\xfb\x4c\x91\x3d\xb8\x02\xce\xee\x45\xc7\xb6\x6c\x8b\x06\x21\x17\x0a\x9c\x66\xd3\xed\xf0\x20\xe0\xac\x49\x42\xea\xd8\xd6\x86\x44\xb1\x8b\x02\xea\xb6\x05\x00\x10\x0a\xec\xd2\xbd\x16\xd4\xa2\xd9\x29\x76\x6b\x91\x22\xdb\x3e\xde\x26\x01\xc6\xb3\x9a\x9e\xe0\xfd\x50\x97\xe4\x13\x0d\xdb\xd2\x4a\xb4\x83\x50\x8b\xcc\x29\x4b\x41\x34\xeb\xda\xf0\x78\x07\x9c\xe4\xe3\xe7\xe9\xf8\xc4\xb4\x5e\xe7\x41\x80\x4c\xc5\x93\xc1\xbb\xe9\xe8\xab\x63\xea\x76\x08\xf3\x7c\x14\x70\x13\x55\x2d\xba\xf7\x07\x71\x8b\x4a\x65\xf8\xa8\xc0\x85\x7a\x3e\xd9\xc6\x27\x0d\x10\xa8\xfa\x82\xc9\xe5\x8c\x6e\x6b\xa3\x6c\x68\x0f\x2a\xc6\x98\x8e\xbf\x24\xcf\x9e\x57\x8c\xb1\xc9\xba\x3c\x37\x46\x8b\x7a\x50\xdf\xf4\xd6\xf3\x75\x7d\x09\xff\xed\xe9\xe4\xe4\xa8\xc0\x2f\x92\xaf\x7a\x5e\x5e\xc9\x3c\x20\x2e\x57\xb6\x9f\x33\x16\x13\x48\x25\x28\xeb\x15\xa0\x93\xc1\x71\xfa\xfe\x53\x05\xf4\x06\xfa\xab\x50\x0f\x7d\x54\xb8\x66\xdb\x52\x56\xf2\xf2\x3c\x3d\x1c\xfe\x13\xf1\x1a\x51\x9d\x9d\xbf\x62\xb7\x75\x85\x06\xcf\x1e\x66\x25\xfc\xe7\x8f\x6f\xc9\xeb\xf3\x0a\xec\xfd\xd0\x23\x0a\x4b\x0c\xee\xff\x8f\xbf\xb1\x6d\xa9\xfd\x50\xdf\xf9\xa5\xcb\xa8\x2f\xbd\x09\xe9\x66\x23\x7d\x87\xf4\x10\xe6\x3f\xca\xd4\xe5\x4b\x00\x5b\x8f\x25\x67\x2d\x27\x24\x3d\xbc\xc0\x43\x45\x39\x23\xbe\x03\x5d\x2e\x82\x62\x70\x0b\x5c\x17\xd2\xd1\xd9\xaf\xe3\xa7\x0b\xb9\xbb\xf4\x00\xd7\xc8\x3d\x92\xf4\xa0\x5c\xb3\x90\xc9\x84\x93\xf1\x30\x1d\x9d\x25\x6f\x4e\xd3\xc3\x61\xb6\x56\xc9\x62\xc6\x87\xd5\xe5\xca\xcb\xcd\xab\x37\xff\x00\xe8\xbf\x66\xf7\x07\x0f\xcb\x64\xcd\xec\x3e\x95\x6a\x65\xec\xe5\x60\x36\x71\xe9\x47\x64\xe1\x4c\x1b\x65\xdf\x57\xf9\x63\xde\x1b\x91\x85\xca\xdd\x29\xe6\x32\xda\x64\xf0\x42\x7b\xff\xfd\x55\x72\xf4\xc1\xb6\xe2\xdf\x01\x00\x00\xff\xff\x5b\x4f\x30\x44\x60\x05\x00\x00")

func tplApiTplBytes() ([]byte, error) {
	return bindataRead(
		_tplApiTpl,
		"tpl/api.tpl",
	)
}

func tplApiTpl() (*asset, error) {
	bytes, err := tplApiTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/api.tpl", size: 1376, mode: os.FileMode(438), modTime: time.Unix(1745815903, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplApi_commonTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xae\xcc\x2b\x49\xac\x50\xb0\x55\x50\x2a\x33\x54\xe2\xe5\xe2\xe5\x2a\xa9\x2c\x48\x55\xd0\xe0\xe5\x52\x50\x50\x50\xf0\x4c\x09\x4a\x2d\x54\xa8\x86\x70\x20\x02\x0a\x99\x79\x25\x66\x26\x0a\x09\x05\x89\x25\x19\x56\x4a\x99\x29\x4a\x09\x0a\xfa\xfa\x0a\x99\x29\x10\x35\xb5\x30\x7d\x3e\x99\xc5\x25\x18\x7a\x41\x82\x0a\xd1\xb1\x50\x13\xb2\x8a\xf3\xf3\x40\x26\xc4\xe7\x64\x16\x97\xe8\xe4\x17\x94\x64\xe6\xe7\x25\xe6\x28\x29\xa4\xe5\x17\xe5\x62\x11\x87\xda\xf3\xb4\x63\xfa\x8b\x85\x2b\x50\x6c\x0b\x48\x4c\x4f\x0d\x4a\x2d\x2e\xcd\x29\x41\xb6\x0e\x24\x0a\x66\xa0\x58\x57\x90\x98\x9e\x8a\x61\x17\xaa\x20\xd8\xa2\x97\x0b\xb7\x3e\x5f\xd0\x88\x6a\x5a\x70\x66\x55\x2a\xa6\x69\xf1\xc5\x99\x55\xd8\x8d\x44\x93\x01\x9b\xfb\x6c\x7d\xff\xcb\x85\x5b\x9f\x4d\xdd\xf0\xb2\xbd\x1f\x61\x7a\x48\x7e\x49\x62\x0e\x86\x5b\x4b\x40\xa2\x18\x26\xa3\x89\x42\x4c\x6d\xd8\xfd\x6c\xea\x06\x78\xa0\x68\x02\x02\x00\x00\xff\xff\x1c\x51\x25\x6b\xd5\x01\x00\x00")

func tplApi_commonTplBytes() ([]byte, error) {
	return bindataRead(
		_tplApi_commonTpl,
		"tpl/api_common.tpl",
	)
}

func tplApi_commonTpl() (*asset, error) {
	bytes, err := tplApi_commonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/api_common.tpl", size: 469, mode: os.FileMode(438), modTime: time.Unix(1731635957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplApi_serviceTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\xb1\x0a\x02\x31\x10\x04\xd0\x3a\xfb\x15\xc3\xd8\x68\xa3\xd8\x1e\xf8\x0b\xfe\xc3\xaa\x2b\x2e\x24\x97\xe3\x6e\x15\x45\xfc\x77\x49\x6c\xec\xde\xc0\xcc\x2c\xaf\x31\xf4\x89\x03\xf8\xd8\x53\xc4\xc7\x6b\xc5\x5a\x52\x78\x64\x1b\x00\x76\x50\xd2\xc5\x96\xf3\x00\x80\x0d\x94\xa4\xf7\xb8\xd5\x79\x00\x7f\xa0\x24\x2b\xea\xb9\x4d\x3a\x28\x1b\x11\x2f\x53\x9d\xa3\xfd\x71\xf5\x0e\x3d\x65\x3b\x6a\xb1\xcf\xee\x3f\x6c\x75\xf2\x5e\xfe\x06\x00\x00\xff\xff\x6f\xee\xa5\x1b\x8a\x00\x00\x00")

func tplApi_serviceTplBytes() ([]byte, error) {
	return bindataRead(
		_tplApi_serviceTpl,
		"tpl/api_service.tpl",
	)
}

func tplApi_serviceTpl() (*asset, error) {
	bytes, err := tplApi_serviceTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/api_service.tpl", size: 138, mode: os.FileMode(438), modTime: time.Unix(1731743545, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplProtoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x6e\xd3\x40\x10\xc6\xef\x96\xfc\x0e\xa3\xe4\x92\x5c\xb0\x4a\x11\x97\x28\x07\x28\x12\x8a\x84\x50\xe4\x96\x33\x5a\xe2\x69\x6a\x61\x7b\x5d\xef\x16\x51\x22\x4b\x54\x08\x1a\x4e\x39\xd0\x82\x10\x28\x34\x15\x20\x4e\x24\x97\x1e\xaa\x1e\xf2\x32\xf5\x12\xbf\x05\xda\x8d\x9d\xd8\x6e\xfe\x48\xbd\x79\x77\x67\x7e\xdf\xce\xf8\x9b\x65\x87\x1e\x27\xaf\xa1\x0e\x25\x3f\xa0\x9c\x6e\x96\x6a\xba\xa6\x6b\x3e\x69\xbd\x24\x6d\x84\x72\x27\xf9\x7a\x4a\x5c\x0c\x6b\xba\x46\x7d\x6e\x53\x0f\xda\xf4\x79\x1a\x52\x87\xd2\x1d\x23\x1f\x37\x65\x18\x06\x6c\x51\xd7\xa5\x9e\xae\xb9\xc8\x98\x8c\x6d\x58\x26\xee\x43\x47\xd7\x00\x00\x6c\x8f\xdf\xbf\x07\x0d\x0b\xea\xb0\x51\x03\x00\xc3\x00\xdb\xd2\xb5\x70\x1e\xbe\xcd\x03\xdb\x6b\x9b\xc8\xd2\x14\xa6\x36\xc0\x44\xe6\xa7\x59\x32\x6d\x32\x3e\x89\xbe\xf5\xaf\xc7\x03\x71\x34\xcc\x01\x1a\xd6\x13\x9b\xf1\x8c\x66\x80\x3e\x12\x8e\xd6\x4c\x5c\x9e\x27\x28\xa5\x1f\x75\xbf\x4c\x06\x7f\x72\x90\x26\x69\xa3\x89\xec\xc0\xe1\xf9\x9b\x37\xa7\xd5\x6f\xd4\x0c\x03\xe2\xc1\xc5\xbf\xb3\xa3\xe2\xe9\xb6\xfd\x06\x01\xea\x70\x57\x86\x88\x61\x2f\x1e\x5c\x88\xd3\x51\x7c\xdc\xcb\x06\xee\x50\x4e\x1c\xa8\xc3\x66\x4d\xde\x40\xbc\xbd\x12\xa7\x23\xa5\xaf\x3a\x58\xee\x70\xf2\xc2\x99\xb6\x35\x5d\xc8\xae\xa2\xc7\x43\xd8\xa2\xfe\x21\x70\x0a\x7c\x0f\xc1\x25\xb6\x07\xea\x17\xc2\xae\xed\xa0\xae\x31\x0c\x5e\xd9\x2d\xcc\x13\x92\x02\x0c\x43\xfc\xf8\x35\x19\x9e\x17\x80\x69\xf1\xaa\x53\x7e\x0b\x1e\x23\x2f\x77\x76\x66\xd9\xb2\x57\x95\xec\x86\x89\xfb\x55\x08\x90\x1f\x04\x1e\x83\x4a\x31\xd4\x44\x56\x55\x46\x58\xa1\x38\x19\xfe\x16\xef\xde\x2f\x55\x6c\x78\xbb\x14\x2a\xca\x35\xd5\x85\x3a\x32\xa0\xa8\xf3\x79\x14\x9d\xf7\x0b\x3a\x73\x85\x07\x96\x95\x05\x2c\xc6\xcd\x8b\x9a\x59\x30\x2b\x11\x75\xcf\xe2\xaf\x3f\x97\x4a\x3c\x42\x27\x2f\x31\x2d\x60\x0d\x54\x7c\xbc\x8c\x8f\x7b\x6b\xd0\x0f\x09\x6f\xed\x2d\xe0\x27\x36\x5f\xa7\x71\x3d\xfe\x2b\x4e\x2e\x97\xd2\x9f\xf9\x16\xe1\x78\xeb\xf6\x28\xcf\xa6\x53\x53\xf0\x89\x74\x5e\xb2\x25\x29\x61\x6e\xd5\x94\xb6\x95\x03\x13\xae\x80\x24\xe2\x37\x41\x4b\x33\x12\x13\xde\x98\xfd\x45\xd4\xfc\x3b\xb0\x62\x2e\x32\xaf\x41\xe6\x33\x99\xf1\xa8\xfb\x41\xbe\x04\x57\x9f\x44\xff\xbb\xba\xd8\xff\x00\x00\x00\xff\xff\x76\x27\x02\x82\x5e\x05\x00\x00")

func tplProtoTplBytes() ([]byte, error) {
	return bindataRead(
		_tplProtoTpl,
		"tpl/proto.tpl",
	)
}

func tplProtoTpl() (*asset, error) {
	bytes, err := tplProtoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/proto.tpl", size: 1374, mode: os.FileMode(438), modTime: time.Unix(1731747525, 0)}
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
	"tpl/api.tpl":         tplApiTpl,
	"tpl/api_common.tpl":  tplApi_commonTpl,
	"tpl/api_service.tpl": tplApi_serviceTpl,
	"tpl/proto.tpl":       tplProtoTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"tpl": &bintree{nil, map[string]*bintree{
		"api.tpl":         &bintree{tplApiTpl, map[string]*bintree{}},
		"api_common.tpl":  &bintree{tplApi_commonTpl, map[string]*bintree{}},
		"api_service.tpl": &bintree{tplApi_serviceTpl, map[string]*bintree{}},
		"proto.tpl":       &bintree{tplProtoTpl, map[string]*bintree{}},
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
