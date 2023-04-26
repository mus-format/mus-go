package mock

import (
	"github.com/ymz-ncnk/mok"
)

func NewUnmarshaler[T any]() Unmarshaler[T] {
	return Unmarshaler[T]{mok.New("Unmarshaler")}
}

type Unmarshaler[T any] struct {
	*mok.Mock
}

func (u Unmarshaler[T]) RegisterUnmarshalMUS(
	fn func(bs []byte) (t T, n int, err error)) Unmarshaler[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u Unmarshaler[T]) RegisterNUnmarshalMUS(n int,
	fn func(bs []byte) (t T, n int, err error)) Unmarshaler[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u Unmarshaler[T]) UnmarshalMUS(bs []byte) (t T, n int, err error) {
	result, err := u.Call("UnmarshalMUS", bs)
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
