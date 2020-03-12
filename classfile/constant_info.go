package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// todo read constant info
	tag := reader.readUint8()
	info := newConstantInfo(tag, cp)
	info.readInfo(reader)
	return info
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{Ref{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{Ref{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{Ref{cp: cp}}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{cp: cp}
	case CONSTANT_Float:
		return &ConstantFloatInfo{cp: cp}
	case CONSTANT_Long:
		return &ConstantLongInfo{cp: cp}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{cp: cp}
	case CONSTANT_NameAndType:
		return &ConstantNameAndType{cp: cp}
	case CONSTANT_Utf8:
		return &ConstantUTF8{cp: cp}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	}

	panic("parse error while read constant info")
}
