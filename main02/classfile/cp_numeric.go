package classfile

import "math"

/**源定义
CONSTANT_Integer_info {
		u1 tag;
		u4 bytes;
}
*/
/**
描述int类型
*/
type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {

	bytes := reader.readUint32()
	this.val = int32(bytes)

}

/*占据两个位置的常量池引用
CONSTANT_Long_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
*/
/**
描述long类型	u8类型
*/
type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {

	val := reader.readUint64()
	this.val = int64(val)
}

/*源定义
CONSTANT_Float_info {
		u1 tag;
		u4 bytes;
}
*/
/**
描述float类型
*/
type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {

	val := reader.readUint32()
	this.val = math.Float32frombits(val)
}

/*源定义
CONSTANT_Double_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
*/
/**
描述double类型	u8类型
*/
type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {

	val := reader.readUint64()
	this.val = math.Float64frombits(val)
}
