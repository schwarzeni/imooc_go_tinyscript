package util

import (
	"go-tinyscript/common"
	"go-tinyscript/lexer"
	"reflect"

	"github.com/pkg/errors"
)

type PeekTokenIterator struct {
	*common.PeekIterator
}

func NewPeekTokenIterator(tokens []*lexer.Token) *PeekTokenIterator {
	return &PeekTokenIterator{common.NewPeekIteratorWithArray(lexer.Tokens2Arrays(tokens))}
}

// MatchNextValue 根据 token 的值匹配下一个 token
func (it *PeekTokenIterator) MatchNextValue(value string) (token *lexer.Token, err error) {
	token = it.Next().(*lexer.Token)
	if token.GetValue() != value {
		return nil, errors.WithStack(NewParseError(token))
	}
	return
}

// MatchNextValue 根据 token 的类型匹配下一个 token
func (it *PeekTokenIterator) MatchNextType(t lexer.TokenType) (token *lexer.Token, err error) {
	token = it.Next().(*lexer.Token)
	if !reflect.DeepEqual(token.GetType(), t) {
		return nil, errors.WithStack(NewParseError(token))
	}
	return
}
