package classfile

/**
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
 */
type DeprecatedAttribute struct {
	EmptyTag
}

func NewDeprecatedAttribute() *DeprecatedAttribute {
	return &DeprecatedAttribute{}
}


/**
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
 */
type SyntheticAttribute struct {
	EmptyTag
}

func NewSyntheticAttribute() *SyntheticAttribute {
	return &SyntheticAttribute{}
}

type EmptyTag struct {
}

func (e *EmptyTag) readAttr(reader *ClassReader) {
}

