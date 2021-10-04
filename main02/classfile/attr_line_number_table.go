package classfile

/*
LineNumberTable_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 line_number_table_length;
	{
		u2 start_pc; u2 line_number;
	}
	line_number_table[line_number_table_length];
}
*/
/*
	linenumbertable表的数据结构
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

/**
linenumbertable表属性数据
*/
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

//加载LineNumberTable表
func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {

	LineNumberTableLength := reader.readUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, LineNumberTableLength)

	for i := range this.lineNumberTable {

		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
