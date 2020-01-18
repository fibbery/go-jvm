package classfile

//CONSTANT_Fieldref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
//CONSTANT_Methodref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
//CONSTANT_InterfaceMethodref_info { u1 tag; u2 class_index; u2 name_and_type_index; }
type ConstantFieldRefInfo struct{ Ref }
type ConstantMethodRefInfo struct{ Ref }
type ConstantInterfaceMethodRefInfo struct{ Ref }

type Ref struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *Ref) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *Ref) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *Ref) NameAndType() *NameAndType {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}
