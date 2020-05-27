package ast

type ExprCall struct {
	*ASTNode
}

func NewExprCall() *ExprCall {
	return &ExprCall{NewASTNode(CALL_EXPTR, string(CALL_EXPTR), nil)}
}
