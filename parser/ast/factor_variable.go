package ast

import "go-tinyscript/lexer"

type Variable struct {
	*Factor
	typeLexeme *lexer.Token
}

func (v *Variable) TypeLexeme() *lexer.Token {
	return v.typeLexeme
}

func (v *Variable) SetTypeLexeme(typeLexeme *lexer.Token) {
	v.typeLexeme = typeLexeme
}

func NewVariable(token *lexer.Token) *Variable {
	v := &Variable{Factor: NewFactor(token)}
	v.nodeType = VARIABLE
	return v
}
