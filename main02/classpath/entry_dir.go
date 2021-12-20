package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//描述目录的绝对路径
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//Abs函数返回path代表的绝对路径，如果path不是绝对路径，会加入当前工作目录以使之成为绝对路径。
	//因为硬链接的存在，不能保证返回的绝对路径是唯一指向该地址的绝对路径。
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (this *DirEntry) readClass(className string) ([]byte, Entry, error) {

	//将目录名和文件名拼接成完整路径
	fileName := filepath.Join(this.absDir, className)
	data, error := ioutil.ReadFile(fileName)
	return data, this, error
}

func (this *DirEntry) String() string {

	return this.absDir
}
