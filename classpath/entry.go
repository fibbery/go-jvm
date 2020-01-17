package classpath

import (
	"os"
	"strings"
)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) Entry {
	if strings.Contains(path, string(os.PathListSeparator)) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(strings.ToLower(path), "zip") || strings.HasSuffix(strings.ToLower(path), "jar") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
