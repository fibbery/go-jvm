package classfile

/**
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;  // 在 code attribute 中的开始index
        u2 line_number; //具体代码行数
    } line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	cp              ConstantPool
	lineNumberTable []*LineNumberEntry
}

func (l *LineNumberTableAttribute) readAttr(reader *ClassReader) {
	l.lineNumberTable = make([]*LineNumberEntry, reader.readUint16())
	for index := range l.lineNumberTable {
		l.lineNumberTable[index] = &LineNumberEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func NewLineNumberTableAttribute(cp ConstantPool) *LineNumberTableAttribute {
	return &LineNumberTableAttribute{cp: cp}
}

type LineNumberEntry struct {
	startPc    uint16
	lineNumber uint16
}
