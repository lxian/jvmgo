package heap

/**
					CLASS	FIELD	METHOD
ACC_PUBLIC			x		x		x
ACC_FINAL			x		x		x
ACC_SUPER			x
ACC_INTERFACE		x
ACC_ABSTRACT		x
ACC_SYNTHETIC		x		x		x
ACC_ANNOTATION		x
ACC_ENUM			x		x
ACC_PRIVATE					x		x
ACC_PROTECTED				x		x
ACC_STATIC					x		x
ACC_VOLATILE				x
ACC_TRANSIENT				x
ACC_SYNCHRONIZED 					x
ACC_BRIDGE							x
ACC_VARARGS							x
ACC_NATIVE							x
ACC_STRICT							x
*/

const (
	ACC_PUBLIC       = 0x0001
	ACC_FINAL        = 0x0010
	ACC_SUPER        = 0x0020
	ACC_INTERFACE    = 0x0200
	ACC_ABSTRACT     = 0x0400
	ACC_SYNTHETIC    = 0x1000
	ACC_ANNOTATION   = 0x2000
	ACC_ENUM         = 0x4000
	ACC_PRIVATE      = 0x0002
	ACC_PROTECTED    = 0x0004
	ACC_STATIC       = 0x0008
	ACC_VOLATILE     = 0x0040
	ACC_TRANSIENT    = 0x0080
	ACC_SYNCHRONIZED = 0x0020
	ACC_BRIDGE       = 0x0040
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	ACC_STRICT       = 0x0800
)

func HasFlag(flags uint16, target ...uint16) bool {
	for _, t := range target {
		if (flags & t) == 0 {
			return false
		}
	}
	return true
}
