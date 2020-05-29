package ast

import (
	"go-tinyscript/parser/util"
)

// StmtReturn 函数返回语句
type StmtReturn struct {
	*Stmt
}

// GetReturnTokens 获取全部的返回值
//func (r *StmtReturn) GetReturnTokens() []Node {
//	return r.GetChildren()
//}

// NewStmtReturn 新建返回语句
func NewStmtReturn() *StmtReturn {
	return &StmtReturn{NewStmt(RETURN_STMT, "return")}
}

// ParseReturn 解析返回语句
func ParseReturn(it *util.PeekTokenIterator) (Node, error) {
	lexeme, err := it.MatchNextValue("return")
	if err != nil {
		return nil, err
	}
	expr, err := ParseExpr(it)
	if err != nil {
		return nil, err
	}
	stmt := NewStmtReturn()
	stmt.SetLexeme(lexeme)
	if expr != nil {
		stmt.AddChild(expr)
	}
	return stmt, nil
}
