package modfs

import (
	"net/http"
	"os"
	"time"
)

type ModTimeFileInfo struct {
	os.FileInfo
	time.Time
}

func (f *ModTimeFileInfo) ModTime() time.Time {
	return f.Time
}

type ModTimeFile struct {
	http.File
	time.Time
}

func (f *ModTimeFile) Stat() (os.FileInfo, error) {
	fileInfo, err := f.File.Stat()
	return &ModTimeFileInfo{FileInfo: fileInfo, Time: f.Time}, err
}

type ModTimeFS struct {
	http.FileSystem
	time.Time
}

func New(fsys http.FileSystem, modTime time.Time) http.FileSystem {
	return &ModTimeFS{FileSystem: fsys, Time: modTime}
}

func (f *ModTimeFS) Open(name string) (http.File, error) {
	file, err := f.FileSystem.Open(name)
	return &ModTimeFile{File: file, Time: f.Time}, err
}
