package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"

	"github.com/pkg/errors"
)

// StmtFunctionDeclare 函数声明语句
type StmtFunctionDeclare struct {
	*Stmt
}

// GetName 获取函数名称
func (fn *StmtFunctionDeclare) GetName() *Variable {
	return fn.GetChild(0).(*Variable)
}

// GetArgs 获取函数变量
func (fn *StmtFunctionDeclare) GetArgs() Node {
	return fn.GetChild(1)
}

// GetReturn 获取函数返回值
func (fn *StmtFunctionDeclare) GetReturn() *lexer.Token {
	return fn.GetName().TypeLexeme()
}

// GetBlock 获取函数体
func (fn *StmtFunctionDeclare) GetBlock() Node {
	return fn.GetChild(2)
}

// NewStmtFunctionDeclare 新建函数声明语句
func NewStmtFunctionDeclare() *StmtFunctionDeclare {
	return &StmtFunctionDeclare{NewStmt(FUNCTION_DECLARE_STMT, string(FUNCTION_DECLARE_STMT))}
}

// ParseFunctionDeclare 解析函数声明
func ParseFunctionDeclare(it *util.PeekTokenIterator) (Node, error) {
	if _, err := it.MatchNextValue("func"); err != nil {
		return nil, err
	}
	fn := NewStmtFunctionDeclare()
	lexeme := it.Peek().(*lexer.Token)
	fv, err := ParseFactor(it)
	if err != nil {
		return nil, err
	}
	funcVariable := fv.(*Variable)
	fn.SetLexeme(lexeme)
	fn.AddChild(funcVariable)
	if _, err := it.MatchNextValue("("); err != nil {
		return nil, err
	}
	args, err := ParseFunctionArgs(it)
	if err != nil {
		return nil, err
	}
	fn.AddChild(args)
	if _, err := it.MatchNextValue(")"); err != nil {
		return nil, err
	}
	keyword, err := it.MatchNextType(lexer.KEYWORD)
	if err != nil {
		return nil, err
	}
	if !keyword.IsType() {
		return nil, errors.WithStack(util.NewParseError(keyword))
	}
	funcVariable.SetTypeLexeme(keyword)
	block, err := ParseBlock(it)
	if err != nil {
		return nil, err
	}
	fn.AddChild(block)
	return fn, nil
}
