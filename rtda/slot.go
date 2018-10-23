package rtda

import "jvmgo/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

func (s *Slot) Ref() *heap.Object {
	return s.ref
}

