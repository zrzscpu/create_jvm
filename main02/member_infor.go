package main

type MemberInfo struct {
	cp              ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

//读取所有的方法和字段
//字段或方法的描述
func reaMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	//有多少个字段或方法
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i, _ := range members {
		members[i] = reaMember(reader, cp)
	}
}

//读取一个方法或字段
func reaMember(reader *ClassReader, cp ConstantPool) *MemberInfo {

	return &MemberInfo{
		cp:              cp,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),

		attributes: readeAttributes(reader, cp),
	}
}

//返回当前表的所代表的字段或方法名
func (this *MemberInfo) Name() string {

	return this.cp.getUtf8(this.nameIndex)
}

//返回当前表的所代表的字段或方法的描述信息
func (this *MemberInfo) Descriptor() string {

	return this.cp.getUtf8(this.descriptorIndex)
}
