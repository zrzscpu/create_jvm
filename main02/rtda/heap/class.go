package heap

import (
	"create_jvm/main02/classfile"
	"strings"
)

//Class对象
type Class struct {
	accessFlags    uint16
	name           string // thisClassName
	superClassName string
	interfaceNames []string

	constantPool *ConstantPool //运行时常量池

	//这个字段并不是真实的值,而是记录每个字段的信息
	fields []*Field

	methods []*Method

	loader *ClassLoader

	superClass *Class

	interfaces []*Class

	//便于创建对象时进行内存的分配
	instanceSlotCount uint //instanceSlotCount字段存放实例变量占据的空间大小

	staticSlotCount uint //staticSlotCount存放类变量占据的空间大小

	//静态字段
	staticVars Slots
}

func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}

func (this *Class) StaticVars() Slots {
	return this.staticVars
}

//将classfile转为Class对象
func newClass(classfile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classfile.Accessflag()
	class.name = classfile.ClassName()
	class.superClassName = classfile.SuperClassName()
	class.interfaceNames = classfile.InterfaceNames()

	//运行时常量池中保留的是对其他类中方法和字段的引用
	class.constantPool = newConstantPool(class, classfile.Constantpool())
	class.fields = newFields(class, classfile.Fields())
	class.methods = newMethods(class, classfile.Methods())
	return class
}

//判断该类的访问类型
func (this *Class) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (this *Class) IsPrivate() bool {
	return 0 != this.accessFlags&ACC_PRIVATE
}
func (this *Class) IsProtected() bool {
	return 0 != this.accessFlags&ACC_PROTECTED
}

func (this *Class) IsStatic() bool {
	return 0 != this.accessFlags&ACC_STATIC
}

func (this *Class) IsFinal() bool {
	return 0 != this.accessFlags&ACC_FINAL
}

func (this *Class) IsSuper() bool {
	return 0 != this.accessFlags&ACC_SUPER
}

func (this *Class) IsSynchronize() bool {
	return 0 != this.accessFlags&ACC_SYNCHRONIZED
}

func (this *Class) IsVolitile() bool {
	return 0 != this.accessFlags&ACC_VOLATILE
}

//是否是个接口
func (this *Class) IsInterface() bool {
	return 0 != this.accessFlags&ACC_INTERFACE
}

//是否是个抽象类
func (this *Class) IsAbstract() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}

//除非是public 或 在同一个包下，否则不能访问
func (this *Class) isAccessibleTo(class *Class) bool {
	return this.IsPublic() || this.getPackageName() == class.getPackageName()
}

//比如类名是java/lang/Object，则它的包名就是java/lang。
func (this *Class) getPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 {
		return this.name[:i]
	}
	return ""
}

//创建这个Class的一个对象
func (this *Class) NewObject() *Object {
	return newObject(this)
}

func (this *Class) isAssignableFrom(class *Class) bool {
	s, t := this, class

	//如果是同一种类型
	if s == t {
		return true
	}
	//如果是t的子类也可也转换
	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else { //实现了t这个接口
		return s.isImplements(t)
	}
}

func (this *Class) isSubClassOf(t *Class) bool {
	if this.superClass == t {
		return true
	}
	return false
}

func (this *Class) isImplements(t *Class) bool {
	for _, class := range this.interfaces {
		//实现了其接口或子接口
		if class == t || class.isSubInterfaceOf(t) {
			return true
		}
	}
	return false
}

func (this *Class) isSubInterfaceOf(t *Class) bool {
	if this.superClass == t {
		return true
	} else {
		return false
	}
}

func (this *Class) GetMainMethod(name, descriptor string) *Method {
	for _, method := range this.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}
