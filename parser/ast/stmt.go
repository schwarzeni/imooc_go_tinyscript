package ast

type Stmt struct {
	*ASTNode
}

func NewStmt(t ASTNodeType, label string) *Stmt {
	return &Stmt{NewASTNode(t, label, nil)}
}
