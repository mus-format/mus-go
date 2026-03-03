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

- Supports data versioning and interface serialization (oneof feature) using DTS.
- Comprehensive pointer support.
- Can encode cyclic graphs and linked lists.
- Offers zero-allocation deserialization.

### Additional Features

- Validation and field skipping during unmarshalling.
- Supports private fields.
- Out-of-order deserialization.

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
    - [unsafe](#unsafe)
    - [pm (pointer mapping)](#pm-pointer-mapping)
  - [Structs Support](#structs-support)
  - [DTS (Data Type metadata Support)](#dts-data-type-metadata-support)
  - [More Features](#more-features)
  - [Testing](#testing)
    - [Fuzz Testing](#fuzz-testing)
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

## Packages

`mus-go` offers several encoding options, each in a separate package.

### varint

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

### raw

This package contains Raw serializers for `byte`, `uint`, `int`, `float`, and
`time.Time` data types.

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

### ord (ordinary)

Contains serializers/constructors for `bool`, `string`, `byte slice`,
`slice`, `map`, and `pointer` types.

Variable-length data types (such as `string`, `slice`, and `map`) are
encoded as `length + data`, with customizable binary representations for both
parts. By default, the length is encoded using `varint.PositiveInt`, which 
limits the length to the maximum value of the `int` type on your system. Such 
encoding works well across different architectures. For example, an attempt to 
unmarshal a string that is too long on a 32-bit system will result in an 
`ErrOverflow`.

For `slice` and `map` types, only constructors are available ([examples](https://github.com/mus-format/examples-go/tree/main/types)).

### unsafe

The `unsafe` package provides maximum performance by using unsafe type 
conversions. This primarily affects the `string` type, where modifying the 
underlying byte slice after unmarshalling will also change the string's contents 
([example](https://github.com/mus-format/examples-go/tree/main/unsafe)).

Provides serializers for the following data types: `byte`, `bool`, `string`,
`array`, `byte slice`, `time.Time` and all `uint`, `int`, `float`.

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

The `pm` package preserves pointer equality after unmarshalling, ensuring that
`ptr1 == ptr2`, while the `ord` package does not. This capability enables the 
serialization of data structures like cyclic graphs or linked lists ([examples](https://github.com/mus-format/examples-go/tree/main/pm)).

## Structs Support

`mus-go` doesn’t support structs out of the box, which means you’ll need to 
implement the `mus.Serializer` interface yourself. Simply deconstruct the struct
into its fields and choose the desired encoding for each ([example](https://github.com/mus-format/examples-go/tree/main/types/struct)).

This approach provides greater flexibility and keeps `mus-go` simple, making it 
easy to implement in other programming languages.

## DTS (Data Type metadata Support)

[dts-go](https://github.com/mus-format/dts-go) enables [typed data serialization](https://ymz-ncnk.medium.com/mus-serialization-format-20f833df12d5) 
to provide data versioning, the oneof feature, and [other capabilities](https://github.com/mus-format/examples-go/tree/main/dts).

## More Features

- **Validation**: Validate data during unmarshalling with a custom function:
  `func(v Type) error` ([examples](https://github.com/mus-format/examples-go/tree/main/validation)).
- **Out-of-Order Deserialization**: Decode fields partially or non-sequentially 
  for greater efficiency ([example](https://github.com/mus-format/examples-go/tree/main/out_of_order)).
- **Zero-Allocation**: Use the `unsafe` package to achieve zero-allocation 
  deserialization.
  
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
difficult to reliably compare serializers. New [benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks)
were created to provide more consistent and reproducible measurements.
