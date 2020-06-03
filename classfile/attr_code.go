package classfile

/**
Code_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 max_stack;
	u2 max_locals;
	u4 code_length;
	u1 code[code_length];
	u2 exception_table_length;
	{   u2 start_pc;
		u2 end_pc;
		u2 handler_pc;
		u2 catch_type;
	} exception_table[exception_table_length];
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}
**/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxlocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributeInfo  []AttributeInfo
}

func NewCodeAttribute(cp ConstantPool) *CodeAttribute {
	return &CodeAttribute{cp: cp}
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}


func (c *CodeAttribute) readAttr(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxlocals = reader.readUint16()
	c.code = reader.readBytes(reader.readUint32())
	c.exceptionTable = readExceptionTable(reader)
	c.attributeInfo = readAttributes(reader, c.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTable := make([]*ExceptionTableEntry, reader.readUint16())
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
