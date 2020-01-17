package classpath

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	info, e := os.Stat(path)
	if e != nil {
		panic(e)
	}
	if !info.IsDir() {
		panic(errors.New(path + "is not a directory"))
	}

	absPath, e := filepath.Abs(path)
	if e != nil {
		panic(e)
	}
	return &DirEntry{absPath}
}

func (t *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(t.absDir, className)
	bytes, e := ioutil.ReadFile(filename)
	return bytes, t, e
}

func (t *DirEntry) String() string {
	return t.absDir
}
