package ast

import "go-tinyscript/lexer"

type Factor struct {
	*ASTNode
}

func NewFactor(token *lexer.Token) *Factor {
	return &Factor{NewASTNode(NONE, token.GetValue(), token)}
}
