package mapops

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Options for the map serializer.
type Options[T, V any] struct {
	LenSer  mus.Serializer[int]
	LenVl   com.Validator[int]
	KeyVl   com.Validator[T]
	ValueVl com.Validator[V]
}

type SetOption[T, V any] func(o *Options[T, V])

func WithLenSer[T, V any](lenSer mus.Serializer[int]) SetOption[T, V] {
	return func(o *Options[T, V]) { o.LenSer = lenSer }
}

func WithLenValidator[T, V any](lenVl com.Validator[int]) SetOption[T, V] {
	return func(o *Options[T, V]) { o.LenVl = lenVl }
}

func WithKeyValidator[T, V any](keyVl com.Validator[T]) SetOption[T, V] {
	return func(o *Options[T, V]) { o.KeyVl = keyVl }
}

func WithValueValidator[T, V any](valVl com.Validator[V]) SetOption[T, V] {
	return func(o *Options[T, V]) { o.ValueVl = valVl }
}

func Apply[T, V any](ops []SetOption[T, V], o *Options[T, V]) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
