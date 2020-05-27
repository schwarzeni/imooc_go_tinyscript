package ast

import "go-tinyscript/lexer"

// Node 定义抽象语法树节点的接口
type Node interface {
	GetChild(index int) Node
	AddChild(node Node)
	GetParent() Node
	SetParent(parent Node)
	GetLexeme() *lexer.Token
	SetLexeme(token *lexer.Token)
	GetChildren() []Node
	GetChildIdx(node Node) int
	SetLabel(label string)
	GetLabel() string
	GetType() ASTNodeType
	SetType(t ASTNodeType)
	Print(indent int)
	ReplaceChild(idx int, node Node)
	Replace(node Node) // 替换当前节点
	GetProp(key string) (interface{}, bool)
	SetProp(key string, val interface{})
	GetProps() map[string]interface{}
	IsValueType() bool
}
