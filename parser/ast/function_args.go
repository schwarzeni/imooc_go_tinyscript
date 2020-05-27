package ast

type FunctionArgs struct {
	*ASTNode
}

func NewFunctionArgs() *FunctionArgs {
	return &FunctionArgs{NewASTNode(FUNCTION_ARGS, string(FUNCTION_ARGS), nil)}
}
