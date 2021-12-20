package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

//用来保存类中的引用和数值
type Slots []Slot

func newSlots(count uint) Slots {
	return make([]Slot, count)
}

func (this Slots) SetInt(index uint, val int32) {

	this[index].num = val
}

func (this Slots) GetInt(index uint) int32 {
	return this[index].num
}

func (this Slots) SetFloat(index uint, val float32) {

	//Float32bits returns the IEEE 754 binary representation of f,
	// with the sign bit of f and the result in the same bit position.
	// Float32bits(Float32frombits(x)) == x.
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}

func (this Slots) GetFloat(index uint) float32 {

	bits := this[index].num
	resval := math.Float32frombits(uint32(bits))
	return resval
}

func (this Slots) SetLong(index uint, val int64) {
	//先保存低32位
	this[index].num = int32(val)
	//再保存高32位
	this[index+1].num = int32(val >> 32)
}

func (this Slots) GetLong(index uint) int64 {
	low := this[index].num
	high := this[index+1].num
	return int64(high<<32) + int64(low)
	//return int64(high)<<32 | int64(low)
}

func (this Slots) SetDouble(index uint, val float64) {

	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}

func (this Slots) GetDouble(index uint) float64 {

	bits := uint64(this.GetLong(index))
	resval := math.Float64frombits(bits)
	return resval
}

func (this Slots) SetRef(index uint, ref *Object) {

	this[index].ref = ref
}

func (this Slots) GetRef(index uint) *Object {

	return this[index].ref
}
