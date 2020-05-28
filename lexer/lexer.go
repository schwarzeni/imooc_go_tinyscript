package lexer

import (
	"go-tinyscript/common"
	"io"

	"github.com/pkg/errors"
)

// Lexer 词法分析器
type Lexer struct{}

func (l *Lexer) analyze(it *common.PeekIterator) (tokens []*Token, err error) {
	for it.HasNext() {
		c, isEnd := it.Next().(string), !it.HasNext()

		if c == " " || c == "\n" {
			continue
		}

		// 删除注释
		if c == "/" && !isEnd {
			lookahead := it.Peek()
			if lookahead == "/" {
				for it.HasNext() && it.Next() != "\n" {
				}
				continue
			} else if lookahead == "*" {
				it.Next()
				valid := false
				for it.HasNext() {
					p := it.Next()
					if p == "*" && it.Peek() == "/" {
						it.Next()
						valid = true
						break
					}
				}
				if !valid {
					return nil, NewLexicalError("comments not match")
				}
				continue
			}
		}

		if c == "{" || c == "}" || c == "(" || c == ")" {
			tokens = append(tokens, NewToken(BRACKET, c))
			continue
		}

		if c == "\"" || c == "'" {
			it.PutBack()
			t, err := MakeString(it)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, t)
			continue
		}

		if common.IsLetter(c) {
			it.PutBack()
			t, err := MakeVarOrKeyword(it)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, t)
			continue
		}

		if common.IsNumber(c) {
			it.PutBack()
			t, err := MakeNumber(it)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, t)
			continue
		}

		if (c == "+" || c == "-" || c == ".") && !isEnd && common.IsNumber(it.Peek().(string)) {
			var lastToken *Token = nil
			if l := len(tokens); l > 0 {
				lastToken = tokens[l-1]
			}
			if lastToken == nil || !lastToken.IsValue() || lastToken.IsOperator() {
				it.PutBack()
				t, err := MakeNumber(it)
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, t)
				continue
			}
		}

		if common.IsOperator(c) {
			it.PutBack()
			t, err := MakeOp(it)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, t)
			continue
		}
		return nil, errors.WithStack(NewLexicalErrorWithChar(c))
	}
	return
}

// Analyze 解析主方程
func (l *Lexer) Analyze(src io.Reader) (tokens []*Token, err error) {
	it := common.NewPeekIteratorWithIOReader(src)
	if tokens, err = l.analyze(it); err != nil {
		return nil, err
	}
	return
}
