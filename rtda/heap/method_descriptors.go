package heap

type MethodDescriptor struct {
	paramTypes []string
	returnType string
}

func (desc *MethodDescriptor) appendParam(paramType string) {
	desc.paramTypes = append(desc.paramTypes, paramType)
}

type MethodDescriptorParser struct {
	idx        int
	descriptor string
	paramCnt   uint
	md         *MethodDescriptor
}

func parseMethodDescriptors(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	parser.descriptor = descriptor
	parser.md = &MethodDescriptor{}

	parser.parse()

	return parser.md
}

func (parser *MethodDescriptorParser) readChar() byte {
	c := parser.descriptor[parser.idx]
	parser.idx += 1
	return c
}

func (parser *MethodDescriptorParser) peekChar() byte {
	return parser.descriptor[parser.idx]
}

func (parser *MethodDescriptorParser) readUntil(char byte) string {
	i := parser.idx
	for ; i < len(parser.descriptor) && parser.descriptor[i] != char; i++ {
	}
	result := parser.descriptor[parser.idx : i+1]
	parser.idx = i + 1
	return result
}

func (parser *MethodDescriptorParser) parse() {
	parser.startParams()
	for !parser.readAndEndParams() {
		parser.parseParam()
	}
	parser.parseReturn()

	parser.md.paramTypes = parser.md.paramTypes[:parser.paramCnt]
}

func (parser *MethodDescriptorParser) startParams() {
	if parser.readChar() != '(' {
		panic("Invalid descriptor")
	}
}

func (parser *MethodDescriptorParser) readAndEndParams() bool {
	if parser.peekChar() == ')' {
		parser.readChar()
		return true
	}
	return false
}

func (parser *MethodDescriptorParser) parseParam() {
	parser.md.appendParam(parseType(parser))
	parser.paramCnt += 1
}

func (parser *MethodDescriptorParser) parseReturn() {
	parser.md.returnType = parseType(parser)
}

func parseType(parser *MethodDescriptorParser) string {
	typeChar := parser.readChar()
	switch typeChar {
	case BYTE, CHAR, DOUBLE, FLOAT, INT, LONG, SHORT, BOOL:
		return string(typeChar)
	case OBJECT:
		return "L" + parser.readUntil(';')
	case ARRAY:
		return parseArr(parser)
	case VOID:
		return "V"
	}
	panic("Unrecognized arg/return type")
}

func parseArr(parser *MethodDescriptorParser) string {
	return "[" + parseType(parser)
}
