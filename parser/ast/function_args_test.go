package ast

import (
	"fmt"
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
	"go-tinyscript/utils/assert"
	"os"
	"strings"
	"testing"
)

func TestParseFunctionArgs(t *testing.T) {
	eq := assert.Eq
	args := createFunctionArgs("int a, float b, string c)")
	eq(args.GetChild(0).GetLexeme().GetValue(), "a")
	eq(args.GetChild(0).(*Variable).TypeLexeme().GetValue(), "int")
	eq(args.GetChild(0).(*Variable).TypeLexeme().GetType(), lexer.KEYWORD)
	eq(args.GetChild(1).GetLexeme().GetValue(), "b")
	eq(args.GetChild(1).(*Variable).TypeLexeme().GetValue(), "float")
	eq(args.GetChild(1).(*Variable).TypeLexeme().GetType(), lexer.KEYWORD)
	eq(args.GetChild(2).GetLexeme().GetValue(), "c")
	eq(args.GetChild(2).(*Variable).TypeLexeme().GetValue(), "string")
	eq(args.GetChild(2).(*Variable).TypeLexeme().GetType(), lexer.KEYWORD)
}

func createFunctionArgs(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseFunctionArgs(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
