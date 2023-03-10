// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package views generated by go-bindata.// sources:
// web/views/article.html
// web/views/categories.html
// web/views/errors/404.html
// web/views/errors/500.html
// web/views/home.html
// web/views/layouts/layout.html
package views

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _articleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x49\xc9\x2c\x53\x48\xce\x49\x2c\x2e\xb6\x55\x2a\xc8\x2f\x2e\xd1\x2d\x48\x4c\x4f\x55\xb2\xe3\x52\x50\xa8\xae\xd6\x73\x2c\x2a\xc9\x4c\xce\x49\xad\xad\xe5\xb2\xd1\x4f\xc9\x2c\xb3\xe3\xe2\x02\x2b\xcf\x4c\xb1\x55\x4a\xcf\x2c\x49\xcc\xc9\x56\xb2\x83\x4a\x00\x02\x00\x00\xff\xff\x68\x41\x05\xf7\x47\x00\x00\x00")

func articleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_articleHtml,
		"article.html",
	)
}

func articleHtml() (*asset, error) {
	bytes, err := articleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "article.html", size: 71, mode: os.FileMode(420), modTime: time.Unix(1673419221, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _categoriesHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x41\x6a\xeb\x30\x10\x86\xf7\x39\xc5\xe0\x7d\xac\x0b\x28\x86\xf0\xde\xe6\x41\x78\x3c\x78\xb9\xc0\xd4\x9e\xd8\x43\x65\xd9\x58\x4a\xb2\x10\xda\xb6\xbd\x4d\xe9\xba\xe7\x29\xb9\x46\xb1\x2c\x05\x27\x6d\xa1\x85\xd2\x9d\x46\xfa\xc7\xf3\xff\xf3\xd9\x39\xde\x41\xfe\xbf\xe9\x8e\x6b\xa5\xbc\x5f\xc8\x8a\x0f\x50\x2a\x34\x66\x95\xe1\x50\x36\x7c\x20\x40\xcd\x2d\x5a\xaa\x60\x87\x15\xfd\xd1\xbf\xbb\xa3\xce\x8a\x05\x80\xec\x93\xb2\xef\x8c\x5d\x5a\xb6\x8a\xb2\xe2\xe5\xe1\xee\xf4\xf4\x0c\x52\xf4\x41\xb3\x57\x49\xa4\xd8\x58\xd6\x75\x68\x05\x70\x6e\x40\x5d\x13\xe4\xdb\x81\x68\x3d\x58\x2e\x15\x19\xef\xc3\x9b\x54\x7c\xd5\xb4\x64\x4b\x6d\xec\x04\x98\x9b\x4c\x82\xd1\xc2\x59\xf0\x91\x37\x89\xd0\x0c\xb4\x5b\x65\xa2\x44\x4b\x75\x37\x30\x19\xe1\x5c\xfe\x2b\x55\x7f\xb1\x25\xef\x33\x08\xfa\x55\xf6\xce\x53\xf1\xf6\x4e\x0a\x9c\x0d\x9e\x72\xc7\xc2\xf4\xa8\x93\x91\x0a\x2d\x41\x4b\x16\x63\x18\xe7\x6a\xb2\x29\xf9\x86\x34\xe4\x1b\x36\xd6\x7b\x38\x3d\xde\x4b\x31\x76\x9e\xf3\x8a\x8a\x0f\x53\x21\x85\xe2\xb4\x40\xd2\x55\x58\x98\x14\x7b\x55\x2c\xa2\xc8\x39\x52\x86\xbe\x93\xe4\x3c\xf0\x18\xf6\x93\x5c\x6d\x43\x53\x9e\x9f\x42\x8a\xd3\x2a\x9d\xcb\x37\xac\x6f\x2f\x29\x6e\xc7\x53\xa4\x17\xcf\x23\xb5\x4b\x58\x33\x0b\x71\x42\x4b\xb3\xf9\x57\x38\x71\x58\x8e\x44\xc3\x27\xff\xed\x6f\x14\x9b\x66\xcb\xd3\xef\x30\x67\x77\x41\xef\xcb\x28\xc3\xf5\x6b\x00\x00\x00\xff\xff\xea\x97\xce\x28\xa4\x03\x00\x00")

func categoriesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_categoriesHtml,
		"categories.html",
	)
}

func categoriesHtml() (*asset, error) {
	bytes, err := categoriesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "categories.html", size: 932, mode: os.FileMode(420), modTime: time.Unix(1673409500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errors404Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x0a\xcd\xe7\x79\x69\xa0\x87\x1d\xec\xc0\xe8\x5a\xd8\x65\xdb\xa1\x85\xed\xe8\xda\x3f\xb5\xc0\xb1\xb3\x54\x4d\xd9\xdb\x8f\x34\x1d\x74\x3b\x09\x7d\xfa\x3f\x81\x64\xee\x9e\xdf\x56\xdb\xcf\xf7\x35\x45\xe9\x52\x5b\x99\xa9\x50\x72\xf9\x60\x15\xb2\x9a\x00\x5c\x68\x2b\x22\x22\xd3\x41\x1c\xf9\xe8\x86\x23\xc4\xaa\xdd\x76\xa3\x1f\xd5\xed\x28\x8a\xf4\x1a\x5f\x27\x1e\xad\xfa\xd0\xbb\x27\xbd\x2a\x5d\xef\x84\xf7\x09\x8a\x7c\xc9\x82\x2c\x56\xbd\xac\x2d\xc2\x01\x7f\xcc\xec\x3a\x58\x35\x32\xce\x7d\x19\xe4\x26\x7c\xe6\x20\xd1\x06\x8c\xec\xa1\x2f\xcd\x3d\x71\x66\x61\x97\xf4\xd1\xbb\x04\xdb\x3c\x2c\x7e\x57\x09\x4b\x42\xbb\x5c\x2c\xe9\xb5\x08\x6d\xca\x29\x07\x53\xcf\xb0\x32\xf5\x7c\x88\xd9\x97\xf0\x7d\xcd\xc7\xa6\x35\x1e\x59\x30\xfc\x97\xae\xd4\xd4\xb1\x99\xd4\xd9\x31\xf5\xe5\x47\x3f\x01\x00\x00\xff\xff\xf0\x70\x71\x97\x33\x01\x00\x00")

func errors404HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors404Html,
		"errors/404.html",
	)
}

func errors404Html() (*asset, error) {
	bytes, err := errors404HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/404.html", size: 307, mode: os.FileMode(420), modTime: time.Unix(1672904546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errors500Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4e\xc3\x30\x0c\x86\xef\x79\x0a\x93\xf3\x42\xd9\x8d\x43\x52\x09\x8d\x21\x71\x1a\x87\x4d\x82\xa3\xd7\x5a\x8d\xa5\xd4\x29\xad\xb7\x6a\x6f\x8f\xba\x0c\x69\xdc\x38\x25\xbf\xfc\xf9\xb3\x6c\xff\xf0\xba\xdb\xec\xbf\x3e\xb6\x10\xb5\x4f\xb5\xf1\xcb\x03\x09\xa5\x0b\x96\xc4\xd6\xc6\xf8\x48\xd8\xd6\x06\x00\xc0\xf7\xa4\x08\x4d\xc4\x71\x22\x0d\xf6\xb0\x7f\x73\xcf\xf6\xbe\x14\x55\x07\x47\xdf\x27\x3e\x07\xfb\xe9\x0e\x2f\x6e\x93\xfb\x01\x95\x8f\x89\x2c\x34\x59\x94\x44\x83\x7d\xdf\x06\x6a\x3b\xfa\xd3\x29\xd8\x53\xb0\x67\xa6\x79\xc8\xa3\xde\xc1\x33\xb7\x1a\x43\x4b\x67\x6e\xc8\x5d\xc3\x0a\x58\x58\x19\x93\x9b\x1a\x4c\x14\xd6\x8f\x4f\xbf\x2a\x65\x4d\x54\xef\x4e\xc3\x04\x53\xee\x49\x23\x4b\x07\x33\x89\xc2\x3c\x66\xe9\x56\xa0\xe3\x05\xb0\x43\x16\x5f\x15\xd6\xf8\xaa\xec\x67\xfc\x31\xb7\x97\x9b\x27\xae\xcb\xe7\x1a\x1a\x12\xa5\xf1\x9f\xd6\x1b\x5c\x34\xd5\xe2\xf1\x55\x11\x2f\x93\x96\x13\xff\x04\x00\x00\xff\xff\x89\x94\x40\x58\x72\x01\x00\x00")

func errors500HtmlBytes() ([]byte, error) {
	return bindataRead(
		_errors500Html,
		"errors/500.html",
	)
}

func errors500Html() (*asset, error) {
	bytes, err := errors500HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "errors/500.html", size: 370, mode: os.FileMode(420), modTime: time.Unix(1672904546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _homeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x31\x4e\xc4\x30\x10\x45\xfb\x9c\x62\x64\xea\x8d\x2f\xe0\xb5\x84\x68\x29\xb6\xd8\x0b\xcc\x26\xc3\xee\x88\xc4\x89\xec\x21\x08\x8d\x5c\xd1\x83\xb8\x0d\xe2\x3c\xec\x3d\x50\x12\x82\x88\x94\xd2\x7e\xfe\x4f\xff\x5b\x35\x62\x38\x13\x94\xb7\x51\xb8\x6a\x28\xe5\x5c\xb8\x9a\x07\xa8\x1a\x4c\x69\x6f\xfa\x2e\xc9\xae\xea\x82\x20\x07\x8a\xc6\x17\x00\xae\x5f\x41\x61\x69\xc8\x78\x87\x70\x89\xf4\xb0\x37\x16\x67\x93\x6a\x79\xcf\xe1\x31\x67\xe3\x55\xcb\xe3\xf8\x2a\x67\x67\xd1\x3b\xdb\x6f\x68\x5a\x12\x9c\xf4\x00\x2e\xf5\x18\x16\x58\xa3\x10\x8c\x70\xc7\x42\xad\xf1\xd7\xf7\x8f\xeb\xe7\xeb\xf7\xd7\x1b\xa8\x96\x87\xa7\x53\xc3\xe9\x72\xe4\x76\x72\x8f\xb9\x0d\xc5\xbf\xf4\x04\x01\x1c\x2e\x0c\x77\x82\x67\xb3\x54\xaf\x50\xe8\xdc\x45\xa6\x64\x55\xcb\xbb\xf9\xf4\x92\xb3\x35\x30\xcd\xdc\x9b\xd5\xb5\xf1\x37\xbf\x42\x80\x15\x18\x67\xce\x3d\xfe\x3a\x6d\xaf\xc6\x53\x92\x88\x95\x4c\x7f\x74\x88\x34\x30\x3d\x8f\xf1\xde\x17\xce\xd6\x3c\xf8\x42\x95\x42\x9d\x73\xf1\x13\x00\x00\xff\xff\x76\x3d\x73\x03\xa9\x01\x00\x00")

func homeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_homeHtml,
		"home.html",
	)
}

func homeHtml() (*asset, error) {
	bytes, err := homeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "home.html", size: 425, mode: os.FileMode(420), modTime: time.Unix(1673409500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4b\x8f\xdc\xc6\x11\xbe\xcf\xaf\x28\xf3\xb2\xbb\xc8\x92\xdc\x07\x10\x24\xda\xe1\x18\x7a\xc1\x16\x10\x4b\x82\xb4\x01\x12\x08\x82\x50\x43\xd6\x90\xbd\x4b\x76\x33\xdd\x3d\x33\x3b\x1e\x0f\x90\x4b\x82\x20\xbe\xe7\xae\x4b\x80\x00\xbe\xf8\x20\xc0\x40\x6e\xf9\x2d\x36\x0c\xfd\x8b\xa0\xbb\x49\x0e\x5f\x2b\xaf\x65\xc0\x27\xf6\xa3\x1e\x5f\x57\x57\x7f\xd5\xec\xe9\x27\x8f\x9e\x3d\xbc\xfc\xf3\xf3\xc7\x90\xe9\x22\x9f\x4d\xa6\xe6\x03\x39\xf2\x34\xf2\x88\x7b\xb3\xc9\x64\x9a\x11\x26\xb3\x09\xc0\xb4\x20\x8d\x10\x67\x28\x15\xe9\xc8\x5b\xea\x85\xff\x3b\x6f\x3f\xc1\xb1\xa0\xc8\xfb\x93\xff\xc7\xfb\xfe\x43\x51\x94\xa8\xd9\x3c\x27\x0f\x62\xc1\x35\x71\x1d\x79\x4f\x1e\x47\x94\xa4\x34\xd0\xc0\xa5\xce\x84\x74\xc3\x9a\xe9\x9c\x66\xdb\x2d\x04\x97\xa6\x05\xbb\xdd\x34\x74\x63\x93\x49\x4f\x2d\x21\x15\x4b\x56\x6a\x26\x78\xcb\xc9\x7d\x78\x90\x8b\x14\x9e\x8b\x35\x49\x4a\xe0\xc1\x06\x3e\xa7\x1b\x31\xf0\x79\x4d\x9b\xb5\x90\x89\x6a\x69\xb6\x64\x9a\xb1\x35\x4b\x74\x16\x25\xb4\x62\x31\xf9\xb6\x73\x0c\x8c\x33\xcd\x30\xf7\x55\x8c\x39\x45\xa7\xc1\xc9\x31\x14\x78\xc3\x8a\x65\xd1\x1e\x5a\x2a\x92\xb6\x8f\xf3\x9c\xa2\x13\xaf\xf2\xbb\x62\xb4\x2e\x85\xd4\x63\xce\x36\xa4\x6a\x31\x2c\xcb\x9c\xfc\x42\xcc\x59\x4e\xfe\x9a\xe6\x3e\x96\xa5\x1f\x63\x69\x8c\x8d\xa9\xce\x73\x8c\xaf\x3f\xa8\xac\x34\xea\xa5\xf2\xe7\x28\x7d\xa5\x37\xe3\x56\x34\xe5\x54\x66\x82\x53\xc4\x45\x6d\x6c\x21\x64\x81\xda\x4f\x48\x53\x6c\x43\xdd\x8f\xa4\x24\x9e\x90\x24\xd9\x8a\xe4\x9a\xe6\xd7\xac\x5a\x62\xce\xf8\x35\x48\xca\x23\x0f\x73\x4d\x92\xa3\x26\x0f\xf4\xa6\xac\x70\xb2\x18\x8d\xd5\x10\xb5\x28\x7e\x73\x53\xe4\x1e\xd8\xdd\x8e\xbc\xfb\x97\xcf\xbe\x80\xd3\xe0\xc4\x83\x4c\xd2\x22\xf2\xac\x44\x60\x24\xfa\x00\x52\xe2\x24\x51\x8b\x36\x02\xb3\xe7\xf0\xdb\xe0\x3c\x38\x31\x09\x0c\x30\xfd\xc4\xf7\x21\x56\x0a\x7c\xbf\x87\xca\x06\x43\x65\x44\xba\xf1\x64\x42\xc5\xe2\x30\x56\x2a\x9c\x0b\xa1\x95\x96\x58\x06\x05\xe3\x41\xac\x54\x7f\x51\x1f\x56\x5f\x08\xae\x7d\x5c\x93\x12\x05\x7d\x9c\x85\x52\x32\x55\xec\xd5\xcc\x32\xee\xac\x6b\x67\xfd\x04\xe5\xb5\x33\xd0\x59\x3c\x4b\x22\x4f\x67\x54\x90\x6f\xe6\xee\x6c\xae\x82\x62\x0f\xe3\x76\x0b\x6c\x01\xc1\x67\x4c\x63\x7e\x1d\x3c\xcc\x19\x71\xfd\xe4\x11\xec\x76\x3f\x67\x81\xa9\xd5\xae\x3e\xcd\x42\xb7\x5b\xe2\xc9\x6e\xe7\xbc\x18\x27\xf7\x39\xe6\x9b\x2f\x49\x06\x0f\x90\x25\x4b\xe7\xc2\x04\xc3\x76\x01\xcd\xac\x66\x71\xb3\xbf\x8e\x19\x4c\x13\x60\x85\x12\xde\x64\x85\x86\xc8\x7d\xbe\xfa\x0a\x5e\xbd\xbe\xb0\x53\x87\x8b\x25\xb7\x59\x0d\x87\x47\xb0\xb5\x43\x4e\x3e\x2b\x20\x82\x44\xc4\xcb\x82\xb8\x0e\x62\x49\xa8\xe9\x71\x4e\xa6\x77\xe8\x39\xe3\xde\xd1\x45\xa5\x90\x15\x81\x92\x31\x44\xe0\x65\x5a\x97\xea\x5e\x18\x66\x45\x30\x37\xc0\x82\x58\x14\xa6\x73\xa5\x3e\xdd\x6e\x07\x6b\xf0\x2e\x5a\x1e\x55\xdb\x61\x4a\xba\xf2\xa6\x1e\x6c\x2e\x31\x7d\x8a\x05\xed\xfd\xbe\x3a\x79\x5d\x6b\xaa\xa0\x44\x49\x5c\x3f\x15\x09\x05\x8c\x2b\x92\xfa\x01\x2d\x84\xa4\xc3\xac\x38\x06\x55\x41\xdc\x1d\x1d\xda\xd6\x34\xdc\xc7\xe5\xf6\x08\x7f\x26\x44\x9a\xd3\x3e\xc4\xae\x0f\x1a\x53\x38\x4c\x35\xa6\xc1\x95\x3a\xea\xc6\x19\x50\x6d\x78\x0c\x4a\xc6\x51\x13\x82\xf5\x7a\x1d\xa4\x56\x53\x63\x5a\x20\xc7\x94\xa4\x0d\x87\x31\x11\x5e\xa9\x4f\x59\x12\xb5\x63\x52\x7b\xf5\x66\x6d\x94\x9d\x8d\x5c\x33\x9e\x88\x75\x90\xa0\xc6\x3f\xe0\x86\x24\x44\xc3\xa1\xd6\xe6\x36\x7b\x6b\x3c\x9a\xfd\x85\x46\x2c\x28\x97\x2a\x3b\x44\x99\xda\x70\xab\xa3\x0b\xd8\x59\x15\x2b\x79\x70\xa5\x0e\x8e\x81\xd3\x1a\x1e\xa1\xa6\xc3\xa3\xa3\x8b\x49\x6b\x32\x16\x7c\xc1\xd2\x83\x63\x38\x18\x43\x7f\x70\x7b\x9c\x27\xd3\xd0\xd5\xd1\xc9\x74\x2e\x92\x8d\x5d\x5d\xc2\x56\x10\xe7\xa8\x54\xe4\x19\xea\x42\xc6\x0d\x8f\x9a\xa3\xa9\x34\x56\x95\xb2\x2b\x26\xc5\xba\x1a\xed\xab\xe7\xbe\x2a\xfc\x73\x30\x8d\x1b\xe5\x9f\x9e\x81\x62\x09\xf9\x8d\x59\x58\x60\x42\x4f\xf8\x0b\x96\x66\xba\x72\x61\xe6\xe7\x28\x1b\x7b\x5d\x8b\x2b\x92\x9a\xc5\x98\xfb\x9a\x6e\x34\x28\xa6\xc9\xb7\xd4\xdc\x12\x07\x98\x66\xe7\xb5\xfc\x5e\xc2\x57\x05\xe6\x86\xc9\x71\xce\x78\x42\x37\x91\xe7\x9f\x7a\xb3\x29\xd6\x92\x58\x19\xaa\xd9\xc0\xdb\xd7\xfb\x33\x5b\xf0\x71\x36\x0d\xb3\xf3\xae\x9f\xd3\x11\x3f\x39\xca\x94\x3e\xca\xcf\xde\xcd\xe9\x6c\xb2\x5f\x7e\x98\xb0\xd5\x6c\x3a\x97\x4d\x0c\x98\x32\x77\x18\x3f\x4f\xa1\x6e\x16\x49\xd3\x54\xc5\x2d\xa1\x6b\x43\x64\xfc\x5a\xd5\xf1\xd6\xe4\x73\x5c\x75\x03\xb8\xcc\xdb\x5d\x4b\x9c\xdd\x01\x73\x08\x4a\xe4\xb5\x6d\x93\x43\x75\x92\xd4\x6d\x29\x4c\xc5\xe4\xb8\x62\x29\x36\x35\xba\x67\xa3\x89\x0a\xc0\x95\xf2\x1d\xf3\xa3\x3b\x1f\xed\x4e\x1d\xac\x2b\x5c\xa1\x4b\xe2\x7b\x17\xde\x6c\xca\x6a\xed\x05\xc2\x02\x7d\xb5\xe4\xbe\x30\x27\x95\xcd\x4c\x18\xfb\x78\x43\x03\xb8\xb7\xac\xb0\xbf\x2e\xb3\xd0\x3d\xaa\xed\x96\xa9\xcf\x45\x41\x4f\x71\x75\x3f\xd6\x6c\x45\x10\x3c\x5c\x4a\x43\x6d\xcf\x51\x67\xe0\x41\xe8\xc1\x6e\xd7\xda\xca\xf7\xff\xfe\xd7\xfb\xb7\xef\xdc\x26\xfe\x52\xd3\x31\x6a\x4a\x85\x64\xe6\xfe\xd5\x72\xd2\x1a\x9e\xfd\xf0\x8f\xbf\xff\xf8\xed\x7f\x6f\x71\x67\x58\x72\xbb\x05\x89\x3c\x35\xc6\x1b\x35\x57\x0a\x07\xb8\x9c\x79\x97\x8d\x4d\x22\xf6\x8d\x56\xbc\x61\x89\x76\xe8\x6d\x9a\x37\x3b\xa2\x44\xec\xed\xad\xd6\xf4\x9b\x32\x9d\x2d\xe7\x96\x70\x0b\x41\x9b\x25\x9a\x63\x22\x53\x73\x6d\x7f\x33\xcf\x91\x5f\x57\x25\x9f\x0b\x51\x9a\xfb\x13\x70\x21\x69\x41\xd2\x5c\xe4\x06\xc9\x83\x92\xa1\x9f\xe3\xdc\x28\x7c\x66\x0d\x0f\x53\x22\xad\xc6\xa1\x4e\x8a\x01\x24\xbd\x66\x5a\x57\x45\xe0\xf1\x4a\x24\x5f\x38\x5c\x03\x6f\x77\xc6\xd9\xc1\x75\xe9\xac\x0f\x81\xe9\x7a\xa2\x85\x6c\xe0\x73\x8f\xf4\x5e\x58\x05\xcc\xe2\x6c\x2e\x9d\x1f\x89\xea\xc5\xcb\x97\x06\xd1\xc0\x1f\x74\x31\x4a\x7b\x3d\x6b\xf0\x99\x64\xa8\x2a\xec\xfe\xfc\x74\x89\xa2\xc3\xd4\x3f\xc1\x52\x50\x33\x8f\xbf\x10\xc2\xc5\x68\x4f\x71\xbd\x39\x7f\x2e\x7b\xec\x31\x75\x13\x83\x43\x5e\xce\xfe\xf7\x1f\x38\x3b\x39\x3b\x37\xdc\xd2\x8f\x4e\x95\xe2\x2e\x4a\x5e\xb5\xaa\x72\xc4\xc6\xed\xca\x23\x99\x9c\x51\x22\x78\xaa\xb2\x65\x98\x0a\xbf\x40\x79\x9d\x88\x35\xf7\xe7\x42\x5c\x37\x9e\x1e\x91\x62\x29\x1f\x21\x25\x80\xf9\xe6\x67\x39\x53\xcb\x82\x15\x78\xcd\xb4\x08\x33\xba\x11\x15\x65\xea\x4d\x29\x52\x89\x65\xb6\x69\x5c\x5e\x36\x43\xa3\x5c\x58\xf6\x89\x70\x18\x4f\x57\x77\x26\xa3\xdd\xaa\xe3\x6e\xdc\xfd\x9d\xaf\xaa\xfe\xef\x5b\x55\xbf\x40\xc6\x07\x55\xff\x52\x94\xae\x62\x74\x67\x6f\x29\x5f\xb8\xd4\xa2\x34\xf7\x35\xf6\x25\xbd\x31\x8d\x37\xe4\x2e\xa3\xde\x6d\x19\x58\xfd\x77\x79\x63\x54\xf5\xc2\xfe\x1f\x9a\x22\x03\xb1\xa3\x5d\xd0\x54\x94\x39\x6a\x82\x8c\x24\xf5\x12\xdd\xfe\x56\x6c\x18\xe5\x09\xb8\xdb\x69\xcf\x5c\xdb\x6d\x89\x29\xe3\xb6\xe2\x01\xb7\x09\x91\x32\x93\x0e\x5a\x8b\x62\x50\x02\xa7\x65\x03\x36\x27\x94\x0b\x76\xe3\xdd\x96\x92\x3f\x7e\xf3\x0d\x9c\xc2\xfb\xb7\xef\x20\x84\x1f\xfe\xf6\x2d\x9c\xd9\xf6\x70\x2b\xad\x74\xa7\x38\x97\x92\xa0\xc4\x74\xbe\xd4\x5a\x70\x35\xe4\x22\xe4\xa9\xf9\x23\x13\x4b\x7b\x5c\x69\xa1\x5d\x0d\x35\x79\x39\xa8\xe1\x4d\xa9\x1b\x21\x8f\x2e\x25\x7f\xff\xdd\x3f\xbf\xff\xee\xaf\xef\xdf\xbe\xf3\x66\x4d\xd3\x9d\x38\x5b\x87\x2d\xc2\x19\xb4\x7b\x35\x2a\x6e\xae\x76\x1d\xc0\x23\x40\x46\xfd\x57\xe0\x4c\x76\x84\x67\xa1\xd7\xc3\xf3\xf5\x1e\xcf\xd7\x7b\x3c\xf0\xc1\x70\x48\x7b\x33\xad\xee\x14\x0e\xea\x78\xc0\x87\x27\xca\x9c\x90\x01\x5f\xde\x7e\xa4\xc6\xd9\xf3\x46\x39\x9e\x74\xf9\xd3\x50\x65\xdb\xe6\x18\x13\xfe\x72\x1e\xfc\x35\x59\xf0\xd7\xe1\xc0\xde\x12\x87\x94\x37\xc6\x70\x9d\x66\xcd\x7a\xfb\x5f\x4c\xfb\x73\x59\x3f\x1c\x5c\xa9\xf0\xea\x2f\x4b\x92\x1b\xff\x3c\x38\x0d\x4e\xec\xd3\xca\x95\x1a\xfd\x75\x1c\x28\x76\x9f\x74\xee\xaa\x55\xb9\x2b\x58\x2a\x51\x93\x7f\x1a\x9c\x05\xa7\x1f\x61\x20\xc0\xb2\x24\x94\x77\xd6\x72\xcf\x3f\x77\x95\xde\x6f\xcc\x9d\x55\x4c\x41\xe8\x09\xff\xe4\xd3\xce\x2d\xa6\xaa\x27\x9c\x0f\x04\xc5\x6d\x72\x2c\xb8\xd2\xe0\xa4\x21\xb2\x3f\xda\xce\xcf\x61\xfd\x02\x13\x57\xfe\xee\xd9\x9f\xec\x1e\x88\xdd\xee\xe0\xb8\x23\xf7\x92\x62\x49\x7a\x44\xd6\x4d\xb4\xe4\x25\x95\xa2\x23\xf7\x82\x4a\xd1\x9a\x17\x6b\x4e\xb2\x23\xf0\xcc\x8c\xb4\x24\x30\x29\x18\xbf\x07\xaf\x46\x44\x5e\xd7\x32\x2c\xe9\x98\x78\x92\xb4\xf4\x2d\x43\x2a\x63\xc0\xad\xff\xe0\x75\xf5\x40\x53\x3d\x31\x38\x15\xf7\xa4\x7a\x58\xcb\x1c\x7d\xe0\x49\xc1\x3d\x25\x4c\xa6\xa1\x7b\xbc\xff\x7f\x00\x00\x00\xff\xff\xd3\x10\x57\x08\xcd\x17\x00\x00")

func layoutsLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_layoutsLayoutHtml,
		"layouts/layout.html",
	)
}

func layoutsLayoutHtml() (*asset, error) {
	bytes, err := layoutsLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/layout.html", size: 6093, mode: os.FileMode(420), modTime: time.Unix(1673409500, 0)}
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
	"article.html":        articleHtml,
	"categories.html":     categoriesHtml,
	"errors/404.html":     errors404Html,
	"errors/500.html":     errors500Html,
	"home.html":           homeHtml,
	"layouts/layout.html": layoutsLayoutHtml,
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
	"article.html":    &bintree{articleHtml, map[string]*bintree{}},
	"categories.html": &bintree{categoriesHtml, map[string]*bintree{}},
	"errors": &bintree{nil, map[string]*bintree{
		"404.html": &bintree{errors404Html, map[string]*bintree{}},
		"500.html": &bintree{errors500Html, map[string]*bintree{}},
	}},
	"home.html": &bintree{homeHtml, map[string]*bintree{}},
	"layouts": &bintree{nil, map[string]*bintree{
		"layout.html": &bintree{layoutsLayoutHtml, map[string]*bintree{}},
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
