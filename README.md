# mus-go: A High-Performance, Flexible Binary Serialization Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-go)](https://goreportcard.com/report/github.com/mus-format/mus-go)
[![codecov](https://codecov.io/gh/mus-format/mus-go/graph/badge.svg?token=WLLZ1MMQDX)](https://codecov.io/gh/mus-format/mus-go)

**mus-go** is a powerful and versatile Go library for efficient binary
serialization.

While `mus-go` was built as a serializer for the [MUS format](https://medium.com/@ymz-ncnk/mus-serialization-format-20f833df12d5),
its minimalist architecture and broad set of serialization primitives also make
it well-suited for implementing other binary formats. Here you can find an
[example](https://github.com/mus-format/examples-go/blob/main/protobuf/main.go)
where it is used to encode data in Protobuf format.

A streaming version is also available: [mus-stream-go](https://github.com/mus-format/mus-stream-go).

To get started quickly, visit the [code generator](https://github.com/mus-format/musgen-go)
page.

## Why mus-go?

### Core Performance & Reliability

- Top-tier performance (see [benchmarks](#benchmarks)).
- Space-efficient data serialization.
- Robust and reliable.
- Cross-architecture compatible (32/64-bit systems).

### Advanced Capabilities

- Supports data versioning.
- Allows interface serialization (oneof feature).
- Comprehensive pointer support.
- Can encode graphs and linked lists.
- Offers zero allocation deserialization.

### Additional Features

- Enables validation and field skipping during unmarshalling.
- Supports private fields.
- Allows out-of-order deserialization.

## mus-go in Action: cmd-stream-go

Want to see it in action? Check out [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go)!
This library, based on the Command Pattern, enables efficient execution of
user-defined Commands on a server. The `cmd-stream/MUS` is about 3 times faster
than `gRPC/Protobuf`.

## Contents

- [mus-go: A High-Performance, Flexible Binary Serialization Library for Go](#mus-go-a-high-performance-flexible-binary-serialization-library-for-go)
  - [Why mus-go?](#why-mus-go)
    - [Core Performance \& Reliability](#core-performance--reliability)
    - [Advanced Capabilities](#advanced-capabilities)
    - [Additional Features](#additional-features)
  - [mus-go in Action: cmd-stream-go](#mus-go-in-action-cmd-stream-go)
  - [Contents](#contents)
  - [Code Generator](#code-generator)
  - [How To](#how-to)
    - [Packages](#packages)
      - [varint](#varint)
      - [raw](#raw)
      - [ord (ordinary)](#ord-ordinary)
        - [Slice](#slice)
      - [Map](#map)
      - [unsafe](#unsafe)
        - [Array](#array)
      - [pm (pointer mapping)](#pm-pointer-mapping)
    - [Structs Support](#structs-support)
    - [DTS (Data Type metadata Support)](#dts-data-type-metadata-support)
    - [Data Versioning](#data-versioning)
    - [MarshallerMUS Interface and MarshalMUS Function](#marshallermus-interface-and-marshalmus-function)
    - [Interface Serialization (oneof feature)](#interface-serialization-oneof-feature)
    - [Validation](#validation)
      - [String](#string)
      - [Slice](#slice-1)
      - [Map](#map-1)
      - [Struct](#struct)
  - [Out of Order Deserialization](#out-of-order-deserialization)
  - [Zero Allocation Deserialization](#zero-allocation-deserialization)
  - [Testing](#testing)
  - [Benchmarks](#benchmarks)

## Code Generator

Manually writing `mus-go` serialization code can be tedious and error-prone. The
[musgen-go](https://github.com/mus-format/musgen-go) code generator offers a
much more efficient and reliable alternative that's simple to use - just provide
a type and call `Generate()`.

## How To

To make a type serializable with `mus-go`, you need to implement the
[mus.Serializer](./mus.go) interface:

```go

import "github.com/mus-format/mus-go"

// YourTypeMUS is a MUS serializer for YourType.
var YourTypeMUS = yourTypeMUS{}

// yourTypeMUS implements the mus.Serializer interface.
type yourTypeMUS struct{}

func (s yourTypeMUS) Marshal(v YourType, bs []byte) (n int)              {...}
func (s yourTypeMUS) Unmarshal(bs []byte) (v YourType, n int, err error) {...}
func (s yourTypeMUS) Size(v YourType) (size int)                         {...}
func (s yourTypeMUS) Skip(bs []byte) (n int, err error)                  {...}
```

Then, you can use it as follows:

```go
var (
  value YourType = ...
  size = YourTypeMUS.Size(value) // The number of bytes required to serialize 
  // the value.
  bs = make([]byte, size)
)

n := YourTypeMUS.Marshal(value, bs) // Returns the number of used bytes.
value, n, err := YourTypeMUS.Unmarshal(bs) // Returns the value, the number of 
// used bytes and any error encountered.

// Instead of unmarshalling the value can be skipped:
n, err := YourTypeMUS.Skip(bs)
```

### Packages

`mus-go` offers several encoding options, each in a separate package.

#### varint

This package provides Varint serializers for all `uint` (e.g., `uint64`,
`uint32`, ...), `int`, `float`, and `byte` data types.

```go
package main

import "github.com/mus-format/mus-go/varint"

func main() {
  var (
    num  = 100
    size = varint.Int.Size(num)
    bs = make([]byte, size)
  )
  n := varint.Int.Marshal(num, bs)
  num, n, err := varint.Int.Unmarshal(bs)
  // ...
}
```

It also includes the `PositiveInt` serializer (Varint without ZigZag) for
efficiently encoding positive `int` values (negative values are supported as
well, though with reduced performance).

#### raw

This package contains Raw serializers for `byte`, `uint`, `int`, `float`, and
`time.Time` data types.

```go
package main

import "github.com/mus-format/mus-go/raw"

func main() {
  var (
    num = 100
    size = raw.Int.Size(num)
    bs  = make([]byte, size)
  )
  n := raw.Int.Marshal(num, bs)
  num, n, err := raw.Int.Unmarshal(bs)
  // ...
}
```

More details about Varint and Raw encodings can be found in the
[MUS format specification](https://github.com/mus-format/specification).
If in doubt, use Varint.

For `time.Time`, there are several serializers:

- `TimeUnix` – encodes a value as a Unix timestamp in seconds.
- `TimeUnixMilli` – encodes a value as a Unix timestamp in milliseconds.
- `TimeUnixMicro` – encodes a value as a Unix timestamp in microseconds.
- `TimeUnixNano` – encodes a value as a Unix timestamp in nanoseconds.

To ensure the deserialized `time.Time` value is in UTC, either set your TZ
environment variable to UTC (e.g., `os.Setenv("TZ", "")`) or use one of the
corresponding UTC serializers (e.g., `TimeUnixUTC`, `TimeUnixMilliUTC`).

#### ord (ordinary)

Contains serializers/constructors for `bool`, `string`, `byte slice`,
`slice`, `map`, and `pointer` types.

Variable-length data types (such as `string`, `slice`, and `map`) are
encoded as `length + data`, with customizable binary representations for both
parts. By default, the length is encoded using a Varint without ZigZag
(`varint.PositiveInt`), which limits the length to the maximum value of the
`int` type on your system. Such encoding works well across different
architectures. For example, an attempt to unmarshal a string that is too long
on a 32-bit system will result in an `ErrOverflow`.

For `slice`, and `map` types, only constructors are available to create
a concrete serializer.

##### Slice

```go
package main

import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
  slops "github.com/mus-format/mus-go/options/slice"
)

func main() {
  var (
    // varint.Int specifies the serializer for the slice's elements.
    ser = ord.NewSliceSer[int](varint.Int)

    // To create a slice serializer with the specific length serializer use:
    // ser = ord.NewSliceSer[int](varint.Int, slops.WithLenSer(lenSer))

    sl = []int{1, 2, 3}
    size = ser.Size(sl)
    bs = make([]byte, size)
  )
  n := ser.Marshal(sl, bs)
  sl, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### Map

```go
package main

import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
  mapops "github.com/mus-format/mus-go/options/map"
)
func main() {
  var (
    // varint.Int specifies the serializer for the map’s keys, and ord.String -
    // the serializer for the map’s values.
    ser = ord.NewMapSer[int, string](varint.Int, ord.String)

    // To create a map serializer with the specific length serializer use:
    // ser = ord.NewMapSer[int, string](varint.Int, ord.String, mapops.WithLenSer(lenSer))

    m    = map[int]string{1: "one", 2: "two", 3: "three"}
    size = ser.Size(m)
    bs   = make([]byte, size)
  )
  n := ser.Marshal(m, bs)
  m, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### unsafe

The unsafe package provides maximum performance, but be careful, it uses an
unsafe type conversion. This warning largely applies to the string type, because
modifying the byte slice after unmarshalling will also change the string’s
contents. Here is an [example](https://github.com/mus-format/examples-go/blob/main/unasafe/main.go)
that demonstrates this behavior more clearly.

Provides serializers for the following data types: `byte`, `bool`, `string`,
`array`, `byte slice`, `time.Time` and all `uint`, `int`, `float`.

##### Array

Unfortunately, Go does not support generic parameterization of array sizes,
as a result, the array serializer constructor looks like:

```go
package main

import (
  "github.com/mus-format/mus-go/unsafe"
  "github.com/mus-format/mus-go/varint"
  arrops "github.com/mus-format/mus-go/options/array"
)

func main() {
   var (
    // The first type parameter of the NewArraySer function represents the 
    // array type, the second represents the type of the array’s elements.
    //
    // As for the function parameters, varint.Int specifies the serializer for 
    // the array’s elements.
    ser = unsafe.NewArraySer[[3]int, int](varint.Int)

    // To create an array serializer with the specific length serializer use:
    // ser = unsafe.NewArraySer[[3]int, int](varint.Int, arrops.WithLenSer(lenSer))

    arr  = [3]int{1, 2, 3}
    size = ser.Size(arr)
    bs   = make([]byte, size)
  )
  n := ser.Marshal(arr, bs)
  arr, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### pm (pointer mapping)

Let's consider two pointers initialized with the same value:

```go
var (
  str = "hello world"
  ptr = &str

  ptr1 *string = ptr
  ptr2 *string = ptr
)
```

The `pm` package preserves pointer equality after unmarshalling `ptr1 == ptr2`,
while the `ord` package does not. This capability enables the serialization of
data structures like graphs or linked lists. You can find corresponding examples
in [examples-go](https://github.com/mus-format/examples-go/tree/main/pm).

### Structs Support

`mus-go` doesn’t support structural data types out of the box, which means
you’ll need to implement the `mus.Serializer` interface yourself. But that’s not
difficult at all. For example:

```go
package main

import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
)

// We will implement the FooMUS serializer for this struct.
type Foo struct {
  str string
  sl  []int
}

// Serializers.
var (
  FooMUS = fooMUS{}

  // IntSliceMUS is used by the FooMUS serializer.
  IntSliceMUS = ord.NewSliceSer[int](varint.Int)
)

// fooMUS implements the mus.Serializer interface.
type fooMUS struct{}

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  n = ord.String.Marshal(v.str, bs)
  return n + IntSliceMUS.Marshal(v.sl, bs[n:])
}

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
  v.str, n, err = ord.String.Unmarshal(bs)
  if err != nil {
    return
  }
  var n1 int
  v.sl, n1, err = IntSliceMUS.Unmarshal(bs[n:])
  n += n1
  return
}

func (s fooMUS) Size(v Foo) (size int) {
  size += ord.String.Size(v.str)
  return size + IntSliceMUS.Size(v.sl)
}

func (s fooMUS) Skip(bs []byte) (n int, err error) {
  n, err = ord.String.Skip(bs)
  if err != nil {
    return
  }
  var n1 int
  n1, err = IntSliceMUS.Skip(bs[n:])
  n += n1
  return
}
```

All you have to do is deconstruct the structure into simpler data types and
choose the desired encoding for each. Of course, this requires some effort.
But, firstly, the code can be generated, secondly, this approach provides
greater flexibility, and thirdly, `mus-go` stays quite simple, making it easy to
implement in other programming languages.

### DTS (Data Type metadata Support)

[dts-go](https://github.com/mus-format/dts-go) enables typed data
serialization using [Data Type Metadata (DTM)](https://medium.com/p/21d7be309e8d).

### Data Versioning

`dts-go` can also be used to implement data versioning. See this [example](https://github.com/mus-format/examples-go/tree/main/versioning).

### MarshallerMUS Interface and MarshalMUS Function

It is often convenient to use the `MarshallerMUS` interface:

```go
type MarshallerMUS interface {
  MarshalMUS(bs []byte) (n int)
  SizeMUS() (size int)
}
```

and `MarshalMUS` function:

```go
func MarshalMUS(v MarshallerMUS) (bs []byte) {
  bs = make([]byte, v.SizeMUS())
  v.MarshalMUS(bs)
  return
}

// Foo implements the MarshallerMUS interface.
type Foo struct {...}
...

func main() {
  // Foo can now be marshalled with a single function call.
  bs := MarshalMUS(Foo{...})
  // ...
}
```

They are already defined in the [ext-go](https://github.com/mus-format/ext-go)
module, which also includes the `MarshallerTypedMUS` interface and the
`MarshalTypedMUS` function for typed data serialization (DTM + data).

The full code of using `MarshalMUS` function can be found [examples-go](https://github.com/mus-format/examples-go/tree/main/marshal_func).

### Interface Serialization (oneof feature)

`dts-go` will also help to create a serializer for an interface. Example:

```go
import (
  "github.com/mus-format/dts-go"
  "github.com/mus-format/ext-go"
)

// Interface to serialize.
type Instruction interface {...}

// Copy implements the Instruction and ext.MarshallerTypedMUS interfaces.
type Copy struct {...}

// MarshalTypedMUS uses CopyDTS.
func (c Copy) MarshalTypedMUS(bs []byte) (n int) {
  return CopyDTS.Marshal(c, bs)
}

// SizeTypedMUS uses CopyDTS.
func (c Copy) SizeTypedMUS() (size int) {
  return CopyDTS.Size(c)
}

// Insert implements the Instruction and ext.MarshallerTypedMUS interfaces.
type Insert struct {...}

// ...

// instructionMUS implements the mus.Serializer interface.
type instructionMUS struct {}

func (s instructionMUS) Marshal(i Instruction, bs []byte) (n int) {
  if m, ok := i.(ext.MarshallerTypedMUS); ok {
    return m.MarshalTypedMUS(bs)
  }
  panic(fmt.Sprintf("%v doesn't implement ext.MarshallerTypedMUS interface", 
    reflect.TypeOf(i)))
}

func (s instructionMUS) Unmarshal(bs []byte) (i Instruction, n int, err error) {
  dtm, n, err := dts.DTMSer.Unmarshal(bs)
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

func (s instructionMUS) Size(i Instruction) (size int) {
  if s, ok := i.(ext.MarshallerTypedMUS); ok {
    return s.SizeTypedMUS()
  }
  panic(fmt.Sprintf("%v doesn't implement ext.MarshallerTypedMUS interface", 
    reflect.TypeOf(i)))
}
```

A full example can be found at [examples-go](https://github.com/mus-format/examples-go/tree/main/oneof).

### Validation

Validation is performed during unmarshalling. Validator is just a function
with the following signature `func (value Type) error`, where `Type` is a type
of the value to which the validator is applied.

#### String

`ord.NewValidStringSer` constructor creates a string serializer with the length
validator.

```go
package main

import (
  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go/ord"
  strops "github.com/mus-format/mus-go/options/string"
)

func main() {
  var (
    // Length validator.
    lenVl = func(length int) (err error) {
      if length > 3 {
        err = com.ErrTooLargeLength
      }
      return
    }
    ser = ord.NewValidStringSer(strops.WithLenValidator(com.ValidatorFn[int](lenVl)))

    // To create a valid string serializer with the specific length serializer
    // use:
    // ser = ord.NewValidStringSer(strops.WithLenSer(lenSer), ...)

    value = "hello world"
    size  = ser.Size(value)
    bs    = make([]byte, size)
  )
  n := ser.Marshal(value, bs)
  // Unmarshalling stops when a validator returns an error. As a result, in
  // this case, we will receive a length validation error.
  value, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### Slice

`ord.NewValidSliceSer` constructor creates a valid slice serializer with the
length and element validators.

```go
package main

import (
  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go/ord"
  slops "github.com/mus-format/mus-go/options/slice"
)

func main() {
  var (
    // Length validator.
    lenVl = func(length int) (err error) {
      if length > 3 {
        err = com.ErrTooLargeLength
      }
      return
    }
    // Element validator.
    elemVl = func(elem string) (err error) {
      if elem == "hello" {
        err = ErrBadElement
      }
      return
    }
    // Each of the validators could be nil.
    ser = ord.NewValidSliceSer[string](ord.String,
      slops.WithLenValidator[string](com.ValidatorFn[int](lenVl)),
      slops.WithElemValidator[string](com.ValidatorFn[string](elemVl)))

    // To create a valid slice serializer with the specific length serializer
    // use:
    // ser = ord.NewValidSliceSer[string](ord.String,
    //   slops.WithLenSer[string](lenSer), ...)

    value = []string{"hello", "world"}
    size  = ser.Size(value)
    bs    = make([]byte, size)
  )
  n := ser.Marshal(value, bs)
  // Unmarshalling stops when any of the validators return an error. As a
  // result, in this case, we will receive an element validation error.
  value, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### Map

`ord.NewValidMapSer` constructor creates a valid map serializer with the
length, key and value validators.

```go
package main

import (
  com "github.com/mus-format/common-go"
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
  mapops "github.com/mus-format/mus-go/options/map"
)

func main() {
  var (
    // Length validator.
    lenVl = func(length int) (err error) {
      if length > 3 {
        err = com.ErrTooLargeLength
      }
      return
    }
    // Key validator.
    keyVl = func(key int) (err error) {
      if key == 1 {
        err = ErrBadKey
      }
      return
    }
    // Value validator.
    valueVl = func(val string) (err error) {
      if val == "hello" {
        err = ErrBadValue
      }
      return
    }
    // Each of the validators could be nil.
    ser = ord.NewValidMapSer[int, string](varint.Int, ord.String,
      mapops.WithLenValidator[int, string](com.ValidatorFn[int](lenVl)),
      mapops.WithKeyValidator[int, string](com.ValidatorFn[int](keyVl)),
      mapops.WithValueValidator[int, string](com.ValidatorFn[string](valueVl)))

    // To create a valid map serializer with the specific length serializer
    // use:
    // ser = ord.NewValidMapSer[int, string](varint.Int, ord.String,
    //   mapops.WithLenSer[int, string](lenSer), ...)

    value = map[int]string{1: "hello", 2: "world"}
    size  = ser.Size(value)
    bs    = make([]byte, size)
  )
  n := ser.Marshal(value, bs)
  // Unmarshalling stops when any of the validators return an error. As a
  // result, in this case, we will receive a key validation error.
  value, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### Struct

Unmarshalling an invalid structure may stop at the first invalid field,
returning a validation error.

```go
package main

import "github.com/mus-format/mus-go/varint"

type fooMUS struct{}

// ...

func (s fooMUS) Unmarshal(bs []byte) (v Foo, n int, err error) {
  // Unmarshal the first field.
  v.str, n, err = ord.String.Unmarshal(bs)
  if err != nil {
    return
  }
  // Validate the first field.
  if err = ValidateFieldA(v.a); err != nil {
    // The rest of the structure remains unmarshaled.
    return
  }
  // ...
}
```

## Out of Order Deserialization

A simple example can be found in [examples-go](https://github.com/mus-format/examples-go/tree/main/out_of_order).

## Zero Allocation Deserialization

Can be achieved using the `unsafe` package.

## Testing

To run all `mus-go` tests, use the following command:

```bash
go test ./...
```

### Fuzz Testing

`mus-go` also includes fuzz tests. To run them, you can use the `fuzz.sh` script:

```bash
./fuzz.sh 10s
```

Or you can run a specific fuzz test using the `go test` command:

```bash
go test -v -fuzz="^FuzzByte$" ./varint -fuzztime 10s
```

## Benchmarks

Performance benchmarks for `mus-go` can be found at:

- [github.com/ymz-ncnk/go-serialization-benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)
- [github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

Why a separate benchmark suite? The existing [benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)
sometimes produce inconsistent results across multiple runs, making it
difficult to reliably compare serializers. A new [benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)
were created to provide more consistent and reproducible measurements.
