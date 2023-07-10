package mus

// Marshaller is the interface that wraps the MarshalMUS method.
//
// MarshalMUS marshals data to the MUS format and returns the number of used
// bytes.
//
// It should panic if receives too small bs.
type Marshaller[T any] interface {
	MarshalMUS(t T, bs []byte) (n int)
}

// MarshallerFn is a functional implementation of the Marshaller interface.
type MarshallerFn[T any] func(t T, bs []byte) (n int)

func (fn MarshallerFn[T]) MarshalMUS(t T, bs []byte) (n int) {
	return fn(t, bs)
}

// UnMarshaller is the interface that wraps the UnmarshalMUS method.
//
// UnmarshalMUS unmarshals data from the MUS format. Returns data, the number of
// used bytes and an error.
type UnMarshaller[T any] interface {
	UnmarshalMUS(bs []byte) (t T, n int, err error)
}

// UnMarshallerFn is a functional implementation of the UnMarshaller interface.
type UnMarshallerFn[T any] func(bs []byte) (t T, n int, err error)

func (fn UnMarshallerFn[T]) UnmarshalMUS(bs []byte) (t T, n int, err error) {
	return fn(bs)
}

// Sizer is the interface that wraps the SizeMUS method.
//
// SizeMUS calculates the size of data in the MUS format.
type Sizer[T any] interface {
	SizeMUS(t T) (size int)
}

// SizerFn is a functional implementation of the Sizer interface.
type SizerFn[T any] func(t T) (size int)

func (fn SizerFn[T]) SizeMUS(t T) (size int) {
	return fn(t)
}

// Skipper is the interface that wraps the SkipMUS method.
//
// SkipMUS skips data in the MUS format. Returns the number of skipped bytes and
// an error.
type Skipper interface {
	SkipMUS(bs []byte) (n int, err error)
}

// SkipperFn is a functional implementation of the Skipper interface.
type SkipperFn func(bs []byte) (n int, err error)

func (fn SkipperFn) SkipMUS(bs []byte) (n int, err error) {
	return fn(bs)
}
