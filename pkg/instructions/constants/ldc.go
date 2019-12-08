package constants

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type LDC struct {
	base.Index8Instruction
}

func (this *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

type LDCW struct {
	base.Index16Instruction
}

func (this *LDCW) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	val := frame.Method().Class().ConstantPool().GetConstant(index)
	operandStack := frame.OperandStack()

	switch val.(type) {
	case int32:
		operandStack.PushInt(val.(int32))
	case float32:
		operandStack.PushFloat(val.(float32))
	// case string:
	// case *heap.ClassSymRef:
	default:
		panic("todo: ldc!")
	}
}

type LDC2W struct {
	base.Index16Instruction
}

func (this *LDC2W) Execute(frame *rtda.Frame) {
	val := frame.Method().Class().ConstantPool().GetConstant(this.Index)
	operandStack := frame.OperandStack()

	switch val.(type) {
	case float64:
		operandStack.PushDouble(val.(float64))
	case int64:
		operandStack.PushLong(val.(int64))
	default:
		panic("java.lang.ClassFormatError")
	}
}