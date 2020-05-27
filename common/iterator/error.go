package iterator

type ErrIteratorEnd struct {
}

func NewErrIteratorEnd() *ErrIteratorEnd {
	return &ErrIteratorEnd{}
}

func (e ErrIteratorEnd) Error() string {
	return "reach to the end"
}

type ErrIteratorNoBegin struct{}

func NewErrIteratorNoBegin() *ErrIteratorNoBegin {
	return &ErrIteratorNoBegin{}
}

func (e ErrIteratorNoBegin) Error() string {
	panic("not begin")
}
