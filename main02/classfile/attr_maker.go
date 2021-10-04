package classfile

/****************************
 */
type DeprecatedAttribute struct {
	MarkerAttribute
}

func (d DeprecatedAttribute) readInfo(reader *ClassReader) {
	panic("implement me")
}

//标记源文件中不存在的由编译器生成的类成员
type SyntheticAttribute struct {
	MarkerAttribute
}

/****************************
以上两种属性已经很少使用，仅仅起到标记作用，不包含任何数据
*/
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing

}
