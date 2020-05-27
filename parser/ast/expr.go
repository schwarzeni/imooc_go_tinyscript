package ast

import "go-tinyscript/lexer"

type Expr struct {
	*ASTNode
}

func NewExpr(t ASTNodeType, lexeme *lexer.Token) *Expr {
	return &Expr{NewASTNode(t, lexeme.GetValue(), lexeme)}
}
