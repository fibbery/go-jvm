package classfile

type ConstantPool []ConstantInfo

func (c ConstantPool) readUTF8(index uint16) string {
	return c.getConstantInfo(index).(*ConstantUTF8).String()
}

func (c ConstantPool) getClassName(index uint16) string {
	return c.getConstantInfo(index).(*ConstantClassInfo).Name()
}

func (c ConstantPool) getNameAndType(index uint16) *NameAndType {
	return c.getConstantInfo(index).(*ConstantNameAndType).getNameAndType()
}

func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := c[index]; info != nil {
		return info
	}
	panic("can't find constantinfo in index " + string(index))
}
