package translator

import (
	"fmt"
	"go-tinyscript/translator/symbol"
	"strings"
)

// TAProgram 三地址程序
type TAProgram struct {
	instructions      []*TAInstruction
	labelCounter      int
	staticSymbolTable *symbol.StaticSymbolTable
}

func NewTAProgram() *TAProgram {
	return &TAProgram{}
}

func (p *TAProgram) SetStaticSymbolTable(staticSymbolTable *symbol.StaticSymbolTable) {
	p.staticSymbolTable = staticSymbolTable
}

func (p *TAProgram) Add(code *TAInstruction) {
	p.instructions = append(p.instructions, code)
}

func (p *TAProgram) Instructions() []*TAInstruction {
	return p.instructions
}

// AddLabel 添加 Label
func (p *TAProgram) AddLabel() *TAInstruction {
	label := fmt.Sprintf("L%d", p.labelCounter)
	p.labelCounter++
	taCode := NewTAInstruction(LABEL, nil, "", nil, "")
	taCode.SetArg1(label)
	p.instructions = append(p.instructions, taCode)
	return taCode
}

func (p *TAProgram) SetStaticSymbols(table *symbol.SymbolTable) {
	for _, s := range table.Symbols() {
		if s.Type() == symbol.IMMEDIATE_SYMBOL {
			p.staticSymbolTable.Add(s)
		}
	}
	for _, c := range table.Children() {
		p.SetStaticSymbols(c)
	}
}

func (p *TAProgram) StaticSymbolTable() *symbol.StaticSymbolTable {
	return p.staticSymbolTable
}

func (p *TAProgram) String() string {
	var lines []string
	for _, v := range p.Instructions() {
		lines = append(lines, v.String())
	}
	return strings.Join(lines, "\n")
}
