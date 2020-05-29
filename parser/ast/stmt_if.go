package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

// StmtIf if 语句
type StmtIf struct {
	*Stmt
}

// NewStmtIf 新建 if stmt
func NewStmtIf() *StmtIf {
	return &StmtIf{NewStmt(IF_STMT, string(IF_STMT))}
}

// GetExpr 获取 if 括号中的表达式
func (sIF *StmtIf) GetExpr() Node {
	return sIF.GetChild(0)
}

// GetBlock 获取 if 语句块
func (sIF *StmtIf) GetBlock() Node {
	return sIF.GetChild(1)
}

// GetElseStmt 获取 else 语句块
func (sIF *StmtIf) GetElseStmt() Node {
	block := sIF.GetChild(2)
	if block.GetType() == BLOCK {
		return block
	}
	return nil
}

// GetElseIfStmt 获取 else if 语句块
func (sIF *StmtIf) GetElseIfStmt() Node {
	stmtIf := sIF.GetChild(2)
	if stmtIf.GetType() == IF_STMT {
		return stmtIf
	}
	return nil
}

// ParseIf IfStmt -> If(Expr) Block Tail 解析 if 语句的首部
func ParseIf(it *util.PeekTokenIterator) (Node, error) {
	lexeme, err := it.MatchNextValue("if")
	if err != nil {
		return nil, err
	}
	if _, err := it.MatchNextValue("("); err != nil {
		return nil, err
	}
	stmtIf := NewStmtIf()
	stmtIf.SetLexeme(lexeme)
	expr, err := ParseExpr(it)
	if err != nil {
		return nil, err
	}
	stmtIf.AddChild(expr)
	if _, err := it.MatchNextValue(")"); err != nil {
		return nil, err
	}
	block, err := ParseBlock(it)
	if err != nil {
		return nil, err
	}
	stmtIf.AddChild(block)
	tail, err := ParseIfTail(it)
	if err != nil {
		return nil, err
	}
	if tail != nil {
		stmtIf.AddChild(tail)
	}
	return stmtIf, nil
}

// ParseIfTail 解析 if 语句的尾部
func ParseIfTail(it *util.PeekTokenIterator) (Node, error) {
	if !it.HasNext() || it.Peek().(*lexer.Token).GetValue() != "else" {
		return nil, nil
	}
	if _, err := it.MatchNextValue("else"); err != nil {
		return nil, err
	}
	lookahead := it.Peek().(*lexer.Token)
	if lookahead.GetValue() == "{" {
		return ParseBlock(it)
	}
	if lookahead.GetValue() == "if" {
		return ParseIf(it)
	}
	return nil, nil
}
