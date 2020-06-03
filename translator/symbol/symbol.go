package symbol

import "go-tinyscript/lexer"

// Symbol 符号，表示值或变量
type Symbol struct {
	parent *SymbolTable
	lexeme *lexer.Token

	// 不同的 symbol 类型. union
	t           SymbolType
	label       string // LABEL_SYMBOL label 的名称
	offset      int    // ADDRESS_SYMBOL
	layerOffset int    // 层级关系  (var a = 2) 3 { 2 { 1 {(a) 0 }}}
}

// Copy 复制当前的符号
func (symbol *Symbol) Copy() *Symbol {
	newS := newSymbol(symbol.t)
	newS.lexeme = symbol.lexeme
	newS.label = symbol.label
	newS.offset = symbol.offset
	newS.layerOffset = symbol.layerOffset
	newS.t = symbol.t
	return newS
}

func newSymbol(t SymbolType) *Symbol {
	return &Symbol{t: t}
}

// NewAddressSymbol 新建地址类型符号
func NewAddressSymbol(lexeme *lexer.Token, offset int) *Symbol {
	addressSymbol := newSymbol(ADDRESS_SYMBOL)
	addressSymbol.lexeme = lexeme
	addressSymbol.offset = offset
	return addressSymbol
}

// NewImmediateSymbol 新建立即数类型符号
func NewImmediateSymbol(lexeme *lexer.Token) *Symbol {
	immediateSymbol := newSymbol(IMMEDIATE_SYMBOL)
	immediateSymbol.lexeme = lexeme
	return immediateSymbol
}

// NewLabelSymbol 新建标签类型符号
func NewLabelSymbol(lexeme *lexer.Token, label string) *Symbol {
	labelSymbol := newSymbol(LABEL_SYMBOL)
	labelSymbol.label = label
	labelSymbol.lexeme = lexeme
	return labelSymbol
}

func (symbol *Symbol) Label() string {
	return symbol.label
}

func (symbol *Symbol) SetLabel(label string) {
	symbol.label = label
}

func (symbol *Symbol) LayerOffset() int {
	return symbol.layerOffset
}

func (symbol *Symbol) SetLayerOffset(layerOffset int) {
	symbol.layerOffset = layerOffset
}

func (symbol *Symbol) Type() SymbolType {
	return symbol.t
}

func (symbol *Symbol) SetType(t SymbolType) {
	symbol.t = t
}

func (symbol *Symbol) Parent() *SymbolTable {
	return symbol.parent
}

func (symbol *Symbol) SetParent(parent *SymbolTable) {
	symbol.parent = parent
}

func (symbol *Symbol) Offset() int {
	return symbol.offset
}

func (symbol *Symbol) SetOffset(offset int) {
	symbol.offset = offset
}

func (symbol *Symbol) Lexeme() *lexer.Token {
	return symbol.lexeme
}

func (symbol *Symbol) SetLexeme(lexeme *lexer.Token) {
	symbol.lexeme = lexeme
}

func (symbol *Symbol) String() string {
	if symbol.t == LABEL_SYMBOL {
		return symbol.Label()
	}
	return symbol.lexeme.GetValue()
}
