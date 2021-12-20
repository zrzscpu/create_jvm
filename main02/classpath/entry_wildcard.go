package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//带*的是指当前目录下的所有文件进行搜索
func newWildcardEntry(path string) CompositeEntry {

	baseDir := path[:len(path)-1] //去掉*
	compositeEntry := []Entry{}

	//Walk函数对每一个文件/目录都会调用WalkFunc函数类型值。
	//func Walk(root string, fn WalkFunc)
	//type WalkFunc func(path string, info fs.FileInfo, err error) error

	//调用时path参数会包含Walk的root参数作为前缀；就是说，如果Walk函数的root为"dir"，
	//该目录下有文件"a"，将会使用"dir/a"调用walkFn参数。

	//walkFunc参数被调用时的info参数是path指定的地址（文件/目录）的文件信息，类型为os.FileInfo。
	//如果遍历path指定的文件或目录时出现了问题，传入的参数err会描述该问题，
	//WalkFunc类型函数可以决定如何去处理该错误（Walk函数将不会深入该目录）；如果该函数返回一个错误，
	//Walk函数的执行会中止,只有一个例外，如果Walk的walkFn返回值是SkipDr，
	//将会跳过该目录的内容而Walk函数照常执行处理下一个文件。

	//放到walk函数中作为walkfunc调用
	walkfunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//如果是一个目录,且该目录不等于baseDir 说明当前目录在baseDir之下
		if info.IsDir() && path != baseDir {
			//这种搜索只搜索当前目录之下的jar包或class文件,而不会进入更深一层的目录搜索
			return filepath.SkipDir
		}
		//在当前路径下且为jar包,进行 添加到 jarEntry中
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix("path", ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	//调用这个函数,遍历baseDir目录下的所有目录和文件
	filepath.Walk(baseDir, walkfunc)
	return compositeEntry
}
