package main

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	//将命令行输入的多个文件名分割成,存入一个entry切片
	for _, path := range strings.Split(pathList, pahtListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry

}

func (this CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	for _, entry := range this {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (this CompositeEntry) String() string {

	strs := make([]string, len(this))
	for i, entry := range this {
		strs[i] = entry.String()
	}
	//将一系列字符串连接为一个字符串，之间用sep来分隔。
	return strings.Join(strs, pahtListSeparator)
}
