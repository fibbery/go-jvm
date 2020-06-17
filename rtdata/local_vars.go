package rtdata

import "math"

type LocalVars []Slot

func NewLocalVars(maxLocals int) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (vars LocalVars) SetInt(index uint, value int32) {
	vars[index].num = value
}

func (vars LocalVars) GetInt(index uint) int32 {
	return vars[index].num
}

func (vars LocalVars) SetFloat(index uint, value float32) {
	vars[index].num = int32(math.Float32bits(value))
}

func (vars LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(vars[index].num))
}

func (vars LocalVars) SetLong(index uint, value int64) {
	low := int32(value)
	high := int32(value >> 32)
	vars[index].num = low
	vars[index+1].num = high
}

func (vars LocalVars) GetLong(index uint) int64 {
	low := vars[index].num
	high := vars[index+1].num
	return int64(high)<<32 | int64(low)
}

func (vars LocalVars) SetDouble(index uint, value float64) {
	vars.SetLong(index, int64(math.Float64bits(value)))
}

func (vars LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(vars.GetLong(index)))
}

func (vars LocalVars) SetRef(index uint, ref *Object) {
	vars[index].ref = ref
}

func (vars LocalVars) GetRef(index uint) *Object {
	return vars[index].ref
}
