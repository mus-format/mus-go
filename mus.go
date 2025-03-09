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
