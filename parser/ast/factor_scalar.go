package ast

import "go-tinyscript/lexer"

type Scalar struct {
	*Factor
}

func NewScalar(token *lexer.Token) *Scalar {
	s := &Scalar{NewFactor(token)}
	s.nodeType = SCALAR
	return s
}
