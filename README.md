# mus-go Serializer
mus-go is fully compatible with the [MUS serialization format](https://ymz-ncnk.medium.com/mus-serialization-format-21d7be309e8d). It is extremely fast and has validation 
support. Also thanks to the minimalist design and a wide range of serialization 
primitives, it can be used to implement other binary serialization formats 
([here](https://github.com/mus-format/mus-examples-go/blob/main/protobuf/main.go) 
is an example where mus-go is used to implement Protobuf encoding).

All of the uses described below produce the correct MUS encoding.

## Brief mus-go Description
- Has a [streaming version](https://github.com/mus-format/mus-stream-go).
- Can run on both 32 and 64-bit systems.
- Variable-length data types (like `string`, `slice`, or `map`) are encoded as: 
  `length + data`. You can choose binary representation for both of these parts. 
  By default, the length is encoded using Varint (actually, Varint without 
  ZigZag). In this case the maximum length is limited by the maximum value of 
  the `int` type on your system. This is ok for use on different architectures, 
  because, if, for example, we try to unmarshal too long string on a 32-bit 
  system, we will get `ErrOverflow`.
- Supports data versioning.
- If invalid data is encountered during deserialization, it returns one
  of the following errors: `ErrOverflow`, `ErrNegativeLength`, `ErrTooSmallByteSlice`, `ErrWrongFormat`.
- Can validate and skip data while unmarshalling.
- Supports pointers.
- Can encode data structures such as graphs or linked lists.
- Supports private fields.
- Supports oneof feature.
- Supports out-of-order deserialization.
- Supports zero allocation deserialization.

# Contents
- [mus-go Serializer](#mus-go-serializer)
  - [Brief mus-go Description](#brief-mus-go-description)
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
  - [pm (pointer mapping) Package](#pm-pointer-mapping-package)
- [Structs Support](#structs-support)
  - [Valid Struct](#valid-struct)
- [Arrays Support](#arrays-support)
- [Generic MarshalMUS Function](#generic-marshalmus-function)
- [Data Type Metadata (DTM) Support](#data-type-metadata-dtm-support)
- [Data Versioning Support](#data-versioning-support)
- [Marshal/Unmarshal interfaces (or oneof feature)](#marshalunmarshal-interfaces-or-oneof-feature)
- [Out of Order Deserialization](#out-of-order-deserialization)
- [Zero Allocation Deserialization](#zero-allocation-deserialization)

# cmd-stream-go library
[cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) - high-performance 
RCX (Remote Command eXecution) library for Golang, which also supports the MUS 
format.

# Tests
Test coverage is 100%.

# Benchmarks
- [github.com/ymz-ncnk/go-serialization-benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks) - contains the results of running serializers in 
  different modes.
- [github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

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
  // UnmarshalValidString accepts:
  // - length unmarshaller (if nil the default value is used)
  // - length validator
  // - skip flag
  // - bs
  str, n, err := ord.UnmarshalValidString(nil, maxLength, true, bs)
  // If skip flag == true and the encoded string str does not meet the 
  // requirements of the validator, then all bytes belonging to this string will
  // be skipped, that is, n will be equal to SizeString(str).
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
  // MarshalSlice accepts:
  // - slice
  // - slice length marshaller (if nil the default value is used)
  // - slice element marshaller
  // - bs
  n := ord.MarshalSlice[int](sl, nil, m, bs)
  // UnmarshalSlice accepts:
  // - slice length unmarshaller (if nil the default value is used)
  // - slice element unmarshaller
  // - bs
  sl, n, err := ord.UnmarshalSlice[int](nil, u, bs)
  // ...
}
```

### Valid Slice
When deserializing a slice, using the `ord.UnmarshalValidSlice()` function, we 
can set length and elements validators as well as `Skipper` that will skip the 
rest of the data if one of the validators returns an error:
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
  // UnmarshalValidSlice accepts:
  // - slice length unmarshaller (if nil the default value is used)
  // - slice length validator
  // - slice element unmarshaller
  // - slice element validator
  // - slice element skipper (if != nil and one of the validators returns an 
  //   error, it will be used to skip the rest of the slice)
  // - bs
  sl, n, err := ord.UnmarshalValidSlice[int](nil, maxLength, u, vl, sk, bs)
  // ...
}
```

### Map
All of the above about the slice type also applies to the map type.

## unsafe Package
You can get maximum performance with this package, but be careful it uses an 
unsafe type conversion.

To a large extent, this warning applies to the `string` type - if we change a 
byte slice, the string obtained from it will also change. In this case, we must 
first process the result, i.e. the string, and only then reuse the byte slice. 
For other types, there is no such behavior. Please visit this 
[example](https://github.com/mus-format/mus-examples-go/blob/main/unasafe/main.go), 
it tries to make things more clear.

Supports the following data types: `bool`, `string`, `byte`, and all `uint`, 
`int`, `float`.

## pm (pointer mapping) Package
Let's consider the following struct:
```go
package main

type TwoPtr struct {
  ptr1 *string
  ptr2 *string
}

func main() {
  str := "the same pointer in two fields"
  ptr := &str
  twoPtr := TwoPtr{
    ptr1: ptr,
    ptr2: ptr,
  }
  // ...
}
```
If we use the `ord` package to serialize this structure, then after unmarshal `twoPtr.ptr1 != twoPtr.ptr2`. But with `pm` package, these fields will be equal.
Unlike the `ord` package, `pm` encodes pointers with the Mapping pointer flag,
described in the [MUS format specification](https://github.com/mus-format/specification#format-features).
Also with its help, we can encode data structures such as graphs or 
linked lists (corresponding examples can be found at 
[mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/pm)).

# Structs Support
In fact, mus-go does not support structural data types, which means that we will
have to implement the `mus.Marshaller`, `mus.Unmarshaller` and `mus.Sizer` 
interfaces ourselves. But it's not difficult at all, for example:
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
  ErrTooBigA                         = errors.New("too big a")
  vl        com.ValidatorFn[int] = func(a int) (err error) {
    // Checks if the Foo.a field is correct.
    if a > 10 {
      err = ErrTooBigA
    }
    return
  }
  sk mus.SkipperFn = func(bs []byte) (n int, err error) { // Skipper for the
    // Foo.a field, skips all subsequent Foo fields.
    n, err = ord.SkipBool(bs)
    if err != nil {
      return
    }
    var n1 int
    n1, err = ord.SkipString(nil, bs[n:])
    n += n1
    return
  }
)

func UnmarshalValidFoo(vl com.Validator[int], sk mus.Skipper, bs []byte) (
  v Foo, n int, err error) {
  v.a, n, err = varint.UnmarshalInt(bs)
  if err != nil {
    return
  }
  var n1 int
  n1, err = validate(v.a, vl, sk, bs[n:])
  n += n1
  if err != nil {
    return
  }
  v.b, n1, err = ord.UnmarshalBool(bs[n:])
  n += n1
  if err != nil {
    return
  }
  v.c, n1, err = ord.UnmarshalString(nil, bs[n:])
  n += n1
  return
}

func validate(field int, vl com.Validator[int], sk mus.Skipper, bs []byte) (
  n int, err error) {
  var skErr error
  err = vl.Validate(field)
  if err != nil && sk != nil { // If Skipper != nil, applies it, otherwise
    // returns a validation error immediately.
    if n, skErr = sk.Skip(bs); skErr != nil {
      err = skErr
    }
  }
  return
}
```

# Arrays Support
Unfortunately, Golang does not support generic parameterization of array sizes. 
Therefore, to serialize an array, we must make a slice of it. Or, for better 
performance, we can implement the necessary `Marshal`, `Unmarshal`, ... 
functions ourselves, as done in the [ord/slice.go](ord/slice.go) file.

# Generic MarshalMUS Function
To define generic `MarshalMUS` function:
```go
package main 

// Define Marshaller interface
type MarshallerMUS[T any] interface {
	MarshalMUS(bs []byte) (n int)
	SizeMUS() (size int)
}

// and the function itself.
func MarshalMUS[T MarshallerMUS[T]](t T) (bs []byte) {
	bs = make([]byte, t.SizeMUS())
	t.MarshalMUS(bs)
	return
}

// Define a structure that implements the MarshallerMUS interface.
type Foo struct {...}

func (f Foo) MarshalMUS(bs []byte) (n int) {
	return MarshalFooMUS(f, bs)
}

func (f Foo) SizeMUS() (size int) {
	return SizeFooMUS(f)
}

func MarshalFooMUS(f Foo, bs []byte) (n int) {...}
func UnmarshalFooMUS(bs []byte) (f Foo, n int, err error)
func SizeFooMUS(f Foo) (size int) {...}
func SkipFooMUS(bs []byte) (n int, err error) {...}

func main() {
  // Now the generic MarshalMUS function can be used like this.
	bs := MarshalMUS(Foo{...})
  // ...
}
```

# Data Type Metadata (DTM) Support
[mus-dts-go](https://github.com/mus-format/mus-dts-go) provides DTM support.

# Data Versioning Support
[mus-dvs-go](https://github.com/mus-format/mus-dvs-go) provides data versioning 
support. 

Using mus-dvs-go imposes almost no restrictions - in the new version of the data
type, we can change the field type, remove a field, and generally do anything we
want as long as we can migrate from one version to another.

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

func MarshalInstruction(instr Instruction, bs []byte) (n int) {
  switch in := instr.(type) {
  case Copy:
    return CopyDTS.Marshal(in, bs)
  case Insert:
    return InsertDTS.Marshal(in, bs)
  default:
    panic(ErrUnexpectedInstructionType)
  }
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
  switch in := instr.(type) {
  case Copy:
    return CopyDTS.Size(in)
  case Insert:
    return InsertDTS.Size(in)
  default:
    panic(ErrUnexpectedInstructionType)
  }
}
```
A full example can be found at [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/oneof).
Take a note, nothing will stop us to Marshal/Unmarshal, for example, a slice of 
interfaces.

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