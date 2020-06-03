package classfile

type ReferenceKind uint8

const (
	REF_getField ReferenceKind = iota + 1
	REF_getStatic
	REF_putField
	REF_putStatic
	REF_invokeVirtual
	REF_invokeStatic
	REF_invokeSpecial
	REF_newInvokeSpecial
	REF_invokeInterface
)

//CONSTANT_MethodHandle_info { u1 tag; u1 reference_kind; u2 reference_index; }
type ConstantMethodHandleInfo struct {
	referenceKind  ReferenceKind
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readConst(reader *ClassReader) {
	c.referenceKind = ReferenceKind(reader.readUint8())
	c.referenceIndex = reader.readUint16()
}
