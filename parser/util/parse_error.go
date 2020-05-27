package util

import (
	"fmt"
	"go-tinyscript/lexer"
)

type ParseError struct {
	msg string
}

func NewParseError(token *lexer.Token) *ParseError {
	return &ParseError{msg: fmt.Sprintf("Syntax Error, unexpected token %s", token.GetValue())}
}

func (e ParseError) Error() string {
	return e.msg
}
