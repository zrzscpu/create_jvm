package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       uint
}

func NewFrame(thread *Thread, maxLocals, maxStack uint16) *Frame {

	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (this *Frame) LocalVars() LocalVars {

	return this.localVars
}

func (this *Frame) OperandStack() *OperandStack {

	return this.operandStack
}

func (this *Frame) Thread() *Thread {
	return this.thread
}

func (this *Frame) NextPc() uint {
	return this.nextPc
}

func (this *Frame) SetNextPc(nextPc uint) {
	this.nextPc = nextPc
}
