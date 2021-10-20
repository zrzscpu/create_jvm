package classfile

//LocalVariableTable_attribute {
//	u2 attribute_name_index;
//	u4 attribute_length;
//	u2 local_variable_table_length;
//	{
//		u2 start_pc;
//		u2 length;
//		u2 name_index;
//		u2 descriptor_index;
//		u2 index;
//	}
//	local_variable_table[local_variable_table_length];
//}

type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

//加载LocalVariableTable
func (this *LocalVariableTableAttribute) readInfo(reader *ClassReader) {

	localVariableTableLength := reader.readUint16()
	this.LocalVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i, _ := range this.LocalVariableTable {
		this.LocalVariableTable[i].StartPc = reader.readUint16()
		this.LocalVariableTable[i].Length = reader.readUint16()
		this.LocalVariableTable[i].NameIndex = reader.readUint16()
		this.LocalVariableTable[i].DescriptorIndex = reader.readUint16()
		this.LocalVariableTable[i].Index = reader.readUint16()
	}

}
