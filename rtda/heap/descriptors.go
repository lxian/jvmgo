package heap

const (
	BOOL   = 'Z'
	BYTE   = 'B'
	CHAR   = 'C'
	SHORT  = 'S'
	INT    = 'I'
	LONG   = 'J'
	FLOAT  = 'F'
	DOUBLE = 'D'
	OBJECT = 'L'
	ARRAY  = '['
	VOID   = 'V'
)

const (
	ARR_BOOL   = "[Z"
	ARR_BYTE   = "[B"
	ARR_CHAR   = "[C"
	ARR_SHORT  = "[S"
	ARR_INT    = "[I"
	ARR_LONG   = "[J"
	ARR_FLOAT  = "[F"
	ARR_DOUBLE = "[D"
)

var primitives_mapping = map[string]rune{
	"boolean": BOOL,
	"byte":    BYTE,
	"char":    CHAR,
	"short":   SHORT,
	"int":     INT,
	"long":    LONG,
	"float":   FLOAT,
	"double":  DOUBLE,
	"void":    VOID,
}
