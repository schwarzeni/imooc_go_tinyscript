package ast

type Block struct {
	*ASTNode
}

func NewBlock() *Block {
	return &Block{NewASTNode(BLOCK, string(BLOCK), nil)}
}
