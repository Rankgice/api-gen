// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package utils generated by go-bindata.// sources:
// tpl/api.tpl
// tpl/api_common.tpl
// tpl/api_service.tpl
// tpl/proto.proto
// tpl/proto.tpl
// tpl/proto_common.tpl
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

var _tplApiTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcf\x6b\x13\x41\x14\xc7\xef\x0b\xfb\x3f\x3c\x36\x97\x04\x24\x8b\x20\x1e\x02\x42\xfd\x01\x12\x10\x91\xa8\x27\x11\x3b\xcd\xbc\xa4\x23\xbb\x33\xeb\xcc\xa4\xb4\x0d\x39\x88\x68\xe3\x29\x07\x7f\x21\x48\x6d\x8a\x8a\x27\xd3\x4b\x0f\xc5\x83\xff\x8c\x59\x93\xff\x42\x66\x77\xda\x6e\x36\x9b\x25\x39\x4d\xde\x7b\xf3\xfd\x7e\xe6\xbd\x9d\x51\x7b\x5c\x93\x5d\xb8\x01\xde\xce\x55\xcf\x75\x5c\x87\x85\x91\x90\x1a\xbc\x7a\xdd\x6f\x8b\x30\x14\xbc\x4e\x22\xe6\xb9\xce\x86\x42\xb9\x83\x12\xaa\xae\x03\x00\x10\x49\xec\xb0\xdd\x06\x54\xfa\xe9\x6a\x90\x86\xbb\x52\xf4\x22\x13\xd5\x64\x2b\xc0\xfb\x24\xc4\x81\xeb\xd4\x5c\xc7\x6c\x66\x6d\x84\x4a\xdf\xae\x92\x14\xf4\xd3\x5d\xbe\x1f\x7f\xfd\x3e\x9b\x1c\xdb\x6d\xb7\x45\x18\x22\xd7\x83\xe9\xf0\xd3\x6c\xfc\x33\x2d\xd9\xd8\x26\x9c\x06\x28\xe1\x2e\xea\x4a\xff\xd1\x85\xfa\x3d\xa6\xb4\xb5\x46\x0d\x7e\xd6\xd8\x0f\x98\xd2\x50\xcd\x56\xb7\xf0\x45\x0d\x24\xea\x9e\xe4\x6a\x31\x63\x74\x5a\xa8\x6a\xa6\x05\x25\x48\xb3\xc9\x8f\xf8\xd5\xeb\x72\xa4\x26\xef\x88\x55\x48\x8c\x77\x84\xdf\x60\x14\xaa\x4d\xba\x1a\xc6\x28\xe4\x60\x3e\x9e\x4c\x8f\x0f\x73\x30\x39\x8c\x9b\x94\x66\x45\xec\x9c\x84\xca\x33\x10\x4a\x8b\xed\x2e\x61\x94\x96\x8c\x77\x33\xfe\xd3\xe1\xd1\xfc\xf3\xb7\x72\xff\x3b\x18\x2c\xfb\x53\x0c\x50\x63\x8e\x20\x0d\x16\xf7\x61\xc9\x3a\x7e\x7b\x36\x3f\x18\xad\x03\x70\x8b\xe8\xf6\xf6\xfa\x14\x5b\xb6\xdc\x64\xaa\x4d\x9a\x7e\x01\xa5\x28\x7f\xff\xfc\x8a\xdf\x9f\x95\x43\x3c\x8e\x28\xd1\x58\x30\x88\x5e\x7e\x0e\xbd\xa4\x72\xed\x51\x0c\x5c\x47\xef\x45\xe6\x0a\x2d\x7c\xcf\xe6\x0e\xd9\x90\xd9\x6c\xdd\x1e\x90\x2e\xc2\xf9\x8f\x71\x7d\xfd\x1a\xc0\xe6\x73\x25\x78\xc3\x8b\x48\x17\xaf\x88\x48\x33\xc1\x49\xe0\x41\x47\xc8\x30\x1f\xdc\x04\xdf\x87\xf9\xf8\xf4\xdf\xd1\xcb\x4b\xb9\x87\x6c\x1f\x57\xc8\x3d\x53\x6c\xbf\x58\x33\x97\x49\x84\xe3\xc9\x68\x3e\x3e\x8d\x3f\x9c\xcc\x0f\x46\xc9\xb1\x0a\x0e\x66\xfb\xb0\x7c\xb8\xe2\x72\x7b\x7b\xcf\xdf\x13\xf3\xd7\x9e\xfd\xc9\xd3\x22\x59\xcb\x6e\x1e\x88\x25\xec\xc5\x60\x42\x5c\xf2\x2e\x99\xce\xb4\x50\xf5\x02\x9d\x5d\x66\x7b\x23\x93\x50\x71\x77\xf2\xb9\xc4\x6d\x3a\x7c\x63\x7a\xff\xfb\x5d\x7c\xf8\xc5\x75\x06\xff\x03\x00\x00\xff\xff\xb8\xa8\xf3\xca\xa2\x05\x00\x00")

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

	info := bindataFileInfo{name: "tpl/api.tpl", size: 1442, mode: os.FileMode(438), modTime: time.Unix(1731636774, 0)}
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

var _tplProtoProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xcf\x8a\xd4\x40\x10\xc6\xef\x81\xbc\x43\x91\xbd\xcc\x5e\xd2\xa8\xc7\x90\x83\x7f\x40\x16\x44\x96\x41\xcf\xd2\x26\x65\x0c\xa6\xff\xd8\x5d\x2b\xca\x90\x93\xa8\xe3\xc9\x83\xa2\x08\xb2\xba\x8b\x8a\x27\xd7\x8b\x87\xc1\x83\x2f\x63\xe2\xcc\x5b\x48\xcf\x64\x9c\x18\x33\x11\xf7\x16\xaa\xaa\x7f\x5f\x7d\xc5\x17\xfb\x40\x12\xbf\x0f\x31\x04\xda\x28\x52\xe7\x82\xc8\xf7\x7c\x4f\xf3\xe4\x0e\xcf\x10\x08\x85\x2e\x38\x61\xe4\x7b\x4a\x53\xae\x24\x64\xea\x46\xd3\x8c\x83\x90\xa9\x04\xb9\x44\x99\xe5\x12\xdd\xc3\x5c\x68\x65\x08\x82\x30\x64\x89\x12\x42\xc9\x70\x09\x75\x2d\x8b\xe6\x5e\x9e\x20\xec\x4c\x88\xdf\x2c\xf0\x2a\x17\x58\xc2\xc4\xf7\x00\x00\x18\xab\xdf\x7e\x98\x9f\x1c\x37\xbd\x8b\x4a\x08\x94\x54\x56\xd3\x57\xf3\xa3\x4f\xab\x11\xa3\x13\xb8\x8c\xb4\x33\xb9\xf6\xfb\xf5\x95\xdc\xd2\xa8\x5d\x18\xe3\xdd\x5d\x30\x48\x07\x46\x5a\x18\x75\x47\xc7\x68\x77\x97\xde\x06\x14\xe7\x27\x1f\xeb\x87\x8f\xb6\x2a\xee\xc9\x5b\x0a\x46\x7b\xa9\x13\xea\xd5\x71\x03\x5d\x9d\x97\x5f\xaa\xe3\xc3\x8e\xce\x46\xe1\x7c\x9a\xb6\x01\xfd\xb8\x8d\x29\x4b\x26\x97\x59\x9b\x5f\x4d\xdf\x2d\x5e\xbf\xdf\xca\xbf\x84\xc5\x9f\xfc\xd5\xf6\x43\xc4\xfa\xe9\x6c\xf1\xe4\xd9\x3f\xb8\x17\x38\x25\xb7\x7b\xe0\xab\x4b\x0f\x0b\xfc\xf8\xfe\xb9\x7e\x31\xdb\x8a\xbe\xae\x53\x4e\x78\xba\xab\x94\xbe\xc7\x58\x27\x13\xbe\x27\xd0\x5a\x97\xe6\x4e\xc3\xc5\xaf\x29\x39\xe6\xbe\x8b\x6a\xf9\x77\x69\x9f\x67\x58\xfe\x0f\xba\x49\xdb\x3a\xdd\x06\x35\x72\xc2\x14\x7a\x4c\x80\x9b\x85\x18\xce\x44\xee\x30\x30\xf0\x03\xb8\x2d\xc6\x68\x0f\x0a\x6a\x7f\xc6\x70\x36\x62\x0c\xaa\xe9\xe3\xc5\xd1\xd7\x9f\xdf\x9e\xd7\x87\x6f\x96\x9b\xf6\xad\xb5\x96\xec\x77\x5d\xfe\x0a\x00\x00\xff\xff\x3a\xf1\x23\xac\x09\x04\x00\x00")

func tplProtoProtoBytes() ([]byte, error) {
	return bindataRead(
		_tplProtoProto,
		"tpl/proto.proto",
	)
}

func tplProtoProto() (*asset, error) {
	bytes, err := tplProtoProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/proto.proto", size: 1033, mode: os.FileMode(438), modTime: time.Unix(1730864879, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplProtoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x6b\x13\x41\x14\xc7\xef\x0b\xfb\x3f\x3c\xd2\x4b\x7a\x0a\xb5\xe2\x65\xd9\x83\x3f\x40\x02\x22\x21\xad\x67\x19\x33\xcf\x38\x98\xec\xac\x33\x53\x51\xc3\x82\x22\xda\x78\xca\x41\xab\x08\x52\x9b\xa2\xe2\xc9\xe4\xd2\x43\xe9\xc1\x7f\xc6\x8c\xd9\xff\x42\x66\x77\x27\xd9\xdd\x66\x13\xf0\xb6\xbc\x79\xef\xfb\x79\xdf\xc7\x37\x91\xcf\x03\x45\x9e\x81\x0f\xb5\x50\x70\xc5\x77\x6b\x9e\xeb\xb8\x4e\x48\x3a\x8f\x49\x17\x61\x6b\x90\x7d\x45\x9e\xeb\xf0\x50\x31\x1e\x40\x97\xdf\xb7\xcf\x3e\xd4\x96\x1d\xe9\xa4\x44\xf1\x94\x75\xcc\xa4\x22\x0f\x7a\x78\x97\xf4\x31\x82\x81\xeb\x00\x00\x34\x1a\xfa\xeb\xf7\xf9\xe4\x34\x7b\xbb\xc9\xfb\x7d\x0c\x54\x34\x1b\x7e\x9a\x8f\x7f\xa6\x2d\x22\xec\xc0\x6d\x54\x5b\x83\xfd\xc5\xf4\x1d\x26\x55\x3d\x5f\x68\xe3\x93\x6d\x10\xa8\x0e\x44\x20\xa1\x5e\x6e\x6d\xa3\xdc\x4e\x36\x59\x43\x9c\x4f\x7e\xe8\xd7\x6f\x2a\x89\xcd\xe0\x21\x87\x7a\x93\x1a\xd0\x4a\x8e\x69\x28\x73\x3e\x4e\x67\xa7\xc7\x25\xce\x92\x70\x9d\xd2\xbc\xc0\x6a\xb9\xa5\xa9\x3d\x25\x58\xd0\x2d\x21\x66\xc3\x93\xf8\xf3\xb7\x4a\xc4\x2d\xec\x15\x11\xa9\x81\x0d\xa2\xfa\xdd\x79\x7c\x38\xda\x20\x7d\x83\xa8\xce\xa3\x15\xfa\xe9\xbd\x37\x32\xfe\xfc\xfe\xa5\x3f\x9c\x57\xaa\xdf\x0b\x29\x51\xf8\xdf\xe7\x89\x0c\xa8\x8f\x52\xa6\x81\x2d\xe4\xc4\x24\x2f\x2b\x19\x95\x96\x89\x78\x74\xb9\xd4\x32\xf9\x5d\xa3\x94\x6d\x50\xa1\x56\x39\x96\xc5\xd1\xc6\x5f\x60\x88\x44\x21\x5d\x29\x6d\x7a\xc1\x87\x1d\xcf\x1c\x0c\xd6\xfc\x42\xcc\xae\x6d\x94\x07\x3d\x95\xff\xf4\xe1\x8a\xd7\x68\xc0\x6c\xf8\x36\x1e\x9f\xfd\xbd\x78\xaf\x8f\xbf\x94\x16\x4b\xd2\x60\x57\x61\x81\xba\x76\x15\x9a\x34\x23\x1a\x24\xa3\xc9\x80\x6d\x5f\x9c\xd8\x8e\xc8\xa4\x00\x02\x65\x98\x4c\x15\xba\x17\x51\xb8\xe4\xd5\x92\x8a\xfe\x18\xb5\x96\x72\x22\x39\x3b\x85\x35\x5b\xe9\x5f\xcd\x8e\x31\x68\xdc\x9d\xbc\x2a\xbf\xee\xb1\x17\x08\xf6\x06\x7a\x32\x8a\xc7\x67\xfa\x68\x1a\x1f\x8e\xf2\x8d\xfb\x5c\x91\x1e\xf8\xb0\xeb\x99\x0d\xf4\xcb\x0b\x7d\x34\x75\x9d\xe8\x5f\x00\x00\x00\xff\xff\x28\xfc\x14\xb3\x01\x05\x00\x00")

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

	info := bindataFileInfo{name: "tpl/proto.tpl", size: 1281, mode: os.FileMode(438), modTime: time.Unix(1731741728, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tplProto_commonTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xae\xcc\x2b\x49\xac\x50\xb0\x55\x50\x2a\x28\xca\x2f\xc9\x37\x56\xb2\xe6\xe5\xe2\xe5\xca\x4d\x2d\x2e\x4e\x4c\x4f\x55\xf0\x4c\x09\x4a\x2d\x54\xa8\xe6\xe5\x52\x50\x50\x50\xc8\xcc\x2b\x31\x33\x51\xf0\x4c\x51\xb0\x55\x30\xb4\x56\x50\x50\xd0\xd7\x57\xc8\x4c\xe1\xe5\xaa\x45\x28\x0f\x2e\x29\xca\xcc\x4b\x0f\x4a\x2d\x86\x69\x29\x06\x0b\x28\x14\xa5\x16\x17\x80\x75\xa1\xa8\xf6\x4c\xf1\xc9\x2c\x2e\x41\xb2\xa0\x28\xb5\x20\x35\xb1\x24\x35\x05\x6e\x13\x48\x1e\x6a\x1b\xd8\xb2\xa7\x1d\xd3\x5f\x2c\x5c\x81\x62\x48\x40\x62\x7a\x6a\x50\x6a\x71\x69\x4e\x09\xaa\x33\x41\xe2\x60\xad\xfa\xfa\x0a\x2f\x17\x6e\x7d\xbe\xa0\x11\x5d\x36\x38\xb3\x2a\x55\x41\xc1\x56\xc1\x08\xa4\xe4\xd9\xfa\xfe\x97\x0b\xb7\x3e\x9b\xba\xe1\x65\x7b\x3f\xb2\xc2\x90\xfc\x92\xc4\x1c\x05\x5b\x05\x63\x6b\x90\x0b\x9e\x35\xec\x7e\x36\x75\x03\x2f\x57\x2d\x20\x00\x00\xff\xff\x26\xb9\x29\x40\x36\x01\x00\x00")

func tplProto_commonTplBytes() ([]byte, error) {
	return bindataRead(
		_tplProto_commonTpl,
		"tpl/proto_common.tpl",
	)
}

func tplProto_commonTpl() (*asset, error) {
	bytes, err := tplProto_commonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/proto_common.tpl", size: 310, mode: os.FileMode(438), modTime: time.Unix(1731741566, 0)}
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
	"tpl/api.tpl":          tplApiTpl,
	"tpl/api_common.tpl":   tplApi_commonTpl,
	"tpl/api_service.tpl":  tplApi_serviceTpl,
	"tpl/proto.proto":      tplProtoProto,
	"tpl/proto.tpl":        tplProtoTpl,
	"tpl/proto_common.tpl": tplProto_commonTpl,
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
		"api.tpl":          &bintree{tplApiTpl, map[string]*bintree{}},
		"api_common.tpl":   &bintree{tplApi_commonTpl, map[string]*bintree{}},
		"api_service.tpl":  &bintree{tplApi_serviceTpl, map[string]*bintree{}},
		"proto.proto":      &bintree{tplProtoProto, map[string]*bintree{}},
		"proto.tpl":        &bintree{tplProtoTpl, map[string]*bintree{}},
		"proto_common.tpl": &bintree{tplProto_commonTpl, map[string]*bintree{}},
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
