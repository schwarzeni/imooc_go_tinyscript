package symbol

import (
	"fmt"
	"strings"
)

// StaticSymbolTable 静态符号表
type StaticSymbolTable struct {
	offsetMap     map[string]*Symbol
	offsetCounter int
	symbols       []*Symbol
}

// Add 添加 symbol
func (staticSymbolTable *StaticSymbolTable) Add(symbol *Symbol) {
	lexval := symbol.Lexeme().GetValue()
	if v, ok := staticSymbolTable.offsetMap[lexval]; ok {
		symbol.SetOffset(v.Offset())
		return
	}
	staticSymbolTable.offsetMap[lexval] = symbol
	symbol.SetOffset(staticSymbolTable.offsetCounter)
	staticSymbolTable.offsetCounter++
	staticSymbolTable.symbols = append(staticSymbolTable.symbols, symbol)
}

// Size 获取 StaticSymbolTable 的大小
func (staticSymbolTable *StaticSymbolTable) Size() int {
	return len(staticSymbolTable.symbols)
}

func (staticSymbolTable *StaticSymbolTable) String() string {
	var list []string
	for idx, v := range staticSymbolTable.symbols {
		list = append(list, fmt.Sprintf("%d: %s", idx, v))
	}
	return strings.Join(list, "\n")
}

func (staticSymbolTable *StaticSymbolTable) Symbols() []*Symbol {
	return staticSymbolTable.symbols
}

func NewStaticSymbolTable() *StaticSymbolTable {
	return &StaticSymbolTable{offsetMap: make(map[string]*Symbol)}
}
