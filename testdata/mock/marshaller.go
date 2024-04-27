package mock

import (
	"github.com/ymz-ncnk/mok"
)

type MarshalMUSFn[T any] func(t T, bs []byte) (n int)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshalMUS(fn MarshalMUSFn[T]) Marshaller[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshalMUS(n int, fn MarshalMUSFn[T]) Marshaller[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaller[T]) MarshalMUS(t T, bs []byte) (n int) {
	result, err := m.Call("MarshalMUS", mok.SafeVal[T](t), bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
