package classfile

import (
	"fmt"
)

type ClassFile struct {
	// magic uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags		uint16
	thisClassIndex	uint16
	superClassIndex	uint16
	interfaces		[]uint16
	fields			[]*MemberInfo
	methods			[]*MemberInfo
	attributes		[]AttributeInfo
}

func ParseClassBytes(classFileBytes []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	classReader := &ClassReader{classFileBytes}
	classFile = &ClassFile{}
	classFile.Read(classReader)
	return
}

func (classFile *ClassFile) Read(reader *ClassReader) {
	classFile.readVersion(reader)
	classFile.constantPool = readConstantPool(reader)
	classFile.readClassInfo(reader)
	classFile.fields = readMembers(reader, classFile.constantPool)
	classFile.methods = readMembers(reader, classFile.constantPool)
	classFile.attributes = readAttributes(reader, classFile.constantPool)
}


func (classFile *ClassFile) readClassInfo(reader *ClassReader) {
	classFile.accessFlags = reader.readUint16()
	classFile.thisClassIndex = reader.readUint16()
	classFile.superClassIndex = reader.readUint16()
	classFile.interfaces = reader.readUint16s()
}

func (classFile *ClassFile) readVersion(reader *ClassReader) {
	// magic number validation
	var magicNum uint32 = reader.readUint32()
	if magicNum != 0xCAFEBABE {
		panic(fmt.Sprintf("Invalid class file with header 0x%X, expected 0xCAFEBABE", magicNum))
	}

	// read version
	classFile.minorVersion = reader.readUint16()
	classFile.majorVersion = reader.readUint16()
	validateVersion(classFile.majorVersion, classFile.minorVersion)
}

func validateVersion(major uint16, minor uint16) {
	switch major {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if minor == 0 {
			return
		}
	}
	panic(fmt.Sprintf("Unsupported class file version %d.%d. Supported versoins: 45.x, 46.0-52.0", major, minor))
}

