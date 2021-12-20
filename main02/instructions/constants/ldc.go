package constants

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//从常量池中加载到操作数栈中
type LDC struct {
	base.Index8Instruction
}

func (this *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (this *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

//取占用两个位置的常量池元素
type LDC2_W struct {
	base.Index16Instruction
}

func (this *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	pool := frame.Method().Class().ConstantPool()
	constant := pool.GetConstant(this.Index)
	switch constant.(type) {
	case int64:
		stack.PushLong(constant.(int64))

	case float64:
		stack.PushDouble(constant.(float64))
	}
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	pool := frame.Method().Class().ConstantPool()
	constant := pool.GetConstant(index)
	switch constant.(type) {
	case int32:
		stack.PushInt(constant.(int32))
	case float32:
		stack.PushLong(constant.(int64))

		//只用classref？？？？？？
	case *heap.ClassRef:

	default:
		panic("todo: ldc!")

	}
}
