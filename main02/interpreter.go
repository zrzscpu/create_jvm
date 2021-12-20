package main

import (
	"create_jvm/main02/instructions"
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
	"fmt"
)

//func interpret(info *classfile.MemberInfo) {
//	//读取当前方法的code属性表
//	codeAttr := info.CodeAttribute()
//	locals := codeAttr.MaxLocals()
//	stack := codeAttr.MaxStack()
//	code := codeAttr.Code()
//
//	//为该方法创建线程
//	thread := rtda.NewThread()
//	frame := thread.NewFrame(locals, stack)
//	thread.PushFrame(frame)
//
//	defer catchErr(frame);
//	loop(thread, code)
//}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

//现在指令码都在bye[]数组中
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//下个要执行的指令
		pc := frame.NextPc()
		thread.SetPC(pc)

		//设置下次读取指令的索引
		reader.Reset(bytecode, pc)
		//操作码
		opcode := reader.ReadUint8()
		//创建指令对象
		instruction := instructions.NewInstructions(opcode)
		//获取操作数
		instruction.FetchOperands(reader)

		//设置下条指令的索引
		frame.SetNextPc(reader.Pc())

		//执行pc :%2d institution:%T %v
		// operand %T
		fmt.Println("%T", frame.OperandStack())
		instruction.Execute(frame)
	}
}

func interpret(method *heap.Method) {

	//为该方法创建线程
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}
