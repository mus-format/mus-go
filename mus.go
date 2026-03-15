// Package mus provides Serializer interface.
package mus

// Serializer is the interface that groups the Marshal, Unmarshal, Size and
// Skip methods.
//
// Marshal fills bs with an encoded value, returning the number of used bytes.
// It should panic if receives too small byte slice.
//
// Unmarshal parses an encoded value from bs, returning the value, the number of
// used bytes and any error encountered.
//
// Size method returns the number of bytes needed to encode the value.
//
// Skip skips an encoded value, returning the number of skipped bytes and any
// error encountered.
type Serializer[T any] interface {
	Marshal(t T, bs []byte) (n int)
	Unmarshal(bs []byte) (t T, n int, err error)
	Size(t T) (size int)
	Skip(bs []byte) (n int, err error)
}

// Marshaller is the interface for types that can marshal themselves into the
// MUS format.
type Marshaller interface {
	MarshalMUS(bs []byte) (n int)
	SizeMUS() (size int)
}

// MarshallerTyped is the interface for types that support typed MUS
// serialization, designed for use with the typed package.
type MarshallerTyped interface {
	MarshalTypedMUS(bs []byte) (n int)
	SizeTypedMUS() (size int)
}

// Marshal creates and returns a byte slice filled with the serialized data.
func Marshal(v Marshaller) (bs []byte) {
	bs = make([]byte, v.SizeMUS())
	v.MarshalMUS(bs)
	return
}

// MarshalTyped creates and returns a byte slice filled with the serialized
// data.
func MarshalTyped(v MarshallerTyped) (bs []byte) {
	bs = make([]byte, v.SizeTypedMUS())
	v.MarshalTypedMUS(bs)
	return
}
