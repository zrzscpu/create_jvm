package math

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//算数右移看符号位,逻辑不看
/**
int 类型左移
*/
type ISHL struct {
	base.NoOperandsInstruction
}

func (this *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2)
	//取后五个bit就可以表示32位以内的移动
	res := val1 << s & 0x1f
	stack.PushInt(res)
}

type ISHR struct {
	base.NoOperandsInstruction
} // int算术右位移
func (this *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	res := val1 >> s
	stack.PushInt(res)
}

type IUSHR struct {
	base.NoOperandsInstruction
} // int逻辑右位移
//逻辑右移动,得先转成无符号数才能实现,因为没有>>>的运算符
func (this *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	res := int32(uint32(val1) >> s)
	stack.PushInt(res)
}

type LSHL struct {
	base.NoOperandsInstruction
} // long算数左移
func (this *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	res := val1 << s
	stack.PushLong(res)
}

type LSHR struct {
	base.NoOperandsInstruction
} // long算术右位移
func (this *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	//只需要后6个字节即可
	s := uint32(val2) & 0x3f
	result := val1 >> s
	stack.PushLong(result)
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (this *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()

	//
	s := uint32(val2) & 0x3f
	result := int64(uint64(val1) >> s)
	stack.PushLong(result)
}
