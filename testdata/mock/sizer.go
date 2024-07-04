package mock

import (
	"github.com/ymz-ncnk/mok"
)

type SizeFn[T any] func(t T) (size int)

func NewSizer[T any]() Sizer[T] {
	return Sizer[T]{mok.New("Sizer")}
}

type Sizer[T any] struct {
	*mok.Mock
}

func (m Sizer[T]) RegisterSize(fn SizeFn[T]) Sizer[T] {
	m.Register("Size", fn)
	return m
}

func (m Sizer[T]) RegisterNSize(n int, fn SizeFn[T]) Sizer[T] {
	m.RegisterN("Size", n, fn)
	return m
}

func (m Sizer[T]) Size(t T) (size int) {
	result, err := m.Call("Size", mok.SafeVal[T](t))
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
