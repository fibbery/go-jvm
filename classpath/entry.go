package classpath

import (
	"os"
	"strings"
)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) Entry {
	if strings.Contains(path, string(os.PathListSeparator)) {
		return newCompositeEntry(path)
	}

}

