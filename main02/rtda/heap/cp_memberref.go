package heap

import "create_jvm/main02/classfile"

//字段方法符号引用的共有信息
type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (this *MemberRef) copyMemberRefInfo(info *classfile.ConstantMemberrefInfo) {
	this.className = info.ClassName()
	this.name, this.descriptor = info.NameAndDescriptor()
}
