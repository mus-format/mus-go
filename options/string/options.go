// Package strops provides options for customizing string serialization.
package strops

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Options for the string serializer.
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

func Apply(ops []SetOption, o *Options) {
	for i := range ops {
		if ops[i] != nil {
			ops[i](o)
		}
	}
}
