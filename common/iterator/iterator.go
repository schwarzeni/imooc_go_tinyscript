package iterator

type BaseIterator interface {
	Value() interface{}
	Scan() bool
}
