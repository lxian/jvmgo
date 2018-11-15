package heap

import "jvmgo/classfile"

type ExceptionEntry struct {
	startPc       int
	endPc         int
	handlerPc     int
	catchClassRef *ClassRef
}

type ExceptionTable []*ExceptionEntry

func (table ExceptionTable) FindExceptionHandler(pc int, expObj *Object) int {
	for _, ee := range table {
		if ee.startPc <= pc && pc < ee.endPc {
			if ee.catchClassRef == nil || expObj.IsInstanceOf(ee.catchClassRef.ResolvedClass()) {
				return ee.handlerPc
			}
		}
	}
	return -1
}

func newExceptionTable(cp *ConstantPool, code *classfile.CodeAttribute) ExceptionTable {
	table := make([]*ExceptionEntry, len(code.ExceptionTable()))
	for i, e := range code.ExceptionTable() {
		expClz := getCatchClass(cp, uint(e.CatchType()))
		table[i] = &ExceptionEntry{
			int(e.StartPc()),
			int(e.EndPc()),
			int(e.HandlerPc()),
			expClz,
		}
	}
	return table
}

func getCatchClass(cp *ConstantPool, catchTypeIdx uint) *ClassRef {
	if catchTypeIdx == 0 {
		return nil
	}
	return cp.GetConstant(catchTypeIdx).(*ClassRef)
}

