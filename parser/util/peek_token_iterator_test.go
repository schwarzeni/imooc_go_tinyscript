package util

import (
	"go-tinyscript/lexer"
	"go-tinyscript/utils/assert"
	"testing"
)

func TestPeekTokenIterator_ALL(t *testing.T) {
	eq := assert.Eq
	notNil := assert.NotNil
	tokens := []*lexer.Token{
		lexer.NewToken(lexer.KEYWORD, "var"),
		lexer.NewToken(lexer.INTEGER, "12"),
		lexer.NewToken(lexer.OPERATOR, "+"),
		lexer.NewToken(lexer.BOOLEAN, "false"),
		lexer.NewToken(lexer.STRING, "+"),
	}
	it := NewPeekTokenIterator(tokens)
	eq(it.Next(), tokens[0])
	token1, err := it.MatchNextType(lexer.INTEGER)
	eq(err, nil)
	eq(token1, tokens[1])
	token2, err := it.MatchNextValue("+")
	eq(err, nil)
	eq(token2, tokens[2])
	_, err = it.MatchNextType(lexer.STRING)
	notNil(err)
	_, err = it.MatchNextValue("2233")
	notNil(err)
	eq(it.HasNext(), false)
}
