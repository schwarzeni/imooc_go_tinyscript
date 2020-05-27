package iterator

type ArrayIterator struct {
	array []interface{}
	idx   int
}

func NewArrayIterator(array []interface{}) *ArrayIterator {
	return &ArrayIterator{array: array, idx: -1}
}

func (a *ArrayIterator) Value() interface{} {
	if a.idx >= len(a.array) {
		panic(NewErrIteratorEnd())
	}
	if a.idx == -1 {
		return nil
	}
	return a.array[a.idx]
}

func (a *ArrayIterator) Scan() bool {
	if a.idx+1 >= len(a.array) {
		return false
	}
	a.idx++
	return true
}
