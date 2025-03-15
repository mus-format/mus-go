package slops

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Options for the slice serializer.
type Options[T any] struct {
	LenSer mus.Serializer[int]
	LenVl  com.Validator[int]
	ElemVl com.Validator[T]
}

type SetOption[T any] func(o *Options[T])

func WithLenSer[T any](lenSer mus.Serializer[int]) SetOption[T] {
	return func(o *Options[T]) { o.LenSer = lenSer }
}

func WithLenValidator[T any](lenVl com.Validator[int]) SetOption[T] {
	return func(o *Options[T]) { o.LenVl = lenVl }
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
