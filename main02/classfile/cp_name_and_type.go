package classfile

/*
CONSTANT_NameAndType_info {
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}
*/

/**
方法或字段的名字和描述符
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {

	this.nameIndex = reader.readUint16()
	this.descriptorIndex = reader.readUint16()
}
