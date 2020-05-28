package ast

type StmtFunctionDeclare struct {
	*Stmt
}

func NewStmtFunctionDeclare() *StmtFunctionDeclare {
	return &StmtFunctionDeclare{NewStmt(FUNCTION_DECLARE_STMT, string(FUNCTION_DECLARE_STMT))}
}
