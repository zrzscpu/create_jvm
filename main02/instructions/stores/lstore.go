package stores

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**
将操作数从栈顶弹出,放到本地变量表的某个位置的指令
*/

/**
加载到第index个槽位
*/
type LSTORE struct {
	base.Index8Instruction
}

func (this *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, this.Index)
}

/**
加载到第0个槽位
*/
type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (this *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (this *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (this *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type ISTORE struct {
	base.Index8Instruction
}

func (this *ISTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, this.Index)
}

/**
加载到第0个槽位
*/
type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (this *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (this *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (this *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
