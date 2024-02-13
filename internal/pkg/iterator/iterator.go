package iterator

type IIterator interface {
	GetNext() interface{}
	HasNext() bool
}
