package translator

import (
	"go-tinyscript/parser"
	"go-tinyscript/translator/symbol"
	"reflect"
	"testing"
)

func assertOpcodes(expected []string, got []*TAInstruction, t *testing.T) {
	for idx, v := range got {
		if !reflect.DeepEqual(expected[idx], v.String()) {
			goto ERR
		}
	}
	return
ERR:
	t.Errorf("expected %v\n, but got %v\n", expected, got)
}

func TestTransExpr(t *testing.T) {
	source := "a+(b-c)+d*(b-c)*2"
	p, err := parser.Parse(source)
	if err != nil {
		panic(err)
	}
	exprNode := p.GetChild(0)

	translator := NewTranslator()
	symbolTable := symbol.NewSymbolTable()
	program := NewTAProgram()
	if _, err := translator.translateExpr(program, exprNode, symbolTable); err != nil {
		panic(err)
	}
	expectedResults := []string{
		"p0 = b - c",
		"p1 = b - c",
		"p2 = p1 * 2",
		"p3 = d * p2",
		"p4 = p0 + p3",
		"p5 = a + p4",
	}
	assertOpcodes(expectedResults, program.Instructions(), t)
}

func TestTransAssign(t *testing.T) {
	source := "a=1.0*2.0*3.0"
	p, err := parser.Parse(source)
	if err != nil {
		panic(err)
	}
	translator := NewTranslator()
	program, err := translator.Translate(p)
	if err != nil {
		panic(err)
	}
	expected := []string{
		"p0 = 2.0 * 3.0",
		"p1 = 1.0 * p0",
		"a = p1",
	}
	assertOpcodes(expected, program.Instructions(), t)
}

func TestTransDeclare(t *testing.T) {
	source := "var a=1.0*2.0*3.0"
	p, err := parser.Parse(source)
	if err != nil {
		panic(err)
	}
	translator := NewTranslator()
	program, err := translator.Translate(p)
	if err != nil {
		panic(err)
	}
	expected := []string{
		"p0 = 2.0 * 3.0",
		"p1 = 1.0 * p0",
		"a = p1",
	}
	assertOpcodes(expected, program.Instructions(), t)
}
