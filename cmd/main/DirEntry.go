package main

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	//转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (this DirEntry) readClass(className string) ([]byte, Entry, error) {

	//将目录名和文件名拼接成完整路径
	fileName := filepath.Join(this.absDir, className)
	data, error := ioutil.ReadFile(fileName)
	return data, this, error
}

func (this DirEntry) String() string {

	return this.absDir
}
