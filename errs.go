package mus

import "errors"

// ErrTooSmallByteSlice means that an Unmarshal requires a longer byte slice
// than was provided.
var ErrTooSmallByteSlice = errors.New("too small byte slice")
