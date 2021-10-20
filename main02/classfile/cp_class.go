package classfile

/*源定义
CONSTANT_Class_info {
	u1 tag;
	u2 name_index;
}
*/
//nameindex 指向一个utf8的常量池索引
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {

	this.nameIndex = reader.readUint16()
}

func (this *ConstantClassInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}
