package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

// FunctionArgs 函数参数列表
type FunctionArgs struct {
	*ASTNode
}

// NewFunctionArgs 新建函数参数列表
func NewFunctionArgs() *FunctionArgs {
	return &FunctionArgs{NewASTNode(FUNCTION_ARGS, string(FUNCTION_ARGS), nil)}
}

// ParseFunctionArgs 解析函数参数列列表
func ParseFunctionArgs(it *util.PeekTokenIterator) (Node, error) {
	args := NewFunctionArgs()

	for {
		t := it.Next().(*lexer.Token)
		v, err := ParseFactor(it)
		if err != nil {
			return nil, err
		}
		variable := v.(*Variable)
		variable.SetTypeLexeme(t)
		args.AddChild(variable)
		if !it.HasNext() || it.Peek().(*lexer.Token).GetValue() == ")" {
			break
		}
		if _, err := it.MatchNextValue(","); err != nil {
			return nil, err
		}
	}

	return args, nil
}
