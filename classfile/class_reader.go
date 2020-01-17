package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}

func (c *ClassReader) readUint16() uint16 {
	var val uint16
	if determineEndian() == 1 {
		val = binary.BigEndian.Uint16(c.data)
	} else {
		val = binary.LittleEndian.Uint16(c.data)
	}
	c.data = c.data[2:]
	return val
}

func (c *ClassReader) readUint32() uint32 {
	var val uint32
	if determineEndian() == 1 {
		val = binary.BigEndian.Uint32(c.data)
	} else {
		val = binary.LittleEndian.Uint32(c.data)
	}
	c.data = c.data[4:]
	return val
}

func (c *ClassReader) readUint64() uint64 {
	var val uint64
	if determineEndian() == 1 {
		val = binary.BigEndian.Uint64(c.data)
	} else {
		val = binary.LittleEndian.Uint64(c.data)
	}
	c.data = c.data[8:]
	return val
}

// readUint16s get a array, length is the first element
func (c *ClassReader) readUint16s() []uint16 {
	length := c.readUint16()
	data := make([]uint16, length)
	for i := range data {
		data[i] = c.readUint16()
	}
	return data
}

func (c *ClassReader) readBytes(length uint32) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}

// determineEndian return 1 while system is bigendian
// return 0 while system is littleendian
func determineEndian() uint8 {
	a := uint16(0x1234)
	b := uint8(a)
	if b == 0x34 {
		return 1
	} else {
		return 0
	}
}
