// Package arrops provides options for customizing array serialization.
package arrops

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Options for the array serializer.
type Options[T any] struct {
	LenSer mus.Serializer[int]
	ElemVl com.Validator[T]
}

type SetOption[T any] func(o *Options[T])

func WithLenSer[T any](lenSer mus.Serializer[int]) SetOption[T] {
	return func(o *Options[T]) { o.LenSer = lenSer }
}

func WithElemValidator[T any](elemVl com.Validator[T]) SetOption[T] {
	return func(o *Options[T]) { o.ElemVl = elemVl }
}

func Apply[T any](ops []SetOption[T], o *Options[T]) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
