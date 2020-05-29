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

func TestParseFunctionDeclare(t *testing.T) {
	eq := assert.Eq
	str := `
    func fn(int a, float bb, string ccc) int {
        var d = a + b +c
        return d
    }`
	node := createFunctionDeclare(str).(*StmtFunctionDeclare)
	eq(node.GetType(), FUNCTION_DECLARE_STMT)
	eq(node.GetLexeme().GetValue(), "fn")
	eq(node.GetArgs().GetChild(0).GetLexeme().GetValue(), "a")
	eq(node.GetArgs().GetChild(0).(*Variable).TypeLexeme().GetValue(), lexer.INTEGER)
	eq(node.GetArgs().GetChild(1).GetLexeme().GetValue(), "bb")
	eq(node.GetArgs().GetChild(1).(*Variable).TypeLexeme().GetValue(), lexer.FLOAT)
	eq(node.GetArgs().GetChild(2).GetLexeme().GetValue(), "ccc")
	eq(node.GetArgs().GetChild(2).(*Variable).TypeLexeme().GetValue(), lexer.STRING)
	eq(node.GetReturn().GetValue(), lexer.INTEGER)
	eq(ToPostfixExpression(node.GetBlock().GetChild(0)), "d a b c + + =")
	eq(node.GetBlock().GetChild(0).GetType(), DECLARE_STMT)
	eq(node.GetBlock().GetChild(1).GetChild(0).GetType(), VARIABLE)
	eq(node.GetBlock().GetChild(1).GetChild(0).GetLexeme().GetValue(), "d")
	eq(node.GetBlock().GetChild(1).GetChild(0).GetLexeme().GetType(), lexer.VARIABLE)
	eq(node.GetBlock().GetChild(1).GetType(), RETURN_STMT)

	str = `
    func fn(int a) int {
        return fn(a+1)
    }`
	node = createFunctionDeclare(str).(*StmtFunctionDeclare)
	eq(node.GetType(), FUNCTION_DECLARE_STMT)
	eq(node.GetLexeme().GetValue(), "fn")
	eq(node.GetArgs().GetChild(0).GetLexeme().GetValue(), "a")
	eq(node.GetArgs().GetChild(0).(*Variable).TypeLexeme().GetValue(), lexer.INTEGER)
	eq(node.GetBlock().GetChild(0).GetType(), RETURN_STMT)
	eq(node.GetBlock().GetChild(0).GetLexeme().GetType(), lexer.KEYWORD)
	eq(node.GetBlock().GetChild(0).GetChild(0).GetType(), CALL_EXPTR)
	call := node.GetBlock().GetChild(0).GetChild(0)
	eq(call.GetChild(0).GetType(), VARIABLE)
	eq(call.GetChild(0).GetLexeme().GetValue(), "fn") // variable
	eq(call.GetChild(1).GetType(), BINARY_EXPR)
	eq(ToPostfixExpression(call.GetChild(1)), "a 1 +")
}
func createFunctionDeclare(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseFunctionDeclare(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
