package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	classpath := &Classpath{}
	classpath.parseBootAndExtClasspath(jreOption)
	classpath.parseUserClasspath(cpOption)
	return classpath
}

func (c *Classpath) readClass(className string) ([]byte, Entry, error) {
	classFilename := strings.ReplaceAll(className, ".", string(os.PathSeparator)) + ".class"
	if bytes, entry, e := c.bootClasspath.readClass(classFilename); e == nil {
		return bytes, entry, e
	}
	if bytes, entry, e := c.extClasspath.readClass(classFilename); e == nil {
		return bytes, entry, e
	}
	return c.userClasspath.readClass(classFilename)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	c.bootClasspath = CompositeEntry(newWildcardEntry(filepath.Join(jreDir, "lib", "*")))
	c.extClasspath = CompositeEntry(newWildcardEntry(filepath.Join(jreDir, "lib", "ext", "*")))
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = NewEntry(cpOption)
}

// getJreDir return a string that descirebe jre file path
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// estimate current path whether contains directory jre
	if exists("./jre") {
		return "./jre"
	}
	// get environment variable
	if jb := os.Getenv("JAVA_HOME"); jb != "" {
		return filepath.Join(jb, "jre")
	}
	panic("can't find jre directory")
}

// exists return a boolean indicate whether the file is exist
func exists(path string) bool {
	if _, e := os.Stat(path); e != nil {
		if os.IsNotExist(e) {
			return false
		}
	}
	return true
}
