package iterator

import (
	"bufio"
	"go-tinyscript/utils/assert"
	"strings"
	"testing"
)

func TestScannerIterator_All(t *testing.T) {
	var Eq = assert.Eq
	str := "1 2 3 4 "
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	it := NewScannerIterator(scanner)
	it.Scan()
	Eq("1", it.Value())
	it.Scan()
	Eq("2", it.Value())
	it.Scan()
	Eq("3", it.Value())
	it.Scan()
	Eq("4", it.Value())
	Eq(it.Scan(), false)
	Eq(it.Scan(), false)
	Eq(it.Scan(), false)
	Eq(it.Scan(), false)
}
