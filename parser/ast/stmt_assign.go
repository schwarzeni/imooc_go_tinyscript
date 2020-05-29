package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"

	"github.com/pkg/errors"
)

// NewStmtAssign 赋值语句 a = 2 + 3 + b
type StmtAssign struct {
	*Stmt
}

// NewStmtAssign 新建赋值语句
func NewStmtAssign() *StmtAssign {
	return &StmtAssign{NewStmt(ASSIGN_STMT, string(ASSIGN_STMT))}
}

// ParseAssign 解析赋值语句
func ParseAssign(it *util.PeekTokenIterator) (Node, error) {
	stmt := NewStmtAssign()
	token := it.Peek().(*lexer.Token)
	factor, err := ParseFactor(it)
	if err != nil {
		return nil, err
	}
	if factor == nil {
		return nil, errors.WithStack(util.NewParseError(token))
	}
	stmt.AddChild(factor)
	lexeme, err := it.MatchNextValue("=")
	if err != nil {
		return nil, err
	}
	expr, err := ParseExpr(it)
	if err != nil {
		return nil, err
	}
	stmt.AddChild(expr)
	stmt.SetLexeme(lexeme)
	return stmt, nil
}
