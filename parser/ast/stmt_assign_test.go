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

func TestParseAssign(t *testing.T) {
	eq := assert.Eq
	eq(ToPostfixExpression(createAssign(`a = 2 +3 + b`)), "a 2 3 b + + =")
	eq(ToPostfixExpression(createAssign(`i = 100 * 2`)), "i 100 2 * =")
}

func createAssign(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseAssign(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
