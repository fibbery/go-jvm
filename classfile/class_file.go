package classfile

import "fmt"

/**
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	magic           uint32
	minorVersion    uint16
	majorVersion    uint16
	constPool       ConstantPool
	accessFlag      uint16
	thisClass       uint16
	superClass      uint16
	interfacesCount uint16
	// interfaces []
	filedsCount uint16
	// field_info []
	methodsCount uint16
	// method_info []
	attributesCount uint16
	// attribute_info []
}

func (file *ClassFile) read(reader *ClassReader) {
	file.readAndCheckMagicNumber(reader)
	file.readAndCheckVersion(reader)
	file.readContantPool(reader)
	file.accessFlag = reader.readUint16()
	file.thisClass = reader.readUint16()
	file.superClass = reader.readUint16()
}

func (file *ClassFile) readAndCheckMagicNumber(reader *ClassReader) {
	u := reader.readUint32()
	if u != 0xCAFEBABE {
		panic("class file format error : magic number!!!")
	}
}

// readAndCheckVersion check is supported jdk version
func (file *ClassFile) readAndCheckVersion(reader *ClassReader) {
	file.majorVersion = reader.readUint16()
	file.minorVersion = reader.readUint16()
	switch file.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if file.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (file *ClassFile) readContantPool(reader *ClassReader) {
	//todo read constant pool
	size := reader.readUint16()
}

func Parse(data []byte) (*ClassFile, error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if !ok {
				fmt.Printf("%v/n", r)
			}
		}
	}()
	reader := ClassReader{data}
	file := &ClassFile{}
	file.read(&reader)
	return file, nil
}