package classfile

type ConstantPool []ConstantInfo

// 读取常量池信息
func (file *ClassFile) readContantPool(reader *ClassReader) ConstantPool {
	size := reader.readUint16()
	cp := make([]ConstantInfo, size)
	for index := 1; index < int(size); index++ {
		cp[index] = readConstantInfo(reader, cp)
		switch cp[index].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			index++
		}
	}
	return cp
}

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
