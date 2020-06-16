package classfile

type AccessFlag uint16

const (
	ACC_PUBLIC    AccessFlag = 0x0001
	ACC_PRIVATE              = ACC_PUBLIC * 2
	ACC_PROTECTED            = ACC_PRIVATE * 2
	ACC_STATIC               = ACC_PROTECTED * 2

	ACC_FINAL        AccessFlag = 0x0010
	ACC_SYNCHRONIZED            = ACC_FINAL * 2
	ACC_BRIDGE                  = ACC_SYNCHRONIZED * 2
	ACC_VARARGS                 = ACC_BRIDGE * 2

	ACC_NATIVE    AccessFlag = 0x0100
	ACC_ABSTRACT             = ACC_NATIVE * 2
	ACC_STRICT               = ACC_ABSTRACT * 2
	ACC_SYNTHETIC            = ACC_STRICT * 2
)

//field_info { u2 access_flags; u2 name_index; u2 descriptor_index; u2 attributes_count; attribute_info attributes[attributes_count]; }
//method_info { u2 access_flags; u2 name_index; u2 descriptor_index; u2 attributes_count; attribute_info attributes[attributes_count]; }
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (m MemberInfo) Name() string{
	return m.cp.readUTF8(m.nameIndex)
}

func readMemberInfo(reader *ClassReader, cp ConstantPool) MemberInfo {
	return MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}
