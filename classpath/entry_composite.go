package classpath

import (
	"errors"
	"strings"
)

type CompisiteEntry []Entry

func newCompositeEntry(pathList string) CompisiteEntry {
	compisiteEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		compisiteEntry = append(compisiteEntry, newEntry(path))
	}
	return compisiteEntry
}

func (self CompisiteEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, e, err := entry.readClass(className)
		if err == nil {
			return data, e, err
		}
	}
	return nil, nil, errors.New("Class not Found")
}

func (self CompisiteEntry) String() string {
	paths := make([]string, len(self))
	for i, entry := range self {
		paths[i] = entry.String()
	}
	return strings.Join(paths, pathListSeparator)
}
