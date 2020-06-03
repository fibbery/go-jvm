package classfile

/**
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/

type ConstantValueAttribute struct {
	cp ConstantPool
	value ConstantInfo
}

func NewConstantValueAttribute(cp ConstantPool) *ConstantValueAttribute {
	return &ConstantValueAttribute{cp: cp}
}


func (c *ConstantValueAttribute) readAttr(reader *ClassReader) {
	c.value = c.cp.getConstantInfo(reader.readUint16())
}

