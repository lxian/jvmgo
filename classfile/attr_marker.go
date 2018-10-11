package classfile

type MarkerAttribute struct{}

type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}

func (_ *MarkerAttribute) readInfo(reader *ClassReader) {
}
