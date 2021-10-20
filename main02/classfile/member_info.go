package classfile

//用于描述方法 字段集合,二者只在属性表上有区别
//field_info { 	u2 access_flags ;
//				u2 name_index;
//				u2 descriptor_index;
//				u2 attributes_count;
//				attribute_info attributes[attributes_count];
//			}
type MemberInfo struct {
	cp              ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

//读取方法表或字段表的所有内容
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	//有多少个字段或方法
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

//读取方法表或字段的中的一个元素的所有信息
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {

	return &MemberInfo{
		cp:              cp,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),

		attributes: readAttributes(reader, cp),
	}
}

//返回当前表的名称
func (this *MemberInfo) Name() string {

	return this.cp.getUtf8(this.nameIndex)
}

//返回当前元素的描述信息
func (this *MemberInfo) Descriptor() string {

	return this.cp.getUtf8(this.descriptorIndex)
}

//返回当前方法的信息
func (this *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attributeInfo := range this.attributes {
		switch attributeInfo.(type) {
		case *CodeAttribute:
			return attributeInfo.(*CodeAttribute)
		}
	}
	return nil
}
