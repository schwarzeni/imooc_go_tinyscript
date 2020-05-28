package ast

import (
	"go-tinyscript/lexer"
	"go-tinyscript/parser/util"
)

type Expr struct {
	*ASTNode
}

func NewExpr(t ASTNodeType, lexeme *lexer.Token) *Expr {
	if lexeme == nil {
		return &Expr{&ASTNode{}}
	}
	return &Expr{NewASTNode(t, lexeme.GetValue(), lexeme)}
}

// ParseExpr 解析表达式
func ParseExpr(it *util.PeekTokenIterator) (Node, error) {
	return e(0, it)
}

var priorityTable = util.NewPrioriryTable()

type exprHof func() (Node, error)

// e
// E(k) -> E(k+1) E_(k)
// E(t) -> F E_(k) | U E_(k)
func e(k int, it *util.PeekTokenIterator) (n Node, err error) {
	if k < priorityTable.Size()-1 {
		return combine(it,
			func() (Node, error) { return e(k+1, it) },
			func() (Node, error) { return e_(k, it) })
	}
	return race(it,
		func() (Node, error) {
			return combine(it,
				func() (Node, error) { return f(it) },
				func() (Node, error) { return e_(k, it) })
		}, func() (Node, error) {
			return combine(it,
				func() (Node, error) { return u(it) },
				func() (Node, error) { return e_(k, it) })
		})
}

// e_  E_(k) -> op(k) E(k+1) E_(k) | ε
func e_(k int, it *util.PeekTokenIterator) (n Node, err error) {
	token := it.Peek().(*lexer.Token)
	value := token.GetValue()
	if priorityTable.Has(k, value) {
		it.Next()
		n = NewExpr(BINARY_EXPR, token)
		var c Node
		if c, err = combine(it,
			func() (Node, error) { return e(k+1, it) },
			func() (Node, error) { return e_(k, it) }); err != nil {
			return nil, err
		}
		n.AddChild(c)
	}
	return
}

// u
func u(it *util.PeekTokenIterator) (n Node, err error) {
	token := it.Peek().(*lexer.Token)
	value := token.GetValue()

	if value == "(" {
		if _, err = it.MatchNextValue("("); err != nil {
			return nil, err
		}
		if n, err = e(0, it); err != nil {
			return nil, err
		}
		if _, err = it.MatchNextValue(")"); err != nil {
			return nil, err
		}
		return
	}

	if value == "++" || value == "--" || value == "!" {
		it.Next()
		n = NewExpr(UNARY_EXPR, token)
		var c Node
		if c, err = e(0, it); err != nil {
			return nil, err
		}
		n.AddChild(c)
		return
	}
	return nil, nil
}

// f factor
func f(it *util.PeekTokenIterator) (factor Node, err error) {
	if factor, err = ParseFactor(it); err != nil {
		return nil, err
	}
	if factor == nil {
		return
	}
	if it.HasNext() && it.Peek().(*lexer.Token).GetValue() == "(" {
		return ParseExperCall(factor, it)
	}
	return
}

// combine 合并 aFunc 和 bFunc 的结果
func combine(it *util.PeekTokenIterator, aFunc, bFunc exprHof) (Node, error) {
	var (
		err  error
		a, b Node
	)
	if a, err = aFunc(); err != nil {
		return nil, err
	}
	if a == nil {
		if !it.HasNext() {
			return nil, nil
		}
		return bFunc()
	}
	if !it.HasNext() {
		return a, nil
	}
	if b, err = bFunc(); err != nil {
		return nil, err
	}
	if b == nil {
		return a, nil
	}

	// 右递归，所以如果 b 不为空的话，则为 Expr ，子树结构
	expr := NewExpr(BINARY_EXPR, b.GetLexeme())
	expr.AddChild(a)
	expr.AddChild(b.GetChild(0))
	return expr, nil
}

// race 只需要 aFunc 或 bFunc 中的一个得出结果即可
// E|F
func race(it *util.PeekTokenIterator, aFunc, bFunc exprHof) (Node, error) {
	var (
		err error
		a   Node
	)
	if !it.HasNext() {
		return nil, nil
	}
	if a, err = aFunc(); err != nil {
		return nil, err
	}
	if a != nil {
		return a, nil
	}
	return bFunc()
}
