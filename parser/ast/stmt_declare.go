package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"

	"github.com/pkg/errors"
)

// StmtDeclare 声明语句 var a = 2
type StmtDeclare struct {
	*Stmt
}

// NewStmtDeclare 新建声明语句
func NewStmtDeclare() *StmtDeclare {
	return &StmtDeclare{NewStmt(DECLARE_STMT, string(DECLARE_STMT))}
}

// ParseDeclare 解析声明语句 var a = 2 + 3
func ParseDeclare(it *util.PeekTokenIterator) (Node, error) {
	stmt := NewStmtDeclare()
	if _, err := it.MatchNextValue("var"); err != nil { // var
		return nil, err
	}
	token := it.Peek().(*lexer.Token) // a
	factor, err := ParseFactor(it)    // a
	if err != nil {
		return nil, err
	}
	if factor == nil {
		return nil, errors.WithStack(util.NewParseError(token))
	}
	stmt.AddChild(factor)
	lexeme, err := it.MatchNextValue("=") // =
	if err != nil {
		return nil, err
	}
	expr, err := ParseExpr(it) // 2 + 3
	if err != nil {
		return nil, err
	}
	stmt.AddChild(expr)
	stmt.SetLexeme(lexeme)
	return stmt, nil
}
