package math

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"math"
)

/**
四条求余数指令
*/
type DREM struct {
	base.NoOperandsInstruction
}

func (this *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double2 := stack.PopDouble()
	double1 := stack.PopDouble()
	if double2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := math.Mod(double1, double2)
	stack.PushDouble(res)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (this *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	float2 := stack.PopFloat()
	float1 := stack.PopFloat()
	if float2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	double2 := float64(float2)
	double1 := float64(float1)
	res := math.Mod(double1, double2)
	stack.PushFloat(float32(res))
}

type LREM struct {
	base.NoOperandsInstruction
}

func (this *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	long2 := stack.PopLong()
	long1 := stack.PopLong()
	if long2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := long2 % long1
	stack.PushLong(res)
}

type IREM struct {
	base.NoOperandsInstruction
}

func (this *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	int2 := stack.PopInt()
	int1 := stack.PopInt()
	if int2 == 0.0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := int2 % int1
	stack.PushInt(res)
}
