// Package typed provides DTM (Data Type Metadata) support for mus-go
// serializer. It wraps a type serializer together with a DTM value,
// enabling typed data serialization.
package typed

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Ser implements the mus.Serializer interface and provides DTM support for the
// mus-go serializer. It helps to serializer DTM + data.
type Ser[T any] struct {
	dtm com.DTM
	ser mus.Serializer[T]
}

// NewSer creates a new TypedSer.
func NewSer[T any](dtm com.DTM, ser mus.Serializer[T]) Ser[T] {
	return Ser[T]{dtm, ser}
}

// DTM returns the initialization value.
func (d Ser[T]) DTM() com.DTM {
	return d.dtm
}

// Marshal marshals DTM + data.
func (d Ser[T]) Marshal(t T, bs []byte) (n int) {
	n = DTMSer.Marshal(d.dtm, bs)
	n += d.ser.Marshal(t, bs[n:])
	return
}

// Unmarshal unmarshals DTM + data.
//
// Returns com.WrongDTMError if the unmarshalled DTM differs from the expected
// one.
func (d Ser[T]) Unmarshal(bs []byte) (t T, n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if dtm != d.dtm {
		err = com.NewWrongDTMError(d.dtm, dtm)
		return
	}
	var n1 int
	t, n1, err = d.UnmarshalData(bs[n:])
	n += n1
	return
}

// Size calculates the size of the DTM + data.
func (d Ser[T]) Size(t T) (size int) {
	size = DTMSer.Size(d.dtm)
	return size + d.ser.Size(t)
}

// Skip skips DTM + data.
//
// Returns com.WrongDTMError if the unmarshalled DTM differs from the expected
// one.
func (d Ser[T]) Skip(bs []byte) (n int, err error) {
	dtm, n, err := DTMSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if dtm != d.dtm {
		err = com.NewWrongDTMError(d.dtm, dtm)
		return
	}
	var n1 int
	n1, err = d.ser.Skip(bs[n:])
	n += n1
	return
}

// UnmarshalData unmarshals only data.
func (d Ser[T]) UnmarshalData(bs []byte) (t T, n int, err error) {
	return d.ser.Unmarshal(bs)
}

// SkipData skips only data.
func (d Ser[T]) SkipData(bs []byte) (n int, err error) {
	return d.ser.Skip(bs)
}
