package classpath

import (
	"errors"
	"os"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	var entries []Entry
	for _, name := range strings.Split(path, string(os.PathListSeparator)) {
		entry := NewEntry(name)
		entries = append(entries, entry)
	}
	return entries
}

func (t CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range t {
		bytes, i, e := entry.readClass(className)
		if e == nil {
			return bytes, i, e
		}
	}
	return nil, nil, errors.New("class not found :" + className)
}

func (t CompositeEntry) String() string {
	var names []string
	for _, entry := range t {
		names = append(names, entry.String())
	}
	return strings.Join(names,string(os.PathListSeparator))
}
