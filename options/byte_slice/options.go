// Package bslopts provides options for customizing byte slice serialization.
package bslopts

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Options for the byte slice serializer.
type Options struct {
	LenSer mus.Serializer[int]
	LenVl  com.Validator[int]
}

type SetOption func(o *Options)

func WithLenSer(lenSer mus.Serializer[int]) SetOption {
	return func(o *Options) { o.LenSer = lenSer }
}

func WithLenValidator(lenVl com.Validator[int]) SetOption {
	return func(o *Options) { o.LenVl = lenVl }
}

func Apply(opts []SetOption, o *Options) {
	for i := range opts {
		if opts[i] != nil {
			opts[i](o)
		}
	}
}
