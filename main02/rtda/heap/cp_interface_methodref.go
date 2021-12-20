package heap

import "create_jvm/main02/classfile"

type InterfaceMethodRef struct {
	MethodRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref

}
