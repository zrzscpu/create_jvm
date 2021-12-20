package classpath

import (
	"os"
	"strings"
)

//const (
//	PathSeparator     = '/' // 操作系统指定的路径分隔符
//	PathListSeparator = ';' // 操作系统指定的表分隔符
//)
//拿到操作系统的分隔符
const pahtListSeparator = string(os.PathListSeparator)

//其实是每个class 文件路径的描述
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	//如果指定了多个目录
	if strings.Contains(path, pahtListSeparator) {
		return newCompositeEntry(path)
	}
	//包含通配符*
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	//精确到文件名(jar包,但不是绝对路径)
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}
	return newDirEntry(path)
}
