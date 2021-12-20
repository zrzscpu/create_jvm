package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//保存一个jar包的地址
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	//转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

//将jar包中所有的class文件读取
func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	//从zip文件或jar文件中提取字节码文件
	reader, err := zip.OpenReader(this.absPath)
	if err != nil {
		panic(nil)
	}
	//读取结束后关闭
	defer reader.Close()

	//对内部的文件进行遍历
	for _, f := range reader.File {
		//找到需要的类了进行读取
		if f.Name == className {
			reader, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer reader.Close()

			data, err := ioutil.ReadAll(reader)

			if err != nil {
				return nil, nil, err
			}
			return data, this, nil
		}
	}

	return nil, nil, errors.New("class not found:" + className)

}

func (this *ZipEntry) String() string {

	return this.absPath
}
