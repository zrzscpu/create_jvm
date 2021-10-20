package classfile

import "encoding/binary"

/*
首先所有的字节码都存储到这个byte数组中,读取多少就剔除多少
*/
type ClassReader struct {
	data []byte
}

/**
读取数据的函数
*/

/*
对应u1类型
*/
func (this *ClassReader) readUint8() uint8 {

	//读取第一个数据，并将第一个数据移除出data数组
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

/*
对应u2类型
*/
func (this *ClassReader) readUint16() uint16 {

	//BigEndian bigEndian 大端字节序的实现。大端模式中字数据的高字节存储在低地址中，而字数据的低字节则存放在高地址中
	// LittleEndian littleEndian 小端字节序的实现
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]
	return val

}

/*
对应u4类型
*/
func (this *ClassReader) readUint32() uint32 {

	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]
	return val
}

/*
对应u8类型
*/
func (this *ClassReader) readUint64() uint64 {

	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]
	return val
}

//读取表格
func (this *ClassReader) readUint16s() []uint16 {
	//表的长度往往都是u2
	length := this.readUint16()
	slice := make([]uint16, length)
	for i, _ := range slice {
		slice[i] = this.readUint16()
	}

	return slice
}

/*
主要对应字符串类型
*/
func (this *ClassReader) readBytes(n uint32) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}
