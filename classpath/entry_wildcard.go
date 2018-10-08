package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompisiteEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		} else if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		//|| strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
			compositeEntry = append(compositeEntry, newZipEntry(path))
			return nil
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
