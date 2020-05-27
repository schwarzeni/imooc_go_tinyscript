package util

import (
	"go-tinyscript/common"
	"go-tinyscript/lexer"
	"reflect"
)

type PeekTokenIterator struct {
	*common.PeekIterator
}

func NewPeekTokenIterator(tokens []*lexer.Token) *PeekTokenIterator {
	arr := make([]interface{}, len(tokens))
	for idx, _ := range arr {
		arr[idx] = tokens[idx]
	}
	return &PeekTokenIterator{common.NewPeekIteratorWithArray(arr)}
}

// MatchNextValue 根据 token 的值匹配下一个 token
func (it *PeekTokenIterator) MatchNextValue(value string) (token *lexer.Token, err error) {
	token = it.Next().(*lexer.Token)
	if token.GetValue() != value {
		return nil, NewParseError(token)
	}
	return
}

// MatchNextValue 根据 token 的类型匹配下一个 token
func (it *PeekTokenIterator) MatchNextType(t lexer.TokenType) (token *lexer.Token, err error) {
	token = it.Next().(*lexer.Token)
	if !reflect.DeepEqual(token.GetType(), t) {
		return nil, NewParseError(token)
	}
	return
}
