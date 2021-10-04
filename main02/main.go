package main

import (
	"create_jvm/main02/classfile"
	"create_jvm/main02/classpath"
	"fmt"
	"strings"
)

//测试类路径解析和类加载(加载进内存)
func main() {

	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println(cmd.versionFlag)

	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	}

	starJvm(cmd)

}

//func starJvm(cmd *Cmd) {
//	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	fmt.Printf("classpath:%v class:%v args:%v ", cp, cmd.class, cmd.args)
//
//	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。将类路径转为文件路径
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	classData, _, err := cp.ReadClass(className)
//
//	if err != nil {
//		fmt.Printf("Could not find or load main class %s\n", cmd.class)
//		return
//	}
//	fmt.Printf("class data:%v\n", classData)
//
//}

func starJvm(cmd *Cmd) {
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, classPath)
	fmt.Println(cmd.class)
	printClassInfo(classFile)
}
func loadClass(class string, cp *classpath.ClassPath) *classfile.ClassFile {

	classData, _, err := cp.ReadClass(class)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {

	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())

	fmt.Printf("constants count: %v\n", len(cf.Constantpool()))
	fmt.Printf("access flags: 0x%x\n", cf.Accessflag())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}

	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
