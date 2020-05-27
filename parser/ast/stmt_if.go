package ast

type StmtIf struct {
	*ASTNode
}

func NewStmtIf() *StmtIf {
	return &StmtIf{NewASTNode(IF_STMT, string(IF_STMT), nil)}
}
