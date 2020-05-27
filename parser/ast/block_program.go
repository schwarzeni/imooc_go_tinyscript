package ast

type Program struct {
	*Block
}

func NewProgram() *Program {
	return &Program{NewBlock()}
}
