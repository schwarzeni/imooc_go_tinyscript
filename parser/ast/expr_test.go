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

func TestParseExperCall(t *testing.T) {
	eq := assert.Eq
	eq(ToPostfixExpression(createExpr("10 * (7 + 4)")), "10 7 4 + *")
	eq(ToPostfixExpression(createExpr("1+2*3")), "1 2 3 * +")
	eq(ToPostfixExpression(createExpr("1*2+3")), "1 2 * 3 +")
	eq(ToPostfixExpression(createExpr("(1*2!=7)==3!=4*5+66")), "1 2 * 7 != 3 4 5 * 66 + != ==")

	eq(ToPostfixExpression(createExpr(`"1" == ""`)), `"1" "" ==`)

	eq(ToPostfixExpression(createExpr(`print(a,b,c)`)), `print a b c`)
	eq(ToPostfixExpression(createExpr(`print(a)`)), `print a`)

	eq(ToPostfixExpression(createExpr(`!(a+b+c)`)), `a b c + + !`)
	eq(ToPostfixExpression(createExpr(`1 + ++i`)), `1 i ++ +`)
	//eq(ToPostfixExpression(createExpr(`1 + (i++)`)), `1 i + ++`)
}

func createExpr(src string) (node Node) {
	l := lexer.Lexer{}
	tokens, err := l.Analyze(strings.NewReader(src))
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	if node, err = ParseExpr(util.NewPeekTokenIterator(tokens)); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(-1)
	}
	return
}
