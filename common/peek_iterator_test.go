package common

import (
	"go-tinyscript/utils/assert"
	"strings"
	"testing"
)

var eq = assert.Eq

func Test_all(t *testing.T) {
	str := "abcd"
	pi := NewPeekIterator(strings.NewReader(str))

	eq(pi.Peek(), "a")
	eq(pi.Peek(), "a")
	eq(pi.Next(), "a")
	eq(pi.Next(), "b")
	eq(pi.Peek(), "c")
	eq(pi.Peek(), "c")
	eq(pi.Peek(), "c")
	eq(pi.Next(), "c")
	eq(pi.Peek(), "d")
	eq(pi.Peek(), "d")
	eq(pi.Next(), "d")
	eq(pi.HasNext(), false)
	pi.PutBack()
	eq(pi.Peek(), "d")
	pi.PutBack()
	eq(pi.Peek(), "c")
	pi.PutBack()
	eq(pi.Peek(), "b")
	eq(pi.Peek(), "b")
	pi.PutBack()
	eq(pi.Peek(), "a")
	eq(pi.Next(), "a")
	eq(pi.Peek(), "b")
	eq(pi.Peek(), "b")
	eq(pi.Next(), "b")
	eq(pi.Next(), "c")
	pi.PutBack()
	eq(pi.Next(), "c")
	eq(pi.Peek(), "d")
	eq(pi.Peek(), "d")
	eq(pi.Next(), "d")
	eq(pi.HasNext(), false)

	// cache size not enough
	tmp := cacheSize
	cacheSize = 2
	pi2 := NewPeekIterator(strings.NewReader("123"))
	pi2.Next()
	pi2.Next()
	pi2.Next()
	pi2.PutBack()
	pi2.PutBack()
	eq(pi2.Next(), "2")
	eq(pi2.Next(), "3")
	eq(pi2.HasNext(), false)
	cacheSize = tmp
}
