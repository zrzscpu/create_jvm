package classpath

import (
	"errors"
	"strings"
)

//为了实现指定多个目录下的加载
//例如 java_home = C:\Program Files\Java\jdk-16.0.1	;C:\Program Files\Java\jdk-16.0.1
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	//按;分隔成不同的entry
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
