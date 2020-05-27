package ast

type StmtFunctionDeclare struct {
	*ASTNode
}

func NewStmtFunctionDeclare() *StmtFunctionDeclare {
	return &StmtFunctionDeclare{NewASTNode(FUNCTION_DECLARE_STMT, string(FUNCTION_DECLARE_STMT), nil)}
}
