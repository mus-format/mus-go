# MUS Format Serializer
mus-go is a [MUS format](https://github.com/mus-format/mus) extremely fast 
serializer with validation support for Golang. Also it supports out of order 
deserialization and zero allocations deserialization.

# Tests
Test coverage is 100%.

# Benchmarks
[github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

# Versioning
[Go to the MUS documentation](https://github.com/mus-format/mus#versioning).

# How To Use
Don't forget to visit [mus-examples-go](https://github.com/mus-format/mus-examples-go).

mus-go offers several encoding options, all of which are located in separate 
packages.

## varint Package
Serializes all `uint` (`uint64`, `uint32`, `uint16`, `uint8`, `uint`), `int`, 
`float`, `byte` data types using Varint encoding. For example:
```go
import "github.com/mus-format/mus-go/varint"

func main() {
  var (
    num  = 1000
    size = varint.SizeInt(num) // The number of bytes required to serialize a
    // given num.
    bs = make([]byte, size)
  )
  n := varint.MarshalInt(num, bs)        // Returns the number of used bytes.
  num, n, err := varint.UnmarshalInt(bs) // In addition to the num, it returns
  // the number of used bytes and an error.
  // ...
}
```

## raw Package
Serializes the same `uint`, `int`, `float`, `byte` data types using Raw 
encoding. For example:
```go
import "github.com/mus-format/mus-go/raw"

func main() {
  var (
    num = 1000
    size = raw.SizeInt(num)
    bs  = make([]byte, size)
  )
  n := raw.MarshalInt(num, bs)
  num, n, err := raw.UnmarshalInt(bs)
  // ...
}
```
More details about Varint and Raw encoding can be found in the 
[MUS Serialization Format](https://github.com/mus-format/mus) documentation.
If in doubt, use Varint.

## ord (ordinary) Package
Supports the following data types: `bool`, `string`, `slice`, `map`, and 
pointers. The principle of serialization of these types is exactly the same as 
in the above examples. Let's consider the features.

### Valid String
When deserializing a string, you can set a limit on its length. This is done 
using the `ord.UnmarshalValidString()` function:
```go
import (
  "errors"

  muscom "github.com/mus-format/mus-common-go"
  "github.com/mus-format/mus-go/ord"
)

func main() {
  var (
    ErrTooLongString                         = errors.New("too long string")
    maxLength        muscom.ValidatorFn[int] = func(length int) (err error) {
      // Checks the length of the string.
      if length > 10 {
        err = ErrTooLongString
      }
      return
    }
  )
  // ...
  str, n, err := ord.UnmarshalValidString(maxLength, bs)
  // If the encoded string str does not meet the requirements of the validator,
  // then all bytes belonging to it will be skipped, that is, n will be equal to
  // SizeString(str).
  // ...
}
```

### Slice
```go
import (
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

func main() {
  var (
    sl = []int{1, 2, 3, 4, 5}
    m  = mus.MarshalerFn[int](varint.MarshalInt) // Implementation of the 
    // mus.Marshaler interface for slice elements.
    u = mus.UnmarshalerFn[int](varint.UnmarshalInt) // Implementation of the
    // mus.Unmarshaler interface for slice elements.
    s = mus.SizerFn[int](varint.SizeInt) // Implementation of the mus.Sizer
    // interface for slice elements.
    size = ord.SizeSlice[int](sl, s)
    bs   = make([]byte, size)
  )
  n := ord.MarshalSlice[int](sl, m, bs)
  sl, n, err := ord.UnmarshalSlice[int](u, bs)
  // ...
}
```

### Valid Slice
When deserializing a slice, we can set a limit on its length, a restriction on 
the elements present in it, as well as `Skipper`. If one of the validators 
returns an error, the rest of the data will be skipped, thanks to `Skipper`.
All this is done using the `ord.UnmarshalValidSlice()` function:
```go
import (
  "errors"

  muscom "github.com/mus-format/mus-common-go"
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

func main() {
  var (
    ErrTooLongSlice    = errors.New("too long slice")
    ErrTooBigSliceElem = errors.New("too big slice elem")
    u                  = mus.UnmarshalerFn[int](varint.UnmarshalInt)
    sk                 = mus.SkipperFn(varint.SkipInt) // Implementation of the
    // mus.Skipper interface for slice elements, may be nil, in which case a
    // validation error will be returned immediately.
    maxLength muscom.ValidatorFn[int] = func(length int) (err error) {
      // Checks the length of the slice.
      if length > 5 {
        err = ErrTooLongSlice
      }
      return
    }
    vl muscom.ValidatorFn[int] = func(e int) (err error) {
      // Checks the slice elements.
      if e > 10 {
        err = ErrTooBigSliceElem
      }
      return
    }
  )
  // ...
  sl, n, err := ord.UnmarshalValidSlice[int](maxLength, u, vl, sk, bs)
  // ...
}
```

### Map and Valid Map
The story is the same as for slice.

## Unsafe Package
With this package you can get maximum performance. But be careful, it uses 
unsafe type conversion.
Supports the following data types: `bool`, `string`, `byte`, and all `uint`, 
`int`, `float`.

# Structs Support
In fact, mus-go does not support structural data types. You will have to 
implement the `mus.Marhsaler`, `mus.Unmarshaler`, `mus.Sizer` interfaces 
yourself, but it is not difficult at all. For example:
```go
import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)
  
type Foo struct {
  a int
  b bool
  c string
}

// MarshalFoo implements the mus.Marshaler interface.
func MarshalFoo(v Foo, bs []byte) (n int) {
  n = varint.MarshalInt(v.a, bs)
  n += ord.MarshalBool(v.b, bs[n:])
  return n + ord.MarshalString(v.c, bs[n:])
}

// UnmarshalFoo implements the mus.Unmarshaler interface.
func UnmarshalFoo(bs []byte) (v Foo, n int, err error) {
  v.a, n, err = varint.UnmarshalInt(bs)
  if err != nil {
    return
  }
  var n1 int
  v.b, n1, err = ord.UnmarshalBool(bs[n:])
  n += n1
  if err != nil {
    return
  }
  v.c, n1, err = ord.UnmarshalString(bs[n:])
  n += n1
  return
}

// SizeFoo implements the mus.Sizer interface.
func SizeFoo(v Foo) (size int) {
  size += varint.SizeInt(v.a)
  size += ord.SizeBool(v.b)
  return size + ord.SizeString(v.c)
}

// SkipFoo implements the mus.Skipper interface.
func SkipFoo(bs []byte) (n int, err error) {
  n, err = varint.SkipInt(bs)
  if err != nil {
    return
  }
  var n1 int
  n1, err = ord.SkipBool(bs[n:])
  n += n1
  if err != nil {
    return
  }
  n1, err = ord.SkipString(bs[n:])
  n += n1
  return
}
```
All you have to do is deconstruct the structure into simpler data types and 
choose the desired encoding for each. Of course, this requires some effort.
But first of all, this code can be generated, secondly, this approach provides 
more flexibility, and thirdly, mus-go remains quite simple, which makes it easy 
to implement for other programming languages.

## Valid Struct
Also, thanks to this approach, we can very quickly find out whether the decoded 
structure is suitable for us or not. And we don't even need to deserialize it 
completely! For example:
```go
import (
  "errors"

  muscom "github.com/mus-format/mus-common-go"
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

// Continuation of the previous section.
var (
  ErrTooBigA                         = errors.New("too bid a")
  avl        muscom.ValidatorFn[int] = func(a int) (err error) {
    // Checks if the Foo.a field is correct.
    if a > 10 {
      err = ErrTooBigA
    }
    return
  }
  ask mus.SkipperFn = func(bs []byte) (n int, err error) { // Skipper for the
    // Foo.a field, skips all subsequent Foo fields.
    n, err = ord.SkipBool(bs)
    if err != nil {
      return
    }
    var n1 int
    n1, err = ord.SkipString(bs[n:])
    n += n1
    return
  }
)

func UnmarshalValidFoo(avl muscom.Validator[int], ask mus.Skipper, bs []byte) (
  v Foo, n int, err error) {
  v.a, n, err = varint.UnmarshalInt(bs)
  if err != nil {
    return
  }
  var (
    n1   int
    err1 error
  )
  if avl != nil {
    err = avl.Validate(v.a)
    if err != nil {
      if ask != nil { // If Skipper != nil, applies it, otherwise returns a
        // validation error immediately.
        n1, err1 = ask.SkipMUS(bs[n:])
        n += n1
        if err1 != nil {
          err = err1
        }
      }
      return
    }
  }
  v.b, n1, err = ord.UnmarshalBool(bs[n:])
  n += n1
  if err != nil {
    return
  }
  v.c, n1, err = ord.UnmarshalString(bs[n:])
  n += n1
  return
}
```

# Arrays Support
Unfortunately, Golang does not support generic parameterization of array sizes.
Therefore, to serialize an array, you must first make a slice of it. Or for 
greater performance, you can implement `Marshal`, `Unmarshal`, ... functions for 
it yourself. This is not very difficult to do, you can also refer to the 
[ord/slice.go](ord/slice.go) file for an example.

# Out of Order Deserialization
A simple example:
```go
package main

import (
  "fmt"

  "github.com/mus-format/mus-go/varint"
)

func main() {
  // We encode three numbers in turn - 5, 10, 15.
  bs := make([]byte, varint.SizeInt(5)+varint.SizeInt(10)+varint.SizeInt(15))
  n1 := varint.MarshalInt(5, bs)
  n2 := varint.MarshalInt(10, bs[n1:])
  varint.MarshalInt(15, bs[n1+n2:])

  // Get them back in the opposite direction. Errors are omitted for simplicity.
  n1, _ = varint.SkipInt(bs)
  n2, _ = varint.SkipInt(bs)
  num, _, _ := varint.UnmarshalInt(bs[n1+n2:])
  fmt.Println(num)
  num, _, _ = varint.UnmarshalInt(bs[n1:])
  fmt.Println(num)
  num, _, _ = varint.UnmarshalInt(bs)
  fmt.Println(num)
  // The output will be:
  // 15
  // 10
  // 5
}
```

# Zero Allocations Deserialization
You can achieve this using `bool`, `byte`, all `uint`, `int`, `float` types and
unsafe package. Please note that the length of variable-length data types 
(such as `string`, `slice` or `map`) is encoded using Varint encoding.