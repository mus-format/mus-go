package mock

import (
	"github.com/ymz-ncnk/mok"
)

type SizeMUSFn[T any] func(t T) (size int)

func NewSizer[T any]() Sizer[T] {
	return Sizer[T]{mok.New("Sizer")}
}

type Sizer[T any] struct {
	*mok.Mock
}

func (m Sizer[T]) RegisterSizeMUS(fn SizeMUSFn[T]) Sizer[T] {
	m.Register("SizeMUS", fn)
	return m
}

func (m Sizer[T]) RegisterNSizeMUS(n int, fn SizeMUSFn[T]) Sizer[T] {
	m.RegisterN("SizeMUS", n, fn)
	return m
}

func (m Sizer[T]) SizeMUS(t T) (size int) {
	result, err := m.Call("SizeMUS", mok.SafeVal[T](t))
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
