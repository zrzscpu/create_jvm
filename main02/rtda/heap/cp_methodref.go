package heap

import "create_jvm/main02/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool,
	info *classfile.ConstantMethodrefInfo) *MethodRef {

	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (this *MethodRef) Descriptor() string {
	return this.descriptor
}

func (this *MethodRef) Name() string {
	return this.name
}
