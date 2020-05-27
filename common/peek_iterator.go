package common

import (
	"bufio"
	"container/list"
	"errors"
	"go-tinyscript/common/iterator"
	"io"
)

var cacheSize = 1000

// ErrEnd 迭代器结束
var ErrEnd = errors.New("iterator reachs to end")

// PeekIterator 迭代器，读取输入字符流
// cache: A B C D
// putBack D C B A
type PeekIterator struct {
	scanner iterator.BaseIterator
	cache   *list.List
	putBack *list.List
}

// PutBack 放回元素
func (it *PeekIterator) PutBack() {
	if it.cache.Len() > 0 {
		elem := it.cache.Back()
		val := elem.Value
		it.cache.Remove(elem)
		it.putBack.PushBack(val)
	}
}

// Peek 查看迭代器中的下一个元素
func (it *PeekIterator) Peek() (val interface{}) {
	if it.putBack.Len() > 0 {
		return it.putBack.Back().Value
	}
	if !it.scanner.Scan() {
		panic(ErrEnd)
	}
	val = it.scanner.Value()
	it.putBack.PushBack(val)
	return
}

// HasNext 判断是否读取完毕
func (it *PeekIterator) HasNext() bool {
	if it.putBack.Len() > 0 {
		return true
	}
	if it.scanner.Scan() {
		it.putBack.PushBack(it.scanner.Value())
		return true
	}
	return false
}

// Next 读取下一个字符
func (it *PeekIterator) Next() (val interface{}) {
	if it.putBack.Len() > 0 {
		elem := it.putBack.Back()
		val = elem.Value
		it.putBack.Remove(elem)
	} else {
		if !it.scanner.Scan() {
			panic(ErrEnd)
		}
		val = it.scanner.Value()
	}
	for it.cache.Len() >= cacheSize {
		it.cache.Remove(it.cache.Front())
	}
	it.cache.PushBack(val)
	return
}

// NewPeekIteratorWithIOReader 初始化迭代器，传入参数需实现接口 io.Reader
func NewPeekIteratorWithIOReader(src io.Reader) (pi *PeekIterator) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanBytes)
	it := iterator.NewScannerIterator(scanner)
	return NewPeekIterator(it)
}

// NewPeekIteratorWithArray 初始化迭代器，需要迭代的为一个数组
func NewPeekIteratorWithArray(src []interface{}) (pi *PeekIterator) {
	it := iterator.NewArrayIterator(src)
	return NewPeekIterator(it)
}

// NewPeekIterator 初始化 PeekIterator 实例
func NewPeekIterator(it iterator.BaseIterator) (pi *PeekIterator) {
	return &PeekIterator{
		scanner: it,
		cache:   list.New(),
		putBack: list.New(),
	}
}
