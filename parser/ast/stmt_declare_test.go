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

func TestParseDeclare(t *testing.T) {
	eq := assert.Eq
	eq(ToPostfixExpression(createDeclare(`var a = 2 +3 + b`)), "a 2 3 b + + =")
	eq(ToPostfixExpression(createDeclare(`var i = 100 * 2`)), "i 100 2 * =")
}

func createDeclare(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseDeclare(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
