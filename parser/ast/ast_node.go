package ast

import (
	"go-tinyscript/lexer"
	print2 "go-tinyscript/utils/print"
	"reflect"
	"strings"
)

var Println = print2.Println

type ASTNode struct {
	// 树
	children []Node
	parent   Node

	// 节点信息
	lexeme   *lexer.Token // 词法单元
	label    string       // 备注（标签）
	nodeType ASTNodeType  // 类型
	props    map[string]interface{}
}

func (A *ASTNode) GetChild(idx int) Node {
	if idx >= len(A.children) || idx < 0 {
		return nil
	}
	return A.children[idx]
}

func (A *ASTNode) AddChild(node Node) {
	node.SetParent(A)
	A.children = append(A.children, node)
}

func (A *ASTNode) GetParent() Node {
	return A.parent
}

func (A *ASTNode) SetParent(parent Node) {
	A.parent = parent
}

func (A *ASTNode) GetLexeme() *lexer.Token {
	return A.lexeme
}

func (A *ASTNode) SetLexeme(token *lexer.Token) {
	A.lexeme = token
}

func (A *ASTNode) GetChildren() []Node {
	return A.children
}

func (A *ASTNode) GetChildIdx(node Node) int {
	for idx, child := range A.children {
		if reflect.DeepEqual(child, node) {
			return idx
		}
	}
	return -1
}

func (A *ASTNode) SetLabel(label string) {
	A.label = label
}

func (A *ASTNode) GetLabel() string {
	return A.label
}

func (A *ASTNode) GetType() ASTNodeType {
	return A.nodeType
}

func (A *ASTNode) SetType(t ASTNodeType) {
	A.nodeType = t
}

func (A *ASTNode) Print(indent int) {
	if indent == 0 {
		Println("print:", A)
	}
	Println(strings.Repeat(" ", indent*2) + A.label)
	for _, child := range A.children {
		child.Print(indent + 1)
	}
}

func (A *ASTNode) ReplaceChild(idx int, node Node) {
	if idx >= len(A.children) || idx < 0 {
		return
	}
	A.children[idx] = node
}

func (A *ASTNode) Replace(node Node) {
	if A.parent != nil {
		idx := A.parent.GetChildIdx(node)
		A.parent.ReplaceChild(idx, node)
	}
}

func (A *ASTNode) GetProp(key string) (val interface{}, ok bool) {
	val, ok = A.props[key]
	return
}

func (A *ASTNode) SetProp(key string, val interface{}) {
	A.props[key] = val
}

func (A *ASTNode) GetProps() map[string]interface{} {
	return A.props
}

func (A *ASTNode) IsValueType() bool {
	return A.nodeType == VARIABLE || A.nodeType == SCALAR
}

func NewASTNode(t ASTNodeType, label string, lexme *lexer.Token) *ASTNode {
	return &ASTNode{
		props:    make(map[string]interface{}),
		nodeType: t,
		label:    label,
		lexeme:   lexme,
	}
}
