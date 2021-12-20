package heap

import "create_jvm/main02/classfile"

//字段的符号引用
type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool,
	info *classfile.ConstantFieldrefInfo) *FieldRef {

	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (this *FieldRef) ResolvedField() *Field {
	if this.field == nil {
		this.resolvedFieldRef()
	}
	return this.field
}

func (this *FieldRef) resolvedFieldRef() {

	//引用当前域的class对象
	class := this.cp.class

	// 域的 class对象(即域的类型)
	resolvedClass := this.ResolvedClass()
	field := lookupField(resolvedClass, this.name, this.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}
	this.field = field
}

func lookupField(class *Class, name string, descriptor string) *Field {

	//首先在域中找
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	//在接口中找
	for _, iface := range class.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	//在父类中找
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}

	return nil
}
