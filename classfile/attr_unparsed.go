package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func NewUnparsedAttribute(name string, length uint32) *UnparsedAttribute {
	return &UnparsedAttribute{name: name, length: length}
}

func (a UnparsedAttribute) readAttr(reader *ClassReader) {
	a.info = reader.readBytes(a.length)
}
