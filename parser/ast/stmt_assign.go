package ast

type StmtAssign struct {
	*Stmt
}

func NewStmtAssign() *StmtAssign {
	return &StmtAssign{NewStmt(ASSIGN_STMT, string(ASSIGN_STMT))}
}
