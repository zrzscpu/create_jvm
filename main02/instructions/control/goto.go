package control

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//无条件跳转
type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.Offset)
}

//func (this *GoToBranch) FetchOperands(reader *base.BytecodeReader) {
//	////
//	//
//	//
//	//val1 := reader.ReadUint8()
//	////this.branchbyte2 = reader.ReadUint8()
//	////this.branchbyte3 = reader.ReadUint8()
//	////this.branchbyte4 = reader.ReadUint8()
//	//
//	////1111 1111 0000 0000 0000 0000 0000 0000
//	//val2 := int32(val1)
//	//val2 <<= 24
//	//val3 := int32(val2)
//	//
//	//println(val1,val2,val3)
//	////vals := val1<< 24
//	////11111111111111111111111111111111 1111 1111 0000 0000 0000 0000 0000 0000
//	////val2 := int32(this.branchbyte2)
//	////vals = val2 << 16
//	////
//	////val3 := int32(this.branchbyte3)
//	////vals = val1 << 8
//	////
//	////val4 := int32(this.branchbyte4)
//	////println(vals)
//	////i := int32(val1 <<24 | val2 <<16 | val3 <<8 | val4)
//	////1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 0011 1011 0010 0000 0000
//	////this.offset = reader.ReadInt32()
//	this.offset = reader.ReadInt32()
//}
