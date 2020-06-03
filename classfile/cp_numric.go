package classfile

import "math"

// CONSTANT_Integer_info { u1 tag; u4 bytes; }
type ConstantIntegerInfo struct {
	cp    ConstantPool
	value int32
}

func (c *ConstantIntegerInfo) readConst(reader *ClassReader) {
	c.value = int32(reader.readUint32())
}

type ConstantFloatInfo struct {
	cp    ConstantPool
	value float32
}

func (c *ConstantFloatInfo) readConst(reader *ClassReader) {
	c.value = math.Float32frombits(reader.readUint32())
}

type ConstantLongInfo struct {
	cp    ConstantPool
	value int64
}

func (c *ConstantLongInfo) readConst(reader *ClassReader) {
	c.value = int64(reader.readUint64())
}

type ConstantDoubleInfo struct {
	cp    ConstantPool
	value float64
}

func (c *ConstantDoubleInfo) readConst(reader *ClassReader) {
	c.value = math.Float64frombits(reader.readUint64())
}
