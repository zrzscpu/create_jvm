package base

type BytecodeReader struct {
	//code存放字节码
	code []byte
	pc   uint
}

//重新设置，使这个对象可以重用
func (this *BytecodeReader) Reset(code []byte, pc uint) {
	this.code = code
	this.pc = pc
}

func (this *BytecodeReader) Pc() uint {
	return this.pc
}

func (this *BytecodeReader) ReadUint8() uint8 {
	i := this.code[this.pc]
	this.pc++
	return i
}

func (this *BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUint8())
}

func (this *BytecodeReader) ReadUint16() uint16 {
	high := uint16(this.ReadUint8())
	low := uint16(this.ReadUint8())
	return (high<<8 | low)
}

func (this *BytecodeReader) ReadInt16() int16 {
	return int16(this.ReadUint16())
}

func (this *BytecodeReader) ReadInt32() int32 {
	//不先转成32位的会有问题
	byte1 := int32(this.ReadUint8())
	byte2 := int32(this.ReadUint8())
	byte3 := int32(this.ReadUint8())
	byte4 := int32(this.ReadUint8())
	res := int32((byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4)
	return res
}

//default offset在字节码中的地址是4的倍数，前面有0-3字节的空白
func (this *BytecodeReader) SkipPadding() {
	for this.pc%4 != 0 {
		this.ReadUint8()
	}
}

func (this *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = this.ReadInt32()
	}
	return ints
}
