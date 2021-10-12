package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {

	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (this Frame) LocalVars() LocalVars {

	return this.localVars
}

func (this Frame) OperandStack() *OperandStack {

	return this.operandStack
}
