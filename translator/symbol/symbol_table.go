package symbol

import (
	"fmt"
	"go-tinyscript/lexer"
)

// SymbolTable 符号表
type SymbolTable struct {
	parent      *SymbolTable
	children    []*SymbolTable
	symbols     []*Symbol
	tempIndex   int // 临时变量的编号，p0、p1、p2 ...
	offsetIndex int // 内存地址的关系
	level       int // 当前节点的高度距离根节点的距离
}

// AddSymbol 添加 Symbol 至 SymbolTable 中
func (symbolTable *SymbolTable) AddSymbol(symbol *Symbol) {
	symbolTable.symbols = append(symbolTable.symbols, symbol)
	symbol.SetParent(symbolTable)
}

// CloneFromSymbolTree 从 SymbolTree 中克隆 lexeme 对应的 Symbol
func (symbolTable *SymbolTable) CloneFromSymbolTree(lexeme *lexer.Token, layerOffset int) *Symbol {
	for _, v := range symbolTable.symbols {
		if v.lexeme.GetValue() == lexeme.GetValue() {
			symbol := v.Copy()
			symbol.SetLayerOffset(layerOffset)
			return symbol
		}
	}
	if symbolTable.parent != nil {
		return symbolTable.parent.CloneFromSymbolTree(lexeme, layerOffset+1)
	}
	return nil
}

// Exists 查找 lexeme 对应的 Token 是否存在
func (symbolTable *SymbolTable) Exists(lexeme *lexer.Token) bool {
	for _, v := range symbolTable.symbols {
		if v.lexeme.GetValue() == lexeme.GetValue() {
			return true
		}
	}
	if symbolTable.parent != nil {
		return symbolTable.parent.Exists(lexeme)
	}
	return false
}

// CreateSymbolByLexeme 构建 Symbol
func (symbolTable *SymbolTable) CreateSymbolByLexeme(lexeme *lexer.Token) (symbol *Symbol) {
	if lexeme.IsScalar() { // 常量，直接创建
		symbol = NewImmediateSymbol(lexeme)
		symbolTable.AddSymbol(symbol)
		return
	}
	for _, v := range symbolTable.symbols {
		if v.lexeme.GetValue() == lexeme.GetValue() { // 已经存在，则直接返回
			return v
		}
	}
	// 验证其是否在父节点以及以上出现过
	if symbol = symbolTable.CloneFromSymbolTree(lexeme, 0); symbol == nil {
		symbol = NewAddressSymbol(lexeme, symbolTable.offsetIndex)
		symbolTable.offsetIndex++
	}
	symbolTable.AddSymbol(symbol)
	return
}

// CreateVariable 创建变量
func (symbolTable *SymbolTable) CreateVariable() (symbol *Symbol) {
	lexeme := lexer.NewToken(lexer.VARIABLE, fmt.Sprintf("p%d", symbolTable.tempIndex))
	symbolTable.tempIndex++
	symbol = NewAddressSymbol(lexeme, symbolTable.offsetIndex)
	symbolTable.offsetIndex++
	symbolTable.AddSymbol(symbol)
	return
}

// AddChild 添加子 SymbolTable 节点
func (symbolTable *SymbolTable) AddChild(child *SymbolTable) {
	child.parent = symbolTable
	child.level = symbolTable.level + 1
	symbolTable.children = append(symbolTable.children, child)
}

// LocalSize 获取本地数据所占大小
func (symbolTable *SymbolTable) LocalSize() int {
	return symbolTable.offsetIndex
}

func (symbolTable *SymbolTable) Children() []*SymbolTable {
	return symbolTable.children
}

func (symbolTable *SymbolTable) Symbols() []*Symbol {
	return symbolTable.symbols
}

func (symbolTable *SymbolTable) CreateLabel(label string, lexeme *lexer.Token) {
	labelSymbol := NewLabelSymbol(lexeme, label)
	symbolTable.AddSymbol(labelSymbol)
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{}
}
