package classfile

/*源定义
CONSTANT_Utf8_info {
	u1 tag;
	u2 length;
	u1 bytes[length];
}
*/
/**
存储字面量
*/
type ConstantUtf8Info struct {
	str string
}

func (this *ConstantUtf8Info) readInfo(reader *ClassReader) {

	//明明是16位的长度，转成32位进行计算，字节留的坑
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	//将byte数组转为utf-8的字符串(粗略的转换，java内部使用的是MUTF8)
	//MUTF-8编码方式和UTF-8大致相同，但并不兼容。
	//差别有两点：一是null字符（代码点U+0000）会被编码成2字节： 0xC0、0x80；
	//二是补充字符（Supplementary Characters， 代码点大于 U+FFFF的Unicode字符）是按UTF-16拆分为代理对（Surrogate Pair） 分别编码的
	this.str = decodeMUTF8(bytes)
}

//这里先直接按utf-8进行解析
func decodeMUTF8(bytes []byte) string {

	return string(bytes)
}
