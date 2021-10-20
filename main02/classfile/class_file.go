package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantpool ConstantPool
	//访问标识
	accessflag uint16
	//类名索引
	thisclass  uint16
	superclass uint16
	interfaces []uint16

	//这三种表的结构相似
	//字段信息
	fileds []*MemberInfo
	//方法信息
	methods []*MemberInfo
	//属性表(attribute是interface)
	attributes []AttributeInfo
}

//将字节码数据存到相应的位置中
func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		//在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止painc过程。
		//若recover在defer的函数之外被调用，它将不会停止painc过程序列。
		//在此情况下，或当该Go程不在painc过程中时，或提供给panic的实参为nil时，recover就会返回nil
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	data := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(data)
	return
}

//将字节文件进行解析进行解析
func (this *ClassFile) read(reader *ClassReader) {
	//魔数判断
	this.readAndCheckMagic(reader)

	//版本判断
	this.readAndCheckVersion(reader)

	//常量池解析
	this.constantpool = readConstantPool(reader)

	//读取这个文件的访问标识
	this.accessflag = reader.readUint16()

	//读取这个类的类名
	this.thisclass = reader.readUint16()

	//读取这个类的父类
	this.superclass = reader.readUint16()

	//读取接口
	this.interfaces = reader.readUint16s()

	//解析域表信息
	this.fileds = readMembers(reader, this.constantpool)

	//解析方法表信息
	this.methods = readMembers(reader, this.constantpool)

	//解析属性信息
	this.attributes = readAttributes(reader, this.constantpool)
}

//魔数判断
func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//版本判断 使用java8 支持版本号为45-52,java版本是向后兼容的
func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {
	this.minorVersion = reader.readUint16()
	this.majorVersion = reader.readUint16()
	switch this.majorVersion {
	case 45:
		return

	case 46, 47, 48, 49, 50, 51, 52, 60:
		if this.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//getter,获得版本信息
func (this *ClassFile) MajorVersion() uint16 {

	return this.majorVersion
}
func (this *ClassFile) MinorVersion() uint16 {

	return this.minorVersion
}
func (this *ClassFile) Constantpool() ConstantPool {

	return this.constantpool
}
func (this *ClassFile) Accessflag() uint16 {

	return this.accessflag
}
func (this *ClassFile) Thisclass() uint16 {

	return this.thisclass
}
func (this *ClassFile) Superclass() uint16 {

	return this.superclass
}
func (this *ClassFile) Interfaces() []uint16 {

	return this.interfaces
}

func (this *ClassFile) Fields() []*MemberInfo {

	return this.fileds
}
func (this *ClassFile) Methods() []*MemberInfo {

	return this.methods
}
func (this *ClassFile) Attributes() []AttributeInfo {

	return this.attributes
}

//从常量池查找类名
func (this *ClassFile) ClassName() string {

	return this.constantpool.getClassName(this.thisclass)
}

//从常量池查找父类类名
func (this *ClassFile) SuperClassName() string {

	if this.superclass > 0 {
		return this.constantpool.getClassName(this.superclass)
	}

	//没有显示的父类
	return ""
}

func (this *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(this.interfaces))
	for i, cpIndex := range this.interfaces {
		interfaceNames[i] = this.constantpool.getClassName(cpIndex)
	}
	return interfaceNames
}
