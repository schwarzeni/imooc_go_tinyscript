package ast

import "go-tinyscript/parser/util"

// Block { ... }
type Block struct {
	*ASTNode
}

// NewBlock 新建一个 Block
func NewBlock() *Block {
	return &Block{NewASTNode(BLOCK, string(BLOCK), nil)}
}

// ParseBlock 解析 { ... }
func ParseBlock(it *util.PeekTokenIterator) (Node, error) {
	if _, err := it.MatchNextValue("{"); err != nil {
		return nil, err
	}
	block := NewBlock()
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
	if _, err := it.MatchNextValue("}"); err != nil {
		return nil, err
	}
	return block, nil
}
