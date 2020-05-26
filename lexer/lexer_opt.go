package lexer

import (
	"go-tinyscript/common"
	"strings"
)

// MakeVarOrKeyword 获取变量或关键词
func MakeVarOrKeyword(it *common.PeekIterator) (*Token, error) {
	s, sBuilder := "", strings.Builder{}

	for it.HasNext() {
		lookahead := it.Peek()
		if !common.IsLiteral(lookahead) {
			s = sBuilder.String()
			break
		}
		sBuilder.WriteString(lookahead)
		it.Next()
	}

	if IsKeyword(s) {
		return NewToken(KEYWORD, s), nil
	}

	if s == "true" || s == "false" {
		return NewToken(BOOLEAN, s), nil
	}

	return NewToken(VARIABLE, s), nil
}

// MakeString 构造字符串
func MakeString(it *common.PeekIterator) (*Token, error) {
	sBuilder, state := strings.Builder{}, 0

	for it.HasNext() {
		c := it.Next()
		sBuilder.WriteString(c)
		switch state {
		case 0:
			switch c {
			case "\"":
				state = 1
			case "'":
				state = 2
			default:
				break
			}
		case 1:
			if c == "\"" {
				return NewToken(STRING, sBuilder.String()), nil
			}
		case 2:
			if c == "'" {
				return NewToken(STRING, sBuilder.String()), nil
			}
		}
	}
	return nil, NewLexicalError("unexpected error")
}

// MakeOp 构造运算符
func MakeOp(it *common.PeekIterator) (*Token, error) {
	state := 0

	for it.HasNext() {
		lookahead := it.Next()
		switch state {
		case 0:
			switch lookahead {
			case "+":
				state = 1
			case "-":
				state = 2
			case "*":
				state = 3
			case "/":
				state = 4
			case ">":
				state = 5
			case "<":
				state = 6
			case "=":
				state = 7
			case "!":
				state = 8
			case "&":
				state = 9
			case "|":
				state = 10
			case "^":
				state = 11
			case "%":
				state = 12
			case ",":
				return NewToken(OPERATOR, ","), nil
			case ";":
				return NewToken(OPERATOR, ";"), nil
			}
			break
		case 1: // +
			if lookahead == "+" {
				return NewToken(OPERATOR, "++"), nil
			} else if lookahead == "=" {
				return NewToken(OPERATOR, "+="), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "+"), nil
			}
		case 2: // -
			if lookahead == "-" {
				return NewToken(OPERATOR, "--"), nil
			} else if lookahead == "=" {
				return NewToken(OPERATOR, "-="), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "-"), nil
			}
		case 3: // *
			if lookahead == "=" {
				return NewToken(OPERATOR, "*="), nil
			}
			it.PutBack()
			return NewToken(OPERATOR, "*"), nil
		case 4: // /
			if lookahead == "=" {
				return NewToken(OPERATOR, "/="), nil
			}
			it.PutBack()
			return NewToken(OPERATOR, "/"), nil
		case 5: // >
			if lookahead == "=" {
				return NewToken(OPERATOR, ">="), nil
			} else if lookahead == ">" {
				return NewToken(OPERATOR, ">>"), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, ">"), nil
			}
		case 6: // <
			if lookahead == "=" {
				return NewToken(OPERATOR, "<="), nil
			} else if lookahead == "<" {
				return NewToken(OPERATOR, "<<"), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "<"), nil
			}
		case 7: // =
			if lookahead == "=" {
				return NewToken(OPERATOR, "=="), nil
			}
			it.PutBack()
			return NewToken(OPERATOR, "="), nil
		case 8: // !
			if lookahead == "=" {
				return NewToken(OPERATOR, "!="), nil
			}
			it.PutBack()
			return NewToken(OPERATOR, "!"), nil
		case 9: // &
			if lookahead == "&" {
				return NewToken(OPERATOR, "&&"), nil
			} else if lookahead == "=" {
				return NewToken(OPERATOR, "&="), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "&"), nil
			}
		case 10: // |
			if lookahead == "|" {
				return NewToken(OPERATOR, "||"), nil
			} else if lookahead == "=" {
				return NewToken(OPERATOR, "|="), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "|"), nil
			}
		case 11: // ^
			if lookahead == "^" {
				return NewToken(OPERATOR, "^^"), nil
			} else if lookahead == "=" {
				return NewToken(OPERATOR, "^="), nil
			} else {
				it.PutBack()
				return NewToken(OPERATOR, "^"), nil
			}
		case 12: // %
			if lookahead == "=" {
				return NewToken(OPERATOR, "%="), nil
			}
			it.PutBack()
			return NewToken(OPERATOR, "%"), nil
		}
	}
	return nil, NewLexicalError("unexpected error")
}

// MakeNumber 构造数字
func MakeNumber(it *common.PeekIterator) (*Token, error) {
	s, state := strings.Builder{}, 0

	for it.HasNext() {
		lookahead := it.Peek()
		switch state {
		case 0:
			if lookahead == "0" {
				state = 1
			} else if common.IsNumber(lookahead) {
				state = 2
			} else if lookahead == "+" || lookahead == "-" {
				state = 3
			} else if lookahead == "." {
				state = 5
			}
		case 1: // 数字以 0 开头
			if lookahead == "0" {
				state = 1
			} else if common.IsNumber(lookahead) {
				state = 2
			} else if lookahead == "." {
				state = 4
			} else {
				return NewToken(INTEGER, s.String()), nil
			}
		case 2: // 1-9
			if common.IsNumber(lookahead) {
				state = 2
			} else if lookahead == "." {
				state = 4
			} else {
				return NewToken(INTEGER, s.String()), nil
			}
		case 3: // +,-
			if common.IsNumber(lookahead) {
				state = 2
			} else if lookahead == "." {
				state = 5
			} else {
				return nil, NewLexicalError("unexpected error")
			}
		case 4: // .
			if lookahead == "." {
				return nil, NewLexicalError("unexpected error")
			} else if common.IsNumber(lookahead) {
				state = 20
			} else {
				return NewToken(FLOAT, s.String()), nil
			}
		case 5: // 起始为 . 的情况
			if common.IsNumber(lookahead) {
				state = 20
			} else {
				return nil, NewLexicalError("unexpected error")
			}
		case 20: // 数字中已经带有 . 了，小数部分
			if common.IsNumber(lookahead) {
				state = 20
			} else if lookahead == "." {
				return nil, NewLexicalError("unexpected error")
			} else {
				return NewToken(FLOAT, s.String()), nil
			}
		}
		it.Next()
		s.WriteString(lookahead)
	}

	return nil, NewLexicalError("unexpected error")
}
