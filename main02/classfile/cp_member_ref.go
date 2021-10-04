package classfile

/*源定义
CONSTANT_Fieldref_info {
	u1 tag;	//tag在这里不用的员
			//在constant中先将类型解析好了
	u2 class_index;
	u2 name_and_type_index;
}
*/
/**
这个是对常量池中以下三种类型的抽象
CONSTANT_Fieldref_info
CONSTANT_Methodref_info
CONSTANT_InterfaceMethodref_info
*/
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMemberrefInfo) readInfo(reader *ClassReader) {

	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

//返回类名
func (this *ConstantMemberrefInfo) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

//返回名称和描述符
func (this *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return this.cp.getNameAndType(this.nameAndTypeIndex)
}

//域信息
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

//方法信息
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

//接口信息
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodTypeInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodHandleInfo struct {
	ConstantMemberrefInfo
}

type ConstantInvokeDynamicInfo struct {
	ConstantMemberrefInfo
}
