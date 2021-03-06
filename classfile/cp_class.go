package classfile

//CONSTANT_Class_info { u1 tag; u2 name_index; }
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (c *ConstantClassInfo) readConst(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.cp.readUTF8(c.nameIndex)
}
