package lexer

// TokenType token类型
type TokenType string

const (
	// KEYWORD 关键词
	KEYWORD TokenType = "keyword"
	// VARIABLE 变量
	VARIABLE TokenType = "variable"
	// OPERATOR 操作符
	OPERATOR TokenType = "operator"
	// BRACKET 括号
	BRACKET TokenType = "bracket"
	// INTEGER 整数
	INTEGER = "integer"
	// STRING 字符串
	STRING = "string"
	// FLOAT 浮点数
	FLOAT = "float"
	// BOOLEAN 布尔值
	BOOLEAN = "boolean"
)

func (t *TokenType) String() string {
	return string(*t)
}
