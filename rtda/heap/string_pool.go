package heap

import "unicode/utf16"

// mapping gostring -> Java String object
var internedStrings = map[string]*Object{}

func InternedJString(gostr string, loader *ClassLoader) *Object {
	if jstr, ok := internedStrings[gostr]; ok {
		return jstr
	}

	chars := stringToUTF16(gostr)
	charArr := &Object{loader.LoadClass("[C"), chars, nil}

	jstr := loader.LoadClass("java/lang/String").NewObject()
	jstr.SetRefVar("value", "[C", charArr)

	internedStrings[gostr] = jstr
	return jstr
}

func GoString(jstr *Object) string {
	if jstr == nil { return ""}
	return string(utf16.Decode(jstr.GetRefVar("value", "[C").Chars()))
}

func JString(goStr string, loader *ClassLoader) *Object {
	return InternedJString(goStr, loader)
}

func stringToUTF16(s string) []uint16 {
	return utf16.Encode([]rune(s))
}
