package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
const (
	Attr_ConstantValue    = "ConstantValue"
	Attr_Code             = "Code"
	Attr_StackMapTable    = "StackMapTable"
	Attr_Exceptions       = "Exceptions"
	Attr_BootstrapMethods = "BootstrapMethods"

	Attr_InnerClasses                         = "InnerClasses"
	Attr_EnclosingMethod                      = "EnclosingMethod"
	Attr_Synthetic                            = "Synthetic"
	Attr_Signature                            = "Signature"
	Attr_RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	Attr_RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	Attr_RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	Attr_RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	Attr_RuntimeVisibleTypeAnnotations        = "RuntimeVisibleTypeAnnotations"
	Attr_RuntimeInvisibleTypeAnnotations      = "RuntimeInvisibleTypeAnnotations"
	Attr_AnnotationDefault                    = "AnnotationDefault"

	Attr_SourceFile             = "SourceFile"
	Attr_SourceDebugExtension   = "SourceDebugExtension"
	Attr_LineNumberTable        = "LineNumberTable"
	Attr_LocalVariableTable     = "LocalVariableTable"
	Attr_LocalVariableTypeTable = "LocalVariableTypeTable"
	Attr_Deprecated             = "Deprecated"
)

type AttributeInfo interface {
	readAttr(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	length := reader.readUint16()
	infos := make([]AttributeInfo, length)
	for index := range infos {
		infos[index] = readAttribute(reader, cp)
	}
	return infos
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	nameIndex := reader.readUint16()
	attrName := cp.readUTF8(nameIndex)
	attrLen := reader.readUint32()

	attribute := newAttributeInfo(cp, attrName, attrLen)
	attribute.readAttr(reader)
	return attribute
}

func newAttributeInfo(cp ConstantPool, name string, length uint32) AttributeInfo {
	switch name {
	case Attr_Code:
		return NewCodeAttribute(cp)
	case Attr_ConstantValue:
		return NewConstantValueAttribute(cp)
	case Attr_Deprecated:
		return NewDeprecatedAttribute()
	case Attr_Exceptions:
		return NewExceptionsAttribute(cp)
	case Attr_LineNumberTable:
		return NewLineNumberTableAttribute(cp)
	case Attr_LocalVariableTable:
		return NewLocalVariableTableAttribute(cp)
	case Attr_SourceFile:
		return NewSourceFileAttribute(cp)
	case Attr_Synthetic:
		return NewSyntheticAttribute()
	}
}
