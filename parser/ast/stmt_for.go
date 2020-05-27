package ast

type StmtFor struct {
	*ASTNode
}

func NewStmtFor() *StmtFor {
	return &StmtFor{NewASTNode(FOR_STMT, string(FOR_STMT), nil)}
}
