package symbol

import (
	"go-tinyscript/lexer"
	"go-tinyscript/utils/assert"
	"testing"
)

func TestSymbolTable(t *testing.T) {
	symbolTable := NewSymbolTable()
	symbolTable.CreateLabel("L0", lexer.NewToken(lexer.VARIABLE, "foo"))
	symbolTable.CreateVariable()
	symbolTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "foo"))
	assert.Eq(1, symbolTable.LocalSize())
}

func TestSymbolTableChain(t *testing.T) {
	symbolTable := NewSymbolTable()
	symbolTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "a"))

	childTable := NewSymbolTable()
	symbolTable.AddChild(childTable)

	childChildTable := NewSymbolTable()
	childTable.AddChild(childChildTable)

	assert.Eq(true, childChildTable.Exists(lexer.NewToken(lexer.VARIABLE, "a")))
	assert.Eq(true, childTable.Exists(lexer.NewToken(lexer.VARIABLE, "a")))
	assert.Eq(false, childChildTable.Exists(lexer.NewToken(lexer.VARIABLE, "c")))
}

func TestSymbolOffSet(t *testing.T) {
	symbolTable := NewSymbolTable()

	symbolTable.CreateSymbolByLexeme(lexer.NewToken(lexer.INTEGER, "100"))

	symbolA := symbolTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "a"))
	symbolB := symbolTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "b"))

	childTable := NewSymbolTable()
	symbolTable.AddChild(childTable)
	anotherSymbolB := childTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "b"))
	symbolC := childTable.CreateSymbolByLexeme(lexer.NewToken(lexer.VARIABLE, "c"))

	assert.Eq(0, symbolA.Offset())
	assert.Eq(1, symbolB.Offset())
	assert.Eq(1, anotherSymbolB.Offset())
	assert.Eq(1, anotherSymbolB.LayerOffset())
	assert.Eq(0, symbolC.Offset())
	assert.Eq(0, symbolC.LayerOffset())
}
