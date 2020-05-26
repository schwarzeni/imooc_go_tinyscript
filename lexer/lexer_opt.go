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
