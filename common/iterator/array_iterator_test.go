package iterator

import (
	"go-tinyscript/utils/assert"
	"testing"
)

func TestArrayIterator_All(t *testing.T) {
	var Eq = assert.Eq
	arr := []int{1, 2, 3, 4}
	middle := make([]interface{}, len(arr))
	for i, d := range arr {
		middle[i] = d
	}
	it := NewArrayIterator(middle)
	it.Scan()
	Eq(it.Value(), 1)
	it.Scan()
	Eq(it.Value(), 2)
	it.Scan()
	Eq(it.Value(), 3)
	it.Scan()
	Eq(it.Value(), 4)
	Eq(it.Scan(), false)
	Eq(it.Scan(), false)
	Eq(it.Scan(), false)
}
