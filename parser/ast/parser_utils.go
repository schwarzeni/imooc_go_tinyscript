package ast

import (
	"strings"
)

// ToPostfixExpression 生成后缀表达式字符串
func ToPostfixExpression(n Node) string {
	if n.GetType() == SCALAR || n.GetType() == VARIABLE {
		return n.GetLexeme().GetValue()
	}

	prts := []string{}
	for _, child := range n.GetChildren() {
		prts = append(prts, ToPostfixExpression(child))
	}
	lexemeStr := ""
	if n.GetLexeme() != nil {
		lexemeStr = n.GetLexeme().GetValue()
	}
	if len(lexemeStr) > 0 {
		return strings.Join(prts, " ") + " " + lexemeStr
	}
	return strings.Join(prts, " ")
}
