package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

// Stmt 语句
type Stmt struct {
	*ASTNode
}

// NewStmt 新建语句
func NewStmt(t ASTNodeType, label string) *Stmt {
	return &Stmt{NewASTNode(t, label, nil)}
}

// ParseStmt 解析语句
func ParseStmt(it *util.PeekTokenIterator) (Node, error) {
	if !it.HasNext() {
		return nil, nil
	}
	token := it.Next().(*lexer.Token)
	var lookahead *lexer.Token
	if it.HasNext() {
		lookahead = it.Peek().(*lexer.Token)
	}
	it.PutBack()

	if token.IsVariable() && lookahead != nil && lookahead.GetValue() == "=" {
		return ParseAssign(it)
	}
	if token.GetValue() == "var" {
		return ParseDeclare(it)
	}
	if token.GetValue() == "{" {
		return ParseBlock(it)
	}
	if token.GetValue() == "if" {
		return ParseIf(it)
	}
	if token.GetValue() == "func" {
		return ParseFunctionDeclare(it)
	}
	if token.GetValue() == "return" {
		return ParseReturn(it)
	}
	return ParseExpr(it)

}
