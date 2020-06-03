package translator

import (
	"errors"
	"go-tinyscript/parser/ast"
	"go-tinyscript/translator/symbol"

	errpkg "github.com/pkg/errors"
)

type Translator struct{}

// Translate 翻译 ast 语法树
func (t *Translator) Translate(astNode ast.Node) (*TAProgram, error) {
	program := NewTAProgram()
	symbolTable := symbol.NewSymbolTable()

	for _, child := range astNode.GetChildren() {
		if err := t.translateStmt(program, child, symbolTable); err != nil {
			return nil, err
		}
	}

	return program, nil
}

// translateStmt 翻译语句
func (t *Translator) translateStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) error {
	var err error
	switch node.GetType() {
	case ast.ASSIGN_STMT:
		err = t.translateAssignStmt(program, node, symbolTable)
	case ast.DECLARE_STMT:
		err = t.translateDeclareStmt(program, node, symbolTable)
	default:
		err = errNotImpl(string(node.GetType()))
	}
	return err
}

// translateAssignStmt 翻译赋值语句
// 2 * 3 + 1; =>
// p0 = 2 * 3;
// p1 = p0 + 1;
func (t *Translator) translateAssignStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) error {
	// var a = expr;
	assigned := symbolTable.CreateSymbolByLexeme(node.GetChild(0).GetLexeme()) // a
	expr := node.GetChild(1)                                                   // expr
	addr, err := t.translateExpr(program, expr, symbolTable)
	if err != nil {
		return err
	}
	program.Add(NewTAInstruction(addr, nil, "=", assigned, ASSIGN))
	return nil
}

// translateExpr 解析表达式
func (t *Translator) translateExpr(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) (sol *symbol.Symbol, err error) {
	if node.IsValueType() {
		addr := symbolTable.CreateSymbolByLexeme(node.GetLexeme())
		node.SetProp("addr", addr)
		return addr, nil
	} else if node.GetType() == ast.CALL_EXPTR {
		return nil, errNotImpl(string(node.GetType()))
	}

	for _, child := range node.GetChildren() {
		if _, err = t.translateExpr(program, child, symbolTable); err != nil {
			return
		}
	}

	if _, ok := node.GetProp("addr"); !ok {
		node.SetProp("addr", symbolTable.CreateVariable())
	}

	arg1, _ := node.GetChild(0).GetProp("addr")
	arg2, _ := node.GetChild(1).GetProp("addr")
	result, _ := node.GetProp("addr")
	instruction := NewTAInstruction(
		arg1.(*symbol.Symbol),
		arg2.(*symbol.Symbol),
		node.GetLexeme().GetValue(),
		result.(*symbol.Symbol),
		ASSIGN,
	)

	program.Add(instruction)
	return instruction.Result(), nil
}

// translateDeclareStmt 解析声明语句
func (t *Translator) translateDeclareStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) error {
	lexeme := node.GetChild(0).GetLexeme()
	if symbolTable.Exists(lexeme) {
		return errpkg.WithStack(errors.New("Syntax Error, Identifier " + lexeme.GetValue() + " is already defined"))
	}
	assigned := symbolTable.CreateSymbolByLexeme(lexeme)
	expr := node.GetChild(1)
	addr, err := t.translateExpr(program, expr, symbolTable)
	if err != nil {
		return err
	}
	program.Add(NewTAInstruction(addr, nil, "=", assigned, ASSIGN))
	return nil
}

func NewTranslator() *Translator {
	return &Translator{}
}

func errNotImpl(msg string) error {
	return errpkg.WithStack(errors.New("Translator not impl for " + msg))
}