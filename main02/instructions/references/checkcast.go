package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (this *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	objref := stack.PopRef()
	stack.PushRef(objref)

	objClassRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.ClassRef)
	if objref == nil {
		return
	}
	class := objClassRef.ResolvedClass()
	if objref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
