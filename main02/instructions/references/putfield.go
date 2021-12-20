package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//放到对象的域中,域的名称由常量池索引指向
//对象 和 要放到对象中的数据 从操作数栈中获取
type PUT_FIELD struct {
	base.Index16Instruction
}

func (this *PUT_FIELD) Execute(frame *rtda.Frame) {
	//获取域对象
	method := frame.Method()
	class := method.Class()
	pool := class.ConstantPool()
	fieldref := pool.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldref.ResolvedField()

	//如果是static的,不能通过对象获取
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//final修饰的实例变量,只能在当前类型的实例内的构造方法中进行修改
	if field.IsFinal() {
		if class != field.Class() || method.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	//通过域对象找相应的set方法
	id := field.SlotId()
	descriptor := field.Descriptor()
	stack := frame.OperandStack()
	switch descriptor[0] {

	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		object := stack.PopRef()
		object.Field().SetInt(id, val)

	case 'F':
		val := stack.PopFloat()
		object := stack.PopRef()
		object.Field().SetFloat(id, val)

	case 'J':
		val := stack.PopLong()
		object := stack.PopRef()
		object.Field().SetLong(id, val)

	case 'D':
		val := stack.PopDouble()
		object := stack.PopRef()
		object.Field().SetDouble(id, val)

	case 'L', '[':
		val := stack.PopRef()
		object := stack.PopRef()
		object.Field().SetRef(id, val)
	}

}
