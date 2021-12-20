package stores

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type A_STORE struct {
	base.Index8Instruction
}

func (this *A_STORE) Execute(frame *rtda.Frame) {
	_astore(frame, this.Index)
}

type A_STORE_0 struct {
	base.NoOperandsInstruction
}

func (this *A_STORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type A_STORE_1 struct {
	base.NoOperandsInstruction
}

func (this *A_STORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type A_STORE_2 struct {
	base.NoOperandsInstruction
}

func (this *A_STORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type A_STORE_3 struct {
	base.NoOperandsInstruction
}

func (this *A_STORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _a_store(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

func _astore(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	frame.LocalVars().SetRef(index, ref)
}
