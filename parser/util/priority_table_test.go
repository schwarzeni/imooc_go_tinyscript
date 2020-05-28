package util

import (
	"go-tinyscript/utils/assert"
	"testing"
)

func TestPrioriryTable_all(t *testing.T) {
	eq := assert.Eq
	pt := NewPrioriryTable()
	eq(pt.Has(1, "=="), true)
	eq(pt.Has(2, "=="), false)
	eq(pt.Size(), 5)
}
