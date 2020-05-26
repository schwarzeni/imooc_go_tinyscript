package common

import (
	"bufio"
	"container/list"
	"errors"
	"io"
)

var cacheSize = 1000

// ErrEnd 迭代器结束
var ErrEnd = errors.New("iterator reachs to end")

// PeekIterator 迭代器，读取输入字符流
// cache: A B C D
// putBack D C B A
type PeekIterator struct {
	scanner *bufio.Scanner
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
func (it *PeekIterator) Peek() (val string) {
	if it.putBack.Len() > 0 {
		return it.putBack.Back().Value.(string)
	}
	if !it.scanner.Scan() {
		panic(ErrEnd)
	}
	val = it.scanner.Text()
	it.putBack.PushBack(val)
	return
}

// HasNext 判断是否读取完毕
func (it *PeekIterator) HasNext() bool {
	if it.putBack.Len() > 0 {
		return true
	}
	if it.scanner.Scan() {
		it.putBack.PushBack(it.scanner.Text())
		return true
	}
	return false
}

// Next 读取下一个字符
func (it *PeekIterator) Next() (val string) {
	if it.putBack.Len() > 0 {
		elem := it.putBack.Back()
		val = elem.Value.(string)
		it.putBack.Remove(elem)
	} else {
		if !it.scanner.Scan() {
			panic(ErrEnd)
		}
		val = it.scanner.Text()
	}
	for it.cache.Len() >= cacheSize {
		it.cache.Remove(it.cache.Front())
	}
	it.cache.PushBack(val)
	return
}

// NewPeekIterator 初始化迭代器，传入参数需实现接口 io.Reader
func NewPeekIterator(src io.Reader) (pi *PeekIterator) {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanBytes)
	pi = &PeekIterator{
		scanner: scanner,
		cache:   list.New(),
		putBack: list.New(),
	}
	return
}
