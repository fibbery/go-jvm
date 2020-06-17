package rtdata

type Frame struct {
	// next frame address
	lower *Frame
	// local variables table
	localVars LocalVars
	// operand stack
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack int) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}
