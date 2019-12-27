// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 0001_app.down.sql (387B)
// 0001_app.up.sql (3.088kB)
// 0002_tokens.down.sql (19B)
// 0002_tokens.up.sql (248B)
// 0003_settings.down.sql (118B)
// 0003_settings.up.sql (1.341kB)
// doc.go (74B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __0001_appDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xcd\x0e\x82\x40\x0c\x84\xef\x3c\x05\xef\xc1\x49\x03\x07\x13\xa3\xc6\x78\xf0\xd6\xac\x4b\x85\x46\xd8\xae\x6d\xf1\xe7\xed\x4d\x4c\xfc\x59\x85\xeb\x37\x93\x6f\xa6\xdc\xae\x37\xf9\x6e\x36\x5f\x56\xb9\xa2\x19\x85\x46\x8b\xec\x0b\x3a\xef\x79\x08\x96\xc2\x83\xf0\x55\x51\xc6\x21\xb4\xa4\xc6\x72\x4f\xc2\xda\xc5\x98\xd6\x23\x4a\x4f\xaa\xc4\x21\xe5\x26\x2e\xe8\xf1\x4f\xde\xb1\x3f\x8d\x3f\x03\x63\x18\x89\x7b\x47\x9d\xa2\x5c\x7e\x4d\x1f\x0e\x82\xe7\x01\xd5\xa0\x71\xef\x6f\x8b\x55\x59\xed\xa7\x3a\xe0\x5b\x67\x40\x35\x50\x7d\x9b\x72\x1a\x47\xf2\x93\x8b\x4f\xc1\x4b\x29\x2e\x34\xa8\x45\xf6\x08\x00\x00\xff\xff\xef\x20\x3b\x16\x83\x01\x00\x00")

func _0001_appDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appDownSql,
		"0001_app.down.sql",
	)
}

func _0001_appDownSql() (*asset, error) {
	bytes, err := _0001_appDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.down.sql", size: 387, mode: os.FileMode(0644), modTime: time.Unix(1575624282, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbc, 0x9c, 0xd2, 0xe1, 0x1d, 0x8, 0x34, 0x6a, 0xc8, 0x37, 0x13, 0xb3, 0x9f, 0x26, 0x23, 0x33, 0xd4, 0x25, 0x8, 0xed, 0x53, 0xe6, 0xd, 0x46, 0xc9, 0xf4, 0x24, 0xf8, 0x1, 0x1f, 0xf5, 0xc8}}
	return a, nil
}

var __0001_appUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc1\x92\xaa\x38\x14\xdd\xf3\x15\x59\xda\x55\x6c\x66\xfd\x56\xa8\xd1\xa6\xc6\x07\x33\x88\xd3\xfd\x56\xa9\x34\x44\x4c\x09\x24\x2f\x09\x6d\xfb\xf7\x53\x21\x09\xa0\x82\xb6\x33\x3b\x92\x7b\x73\x3c\xe7\xdc\xeb\x4d\x16\x09\x0c\x52\x08\xd2\x60\xbe\x81\x20\x5c\x81\x28\x4e\x01\x7c\x0f\xb7\xe9\x16\x48\xa2\x14\xad\x0b\x09\x66\x9e\x3a\x73\x02\xfe\x09\x92\xc5\x6b\x90\x80\xbf\x92\xf0\x67\x90\xfc\x02\x7f\xc2\x5f\xbe\xf7\x89\xcb\x86\x80\xf9\x26\x9e\x7b\x2f\xe0\x2d\x4c\x5f\xe3\x5d\x0a\x92\xf8\x2d\x5c\xfe\xf0\xbc\x3b\xe0\x38\xcb\x58\x53\x2b\x0d\x8e\xf3\x5c\x10\x29\xc7\xf1\x4f\xb8\x2c\x89\x02\xf3\x38\xde\xc0\x20\xf2\xbd\xec\x80\x07\xab\x96\x57\x0a\xdf\x53\xdf\x93\x8a\x09\x5c\xb8\x15\x6f\x3e\x8e\xe4\xdc\xf2\xf2\x3d\x8e\xd5\xc1\xee\xd7\xb8\x72\x29\x19\x2b\x99\x70\xdf\x82\x60\x45\x72\x84\x15\x58\x06\x29\x4c\xc3\x9f\xb0\x25\x1b\xed\x36\x1b\xdf\x6b\x78\x3e\x19\x9d\x56\xbd\x8b\xc2\xbf\x77\x10\x84\xd1\x12\xbe\x83\xa6\xa6\xbf\x1b\x82\x8c\x1a\xe4\x14\xc7\xd1\xc0\x07\x13\x7b\x01\x6f\xaf\x30\x81\xdd\xf2\xc7\x3d\x38\x6d\xc6\x38\x98\x8e\x74\x50\xed\xa2\x03\x32\x08\xbd\x62\x64\x4f\x5d\x01\x74\xf1\x1e\xa6\xdf\xba\x5f\xdb\x0f\xc1\x4e\x92\x08\x5d\x5b\x9a\xb7\x0e\x5f\xd6\xb4\x2b\xc2\xc0\x63\x45\x2b\x22\x15\xae\x38\xd8\x6d\xd7\xe1\x3a\x82\x4b\x30\x0f\xd7\x61\x94\xfa\x5e\x8e\x39\x77\x25\x07\x4b\xb8\x0a\x76\x9b\x14\xec\x71\x29\x89\xef\x1d\xa8\xae\xfb\x39\xac\x73\xf2\x05\x76\xd1\xd6\x9c\x0c\xa3\xf4\xb9\x6e\x74\x8c\x91\xc5\x03\x33\xcf\x6e\x21\xa7\xa0\xa7\xea\x72\x4c\xeb\xac\xe2\x04\x86\xeb\x48\x2b\x9b\xf5\x67\x5e\x40\x02\x57\x30\x81\xd1\x02\xf6\xe8\x33\xbd\x1f\x6b\x0d\x1b\x98\x42\xb0\x08\xb6\x8b\x60\x09\xbd\x07\x6e\x6a\xf9\xda\xca\xde\xb5\x81\x99\xcf\xc9\xe4\x44\x54\x54\x4a\xca\x6a\x0d\xa8\x81\xd1\x58\x2d\xfa\xb4\xeb\xc8\x50\x6c\x77\xfc\x42\x6b\xcb\x76\x66\xb6\xc7\xa5\xde\x23\xa8\x04\xae\xe5\xde\xb4\x4e\x4d\xd4\x89\x89\xa3\x2e\x40\x57\x58\xd3\x12\xc3\x5a\x60\x79\xe8\x06\x47\xbf\x7d\x3d\x52\xfa\xc8\x47\x79\x44\x13\x87\xd4\x97\x9d\x17\x92\xd4\x39\x11\x23\x19\x82\x64\x84\x72\x65\xd3\x4a\x56\xd8\xaf\x8b\xf1\x38\xee\x56\xaf\xc6\x77\x14\x2e\x7b\xa4\x64\xd9\x51\x0e\xd3\x4c\xca\x8d\x87\xbe\xb7\x88\xa3\x6d\x9a\x04\xda\x08\x3b\x07\x9c\x6d\x88\x13\xe1\xe6\x41\xfb\x6d\xe1\xdc\xf0\x98\x69\x4c\xdf\x26\xf8\xfd\x6f\xbd\x3c\xea\x41\xc3\xee\x7f\x16\xa5\x6e\xaa\x0f\x22\x6e\xd3\x07\x7f\xfd\x69\x48\x82\xf3\x76\x06\x74\x03\x60\x15\x6c\xb6\xa3\x66\xb4\x5c\x47\xd5\x5f\x9b\x3b\x79\xd8\x30\x7d\x84\x61\xb2\x1e\x7a\xe7\x66\x2a\x52\x0c\x3d\xe7\xe3\xfd\x2e\x9e\xb2\x53\x9e\xeb\x0c\xb4\x83\xf3\x4e\xff\x59\xee\xf7\x3b\xd0\x25\x7d\xab\x07\x2b\xcc\x39\xad\x0b\xb4\x67\xc2\xdd\x28\x9d\xe2\x51\x27\x5d\x1b\xf6\x74\x9e\xe9\xc8\x0a\xd3\x52\x12\xf1\x69\x66\x05\x00\x00\xd0\x7c\xfc\x05\xa1\x63\xed\x94\xbb\xb5\x51\x87\xa6\x4d\xd6\x51\x8e\xa5\x3c\x31\xd1\x41\x9b\xdd\x7d\x49\x88\xba\x39\xf1\xdc\x2c\xee\x05\x20\x41\x7e\x37\x44\x2a\x54\x60\xee\xc4\x14\x98\xa3\xbd\x60\xd5\xc5\x9d\x06\xd7\xf0\x9a\x9f\xce\x53\xec\x51\xd6\xe8\x2d\xac\x03\xed\x03\xe2\xfa\x86\x9b\xd6\x61\x9e\x0e\x13\xcc\x91\x05\x43\x34\xff\xd2\x2d\x33\x29\xd0\xe6\x7d\xbb\xc0\x48\x31\x4e\x33\xe7\x4c\xbb\x98\xae\xb4\x05\x97\x97\x05\x2b\xb1\x54\x8e\x45\xe7\x91\x9b\x24\x7f\x98\x9c\x9c\xca\x8c\x7d\x12\x71\xbe\x79\x6b\xd8\x51\xd3\x36\x12\x29\x98\xa2\xfa\x19\x34\x9e\xf5\x9f\x7b\xa0\xe5\xed\x7c\x12\xb8\x2e\x88\x13\xec\x6a\x34\x29\xb9\x64\x27\xd2\xcb\x33\x6d\x63\x35\x9a\x84\x03\x2d\x0e\xc3\x0c\xc5\x5c\xfc\x96\xee\xbf\x01\x00\x00\xff\xff\x5a\xea\xe5\xa6\x10\x0c\x00\x00")

func _0001_appUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_appUpSql,
		"0001_app.up.sql",
	)
}

func _0001_appUpSql() (*asset, error) {
	bytes, err := _0001_appUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_app.up.sql", size: 3088, mode: os.FileMode(0644), modTime: time.Unix(1576607640, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0xb8, 0x68, 0x17, 0x49, 0x51, 0xc0, 0xe8, 0xbc, 0x36, 0xa4, 0x29, 0xc9, 0x93, 0x6c, 0x3e, 0xdf, 0x3d, 0x23, 0x22, 0xab, 0x18, 0x49, 0xbd, 0x6, 0xf, 0xc5, 0xec, 0xf8, 0xcf, 0x1b, 0x6a}}
	return a, nil
}

var __0002_tokensDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x28\xc9\xcf\x4e\xcd\x2b\xb6\xe6\x02\x04\x00\x00\xff\xff\xf0\xdb\x32\xa7\x13\x00\x00\x00")

func _0002_tokensDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_tokensDownSql,
		"0002_tokens.down.sql",
	)
}

func _0002_tokensDownSql() (*asset, error) {
	bytes, err := _0002_tokensDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_tokens.down.sql", size: 19, mode: os.FileMode(0644), modTime: time.Unix(1576607640, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd1, 0x31, 0x2, 0xcc, 0x2f, 0x38, 0x90, 0xf7, 0x58, 0x37, 0x47, 0xf4, 0x18, 0xf7, 0x72, 0x74, 0x67, 0x14, 0x7e, 0xf3, 0xb1, 0xd6, 0x5f, 0xb0, 0xd5, 0xe7, 0x91, 0xf4, 0x26, 0x77, 0x8e, 0x68}}
	return a, nil
}

var __0002_tokensUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xcd\x6a\x85\x30\x10\x46\xf7\x79\x8a\x6f\x79\x05\xdf\xa0\xab\xa8\xa9\x0e\xb5\xb1\xc4\xb1\xea\xaa\x58\x93\x85\xf8\x13\x30\x42\xe9\xdb\x17\x4b\x4b\x2b\xdc\xed\x37\x67\x0e\x27\x35\x4a\xb2\x02\xcb\xa4\x54\xa0\x47\xe8\x8a\xa1\x3a\xaa\xb9\xc6\xe1\x67\xb7\x05\xdc\x04\x30\x58\xbb\xbb\x10\xf0\x2a\x4d\x5a\x48\xf3\x4d\xe9\xa6\x2c\x63\x01\x6c\xee\xf8\xf0\xfb\xfc\x36\x59\x34\xba\xa6\x5c\xab\x0c\x09\xe5\xa4\xf9\x8a\x0d\xab\x03\xab\xee\xba\x86\xcf\xf5\xdd\x2f\x77\xbd\xd6\x8d\xd3\x3a\x2c\xe1\xcf\x4a\x9a\xcf\xc3\xe8\x17\xbf\xff\xbe\x9c\xc3\x8b\xa1\x67\x69\x7a\x3c\xa9\x1e\xb7\x9f\xd4\xf8\x5f\x57\x24\x22\xb4\xc4\x45\xd5\x30\x4c\xd5\x52\xf6\x20\xc4\x57\x00\x00\x00\xff\xff\x73\xf3\x87\xe5\xf8\x00\x00\x00")

func _0002_tokensUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_tokensUpSql,
		"0002_tokens.up.sql",
	)
}

func _0002_tokensUpSql() (*asset, error) {
	bytes, err := _0002_tokensUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_tokens.up.sql", size: 248, mode: os.FileMode(0644), modTime: time.Unix(1576607640, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0xd6, 0xde, 0xd3, 0x7b, 0xee, 0x92, 0x11, 0x38, 0xa4, 0xeb, 0x84, 0xca, 0xcb, 0x37, 0x75, 0x5, 0x77, 0x7f, 0x14, 0x39, 0xee, 0xa1, 0x8b, 0xd4, 0x5c, 0x6e, 0x55, 0x6, 0x50, 0x16, 0xd4}}
	return a, nil
}

var __0003_settingsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xca\xb1\x0a\xc2\x40\x0c\x06\xe0\x3d\x4f\xf1\x8f\x0a\xbe\x41\xa7\x5c\x1b\x69\xb0\x9a\x92\x46\x6b\x47\x87\x43\x04\x11\xe1\x4e\xc1\xb7\x77\x11\xd7\x8f\xaf\x73\x1b\x11\x9c\x06\x41\xc9\xb5\xde\x1e\xd7\xd2\x50\xeb\xc2\x21\x3f\xd6\x2d\x0e\x16\x90\xb3\x4e\x31\xfd\x13\x56\x04\xd4\xcf\x33\xe3\xc4\xde\xf6\xec\x18\x5d\xf7\xec\x0b\x76\xb2\x6c\x08\x78\x5f\xee\xaf\x8c\x34\x58\xa2\x35\x66\x8d\xde\x8e\x01\xb7\x59\xbb\x86\xe8\x1b\x00\x00\xff\xff\x49\x2e\x16\x6c\x76\x00\x00\x00")

func _0003_settingsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0003_settingsDownSql,
		"0003_settings.down.sql",
	)
}

func _0003_settingsDownSql() (*asset, error) {
	bytes, err := _0003_settingsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0003_settings.down.sql", size: 118, mode: os.FileMode(0644), modTime: time.Unix(1576866163, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe5, 0xa6, 0xf5, 0xc0, 0x60, 0x64, 0x77, 0xe2, 0xe7, 0x3c, 0x9b, 0xb1, 0x52, 0xa9, 0x95, 0x16, 0xf8, 0x60, 0x2f, 0xa5, 0xeb, 0x46, 0xb9, 0xb9, 0x8f, 0x4c, 0xf4, 0xfd, 0xbb, 0xe7, 0xe5, 0xe5}}
	return a, nil
}

var __0003_settingsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\xbd\x6e\xdb\x30\x10\xc7\x77\x3f\x05\xb7\xb4\x40\x87\x66\x28\x50\x20\x93\x1c\xab\x89\x50\x57\x0a\x54\xb9\x41\x26\x82\x26\xcf\xd6\xc1\x14\x8f\xe0\x51\x36\xfc\xf6\x85\x1c\x59\x56\x53\xd9\xa3\x78\xbf\xfb\xbe\xbf\x16\x65\xf1\x22\xaa\x64\xbe\x4c\x05\x43\x8c\xe8\xb6\xfc\x30\x7b\x2c\xd3\xa4\x4a\x3f\x3c\x8b\x4f\x33\x21\x94\x31\x01\x98\xc5\x9f\xa4\x7c\x7c\x4e\x4a\x91\x17\x95\xc8\x57\xcb\xe5\x97\x99\x10\xba\x56\xc4\xb2\x21\x03\x62\x5e\x14\xcb\x34\xc9\xc5\x22\xfd\x91\xac\x96\x95\xd8\x28\xcb\x70\x62\xda\x10\xc0\xe9\xe3\x10\xe0\x4c\xdc\xb5\x6c\xee\x2e\x44\x94\x0e\xe2\x81\xc2\x6e\x3a\x53\xcb\x91\x1a\xb9\x26\x8a\x8e\x0c\xb0\x98\x2f\x8b\xf9\x94\x41\x82\x53\x6b\x0b\x66\x00\x8c\xf2\x9e\xe5\xad\x2e\x00\xfd\xfd\xb7\xef\xf7\x37\x99\x8d\x05\x88\x67\x4b\xf7\x50\xa3\x01\x59\x53\x03\x32\x12\xd9\x88\xfe\xfa\x04\xd0\x71\x54\xd6\xaa\x88\xe4\x24\x9a\xc9\xf8\x3b\x38\xca\xf6\xba\x4d\xab\x60\xe4\x29\x8e\xd3\x30\x06\xc7\x76\xaf\x30\x80\x91\xe4\xc4\x2a\xff\x9d\x3d\xe5\xe9\x42\xcc\xb3\xa7\x2c\xaf\x3e\x42\xe8\xb6\x63\x7f\xab\x38\xca\xd6\x1b\x15\xc1\x4c\xb9\x5a\x15\x81\xa3\x34\x10\x70\x0f\x5d\x84\x58\x5f\xb0\x2c\xaf\x86\x8e\xbf\x9e\x68\xda\x4a\x0b\x7b\xb0\xe3\x14\x8d\x83\x86\x1c\xea\xf1\x9b\x53\x0d\x4c\xf6\xdb\xdf\xc1\xfb\x8e\xff\xb5\x90\x01\xa9\xc9\x6d\x70\x3b\xec\xd7\x51\xc4\x0d\xea\xd3\x74\x47\xdb\xbf\xb6\x0c\x5f\x53\xa4\xf7\x1e\xfe\x0b\xef\xd1\x39\x30\xb2\x51\x68\x19\xc2\x1e\xc2\xe5\xcc\x7c\x80\x0d\x84\x6e\xbc\xe3\xb2\x7b\xcb\x1e\xe1\x20\x7d\xc0\xbd\xd2\xc7\x1b\x99\xdb\xb5\x45\x2d\x77\x70\x9c\xec\x3a\x40\x03\xcd\x1a\x82\xe4\xa3\xd3\xe8\xb6\x52\xd7\x84\xfa\x86\xb0\xb8\xa6\xc3\x7b\x35\x3d\x72\x7a\xc4\xad\xeb\x9c\x7d\x1d\x14\x4f\x8f\x97\x23\xea\x1d\x04\x96\x5e\xe9\x1d\xcb\xfe\x3a\x47\x8a\x19\x80\x00\xba\xd3\xe5\xf9\xfb\x02\xf4\x15\x92\x93\x0d\xad\xd1\xc2\x20\xdd\xeb\xc5\x1e\x5d\xac\x21\xa2\x1e\x0b\x60\xf8\x13\xa0\xb9\x13\x2f\x65\xf6\x2b\x29\xdf\xc4\xcf\xf4\xad\x73\x68\x19\x42\xd7\xdc\x25\xeb\xa1\xab\x32\xca\x40\x14\x6f\x4a\xb5\xe7\x18\xba\xa3\x96\x5e\x31\xdf\xba\x87\x9e\xde\x23\xe3\xda\x76\x62\xde\x81\x1b\xe2\xce\x3e\x8b\xd7\xac\x7a\x2e\x56\x95\x28\x8b\xd7\x6c\xf1\x30\xfb\x1b\x00\x00\xff\xff\xa5\xef\xfe\xe7\x3d\x05\x00\x00")

func _0003_settingsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0003_settingsUpSql,
		"0003_settings.up.sql",
	)
}

func _0003_settingsUpSql() (*asset, error) {
	bytes, err := _0003_settingsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0003_settings.up.sql", size: 1341, mode: os.FileMode(0644), modTime: time.Unix(1576866163, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcb, 0x11, 0x7, 0x32, 0x73, 0xa0, 0x71, 0x99, 0x31, 0x49, 0xd0, 0x8, 0x34, 0x54, 0xc7, 0x8c, 0x8, 0x2e, 0x27, 0x91, 0xc7, 0x9d, 0x33, 0x32, 0xe, 0xd8, 0x4e, 0x3f, 0x3f, 0x63, 0x1e, 0x44}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x0d\xc4\x20\x0c\x05\xd0\x9e\x29\xfe\x02\xd8\xfd\x6d\xe3\x4b\xac\x2f\x44\x82\x09\x78\x7f\xa5\x49\xfd\xa6\x1d\xdd\xe8\xd8\xcf\x55\x8a\x2a\xe3\x47\x1f\xbe\x2c\x1d\x8c\xfa\x6f\xe3\xb4\x34\xd4\xd9\x89\xbb\x71\x59\xb6\x18\x1b\x35\x20\xa2\x9f\x0a\x03\xa2\xe5\x0d\x00\x00\xff\xff\x60\xcd\x06\xbe\x4a\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 74, mode: os.FileMode(0644), modTime: time.Unix(1575624282, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xde, 0x7c, 0x28, 0xcd, 0x47, 0xf2, 0xfa, 0x7c, 0x51, 0x2d, 0xd8, 0x38, 0xb, 0xb0, 0x34, 0x9d, 0x4c, 0x62, 0xa, 0x9e, 0x28, 0xc3, 0x31, 0x23, 0xd9, 0xbb, 0x89, 0x9f, 0xa0, 0x89, 0x1f, 0xe8}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"0001_app.down.sql": _0001_appDownSql,

	"0001_app.up.sql": _0001_appUpSql,

	"0002_tokens.down.sql": _0002_tokensDownSql,

	"0002_tokens.up.sql": _0002_tokensUpSql,

	"0003_settings.down.sql": _0003_settingsDownSql,

	"0003_settings.up.sql": _0003_settingsUpSql,

	"doc.go": docGo,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"0001_app.down.sql":      &bintree{_0001_appDownSql, map[string]*bintree{}},
	"0001_app.up.sql":        &bintree{_0001_appUpSql, map[string]*bintree{}},
	"0002_tokens.down.sql":   &bintree{_0002_tokensDownSql, map[string]*bintree{}},
	"0002_tokens.up.sql":     &bintree{_0002_tokensUpSql, map[string]*bintree{}},
	"0003_settings.down.sql": &bintree{_0003_settingsDownSql, map[string]*bintree{}},
	"0003_settings.up.sql":   &bintree{_0003_settingsUpSql, map[string]*bintree{}},
	"doc.go":                 &bintree{docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
