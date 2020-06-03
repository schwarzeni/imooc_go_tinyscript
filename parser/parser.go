package parser

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/ast"
	"go-tinyscript/parser/util"
	"strings"
)

// Parse 解析字符串输入
func Parse(src string) (ast.Node, error) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		return nil, err
	}
	return ast.ParseProgram(util.NewPeekTokenIterator(tokens))
}
