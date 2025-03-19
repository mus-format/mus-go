# mus-go Serializer

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-go)](https://goreportcard.com/report/github.com/mus-format/mus-go)
[![codecov](https://codecov.io/github/mus-format/mus-go/graph/badge.svg?token=WLLZ1MMQDX)](https://codecov.io/github/mus-format/mus-go)

mus-go is a [MUS format](https://medium.com/p/21d7be309e8d) serializer. However,
due to its minimalist design and a wide range of serialization primitives, it 
can also be used to implement other binary serialization formats 
([here](https://github.com/mus-format/mus-examples-go/blob/main/protobuf/main.go)
is an example where mus-go is utilized to implement Protobuf encoding).

To get started quickly, go to the [code generator](https://github.com/mus-format/musgen-go) page.

## Why mus-go?
It is lightning fast, space efficient and well tested.

## Description
- Has a [streaming version](https://github.com/mus-format/mus-stream-go).
- Can run on both 32 and 64-bit systems.
- Variable-length data types (like `string`, `array`, `slice`, or `map`) are 
  encoded as: `length + data`. You can choose binary representation for both of 
  these parts. 
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
  - [Description](#description)
- [Contents](#contents)
- [cmd-stream-go](#cmd-stream-go)
- [musgen-go](#musgen-go)
- [Benchmarks](#benchmarks)
- [How To](#how-to)
  - [Packages](#packages)
    - [varint](#varint)
    - [raw](#raw)
    - [ord (ordinary)](#ord-ordinary)
      - [Array](#array)
      - [Slice](#slice)
      - [Map](#map)
    - [unsafe](#unsafe)
    - [pm (pointer mapping)](#pm-pointer-mapping)
  - [Structs Support](#structs-support)
  - [MarshallerMUS Interface](#marshallermus-interface)
  - [Generic MarshalMUS Function](#generic-marshalmus-function)
  - [DTM (Data Type Metadata) Support](#dtm-data-type-metadata-support)
  - [Data Versioning](#data-versioning)
  - [Interface Serialization (oneof feature)](#interface-serialization-oneof-feature)
  - [Validation](#validation)
    - [String](#string)
    - [Slice](#slice-1)
    - [Map](#map-1)
    - [Struct](#struct)
- [Out of Order Deserialization](#out-of-order-deserialization)
- [Zero Allocation Deserialization](#zero-allocation-deserialization)

# cmd-stream-go
[cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go) allows to execute 
commands on the server. cmd-stream-go/MUS is about [3 times faster](https://github.com/ymz-ncnk/go-client-server-communication-benchmarks) 
than gRPC/Protobuf.

# musgen-go
Writing mus-go code manually can be tedious and error-prone. A better approach 
is to use a [code generator](https://github.com/mus-format/musgen-go), it's also
incredibly easy to use - just provide a type and call `Generate()`.

# Benchmarks
- [github.com/ymz-ncnk/go-serialization-benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)
- [github.com/alecthomas/go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks)

Why did I create another [benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)?  
The existing [benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) 
have some notable issues - try running them several times, and you'll likely get
inconsistent results, making it difficult to determine which serializer is truly 
faster. That was one of the reasons, and basically I made them for my own use.

# How To
With mus-go, to make a type serializable, you need to implement the [Serializer](./mus.go) 
interface:
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

And than use it like:
```go
var (
  value YourType = ...
  size = YourTypeMUS.Size(value) // The number of bytes required to serialize the value.
  bs = make([]byte, size)
)

n := YourTypeMUS.Marshal(value, bs) // Returns the number of used bytes.
value, n, err := YourTypeMUS.Unmarshal(bs) // Returns the value, the number of 
// used bytes and any error encountered.

// Instead of unmarshalling the value can be skipped:
n, err := YourTypeMUS.Skip(bs)
```

## Packages
mus-go offers several encoding options, each of which is in a separate package.

### varint
Contains Varint serialzers for all `uint` (`uint64`, `uint32`, `uint16`, 
`uint8`, `uint`), `int`, `float`, `byte` data types. Example:
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

Also includes the `PositiveInt` serializer (Varint without ZigZag) for positive 
`int` values. It can handle negative values as well, but with lower performance.

### raw
Contains Raw serializers for the same `byte`, `uint`, `int`, `float`, `time.Time` 
data types. Example:
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

To ensure the deserialized value is in UTC, make sure your TZ environment 
variable is set to UTC. This can be done as follows:
```go
os.Setenv("TZ", "")
```

Alternatively, you can use one of the corresponding UTC serializers, e.g., 
`TimeUnixUTC`, `TimeUnixMilliUTC`, etc.

### ord (ordinary)
Contains serializers/constructors for `bool`, `string`, `array`, `byte slice`,
`slice`, `map`, and pointer types.

Variable-length data types (such as `string`, `array`, `slice`, or `map`) are 
encoded as `length + data`. You can choose the binary representation for both 
parts. By default, the length is encoded using a Varint without ZigZag 
(`varint.PositiveInt`). In this case, the maximum length is limited by the 
maximum value of the `int` type on your system. This works well across different 
architectures - for example, an attempt to unmarshal a string that is too long 
on a 32-bit system will result in an `ErrOverflow`.

For `array`, `slice`, and `map` types, there are only constructors available to 
create a concrete serializer.

#### Array
Unfortunately, Go does not support generic parameterization of array sizes,
as a result, the array serializer constructor looks like:
```go
package main

import (
  "github.com/mus-format/mus-go/ord"
  "github.com/mus-format/mus-go/varint"
  arrops "github.com/mus-format/mus-go/options/array"
)

func main() {
   var (
    // The first type parameter of the NewArraySer function represents the array
    // type, and the second - the type of the array’s elements.
    //
    // As for the function parameters, the number 3 specifies the length of the
    // array, and varint.Int - the serializer for the array’s elements.
    ser = ord.NewArraySer[[3]int, int](3, varint.Int)

    // To create an array serializer with the specific length serializer use:
    // ser = ord.NewArraySer[[3]int, int](3, varint.Int, arrops.WithLenSer(lenSer))

    arr  = [3]int{1, 2, 3}
    size = ser.Size(arr)
    bs   = make([]byte, size)
  )
  n := ser.Marshal(arr, bs)
  arr, n, err := ser.Unmarshal(bs)
  // ...
}
```

#### Slice
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

### unsafe
The unsafe package provides maximum performance, but be careful - it uses an 
unsafe type conversion. This warning largely applies to the string type because 
modifying the byte slice after unmarshalling will also change the string’s 
contents. Here is an [example](https://github.com/mus-format/mus-examples-go/blob/main/unasafe/main.go) 
that demonstrates this behavior more clearly.

Provides serializers for the following data types: `byte`, `bool`, `string`, 
`byte slice`, `time.Time` and all `uint`, `int`, `float`.

### pm (pointer mapping)
Let's consider two pointers initialized with the same value:
```go
var (
  str = "hello world"
  ptr = &str

  ptr1 *string = ptr
  ptr2 *string = ptr
)
```

The `pm` package ensures that these pointers are serialized in such a way that 
after unmarshalling, they remain equal - `ptr1 == ptr2`. This behavior differs 
from the `ord` package, where the pointers would no longer be equal.

The `pm` package enables the serialization of data structures like graphs or 
linked lists. You can find corresponding examples in [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/pm).

## Structs Support
mus-go doesn’t support structural data types out of the box, which means you’ll 
need to implement the `mus.Serializer` interface yourself. But that’s not 
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
greater flexibility, and thirdly, mus-go stays quite simple, making it easy to 
implement in other programming languages.

## MarshallerMUS Interface
It is often convenient to define the `MarshallerMUS` interface:
```go
type MarshallerMUS interface {
  MarshalMUS(bs []byte) (n int)
  SizeMUS() (size int)
}

// Foo implements the MarshallerMUS interface.
type Foo struct {...}

func (f Foo) MarshalMUS(bs []byte) (n int) {
  return FooMUS.Marshal(f, bs) // or FooDTS.Marshal(f, bs)
}

func (f Foo) SizeMUS() (size int) {
  return FooMUS.Size(f) // or FooDTS.Size(f)
}
```

## Generic MarshalMUS Function
To define generic `MarshalMUS` function:
```go
package main

// Define the MarshallerMUS interface ...
type MarshallerMUS interface {
  MarshalMUS(bs []byte) (n int)
  SizeMUS() (size int)
}

// ... and the function itself.
func MarshalMUS(v MarshallerMUS) (bs []byte) {
  bs = make([]byte, v.SizeMUS())
  v.MarshalMUS(bs)
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

## DTM (Data Type Metadata) Support
[mus-dts-go](https://github.com/mus-format/mus-dts-go) provides [DTM](https://medium.com/p/21d7be309e8d) 
support.

## Data Versioning
mus-dts-go can be used to implement data versioning. [Here](https://github.com/mus-format/mus-examples-go/tree/main/versioning)
is an example.

## Interface Serialization (oneof feature)
mus-dts-go will also help to create a serializer for an interface. Example:
```go
import dts "github.com/mus-format/mus-dts-go"

// Interface to serializer.
type Instruction interface {...}

// Copy implements the Instruction and MarshallerMUS interfaces.
type Copy struct {...}

// MarshalMUS uses CopyDTS.
func (c Copy) MarshalMUS(bs []byte) (n int) {
  return CopyDTS.Marshal(c, bs)
}

// SizeMUS uses CopyDTS.
func (c Copy) SizeMUS() (size int) {
  return CopyDTS.Size(c, bs)
}

// Insert implements the Instruction and MarshallerMUS interfaces.
type Insert struct {...}

// MarshalMUS uses InsertDTS.
func (i Insert) MarshalMUS(bs []byte) (n int) {
  return InsertDTS.Marshal(c, bs)
}

// SizeMUS uses InsertDTS.
func (i Insert) SizeMUS() (size int) {
  return InsertDTS.Size(c, bs)
}

// instructionMUS implements the mus.Serializer interface.
type instructionMUS struct {}

func (s instructionMUS) Marshal(i Instruction, bs []byte) (n int) {
  if m, ok := i.(MarshallerMUS); ok {
    return m.MarshalMUS(bs)
  }
  panic("i doesn't implement the MarshallerMUS interface")
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
  if s, ok := i.(MarshallerMUS); ok {
    return s.SizeMUS()
  }
  panic("i doesn't implement the MarshallerMUS interface")
}
```

A full example can be found at [mus-examples-go](https://github.com/mus-format/mus-examples-go/tree/main/oneof).

## Validation
Validation is performed during unmarshalling. Validator is just a function 
with the following signature `func (value Type) error`, where `Type` is a type 
of the value to which the validator is applied.

### String
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

### Slice
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

### Map
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

### Struct
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

# Out of Order Deserialization
A simple example can be found [here](https://github.com/mus-format/mus-examples-go/tree/main/out_of_order).

# Zero Allocation Deserialization
Can be achieved using the `unsafe` package.
