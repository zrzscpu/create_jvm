package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//操作数栈中放着引用,索引是一个对象引用,即比较的是两个引用
type INSTANCEOF struct {
	base.Index16Instruction
}

func (this *INSTANCEOF) Execute(frame *rtda.Frame) {
	pool := frame.Method().Class().ConstantPool()
	stack := frame.OperandStack()
	ref1 := stack.PopRef()
	if ref1 == nil {
		stack.PushInt(0)
	}
	ref2 := pool.GetConstant(this.Index).(*heap.ClassRef)

	if ref1.IsInstanceOf(ref2.ResolvedClass()) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}

}
