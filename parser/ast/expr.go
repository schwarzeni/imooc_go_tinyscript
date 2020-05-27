package ast

import "go-tinyscript/lexer"

type Expr struct {
	*ASTNode
}

func NewExpr(t ASTNodeType, lexeme *lexer.Token) *Expr {
	if lexeme == nil {
		return &Expr{&ASTNode{}}
	}
	return &Expr{NewASTNode(t, lexeme.GetValue(), lexeme)}
}
