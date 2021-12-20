package heap

import "create_jvm/main02/classfile"

//Class的方法属性
type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	code      []byte
}

func (this *Method) copyAttributes(memberInfo *classfile.MemberInfo) {
	if codeAttribute := memberInfo.CodeAttribute(); codeAttribute != nil {
		this.maxStack = codeAttribute.MaxStack()
		this.maxLocals = codeAttribute.MaxLocals()
		this.code = codeAttribute.Code()
	}
}

//创建类的方法集合
func newMethods(class *Class, classfileMethod []*classfile.MemberInfo) []*Method {
	//类文件的 方法表长度
	methods := make([]*Method, len(classfileMethod))
	for index, val := range classfileMethod {
		methods[index] = &Method{}
		methods[index].class = class
		methods[index].copyMemberInfo(val)
		//方法表元素的属性表主要就是code,maxloacals,maxstack
		methods[index].copyAttributes(val)
	}
	return methods
}

func (this *Method) MaxStack() uint16 {
	return this.maxStack
}
func (this *Method) MaxLocals() uint16 {
	return this.maxLocals
}
func (this *Method) Code() []byte {
	return this.code
}
