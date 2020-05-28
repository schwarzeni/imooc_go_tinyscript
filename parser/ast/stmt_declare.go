package ast

type StmtDeclare struct {
	*Stmt
}

func NewStmtDeclare() *StmtDeclare {
	return &StmtDeclare{NewStmt(DECLARE_STMT, string(DECLARE_STMT))}
}
