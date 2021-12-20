package heap

import "create_jvm/main02/classfile"

//Class的字段
type Field struct {
	ClassMember

	//为了便于创建类时候所需内存大小的计算,给每个字段设置一个slotid用于技术
	slotId uint

	constValueIndex uint
}

//如果 域 是final 修饰的，那它下面会有一个 <ConstantValue>属性表，
//里面存放一个指向常量池中一个常量的索引
func (this *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		this.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

//创建字段属性
func newFields(class *Class, classfileFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(classfileFields))
	//遍历每个fields属性
	for index, val := range classfileFields {
		fields[index] = &Field{}
		fields[index].class = class
		//名称和描述符
		fields[index].copyMemberInfo(val)
		//final修饰的下有个constantattribute
		fields[index].copyAttributes(val)
	}
	return fields
}

func (this *Field) IsStatic() bool {
	return 0 != this.accessFlag&ACC_STATIC
}

func (this *Field) IsLongOrDouble() bool {
	return this.descriptor == "J" || this.descriptor == "D"
}

func (this *Field) IsFinal() bool {
	return 0 != this.accessFlag&ACC_FINAL
}

//运行时常量池索引
func (this *Field) ConstantValueIndex() uint {
	return this.constValueIndex
}

//在slot中索引
func (this *Field) SlotId() uint {
	return this.slotId
}
