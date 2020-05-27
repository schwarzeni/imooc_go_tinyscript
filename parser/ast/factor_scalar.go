package ast

import "go-tinyscript/lexer"

type Scalar struct {
	*Factor
}

func NewScalar(token *lexer.Token) *Scalar {
	return &Scalar{NewFactor(token)}
}
