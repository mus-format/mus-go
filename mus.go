package mus

// Marshaller is the interface that wraps the Marshal method.
//
// Marshal method returns the number of used bytes. It should panic if receives
// too small bs.
type Marshaller[T any] interface {
	Marshal(t T, bs []byte) (n int)
}

// MarshallerFn is a functional implementation of the Marshaller interface.
type MarshallerFn[T any] func(t T, bs []byte) (n int)

func (fn MarshallerFn[T]) Marshal(t T, bs []byte) (n int) {
	return fn(t, bs)
}

// Unmarshaller is the interface that wraps the Unmarshal method.
//
// Unmarshal method returns data, the number of used bytes and an error.
type Unmarshaller[T any] interface {
	Unmarshal(bs []byte) (t T, n int, err error)
}

// UnmarshallerFn is a functional implementation of the Unmarshaller interface.
type UnmarshallerFn[T any] func(bs []byte) (t T, n int, err error)

func (fn UnmarshallerFn[T]) Unmarshal(bs []byte) (t T, n int, err error) {
	return fn(bs)
}

// Sizer is the interface that wraps the Size method.
type Sizer[T any] interface {
	Size(t T) (size int)
}

// SizerFn is a functional implementation of the Sizer interface.
type SizerFn[T any] func(t T) (size int)

func (fn SizerFn[T]) Size(t T) (size int) {
	return fn(t)
}

// Skipper is the interface that wraps the Skip method.
//
// Skip method returns the number of skipped bytes and an error.
type Skipper interface {
	Skip(bs []byte) (n int, err error)
}

// SkipperFn is a functional implementation of the Skipper interface.
type SkipperFn func(bs []byte) (n int, err error)

func (fn SkipperFn) Skip(bs []byte) (n int, err error) {
	return fn(bs)
}
