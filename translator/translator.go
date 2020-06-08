package translator

import (
	"errors"
	"fmt"
	"go-tinyscript/lexer"
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

	program.SetStaticSymbols(symbolTable)

	main := lexer.NewToken(lexer.VARIABLE, "main")
	if symbolTable.Exists(main) {
		symbolTable.CreateVariable() // 返回值
		program.Add(NewTAInstruction(-symbolTable.LocalSize(), nil, "", nil, SP))
		program.Add(NewTAInstruction(symbolTable.CloneFromSymbolTree(main, 0), nil, "", nil, CALL))
		program.Add(NewTAInstruction(symbolTable.LocalSize(), nil, "", nil, SP))
	}

	return program, nil
}

// translateStmt 翻译语句
func (t *Translator) translateStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) error {
	var err error
	switch node.GetType() {
	case ast.BLOCK:
		err = t.translateBlock(program, node.(*ast.Block), symbolTable)
	case ast.IF_STMT:
		err = t.translateIfStmt(program, node.(*ast.StmtIf), symbolTable)
	case ast.ASSIGN_STMT:
		err = t.translateAssignStmt(program, node, symbolTable)
	case ast.DECLARE_STMT:
		err = t.translateDeclareStmt(program, node, symbolTable)
	case ast.FUNCTION_DECLARE_STMT:
		err = t.translateFunctionDeclareStmt(program, node, symbolTable)
	case ast.CALL_EXPTR:
		_, err = t.translateCallStmt(program, node, symbolTable)
	case ast.RETURN_STMT:
		err = t.translateReturnStmt(program, node, symbolTable)
	default:
		err = errpkg.WithStack(errors.New("Translator not impl for " + string(node.GetType())))
	}
	return err
}

// translateBlock 翻译 Block
func (t *Translator) translateBlock(program *TAProgram, block *ast.Block, parentTable *symbol.SymbolTable) error {
	symbolTable := symbol.NewSymbolTable()
	// 每个Block增加一个作用域链
	parentOffset := symbolTable.CreateVariable()
	parentOffset.SetLexeme(lexer.NewToken(lexer.INTEGER, fmt.Sprintf("%d", symbolTable.LocalSize())))
	parentTable.AddChild(symbolTable)
	for _, child := range block.GetChildren() {
		if err := t.translateStmt(program, child, symbolTable); err != nil {
			return err
		}
	}
	return nil
}

// translateIfStmt 翻译 if 语句
func (t *Translator) translateIfStmt(program *TAProgram, node *ast.StmtIf, symbolTable *symbol.SymbolTable) error {
	expr := node.GetExpr()
	exprAddr, err := t.translateExpr(program, expr, symbolTable)
	if err != nil {
		return err
	}
	ifOpCode := NewTAInstruction(exprAddr, nil, "", nil, IF)
	program.Add(ifOpCode)

	if err = t.translateBlock(program, node.GetBlock().(*ast.Block), symbolTable); err != nil {
		return err
	}

	var gotoInstruction *TAInstruction
	// IF ... ElSE ...
	if node.GetChild(2) != nil {
		gotoInstruction = NewTAInstruction(nil, nil, "", nil, GOTO)
		program.Add(gotoInstruction)
		labelEndIf := program.AddLabel()
		ifOpCode.SetArg2(labelEndIf.Arg1())
	}

	if node.GetElseStmt() != nil { // ELSE {
		err = t.translateBlock(program, node.GetElseStmt().(*ast.Block), symbolTable)
	} else if node.GetElseIfStmt() != nil { // ELSE IF ... {
		err = t.translateIfStmt(program, node.GetElseIfStmt().(*ast.StmtIf), symbolTable)
	}
	if err != nil {
		return err
	}

	labelEnd := program.AddLabel()
	if node.GetChild(2) == nil {
		ifOpCode.SetArg2(labelEnd.Arg1())
	} else {
		gotoInstruction.SetArg1(labelEnd.Arg1())
	}
	return nil
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
		addr, err := t.translateCallStmt(program, node, symbolTable)
		if err != nil {
			return nil, err
		}
		node.SetProp("addr", addr)
		return addr, nil
	} else if _, ok := node.(*ast.Expr); ok {
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
	return nil, errpkg.WithStack(errors.New("Unexpected node type: " + string(node.GetType())))
}

// translateDeclareStmt 解析声明语句
func (t *Translator) translateDeclareStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) error {
	lexeme := node.GetChild(0).GetLexeme()
	// 不能重复声明变量
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

// translateReturnStmt 解析 return 语句
func (t *Translator) translateReturnStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) (err error) {
	var resultValue *symbol.Symbol
	if node.GetChild(0) != nil {
		if resultValue, err = t.translateExpr(program, node.GetChild(0), symbolTable); err != nil {
			return err
		}
	}
	program.Add(NewTAInstruction(resultValue, nil, "", nil, RETURN))
	return nil
}

// translateFunctionDeclareStmt 解析函数声明
func (t *Translator) translateFunctionDeclareStmt(program *TAProgram, node ast.Node, parent *symbol.SymbolTable) (err error) {
	label := program.AddLabel()
	symbolTable := symbol.NewSymbolTable()

	program.Add(NewTAInstruction(nil, nil, "", nil, FUNC_BEGIN))
	symbolTable.CreateVariable() // 返回地址

	label.SetArg2(node.GetLexeme().GetValue())
	fn := node.(*ast.StmtFunctionDeclare)
	args := fn.GetArgs()
	parent.AddChild(symbolTable)
	parent.CreateLabel(label.Arg1().(string), node.GetLexeme())
	for _, arg := range args.GetChildren() {
		symbolTable.CreateSymbolByLexeme(arg.GetLexeme())
	}

	for _, child := range fn.GetBlock().GetChildren() {
		if err := t.translateStmt(program, child, symbolTable); err != nil {
			return err
		}
	}
	return nil
}

// translateCallStmt 函数调用
func (t *Translator) translateCallStmt(program *TAProgram, node ast.Node, symbolTable *symbol.SymbolTable) (*symbol.Symbol, error) {
	factor := node.GetChild(0)
	var list []*TAInstruction
	for i := 1; i < len(node.GetChildren()); i++ {
		expr := node.GetChildren()[i]
		addr, err := t.translateExpr(program, expr, symbolTable)
		if err != nil {
			return nil, err
		}
		list = append(list, NewTAInstruction(addr, i-1, "", nil, PARAM))
	}

	for _, instruction := range list {
		instruction.SetArg2(symbolTable.LocalSize() + instruction.Arg2().(int) + 2)
		program.Add(instruction)
	}

	returnValue := symbolTable.CreateVariable()
	fnAddr := symbolTable.CloneFromSymbolTree(factor.GetLexeme(), 0)
	if fnAddr == nil {
		return nil, errpkg.WithStack(errors.New("function " + factor.GetLexeme().GetValue() + " not found"))
	}
	program.Add(NewTAInstruction(-symbolTable.LocalSize(), nil, "", nil, SP))
	program.Add(NewTAInstruction(fnAddr, nil, "", nil, CALL))
	program.Add(NewTAInstruction(symbolTable.LocalSize(), nil, "", nil, SP))

	return returnValue, nil
}

func NewTranslator() *Translator {
	return &Translator{}
}
