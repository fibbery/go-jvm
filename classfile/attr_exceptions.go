package classfile

/**
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	cp ConstantPool
	classnames []string
}

func NewExceptionsAttribute(cp ConstantPool) *ExceptionsAttribute {
	return &ExceptionsAttribute{cp: cp}
}

func (e *ExceptionsAttribute) readAttr(reader *ClassReader) {
	exceptionIndexTable := reader.readUint16s()
	e.classnames = make([]string,len(exceptionIndexTable))
	for index,value := range exceptionIndexTable {
		e.classnames[index] = e.cp.getClassName(value)
	}
}


