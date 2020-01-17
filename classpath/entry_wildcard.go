package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type WildcardEntry CompositeEntry

func newWildcardEntry(path string) CompositeEntry{
	baseDir := path[:len(path)-1] //remove *
	var entries []Entry
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && baseDir != path{
			return filepath.SkipDir
		}
		name := strings.ToLower(info.Name())
		if strings.HasSuffix(name, ".zip") || strings.HasSuffix(name, ".jar") {
			entries = append(entries, newZipEntry(path))
		}
		return nil
	}
	_ = filepath.Walk(baseDir, walkFunc)
	return entries
}
