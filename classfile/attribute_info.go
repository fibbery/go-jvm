package classfile

//attribute_info { u2 attribute_name_index; u4 attribute_length; u1 info[attribute_length]; }
type AttributeInfo struct {
	nameIndex uint16
}
