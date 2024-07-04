package mock

import (
	"github.com/ymz-ncnk/mok"
)

type MarshalFn[T any] func(t T, bs []byte) (n int)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshal(fn MarshalFn[T]) Marshaller[T] {
	m.Register("Marshal", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshal(n int, fn MarshalFn[T]) Marshaller[T] {
	m.RegisterN("Marshal", n, fn)
	return m
}

func (m Marshaller[T]) Marshal(t T, bs []byte) (n int) {
	result, err := m.Call("Marshal", mok.SafeVal[T](t), bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
