package main

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

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

func (this ZipEntry) readClass(className string) ([]byte, Entry, error) {

	reader, err := zip.OpenReader(this.absPath)
	if err != nil {
		panic(nil)
	}
	//读取结束后关闭
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == className {
			reader, err := f.Open()
			if nil != nil {
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

func (this ZipEntry) String() string {

	return this.absPath
}
