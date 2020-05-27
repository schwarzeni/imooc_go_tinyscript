package parser

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/ast"
	"go-tinyscript/parser/util"
	"go-tinyscript/utils/assert"
	"strings"
	"testing"
)

func TestSimpleParser(t *testing.T) {
	eq := assert.Eq
	var err error
	source := "1+2+3+4"
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(source))
	if err != nil {
		panic(err)
	}

	expr, err := SimpleParser(util.NewPeekTokenIterator(tokens))
	if err != nil {
		panic(err)
	}
	eq(2, len(expr.GetChildren()))

	v1 := expr.GetChild(0).(*ast.Scalar)
	eq("1", v1.GetLexeme().GetValue())
	eq("+", expr.GetLexeme().GetValue())

	e2 := expr.GetChild(1).(*ast.Expr)
	v2 := e2.GetChild(0).(*ast.Scalar)
	eq("2", v2.GetLexeme().GetValue())
	eq("+", e2.GetLexeme().GetValue())

	e3 := e2.GetChild(1).(*ast.Expr)
	v3 := e3.GetChild(0).(*ast.Scalar)
	v4 := e3.GetChild(1).(*ast.Scalar)
	eq("3", v3.GetLexeme().GetValue())
	eq("+", e3.GetLexeme().GetValue())
	eq("4", v4.GetLexeme().GetValue())
}
