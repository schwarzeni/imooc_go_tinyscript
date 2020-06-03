package ast

import "go-tinyscript/parser/util"

type Program struct {
	*Block
}

func NewProgram() *Program {
	return &Program{NewBlock()}
}

// ParseProgram 解析程序
func ParseProgram(it *util.PeekTokenIterator) (Node, error) {
	block := NewProgram()
	for {
		stmt, err := ParseStmt(it)
		if err != nil {
			return nil, err
		}
		if stmt == nil {
			break
		}
		block.AddChild(stmt)
	}
	return block, nil
}
