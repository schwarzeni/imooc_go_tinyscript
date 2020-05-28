package ast

type StmtIf struct {
	*Stmt
}

func NewStmtIf() *StmtIf {
	return &StmtIf{NewStmt(IF_STMT, string(IF_STMT))}
}
