package mock

import (
	"github.com/ymz-ncnk/mok"
)

func NewUnMarshaller[T any]() UnMarshaller[T] {
	return UnMarshaller[T]{mok.New("UnMarshaller")}
}

type UnMarshaller[T any] struct {
	*mok.Mock
}

func (u UnMarshaller[T]) RegisterUnmarshalMUS(
	fn func(bs []byte) (t T, n int, err error)) UnMarshaller[T] {
	u.Register("UnmarshalMUS", fn)
	return u
}

func (u UnMarshaller[T]) RegisterNUnmarshalMUS(n int,
	fn func(bs []byte) (t T, n int, err error)) UnMarshaller[T] {
	u.RegisterN("UnmarshalMUS", n, fn)
	return u
}

func (u UnMarshaller[T]) UnmarshalMUS(bs []byte) (t T, n int, err error) {
	result, err := u.Call("UnmarshalMUS", bs)
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}
