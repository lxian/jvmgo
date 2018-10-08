package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type Classpath struct {
	bootClassPathEntry Entry
	extClassPathEntry  Entry
	userClassPathEntry Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreBase := getJreDir(jreOption)

	self.bootClassPathEntry = newWildcardEntry(filepath.Join(jreBase, "lib", "*"))
	self.extClassPathEntry = newWildcardEntry(filepath.Join(jreBase, "lib", "ext", "*"))
}

func (self *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPathEntry = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	} else if exists("./jre") {
		return "./jre"
	} else if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	} else {
		panic("No JRE path specified")
	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	if !strings.HasSuffix(className, ".class") {
		className = className + ".class"
	}

	if data, entry, err := self.bootClassPathEntry.readClass(className); err == nil {
		return data, entry, err
	} else if data, entry, err := self.extClassPathEntry.readClass(className); err == nil {
		return data, entry, err
	} else {
		return self.userClassPathEntry.readClass(className)
	}
}

func (self *Classpath) String() string {
	return strings.Join(
		[]string{self.bootClassPathEntry.String(), self.extClassPathEntry.String(), self.userClassPathEntry.String()},
		pathListSeparator)
}
