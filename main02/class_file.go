package main

import (
	"fmt"
	"go/constant"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantpool Constantpool
	//访问标识
	accessflag uint16
	//类名索引
	thisclass  uint16
	superclass uint16
	interfaces []uint16

	fileds     []*MemberInfor
	methods    []*MemberInfor
	attributes []*MemberInfor
}

func parse(classData []byte) (cf *ClassFile, err error) {

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
	classfile := &ClassFile{}
	classfile.read(data)

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

	case 46, 47, 48, 49, 50, 51, 52:
		if this.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//getter
func (this *ClassFile) MajorVersion() uint16 {

	return this.majorVersion
}
func (this *ClassFile) MinorVersion() uint16 {

	return this.minorVersion
}
func (this *ClassFile) Constantpool() Constantpool {

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

func (this *ClassFile) Fileds() []*MemberInfor {

	return this.fileds
}
func (this *ClassFile) Methods() []*MemberInfor {

	return this.methods
}
func (this *ClassFile) Attributes() []*MemberInfor {

	return this.attributes
}

//从常量池查找类名
func (this *ClassFile) ClassName() string {

	return this.constantpool.getClassName(this.thisclass)
}

//从常量池查找父类类名
func (this *ClassFile) SuperClassName() string {

	if this.superclass > 0 {
		return this.constantpool.getClassName(this.thisclass)
	}

	//没有显示的父类
	return ""
}

//从常量池查找类名
func (this *ClassFile) InterfaceName() []string {
	length := len(this.interfaces)
	if length > 0 {
		strings := make([]string, length)
		for i, interfaceindex := range this.interfaces {
			strings[i] = this.Constantpool.getInterfaceName(interfaceindex)
		}
		return strings
	}
	return make([]string, 0)
}
