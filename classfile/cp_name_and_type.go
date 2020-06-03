package classfile

type NameAndType struct {
	Name       string
	Descriptor string
}

//CONSTANT_NameAndType_info { u1 tag; u2 name_index; u2 descriptor_index; }
type ConstantNameAndType struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndType) readConst(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}

func (c *ConstantNameAndType) getNameAndType() *NameAndType {
	name := c.cp.readUTF8(c.nameIndex)
	descriptor := c.cp.readUTF8(c.descriptorIndex)
	return &NameAndType{name, descriptor}
}
