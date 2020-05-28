package parser

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/ast"
	"go-tinyscript/parser/util"
)

// SimpleParser 简答的 parser
// Expr -> digit + Expr | digit
// digit -> 0|1|2|3|4|5|...|9
func SimpleParser(it *util.PeekTokenIterator) (ast.Node, error) {
	expr := ast.NewExpr(ast.NONE, nil)
	scalar, err := ast.ParseFactor(it)
	if err != nil {
		return nil, err
	}
	if !it.HasNext() {
		return scalar, nil
	}
	expr.SetLexeme(it.Peek().(*lexer.Token))
	if _, err := it.MatchNextValue("+"); err != nil {
		return nil, err
	}
	expr.SetLabel("+")
	expr.AddChild(scalar)
	expr.SetType(ast.BINARY_EXPR)
	rightNode, err := SimpleParser(it)
	if err != nil {
		return nil, err
	}
	expr.AddChild(rightNode)
	return expr, nil
}
