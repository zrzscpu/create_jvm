package main

import (
	"fmt"
	"strings"
)

//测试命令行模块
//func main() {
//
//	cmd := parseCmd()
//
//	if cmd.versionFlag {
//		fmt.Println(cmd.versionFlag)
//
//	} else if cmd.helpFlag || cmd.class == "" {
//		printUsage()
//	} else {
//		startjvm(cmd)
//	}
//
//	fmt.Printf(cmd.class)
//}
//func startjvm(cmd *Cmd) {
//
//	fmt.Printf("classpath: %s args: %s", cmd.class, cmd.args)
//}

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

func starJvm(cmd *Cmd) {
	cp := Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v ", cp, cmd.class, cmd.args)

	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。将类路径转为文件路径
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("could not find or load main class")
		return
	}

	fmt.Printf("classdata:%v", classData)

}

//func main() {
//	readFile, err := ioutil.ReadFile("E:/onedriver/OneDrive/桌面/新建文本文档 (2).txt")
//	if err!=nil {
//		fmt.Printf("%v", err)
//	}
//	//需要将byte切片转为string类型
//	fmt.Printf("%v", string(readFile))
//}
