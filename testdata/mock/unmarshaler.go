package mock

import (
	"github.com/ymz-ncnk/mok"
)

func NewUnmarshaller[T any]() Unmarshaller[T] {
	return Unmarshaller[T]{mok.New("Unmarshaller")}
}

type Unmarshaller[T any] struct {
	*mok.Mock
}

func (u Unmarshaller[T]) RegisterUnmarshalMUS(
	fn func(bs []byte) (t T, n int, err error)) Unmarshaller[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u Unmarshaller[T]) RegisterNUnmarshalMUS(n int,
	fn func(bs []byte) (t T, n int, err error)) Unmarshaller[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u Unmarshaller[T]) UnmarshalMUS(bs []byte) (t T, n int, err error) {
	result, err := u.Call("UnmarshalMUS", bs)
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
