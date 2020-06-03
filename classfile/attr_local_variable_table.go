package classfile
/**
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;  code里面变量开始生效的index
        u2 length;   code里面变量结束有效期的index
        u2 name_index;
        u2 descriptor_index;
        u2 index; 在本地变量表的槽位
    } local_variable_table[local_variable_table_length];
}
 */

type LocalVariableTableAttribute struct {
	cp ConstantPool
	localVariableTable []*LocalVariableEntry
}

func (l *LocalVariableTableAttribute) readAttr(reader *ClassReader) {
	l.localVariableTable= make([]*LocalVariableEntry,reader.readUint16())
	for index := range l.localVariableTable {
		l.localVariableTable[index] = &LocalVariableEntry{
			startPc:    reader.readUint16(),
			length:     reader.readUint16(),
			name:       l.cp.readUTF8(reader.readUint16()),
			descriptor: l.cp.readUTF8(reader.readUint16()),
			slot:       reader.readUint16(),
		}
	}
}

func NewLocalVariableTableAttribute(cp ConstantPool) *LocalVariableTableAttribute {
	return &LocalVariableTableAttribute{cp: cp}
}

type LocalVariableEntry struct {
	startPc uint16
	length uint16
	name string
	descriptor string
	slot uint16
}
