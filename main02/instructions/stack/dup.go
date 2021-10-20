package stack

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**
dup系列指令复制栈顶的n个变量，然后还是放到操作数栈中
*/

type DUP struct {
	base.NoOperandsInstruction
}

func (this *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
}

/**
下面五条指令先空着
*/
/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
type DUP_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
type DUP_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
type DUP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
type DUP2_X2 struct {
	base.NoOperandsInstruction
}
