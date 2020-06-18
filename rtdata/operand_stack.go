package rtdata

import "math"

type OperandStack struct {
	size  int
	slots []Slot
}

func NewOperandStack(maxStack int) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (o *OperandStack) PushInt(value int32) {
	o.slots[o.size].num = value
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

func (o *OperandStack) PutFloat(value float32) {
	o.slots[o.size].num = int32(math.Float32bits(value))
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	bits := (uint32)(o.slots[o.size].num)
	return math.Float32frombits(bits)
}

func (o *OperandStack) PutLong(value int64) {
	//low
	o.slots[o.size].num = int32(value)
	//high
	o.slots[o.size+1].num = int32(value >> 32)
	o.size = o.size + 2
}

func (o *OperandStack) PopLong() int64 {
	o.size = o.size - 2
	low := int64(o.slots[o.size].num)
	high := int64(o.slots[o.size+1].num << 32)
	return high | low
}

func (o *OperandStack) PushDouble(value float64) {
	bits := math.Float64bits(value)
	o.PutLong(int64(bits))
}

func (o *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(o.PopFloat()))
}

func (o *OperandStack) PushRef(ref *Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	return o.slots[o.size].ref
}
