package rtdata

type OperandStack struct {
	max_stack int
}

func NewOperandStack(maxStack int) *OperandStack {
	return &OperandStack{max_stack: maxStack}
}
