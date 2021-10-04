package classfile

/*
默认情况下的属性表就是解析类型失败给生成一张这样的表
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (this *UnparsedAttribute) readInfo(reader *ClassReader) {

	this.info = reader.readBytes(this.length)
}
