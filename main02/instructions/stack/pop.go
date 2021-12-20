package stack

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (this *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

//弹出占用两个槽位的元素
type POP2 struct {
	base.NoOperandsInstruction
}

func (this *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
