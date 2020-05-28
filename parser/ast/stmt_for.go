package ast

type StmtFor struct {
	*Stmt
}

func NewStmtFor() *StmtFor {
	return &StmtFor{NewStmt(FOR_STMT, string(FOR_STMT))}
}
