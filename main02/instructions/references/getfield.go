package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//常量池索引,操作数栈中保存了要从哪个对象中取字段的这个对象。取出来的字段放回操作数栈
type GET_FIELD struct {
	base.Index16Instruction
}

func (this *GET_FIELD) Execute(frame *rtda.Frame) {
	method := frame.Method()
	class := method.Class()
	pool := class.ConstantPool()

	//对于jvm而言是不知道实例对象中各个字段的类型的，需要根据filed来进行确定
	fieldRef := pool.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.OperandStack()
	ref := stack.PopRef()

	if ref != nil {
		id := field.SlotId()
		descriptor := field.Descriptor()
		stack := frame.OperandStack()
		switch descriptor[0] {

		case 'Z', 'B', 'C', 'S', 'I':
			stack.PushInt(ref.Field().GetInt(id))

		case 'F':
			stack.PushFloat(ref.Field().GetFloat(id))

		case 'J':
			stack.PushLong(ref.Field().GetLong(id))

		case 'D':
			stack.PushDouble(ref.Field().GetDouble(id))

		case 'L', '[':
			stack.PushRef(ref.Field().GetRef(id))
		}
	} else {
		panic("java.lang.NullPointerException")
	}
}
