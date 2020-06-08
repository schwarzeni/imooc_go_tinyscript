package translator

import (
	"go-tinyscript/parser"
	"go-tinyscript/translator/symbol"
	"go-tinyscript/utils/assert"
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

func getProgram(source string) *TAProgram {
	p, err := parser.Parse(source)
	if err != nil {
		panic(err)
	}
	translator := NewTranslator()
	program, err := translator.Translate(p)
	if err != nil {
		panic(err)
	}
	return program
}

func TestTransBlock(t *testing.T) {
	source := `
var a = 1
{
    var b = a * 100
}
{
    var b = a * 10
    {
        var c = b * a + 10
    }
}
`

	assert.Eq(`a = 1
p1 = a * 100
b = p1
p1 = a * 10
b = p1
p1 = b * a
p2 = p1 + 10
c = p2`, getProgram(source).String())
}

func TestTransIf(t *testing.T) {
	source := `
if (a > 2) {
    var c = 2
}
`
	expected := `p0 = a > 2
IF p0 ELSE L0
c = 2
L0:`

	assert.Eq(expected, getProgram(source).String())
}

func TestTransIfElse(t *testing.T) {
	source := `
if ( a > b +2 ) {
    var c = 3
} else {
    d = 4
}
`
	expected := `p0 = b + 2
p1 = a > p0
IF p1 ELSE L0
c = 3
GOTO L1
L0:
d = 4
L1:`
	assert.Eq(expected, getProgram(source).String())
}

func TestTransIfElseIf(t *testing.T) {
	source := `
if(a == 1) {
  b = 100
} else if(a == 2) {
  b = 500
} else if(a == 3) {
  b = a * 1000
} else {
  b = -1
}`
	expected := `p0 = a == 1
IF p0 ELSE L0
b = 100
GOTO L5
L0:
p1 = a == 2
IF p1 ELSE L1
b = 500
GOTO L4
L1:
p2 = a == 3
IF p2 ELSE L2
p1 = a * 1000
b = p1
GOTO L3
L2:
b = -1
L3:
L4:
L5:`
	assert.Eq(expected, getProgram(source).String())
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
	expected := []string{
		"p0 = 2.0 * 3.0",
		"p1 = 1.0 * p0",
		"a = p1",
	}
	assertOpcodes(expected, getProgram(source).Instructions(), t)
}

func TestTransDeclare(t *testing.T) {
	source := "var a=1.0*2.0*3.0"
	expected := []string{
		"p0 = 2.0 * 3.0",
		"p1 = 1.0 * p0",
		"a = p1",
	}
	assertOpcodes(expected, getProgram(source).Instructions(), t)
}

func TestFunctionDeclare(t *testing.T) {
	source := `
func add(int a, int b) int {
  return a + b
}
`
	expected := `L0:
FUNC_BEGIN
p1 = a + b
RETURN p1`
	assert.Eq(expected, getProgram(source).String())
}

func TestFunctionDeclareAndCall(t *testing.T) {
	source := `
func fact(int n)  int {
 if(n == 0) {
   return 1
 }
 return fact(n-1) * n
}
`
	expected := `L0:
FUNC_BEGIN
p1 = n == 0
IF p1 ELSE L1
RETURN 1
L1:
p2 = n - 1
PARAM p2 6
SP -5
CALL L0
SP 5
p4 = p3 * n
RETURN p4`
	assert.Eq(expected, getProgram(source).String())
}
