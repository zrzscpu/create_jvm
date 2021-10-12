package rtda

import "math"

type OperandStack struct {
	index uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
		index: 0,
	}
}

func (this *OperandStack) PushInt(val int32) {

	this.slots[this.index].num = val
	this.index++
}

func (this *OperandStack) PopInt() int32 {
	this.index--
	return this.slots[this.index].num
}

func (this *OperandStack) PushFloat(val float32) {

	bits := math.Float32bits(val)
	this.slots[this.index].num = int32(bits)
	this.index++
}

func (this *OperandStack) PopFloat() float32 {

	this.index--
	bits := this.slots[this.index].num
	return math.Float32frombits(uint32(bits))
}

func (this *OperandStack) PushLong(val int64) {

	low := int32(val)
	this.slots[this.index].num = low
	this.index++
	high := int32(val >> 32)
	this.slots[this.index].num = high
	this.index++
}

func (this *OperandStack) PopLong() int64 {

	this.index--
	high := this.slots[this.index].num
	this.index--
	low := this.slots[this.index].num
	return int64(high<<32) + int64(low)
}

func (this *OperandStack) PushDouble(val float64) {

	this.PushLong(int64(math.Float64bits(val)))
}

func (this *OperandStack) PopDouble() float64 {

	bits := uint64(this.PopLong())
	return math.Float64frombits(bits)
}

func (this *OperandStack) PushRef(ref *Object) {
	this.slots[this.index].ref = ref
}

func (this *OperandStack) PopRef() *Object {
	this.index--
	ref := this.slots[this.index].ref
	this.slots[this.index].ref = nil
	return ref
}
