package loads

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**
load指令，从局部变量表中获取变量放入操作数栈
*/

/*
先是加载int类型的
*/

/**
先从当前的指令中解析得到应该从局部变量表的那个槽位进行加载
*/
type ILOAD struct {
	base.Index8Instruction
}

func (this *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, this.Index)
}

/**
0号槽位的加载
*/
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (this *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

/**
一号槽位
*/
type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (this *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

/**
2号槽位
*/
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (this *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

/**
提高代码的复用性
*/
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
