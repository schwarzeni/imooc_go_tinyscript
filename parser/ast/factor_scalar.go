package ast

type Scalar struct {
	*ASTNode
}

func NewScalar() *Scalar {
	return &Scalar{NewASTNode(SCALAR, string(SCALAR), nil)}
}
