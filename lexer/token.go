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

// IsNumber 判断是否为数字：整型或浮点型
func (t *Token) IsNumber() bool {
	return t._type == INTEGER || t._type == FLOAT
}

// IsValue 是否是一个值
func (t *Token) IsValue() bool {
	return t.IsVariable() || t.IsScalar()
}

// IsOperator 判断是否为运算符
func (t *Token) IsOperator() bool {
	return t._type == OPERATOR
}

// NewToken 新建 Token 实例
func NewToken(_type TokenType, _value string) *Token {
	return &Token{_type: _type, _value: _value}
}
