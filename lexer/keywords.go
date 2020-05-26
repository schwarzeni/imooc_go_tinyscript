package lexer

// keywords 关键词列表
var keywords = map[string]struct{}{
	"var":    {},
	"int":    {},
	"float":  {},
	"bool":   {},
	"void":   {},
	"string": {},
	"if":     {},
	"else":   {},
	"for":    {},
	"while":  {},
	"break":  {},
	"func":   {},
	"return": {},
}

// IsKeyword 判断传入参数是否为关键词
func IsKeyword(word string) bool {
	_, ok := keywords[word]
	return ok
}
