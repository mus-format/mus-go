# mus-go Serializer

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-go)](https://goreportcard.com/report/github.com/mus-format/mus-go)
[![codecov](https://codecov.io/gh/mus-format/mus-go/graph/badge.svg?token=WLLZ1MMQDX)](https://codecov.io/gh/mus-format/mus-go)

mus-go is a [MUS format](https://medium.com/p/21d7be309e8d) serializer. However,
due to its minimalist design and a wide range of serialization primitives, it 
can also be used to implement other binary serialization formats 
([here](https://github.com/mus-format/mus-examples-go/blob/main/protobuf/main.go)
is an example where mus-go is utilized to implement Protobuf encoding).

To get started quickly, go to the [code generator](https://github.com/mus-format/musgen-go) page.

## Why mus-go?
It is lightning fast and space efficient.

## Brief mus-go Description
- Has a [streaming version](https://github.com/mus-format/mus-stream-go).
- Can run on both 32 and 64-bit systems.
- Variable-length data types (like `string`, `slice`, or `map`) are encoded as: 
  `length + data`. You can choose binary representation for both of these parts. 
- Supports data versioning.
- Deserialization may fail with one of the following errors: `ErrOverflow`, 
  `ErrNegativeLength`, `ErrTooSmallByteSlice`, `ErrWrongFormat`.
- Can validate and skip data while unmarshalling.
- Supports pointers.
- Can encode data structures such as graphs or linked lists.
- Supports oneof feature.
- Supports private fields.
- Supports out-of-order deserialization.
- Supports zero allocation deserialization.

# Contents
- [mus-go Serializer](#mus-go-serializer)
  - [Why mus-go?](#why-mus-go)
  - [Brief mus-go Description](#brief-mus-go-description)
- [Contents](#contents)
- [cmd-stream-go](#cmd-stream-go)
- [musgen-go](#musgen-go)
- [Benchmarks](#benchmarks)
- [How To Use](#how-to-use)
  - [varint Package](#varint-package)
  - [raw Package](#raw-package)
  - [ord (ordinary) Package](#ord-ordinary-package)
  - [unsafe Package](#unsafe-package)
  - [pm (pointer mapping) Package](#pm-pointer-mapping-package)
- [Structs Support](#structs-support)
- [Arrays Support](#arrays-support)
- [MarshallerMUS Interface](#marshallermus-interface)
- [Generic MarshalMUS Function](#generic-marshalmus-function)
- [Data Type Metadata (DTM) Support](#data-type-metadata-dtm-support)
- [Data Versioning](#data-versioning)
- [Marshal/Unmarshal interfaces (or oneof feature)](#marshalunmarshal-interfaces-or-oneof-feature)
- [Validation](#validation)
  - [String](#string)
  - [Slice](#slice)
  - [Map](#map)
  - [Struct](#struct)
- [Out of Order Deserialization](#out-of-order-deserialization)
- [Zero Allocation Deserialization](#zero-allocation-deserialization)

# cmd-stream-go
[cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) - high-performance
client-server library for Golang that implements the Command pattern. 
cmd-stream-go/MUS is about [3 times faster](https://github.com/ymz-ncnk/go-client-server-communication-benchmarks) 
than gRPC/Protobuf.

# musgen-go
Writing mus-go code manually can be quite tedious and error-prone. It's much 
better to use a [code generator](https://github.com/mus-format/musgen-go) that 
can produce `Marshal`, `Unmarshal`, `Size`, and `Skip` functions for you. Also 
it is very easy to use - just give it a type and call `Generate()`.

# Benchmarks
- [github.com/ymz-ncnk/go-serialization-benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks) - 
  contains the results of running serializers in different modes.
- [github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

Why did I write another [benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)?
Existing [benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) 
currently have some issues - just try to run them several times, you will most 
likely get different results, such that it's impossible to determine which 
serializer is faster. Having done so, I simply did not know which one to
choose. That was one of the reasons, and basically I made them for my own use.

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
    size = varint.SizeInt(num) // The number of bytes required to serialize num.
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
pointers. 

Variable-length data types (like `string`, `slice`, or `map`) are encoded as: 
`length + data`. You can choose binary representation for both of these parts. 
By default, the length is encoded using Varint without ZigZag (see 
`...PositiveInt()` functions from the varint package). In this case the 
maximum length is limited by the maximum value of the `int` type on your system.
This is ok for different architectures - an attempt to unmarshal, for example, 
too long string on a 32-bit system, will result in `ErrOverflow`.

Let's look at the serialization of the slice type:
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
    lenM mus.Marshaller // Length marshaller, if nil varint.MarshalPositiveInt() is used.
    m  = mus.MarshallerFn[int](varint.MarshalInt) // Implementation of the 
    // mus.Marshaller interface for slice elements.
    lenU mus.Unmarshaller // Length unmarshaller, if nil varint.UnmarshalPositiveInt() is used.
    u = mus.UnmarshallerFn[int](varint.UnmarshalInt) // Implementation of the
    // mus.Unmarshaller interface for slice elements.
    s = mus.SizerFn[int](varint.SizeInt) // Implementation of the mus.Sizer
    // interface for slice elements.
    size = ord.SizeSlice[int](sl, s)
    bs   = make([]byte, size)
  )
  ord.MarshalSlice[int](sl, lenM, m, bs)
  sl, n, err := ord.UnmarshalSlice[int](lenU, u, bs)
  // ...
}
```
Maps are serialized in the same way.

## unsafe Package
unsafe package provides maximum performance, but be careful it uses an unsafe 
type conversion.

This warning largely applies to the string type - modifying the byte slice after 
unmarshalling will also change the string’s contents. Here is an 
[example](https://github.com/mus-format/mus-examples-go/blob/main/unasafe/main.go) 
that demonstrates this behavior more clearly.

Supports the following data types: `bool`, `string`, `byte`, and all `uint`, 
`int`, `float`.

## pm (pointer mapping) Package
Let's consider the following struct:
```go
type TwoPtr struct {
  ptr1 *string
  ptr2 *string
}
```
With the `ord` package after unmarshal `twoPtr.ptr1 != twoPtr.ptr2`, and with 
the `pm` package, they will be equal. This feature allows to serialize data 
structures such as graphs or linked lists. Corresponding examples can be found 
at [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/pm).

# Structs Support
In fact, mus-go does not support structural data types, which means that you will
have to implement the `mus.Marshaller`, `mus.Unmarshaller`, `mus.Sizer` and
`mus.Skipper` interfaces yourself. But it's not difficult at all, for example:
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
  return n + ord.MarshalString(v.c, nil, bs[n:])
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
  v.c, n1, err = ord.UnmarshalString(nil, bs[n:])
  n += n1
  return
}

// SizeFoo implements the mus.Sizer interface.
func SizeFoo(v Foo) (size int) {
  size += varint.SizeInt(v.a)
  size += ord.SizeBool(v.b)
  return size + ord.SizeString(v.c, nil)
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
  n1, err = ord.SkipString(nil, bs[n:])
  n += n1
  return
}
```
All you have to do is deconstruct the structure into simpler data types and 
choose the desired encoding for each. Of course, this requires some effort.
But, firstly, this code can be generated, secondly, this approach provides 
more flexibility, and thirdly, mus-go remains quite simple, which makes it easy 
to implement for other programming languages.

# Arrays Support
Unfortunately, Golang does not support generic parameterization of array sizes. 
Therefore, to serialize an array - make a slice of it and use the `ord` package. 
Or, for better performance, implement the necessary `Marshal`, `Unmarshal`, ... 
functions, as done in the [ord/slice.go](ord/slice.go) file.

# MarshallerMUS Interface
It is often convenient to define the `MarshallerMUS` interface:
```go
type MarshallerMUS interface {
  MarshalMUS(bs []byte) (n int)
  SizeMUS() (size int)
}

// Foo implements the MarshallerMUS interface.
type Foo struct {...}

func (f Foo) MarshalMUS(bs []byte) (n int) {
  return MarshalFooMUS(f, bs) // or FooDTS.Marshal(f, bs)
}

func (f Foo) SizeMUS() (size int) {
  return SizeFooMUS(f) // or FooDTS.Size(f)
}
...
```

# Generic MarshalMUS Function
To define generic `MarshalMUS` function:
```go
package main 

// Define MarshallerMUS interface ...
type MarshallerMUS interface {
  MarshalMUS(bs []byte) (n int)
  SizeMUS() (size int)
}

// ... and the function itself.
func MarshalMUS[T MarshallerMUS](t T) (bs []byte) {
  bs = make([]byte, t.SizeMUS())
  t.MarshalMUS(bs)
  return
}

// Define a structure that implements the MarshallerMUS interface.
type Foo struct {...}
...

func main() {
  // Now the generic MarshalMUS function can be used like this.
  bs := MarshalMUS(Foo{...})
  // ...
}
```

The full code can be found [here](https://github.com/mus-format/mus-examples-go/tree/main/generic_marshal).

# Data Type Metadata (DTM) Support
[mus-dts-go](https://github.com/mus-format/mus-dts-go) provides DTM support.

# Data Versioning
mus-dts-go can be used to implement data versioning. [Here](https://github.com/mus-format/mus-examples-go/tree/main/versionings) is an example.

# Marshal/Unmarshal interfaces (or oneof feature)
You should read the [mus-dts-go](https://github.com/mus-format/mus-dts-go)
documentation first.

A simple example:
```go
// Interface to Marshal/Unmarshal.
type Instruction interface {...}

// Copy implements the Instruction and MarshallerMUS interfaces.
type Copy struct {...}

// Insert implements the Instruction and MarshallerMUS interfaces.
type Insert struct {...}

var (
  CopyDTS = ...
  InsertDTS = ...
)

func MarshalInstruction(instr Instruction, bs []byte) (n int) {
  if m, ok := instr.(MarshallerMUS); ok {
    return m.MarshalMUS(bs)
  }
  panic("instr doesn't implement the MarshallerMUS interface")
}

func UnmarshalInstruction(bs []byte) (instr Instruction, n int, err error) {
  dtm, n, err := dts.UnmarshalDTM(bs)
  if err != nil {
    return
  }
  switch dtm {
  case CopyDTM:
    return CopyDTS.UnmarshalData(bs[n:])
  case InsertDTM:
    return InsertDTS.UnmarshalData(bs[n:])
  default:
    err = ErrUnexpectedDTM
    return
  }
}

func SizeInstruction(instr Instruction) (size int) {
  if s, ok := instr.(MarshallerMUS); ok {
    return s.SizeMUS()
  }
  panic("instr doesn't implement the MarshallerMUS interface")
}
```
A full example can be found at [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/oneof).
Take a note, nothing will stop us to Marshal/Unmarshal, for example, a slice of 
interfaces.

# Validation
Validation is performed during deserialization.

## String
With the `ord.UnmarshalValidString()` function, you can validate the length of
a string.
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
    lenU mus.Unmarshaller // Length unmarshaller, if nil varint.UnmarshalPositiveInt() is used.
    lenVl        com.ValidatorFn[int] = func(length int) (err error) {  // Length validator.
      if length > 10 {
        err = ErrTooLongString
      }
      return
    }
    skip = true // If true and the encoded string does not meet the requirements
    // of the validator, all bytes belonging to it will be skipped, n will be 
    // equal to SizeString(str).
  )
  // ...
  str, n, err := ord.UnmarshalValidString(lenU, lenVl, skip, bs)
  // ...
}
```

## Slice
With the `ord.UnmarshalValidSlice()` function, you can validate the length and
elements of a slice. Also it provides an option to skip the rest of the data
if one of the validators returns an error.
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

    lenU mus.Unmarshaller // Length unmarshaller, if nil varint.UnmarshalPositiveInt() is used.
    lenVl com.ValidatorFn[int] = func(length int) (err error) { // Length validator.
      if length > 5 {
        err = ErrTooLongSlice
      }
      return
    }
    u                  = mus.UnmarshallerFn[int](varint.UnmarshalInt)
    vl com.ValidatorFn[int] = func(e int) (err error) { // Elements validator.
      if e > 10 {
        err = ErrTooBigSliceElem
      }
      return
    }
    sk                 = mus.SkipperFn(varint.SkipInt) // If nil, a validation 
    // error will be returned immediately. If != nil and one of the validators 
    // returns an error, will be used to skip the rest of the slice.
  )
  // ...
  sl, n, err := ord.UnmarshalValidSlice[int](lenU, lenVl, u, vl, sk, bs)
  // ...
}
```

## Map
Validation works in the same way as for the slice type.

## Struct
Unmarshalling an invalid structure can stop at the first invalid field with a 
validation error.
```go
package main

import (
  "errors"

  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

func UnmarshalValidFoo(vl com.Validator[int], bs []byte) (v Foo, n int, err error) {
  v.a, n, err = varint.UnmarshalInt(bs)
  if err != nil {
    return
  }
  // There is no need to deserialize the entire structure to find out that it is 
  // invalid.
  if err = vl.Validate(v.a); err != nil {
    err = fmt.Errorf("incorrect field 'a': %w", err)
    return // The rest of the structure remains unmarshaled.
  }
  // ...
}

// vl can be used to check Foo.a field.
var vl com.ValidatorFn[int] = func(n int) (err error) {
  if n > 10 {
    return errors.New("bigger than 10")
  }
  return
}
```

# Out of Order Deserialization
A simple example:
```go
package main

import (
  "fmt"

  "github.com/mus-format/mus-go/varint"
)

func main() {
  // Encode three numbers in turn - 5, 10, 15.
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
Can be achieved using the unsafe package.