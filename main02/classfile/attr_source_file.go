package classfile

/*源定义
SourceFile_attribute {
	u2 attribute_name_index;	属性名的索引
	u4 attribute_length;
	u2 sourcefile_index;		源文件名
}
*/
type SourceFileAttribute struct {
	cp            ConstantPool
	sourFileIndex uint16
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourFileIndex = reader.readUint16()
}

func (this *SourceFileAttribute) FileName() string {

	return this.cp.getUtf8(this.sourFileIndex)
}
