package mus

import (
	"errors"

	com "github.com/mus-format/common-go"
)

// ErrTooSmallByteSlice means that an Unmarshal requires a longer byte slice
// than was provided.
var ErrTooSmallByteSlice = errors.New(com.ErrorPrefix + "too small byte slice")
