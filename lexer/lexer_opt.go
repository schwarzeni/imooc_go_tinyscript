package lexer

import (
	"go-tinyscript/common"
	"strings"
)

// MakeVarOrKeyword 获取变量或关键词
func MakeVarOrKeyword(it *common.PeekIterator) (*Token, error) {
	s, sBuilder := "", strings.Builder{}

	for it.HasNext() {
		lookahead := it.Peek()
		if !common.IsLiteral(lookahead) {
			s = sBuilder.String()
			break
		}
		sBuilder.WriteString(lookahead)
		it.Next()
	}

	if IsKeyword(s) {
		return NewToken(KEYWORD, s), nil
	}

	if s == "true" || s == "false" {
		return NewToken(BOOLEAN, s), nil
	}

	return NewToken(VARIABLE, s), nil
}

// MakeString 构造字符串
func MakeString(it *common.PeekIterator) (*Token, error) {
	sBuilder, state := strings.Builder{}, 0

	for it.HasNext() {
		c := it.Next()
		sBuilder.WriteString(c)
		switch state {
		case 0:
			switch c {
			case "\"":
				state = 1
			case "'":
				state = 2
			default:
				break
			}
		case 1:
			if c == "\"" {
				return NewToken(STRING, sBuilder.String()), nil
			}
		case 2:
			if c == "'" {
				return NewToken(STRING, sBuilder.String()), nil
			}
		}
	}
	return nil, NewLexicalError("unexpected error")
}
