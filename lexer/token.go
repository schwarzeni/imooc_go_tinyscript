package lexer

import (
	"fmt"
)

// Token 每一个字段块的信息
type Token struct {
	_type  TokenType
	_value string
}

func (t *Token) getType() TokenType {
	return t._type
}

func (t *Token) getValue() string {
	return t._value
}

func (t *Token) String() string {
	return fmt.Sprintf("type %s, value %s", t._type, t._value)
}

// IsVariable 判断是否为变量类型
func (t *Token) IsVariable() bool {
	return t._type == VARIABLE
}

// IsScalar 判断是否为变量值
func (t *Token) IsScalar() bool {
	return t._type == INTEGER || t._type == STRING ||
		t._type == FLOAT || t._type == BOOLEAN
}

// NewToken 新建 Token 实例
func NewToken(_type TokenType, _value string) *Token {
	return &Token{_type: _type, _value: _value}
}
