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

	startJVM(cmd)

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

//解析字节码测试
//func starJvm(cmd *Cmd) {
//	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
//	className := strings.Replace(cmd.class, ".", "/", -1)
//	classFile := loadClass(className, classPath)
//	fmt.Println(cmd.class)
//	printClassInfo(classFile)
//}
//func loadClass(class string, cp *classpath.ClassPath) *classfile.ClassFile {
//
//	classData, _, err := cp.ReadClass(class)
//	if err != nil {
//		panic(err)
//	}
//	cf, err := classfile.Parse(classData)
//	if err != nil {
//		panic(err)
//	}
//
//	return cf
//}
//
//func printClassInfo(cf *classfile.ClassFile) {
//
//	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
//
//	fmt.Printf("constants count: %v\n", len(cf.Constantpool()))
//	fmt.Printf("access flags: 0x%x\n", cf.Accessflag())
//	fmt.Printf("this class: %v\n", cf.ClassName())
//	fmt.Printf("super class: %v\n", cf.SuperClassName())
//	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
//	fmt.Printf("fields count: %v\n", len(cf.Fields()))
//	for _, f := range cf.Fields() {
//		fmt.Printf(" %s\n", f.Name())
//	}
//
//	fmt.Printf("methods count: %v\n", len(cf.Methods()))
//	for _, m := range cf.Methods() {
//		fmt.Printf(" %s\n", m.Name())
//	}
//}

//对运行时数据区的测试
//func starJvm(cmd *Cmd) {
//	frame := rtda.NewFrame(100, 100)
//	testLocalVars(frame.LocalVars())
//	//testOperandStack(frame.OperandStack())
//}
//
//func testOperandStack(stack *rtda.OperandStack) {
//
//	stack.PushInt(123)
//	println(stack.PopInt())
//
//	stack.PushLong(123333333333)
//	println(stack.PopInt())
//
//	stack.PushFloat(3.14)
//	println(stack.PopInt())
//
//	stack.PushDouble(3.1415151431413124)
//	println(stack.PopDouble())
//
//	stack.PushRef(nil)
//	println(stack.PopRef())
//
//}
//
//func testLocalVars(LocalVars rtda.LocalVars) {
//	LocalVars.SetInt(0, 100)
//	LocalVars.SetInt(1, -100)
//	LocalVars.SetLong(2, 2997924570)
//	LocalVars.SetLong(4, 2997924570)
//
//	LocalVars.SetFloat(6, 3.141592653)
//	LocalVars.SetDouble(7, 2.71828182845)
//	LocalVars.SetRef(9, nil)
//
//	println(LocalVars.GetInt(0))
//	println(LocalVars.GetInt(1))
//	println(LocalVars.GetLong(2))
//	println(LocalVars.GetLong(4))
//	println(LocalVars.GetFloat(6))
//	println(LocalVars.GetDouble(7))
//	println(LocalVars.GetRef(9))
//}

func loadClass(className string, cp *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func getMainMethod(file *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range file.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)

	cf := loadClass(className, cp)

	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
