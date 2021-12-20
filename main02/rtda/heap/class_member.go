package heap

import "create_jvm/main02/classfile"

//存在class中字段或方法
type ClassMember struct {
	accessFlag uint16
	//这三项原本在常量池中都是索引,这里直接解析位字符串,class更近一部解析为Class 对象
	name       string
	descriptor string
	class      *Class
}

//从classfile中 将 这个引用(类,方法,字段)的	name和描述符填充到这个引用里
func (this *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	this.accessFlag = info.AccessFlags()
	this.name = info.Name()
	this.descriptor = info.Descriptor()
}

//当前引用能否被另一个Class 对象访问
func (this *ClassMember) isAccessibleTo(class *Class) bool {
	if this.isPublic() {
		return true
	}
	c := this.class

	//如果是protected的
	if this.isProtected() {
		//要么在同一个包下,要么可以??????????????????
		return c == class || class.isAccessibleTo(c) ||
			c.getPackageName() == class.getPackageName()
	}

	if !this.isPrivate() {
		return c.getPackageName() == class.getPackageName()
	}

	return c == class

}

//判断当前字段的访问标识
func (this *ClassMember) isPublic() bool {
	return 0 != this.accessFlag&ACC_PUBLIC
}
func (this *ClassMember) isPrivate() bool {
	return 0 != this.accessFlag&ACC_PRIVATE
}
func (this *ClassMember) isProtected() bool {
	return 0 != this.accessFlag&ACC_PROTECTED
}

func (this *ClassMember) isStatic() bool {
	return 0 != this.accessFlag&ACC_STATIC
}

func (this *ClassMember) isFinal() bool {
	return 0 != this.accessFlag&ACC_FINAL
}

func (this *ClassMember) isSuper() bool {
	return 0 != this.accessFlag&ACC_SUPER
}

func (this *ClassMember) isSynchronize() bool {
	return 0 != this.accessFlag&ACC_SYNCHRONIZED
}

func (this *ClassMember) isVolitile() bool {
	return 0 != this.accessFlag&ACC_VOLATILE
}

//get方法
func (this *ClassMember) Class() *Class {
	return this.class
}

func (this *ClassMember) Name() string {
	return this.name
}

func (this *ClassMember) Descriptor() string {
	return this.descriptor
}
