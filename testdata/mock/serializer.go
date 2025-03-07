package mock

import "github.com/ymz-ncnk/mok"

type MarshalFn[T any] func(t T, bs []byte) (n int)
type UnmarshalFn[T any] func(bs []byte) (t T, n int, err error)
type SizeFn[T any] func(t T) (size int)
type SkipFn func(bs []byte) (n int, err error)

func NewSerializer[T any]() Serializer[T] {
	return Serializer[T]{mok.New("Serializer")}
}

type Serializer[T any] struct {
	*mok.Mock
}

func (m Serializer[T]) RegisterMarshal(fn MarshalFn[T]) Serializer[T] {
	m.Register("Marshal", fn)
	return m
}

func (m Serializer[T]) RegisterMarshalN(n int, fn MarshalFn[T]) Serializer[T] {
	m.RegisterN("Marshal", n, fn)
	return m
}

func (u Serializer[T]) RegisterUnmarshal(fn UnmarshalFn[T]) Serializer[T] {
	u.Register("Unmarshal", fn)
	return u
}

func (u Serializer[T]) RegisterUnmarshalN(n int,
	fn func(bs []byte) (t T, n int, err error)) Serializer[T] {
	u.RegisterN("Unmarshal", n, fn)
	return u
}

func (m Serializer[T]) RegisterSize(fn SizeFn[T]) Serializer[T] {
	m.Register("Size", fn)
	return m
}

func (m Serializer[T]) RegisterSizeN(n int, fn SizeFn[T]) Serializer[T] {
	m.RegisterN("Size", n, fn)
	return m
}

func (u Serializer[T]) RegisterSkip(fn SkipFn) Serializer[T] {
	u.Register("Skip", fn)
	return u
}

func (u Serializer[T]) RegisterSkipN(n int, fn SkipFn) Serializer[T] {
	u.RegisterN("Skip", n, fn)
	return u
}

func (m Serializer[T]) Marshal(t T, bs []byte) (n int) {
	result, err := m.Call("Marshal", mok.SafeVal[T](t), bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}

func (u Serializer[T]) Unmarshal(bs []byte) (t T, n int, err error) {
	result, err := u.Call("Unmarshal", bs)
	if err != nil {
		panic(err)
	}
	t, _ = result[0].(T)
	n = result[1].(int)
	err, _ = result[2].(error)
	return
}

func (m Serializer[T]) Size(t T) (size int) {
	result, err := m.Call("Size", mok.SafeVal[T](t))
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}

func (u Serializer[T]) Skip(bs []byte) (n int, err error) {
	result, err := u.Call("Skip", bs)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
