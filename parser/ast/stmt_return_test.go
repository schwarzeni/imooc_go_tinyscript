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

func TestParseReturn(t *testing.T) {
	eq := assert.Eq
	eq(createReturn("return 2").GetChild(0).GetLexeme().GetValue(), "2")
	eq(createReturn("return a").GetChild(0).GetLexeme().GetValue(), "a")
}

func createReturn(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseReturn(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
