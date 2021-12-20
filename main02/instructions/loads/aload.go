package loads

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type A_LOAD struct {
	base.Index8Instruction
}

func (this *A_LOAD) Execute(frame *rtda.Frame) {
	_aload(frame, this.Index)
}

type A_LOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *A_LOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type A_LOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *A_LOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type A_LOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *A_LOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type A_LOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *A_LOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

/**
提高代码的复用性
*/
func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}
