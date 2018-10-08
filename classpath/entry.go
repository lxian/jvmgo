package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// return entry composite
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		// return zip entry
		return newZipEntry(path)
	} else if strings.HasSuffix(path, "*") {
		// return wildcard entry
		return newWildcardEntry(path)
	} else {
		// return dir entry
		return newDirEntry(path)
	}
}
