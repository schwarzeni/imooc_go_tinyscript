package ast

type StmtDeclare struct {
	*ASTNode
}

func NewStmtDeclare() *StmtDeclare {
	return &StmtDeclare{NewASTNode(DECLARE_STMT, string(DECLARE_STMT), nil)}
}
