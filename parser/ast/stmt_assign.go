package ast

type StmtAssign struct {
	*ASTNode
}

func NewStmtAssign() *StmtAssign {
	return &StmtAssign{NewASTNode(ASSIGN_STMT, string(ASSIGN_STMT), nil)}
}
