package stack

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//交换操作数栈栈顶的两个操作数
type SWAP struct {
	base.NoOperandsInstruction
}

func (this *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
