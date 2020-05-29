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

func TestParseIf(t *testing.T) {
	eq := assert.Eq
	str := `
if (a ==3) {
    var b = 2 + 33
} else if (a != 6) {
    c = c +1
} else if (a >= 4) {
    d = d * 2 + 2
} else {
    var ff = 2 +3
    var ged = ff * 3
}`
	// if (a ==3) {
	//     var b = 2 + 33
	// } else if ...
	node := createIf(str).(*StmtIf)
	eq(node.GetType(), IF_STMT)
	eq(node.GetElseIfStmt().GetType(), IF_STMT)
	eq(node.GetExpr().GetLexeme().GetValue(), "==")
	eq(node.GetExpr().GetChild(0).GetLexeme().GetValue(), "a")
	eq(node.GetExpr().GetChild(1).GetLexeme().GetValue(), "3")
	eq(node.GetBlock().GetChild(0).GetType(), DECLARE_STMT)
	eq(ToPostfixExpression(node.GetBlock()), "b 2 33 + =")

	// } else if (a != 6) {
	//      c = c +1
	// } else if ...
	node = node.GetElseIfStmt().(*StmtIf)
	eq(node.GetType(), IF_STMT)
	eq(node.GetElseIfStmt().GetType(), IF_STMT)
	eq(node.GetExpr().GetLexeme().GetValue(), "!=")
	eq(node.GetExpr().GetChild(0).GetLexeme().GetValue(), "a")
	eq(node.GetExpr().GetChild(1).GetLexeme().GetValue(), "6")
	eq(node.GetBlock().GetChild(0).GetType(), ASSIGN_STMT)
	eq(ToPostfixExpression(node.GetBlock()), "c c 1 + =")

	// } else if (a >= 4) {
	//    d = d * 2 + 2
	// } else { ....
	node = node.GetElseIfStmt().(*StmtIf)
	eq(node.GetType(), IF_STMT)
	eq(node.GetElseStmt().GetType(), BLOCK)
	eq(node.GetExpr().GetLexeme().GetValue(), ">=")
	eq(node.GetExpr().GetChild(0).GetLexeme().GetValue(), "a")
	eq(node.GetExpr().GetChild(1).GetLexeme().GetValue(), "4")
	eq(node.GetBlock().GetChild(0).GetType(), ASSIGN_STMT)
	eq(ToPostfixExpression(node.GetBlock()), "d d 2 * 2 + =")

	// } else {
	//    var ff = 2 +3
	//    var ged = ff * 3
	// }
	block := node.GetElseStmt().(*Block)
	eq(block.GetType(), BLOCK)
	eq(ToPostfixExpression(block.GetChild(0)), "ff 2 3 + =")
	eq(block.GetChild(0).GetType(), DECLARE_STMT)
	eq(ToPostfixExpression(block.GetChild(1)), "ged ff 3 * =")
	eq(block.GetChild(1).GetType(), DECLARE_STMT)
}

func createIf(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseIf(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
