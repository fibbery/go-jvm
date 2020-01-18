package classfile

type ConstantUTF8 struct {
	cp    ConstantPool
	value string
}

func (c *ConstantUTF8) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	data := reader.readBytes(length)
	c.value = string(data)
}

func (c *ConstantUTF8) String() string {
	return c.value
}
