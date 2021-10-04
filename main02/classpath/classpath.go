package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {

	cp := &ClassPath{}
	cp.parseBootAndExtClassthPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

//将用户的类进行加载
func (this *ClassPath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "." //有点疑惑
	}
	this.userClassPath = newDirEntry(cpOption)
}

//将jre的配置信息转为路径信息存储到	bootClassPath Entry 和 extClassPath Entry
func (this *ClassPath) parseBootAndExtClassthPath(jreOption string) {

	jreDir := getJreDir(jreOption)

	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	//加载启动类加载器下的所有class文件
	this.bootClassPath = newWildcardEntry(jreLibPath)

	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	//记载扩展类加载器下的所有class文件
	this.extClassPath = newWildcardEntry(jreExtPath)

}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		//目录不存在的报错
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//将jreoption转为地址
func getJreDir(jreOption string) string {

	//exists判断目录是否存在
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	//如果用户没有输入jreOption选项，或者该选项的路径不存在，在当前工作目录下寻找
	if exists("./jre") {
		return "./jre"
	}
	//从javahome环境变量下查找
	if env := os.Getenv("JAVA_HOME"); env != "" {
		//jdk1.8的结构
		return filepath.Join(env, "jre")

		//jdk16 的jre
	}
	//如果都找不到报错
	panic("找不到jre文件夹")
}

func (this *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	//双亲委派
	if data, entry, err := this.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := this.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return this.userClassPath.readClass(className)
}

func (this *ClassPath) String() string {
	return this.userClassPath.String()
}
