package lexer

import "fmt"

// LexicalError 词法解析器报错
type LexicalError struct {
	msg string
}

func (e LexicalError) Error() string {
	return e.msg
}

// NewLexicalError 新建一个错误
func NewLexicalError(msg string) LexicalError {
	return LexicalError{msg: msg}
}

// NewLexicalErrorWithChar 新建一个由于出现异常字符所造成的词法解析器错误
func NewLexicalErrorWithChar(char string) LexicalError {
	return NewLexicalError(fmt.Sprintf("Unexpected character %s", char))
}
