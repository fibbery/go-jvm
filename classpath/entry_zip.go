package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry parse zip or jar file
type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, e := filepath.Abs(path)
	if e != nil {
		panic(e)
	}
	return &ZipEntry{absDir}
}

func (entry *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	z, e := zip.OpenReader(entry.absDir)
	if e != nil {
		panic(e)
	}
	for _, file := range z.File {
		if file.Name != className {
			continue
		}
		rc, e := file.Open()
		if e != nil {
			return nil, nil, e
		}
		bytes, e := ioutil.ReadAll(rc)
		if e != nil {
			return nil, nil, e
		}
		_ = rc.Close()
		return bytes, entry, nil
	}
	return nil, nil, errors.New("class not found : " + className)
}

func (entry *ZipEntry) String() string {
	return entry.absDir
}
