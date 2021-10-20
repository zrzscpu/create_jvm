package extended

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/instructions/loads"
	"create_jvm/main02/instructions/math"
	"create_jvm/main02/rtda"
)

//扩展了操作数索引（局部变量表索引）的指令
type WIDE struct {
	modifiedInstruction base.Instruction
}

/**
都是在原来的任意位置上加载存储指令上进行扩展
*/
func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	//iload
	case 0x15:
		intst := &loads.ILOAD{}
		//原本这个指令是reader.readuint8
		intst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = intst

	//lload
	case 0x16:

	//fload
	case 0x17:

	//dload
	case 0x18:

	//aload
	case 0x19:

	//istore
	case 0x36:

	//lstore
	case 0x37:

	//fstore
	case 0x38:

	// dstore
	case 0x39:

	// astore
	case 0x3a:

	case 0x84:
		inst := &math.IINNC{}
		inst.Index = uint(reader.ReadInt16())
		inst.Const = int32(reader.ReadUint16())
		this.modifiedInstruction = inst

	// ret
	case 0xa9:
		panic("Unsupported opcode: 0xa9!")

	}

}

func (this *WIDE) Execute(frame *rtda.Frame) {

}
