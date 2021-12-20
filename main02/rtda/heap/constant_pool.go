package heap

import (
	"create_jvm/main02/classfile"
	"fmt"
)

//定义这样的接口,默认所有类型都实现了这个接口
type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfpool classfile.ConstantPool) *ConstantPool {

	//classfile中静态常量池的长度
	cpCount := len(cfpool)
	consts := make([]Constant, cpCount)

	//创建class的运行时常量池
	rtCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfpool[i]

		//静态常量池中的constantinfo 具体的类型：
		//CONSTANT_Integer
		//CONSTANT_Float
		//CONSTANT_Long:
		//CONSTANT_Double:
		//CONSTANT_Utf8:
		//ConstantUtf8Info
		//CONSTANT_String:
		//CONSTANT_Class:
		//CONSTANT_Fieldref:
		//CONSTANT_Methodref
		//CONSTANT_InterfaceMethodref
		//CONSTANT_NameAndType:
		//CONSTANT_MethodType:
		//CONSTANT_MethodHandle:
		//CONSTANT_InvokeDynamic:

		//转为运行时常量池后变为了:
		//CONSTANT_Integer
		//CONSTANT_Float
		//CONSTANT_Long:
		//CONSTANT_Double:
		//CONSTANT_Utf8:(string)
		//这三个是将原先这些引用中的name , class 等属性直接解析了
		//CONSTANT_Fieldref:
		//CONSTANT_Methodref:
		//ConstantClassInfo
		//CONSTANT_InterfaceMethodref

		switch cpInfo.(type) {

		//只保留类文件常量池表中的常量信息：被final 修饰的field
		//不被final修饰的是变成了一些方法字段接口信息:
		//描述符,名称
		//所属类

		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()

		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()

		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++

		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++

		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()

			//以下为符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)

		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)

		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)

		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)

		}
		//多出来的那几种类型不理会
	}
	return rtCp
}

func (this *ConstantPool) GetConstant(index uint) Constant {
	if c := this.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
