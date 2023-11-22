# MUS Format Serializer
mus-go is a [MUS format](https://ymz-ncnk.medium.com/mus-serialization-format-21d7be309e8d)
extremely fast serializer with validation support for Golang. It supports out of 
order deserialization, zero allocation deserialization, can encode data 
structures such as a graph or a linked list, and also has a
[streaming version](https://github.com/mus-format/mus-stream-go).

# Contents
- [MUS Format Serializer](#mus-format-serializer)
- [Contents](#contents)
- [cmd-stream-go library](#cmd-stream-go-library)
- [Tests](#tests)
- [Benchmarks](#benchmarks)
- [How To Use](#how-to-use)
  - [varint Package](#varint-package)
  - [raw Package](#raw-package)
  - [ord (ordinary) Package](#ord-ordinary-package)
    - [Valid String](#valid-string)
    - [Slice](#slice)
    - [Valid Slice](#valid-slice)
    - [Map](#map)
  - [unsafe Package](#unsafe-package)
  - [pm (pointer mapping) package](#pm-pointer-mapping-package)
- [Structs Support](#structs-support)
  - [Valid Struct](#valid-struct)
- [Arrays Support](#arrays-support)
- [Data Type Metadata (DTM) Support](#data-type-metadata-dtm-support)
- [Data Versioning Support](#data-versioning-support)
- [Marshal/Unmarshal interfaces (or oneof feature)](#marshalunmarshal-interfaces-or-oneof-feature)
- [Out of Order Deserialization](#out-of-order-deserialization)
- [Zero Allocation Deserialization](#zero-allocation-deserialization)

# cmd-stream-go library
If you're looking for a client-server communication library that supports the 
MUS format, try [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go). It 
also has excellent performance.

# Tests
Test coverage is 100%.

# Benchmarks
[github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

# How To Use
Don't forget to visit [mus-examples-go](https://github.com/mus-format/mus-examples-go).

mus-go offers several encoding options, each of which is in a separate package.

## varint Package
Serializes all `uint` (`uint64`, `uint32`, `uint16`, `uint8`, `uint`), `int`, 
`float`, `byte` data types using Varint encoding. For example:
```go
package main

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
package main

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
More details about Varint and Raw encodings can be found in the 
[MUS format specification](https://github.com/mus-format/specification).
If in doubt, use Varint.

## ord (ordinary) Package
Supports the following data types: `bool`, `string`, `slice`, `map`, and 
pointers. The principle of serialization of these types is exactly the same as 
in the above examples. Let's consider the features.

### Valid String
When deserializing a string, you can set a limit on its length. This is done 
using the `ord.UnmarshalValidString()` function:
```go
package main

import (
  "errors"

  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go/ord"
)

func main() {
  var (
    ErrTooLongString                         = errors.New("too long string")
    maxLength        com.ValidatorFn[int] = func(length int) (err error) {
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
package main

import (
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

func main() {
  var (
    sl = []int{1, 2, 3, 4, 5}
    m  = mus.MarshallerFn[int](varint.MarshalInt) // Implementation of the 
    // mus.Marshaller interface for slice elements.
    u = mus.UnmarshallerFn[int](varint.UnmarshalInt) // Implementation of the
    // mus.Unmarshaller interface for slice elements.
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
the elements present in it, as well as a `Skipper`. If one of the validators 
returns an error, the rest of the data will be skipped, thanks to the `Skipper`.
All this is done using the `ord.UnmarshalValidSlice()` function:
```go
package main

import (
  "errors"

  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

func main() {
  var (
    ErrTooLongSlice    = errors.New("too long slice")
    ErrTooBigSliceElem = errors.New("too big slice elem")
    u                  = mus.UnmarshallerFn[int](varint.UnmarshalInt)
    sk                 = mus.SkipperFn(varint.SkipInt) // Implementation of the
    // mus.Skipper interface for the slice elements, may be nil, in which case 
    // a validation error will be returned immediately.
    maxLength com.ValidatorFn[int] = func(length int) (err error) {
      // Checks the length of the slice.
      if length > 5 {
        err = ErrTooLongSlice
      }
      return
    }
    vl com.ValidatorFn[int] = func(e int) (err error) {
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

### Map
All of the above about slice applies to map.

## unsafe Package
With this package you can get maximum performance. But be careful, it uses 
unsafe type conversion.

To a large extent, this warning applies to the `string` type. With this package, 
we can unmarshal a byte slice into a string, so that if we than change it, the 
string will change as well. This allows us to unmarshal strings very quickly. 
However, the slice in this case should be reused only after processing the 
result. There is no such behavior for other types.

Supports the following data types: `bool`, `string`, `byte`, and all `uint`, 
`int`, `float`.

## pm (pointer mapping) package
Unlike the `ord` package, `pm` encodes pointers with the Mapping pointer flag,
which is described in the 
[MUS format specification](https://github.com/mus-format/specification#format-features).
Thanks to this package, you can encode data structures such as a graph or a 
linked list (corresponding examples can be found at 
[mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/pm)).

# Structs Support
In fact, mus-go does not support structural data types. You will have to 
implement the `mus.Marshaller`, `mus.Unmarshaller`, `mus.Sizer` interfaces 
yourself, but it is not difficult at all. For example:
```go
package main

import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)
  
type Foo struct {
  a int
  b bool
  c string
}

// MarshalFoo implements the mus.Marshaller interface.
func MarshalFoo(v Foo, bs []byte) (n int) {
  n = varint.MarshalInt(v.a, bs)
  n += ord.MarshalBool(v.b, bs[n:])
  return n + ord.MarshalString(v.c, bs[n:])
}

// UnmarshalFoo implements the mus.Unmarshaller interface.
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
package main

import (
  "errors"

  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

// Continuation of the previous section.
var (
  ErrTooBigA                         = errors.New("too bid a")
  avl        com.ValidatorFn[int] = func(a int) (err error) {
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

func UnmarshalValidFoo(avl com.Validator[int], ask mus.Skipper, bs []byte) (
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

# Data Type Metadata (DTM) Support
[mus-dts-go](https://github.com/mus-format/mus-dts-go) provides DTM support.

# Data Versioning Support
[mus-dvs-go](https://github.com/mus-format/mus-dvs-go) provides data versioning 
support. Using mus-dvs-go imposes almost no restrictions. In the new version of 
the data, you can change the field type, remove a field, and generally do 
anything you want as long as you can migrate from one version to another.

# Marshal/Unmarshal interfaces (or oneof feature)
You should read the [mus-dts-go](https://github.com/mus-format/mus-dts-go)
documentation first.

A simple example:
```go
// Interface to Marshal/Unmarshal.
type Instruction interface {...}

// Copy implements the Instruction interface.
type Copy struct {...}

// Insert implements the Instruction interface.
type Insert struct {...}

var (
  CopyDTS = ...
  InsertDTS = ...
)

// With help of the type switch and regular switch we can implement 
// Marshal/Unmarshal/Size functions for the Instruction interface.

func MarshalInstructionMUS(instr Instruction, bs []byte) (n int) {
  switch in := instr.(type) {
  case Copy:
    return CopyDTS.MarshalMUS(in, bs)
  case Insert:
    return InsertDTS.MarshalMUS(in, bs)
  default:
    panic(ErrUnexpectedInstructionType)
  }
}

func UnmarshalInstructionMUS(bs []byte) (instr Instruction, n int, err error) {
  dtm, n, err := dts.UnmarshalDTMUS(bs)
  if err != nil {
    return
  }
  switch dtm {
  case CopyDTM:
    return CopyDTS.UnmarshalDataMUS(bs[n:])
  case InsertDTM:
    return InsertDTS.UnmarshalDataMUS(bs[n:])
  default:
    err = ErrUnexpectedDTM
    return
  }
}

func SizeInstructionMUS(instr Instruction) (size int) {
  switch in := instr.(type) {
  case Copy:
    return CopyDTS.SizeMUS(in)
  case Insert:
    return InsertDTS.SizeMUS(in)
  default:
    panic(ErrUnexpectedInstructionType)
  }
}
```
A full example you can find at [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/oneof).
Take a note, nothing will stop you to Marshal/Unmarshal, for example, a slice of
Instructions.

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

# Zero Allocation Deserialization
You can achieve this using `bool`, `byte`, all `uint`, `int`, `float` types and
unsafe package. Please note that the length of variable-length data types 
(such as `string`, `slice` or `map`) is encoded using Varint encoding.