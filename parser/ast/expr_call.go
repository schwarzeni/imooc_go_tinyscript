package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

// ExprCall 函数调用语句
type ExprCall struct {
	*ASTNode
}

// NewExprCall 新建一个函数调用语句
func NewExprCall() *ExprCall {
	return &ExprCall{NewASTNode(CALL_EXPTR, string(CALL_EXPTR), nil)}
}

// ParseExperCall 解析括号中的表达式
func ParseExperCall(factor Node, it *util.PeekTokenIterator) (expr Node, err error) {
	expr = NewExprCall()
	expr.AddChild(factor)

	if _, err := it.MatchNextValue("("); err != nil {
		return nil, err
	}

	var p Node = nil
	for {
		if p, err = ParseExpr(it); err != nil {
			return nil, err
		}
		if p == nil {
			break
		}
		expr.AddChild(p)
		if it.Peek().(*lexer.Token).GetValue() != ")" {
			if _, err = it.MatchNextValue(","); err != nil {
				return nil, err
			}
		}
	}

	if _, err := it.MatchNextValue(")"); err != nil {
		return nil, err
	}
	return
}
