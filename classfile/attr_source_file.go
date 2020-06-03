package classfile

/**
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

type SourceFileAttribute struct {
	cp             ConstantPool
	sourceFileName string
}

func NewSourceFileAttribute(cp ConstantPool) *SourceFileAttribute {
	return &SourceFileAttribute{cp: cp}
}

func (s *SourceFileAttribute) readAttr(reader *ClassReader) {
	s.sourceFileName = s.cp.readUTF8(reader.readUint16())
}
