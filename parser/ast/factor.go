package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

type Factor struct {
	*ASTNode
}

func NewFactor(token *lexer.Token) *Factor {
	return &Factor{NewASTNode(NONE, token.GetValue(), token)}
}

// 解析 Factor 表达式
func ParseFactor(it *util.PeekTokenIterator) (Node, error) {
	token := it.Peek().(*lexer.Token)
	t := token.GetType()

	if t == lexer.VARIABLE {
		it.Next()
		return NewVariable(token), nil
	} else if token.IsScalar() {
		it.Next()
		return NewScalar(token), nil
	}
	return nil, util.NewParseError(token)
}
