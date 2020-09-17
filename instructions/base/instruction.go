package base

import "github.com/fibbery/go-jvm/rtdata"

type Instruction interface {
	FetchOperand(reader *BytecodeReader)
	Execute(frame *rtdata.Frame)
}