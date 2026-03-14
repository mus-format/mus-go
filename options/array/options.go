// Package arropts provides options for customizing array serialization.
package arropts

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

func Apply[T any](opts []SetOption[T], o *Options[T]) {
	for i := range opts {
		if opts[i] != nil {
			opts[i](o)
		}
	}
}
