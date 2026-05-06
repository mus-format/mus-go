# mus: A High-Performance, Flexible Binary Serialization Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mus-format/mus-go.svg)](https://pkg.go.dev/github.com/mus-format/mus-go)
[![GoReportCard](https://goreportcard.com/badge/mus-format/mus-go)](https://goreportcard.com/report/github.com/mus-format/mus-go)
[![codecov](https://codecov.io/gh/mus-format/mus-go/graph/badge.svg?token=WLLZ1MMQDX)](https://codecov.io/gh/mus-format/mus-go)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/mus-format/mus-go/badge)](https://scorecard.dev/viewer/?uri=github.com/mus-format/mus-go)

**mus** is a powerful and versatile Go library for efficient binary
serialization.

While originally built for the [MUS format](https://medium.com/@ymz-ncnk/mus-serialization-format-20f833df12d5), 
its minimalist architecture and extensive set of serialization primitives make 
it ideal for implementing other binary formats (including [Protobuf-style encoding](https://github.com/mus-format/examples-go/tree/main/protobuf)).

A streaming version is also available: [mus-stream](https://github.com/mus-format/mus-stream-go).

## Why mus?

### Performance

- Optimized for maximum speed (see [benchmarks](#benchmarks)).
- Offers space-efficient binary representation.
- Has zero-allocation mode.

### Capabilities

In addition to 32/64-bit compatibility and private field access, the library 
provides:

- Typed serialization (versioning, `oneof`).
- Comprehensive pointers support (handles cyclic graphs and linked lists).
- Validation during unmarshalling.
- Data skipping (enables out-of-order and partial unmarshalling).

## Contents

- [mus: A High-Performance, Flexible Binary Serialization Library for Go](#mus-a-high-performance-flexible-binary-serialization-library-for-go)
  - [Why mus?](#why-mus)
    - [Performance](#performance)
    - [Capabilities](#capabilities)
  - [Contents](#contents)
  - [Quick Start](#quick-start)
  - [Code Generation (Recommended)](#code-generation-recommended)
  - [How To](#how-to)
  - [Packages](#packages)
    - [varint](#varint)
    - [raw](#raw)
    - [ord (ordinary)](#ord-ordinary)
    - [unsafe](#unsafe)
    - [pm (pointer mapping)](#pm-pointer-mapping)
    - [typed (data type metadata support)](#typed-data-type-metadata-support)
  - [Structs Support](#structs-support)
  - [More Features](#more-features)
  - [Testing](#testing)
    - [Fuzz Testing](#fuzz-testing)
  - [Benchmarks](#benchmarks)
  - [Contributing \& Security](#contributing--security)
  - [Version Compatibility](#version-compatibility)
  - [Used By](#used-by)
  
## Quick Start

Here's an example of how to use `mus` to serialize a number.

```go
package main

import (
  "fmt"
  "github.com/mus-format/mus-go/varint"
)

func main() {
  num  := 100

  // Pre-allocate a buffer.
  var (
    size = varint.Int.Size(num)
    bs   = make([]byte, size)
  )
  
  // Marshal.
  varint.Int.Marshal(num, bs)
  
  // Unmarshal.
  num, _, err := varint.Int.Unmarshal(bs)
  if err != nil {
    panic(err)
  }
}
```

## Code Generation (Recommended)

Implementing `mus` serializers manually can be tedious and error-prone. There
are two ways to automate this:

- [mus-gen](https://github.com/mus-format/mus-gen-go) — Traditional code 
  generator.
- [mus-skill](https://github.com/mus-format/mus-skill-go) — AI agent skill.

## How To

To make a type serializable with `mus`, you need to implement the
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
value := YourType{...}

var (
  size = YourTypeMUS.Size(value) // The number of bytes required to serialize 
  // the value.
  bs = make([]byte, size)
)

n := YourTypeMUS.Marshal(value, bs) // Returns the number of used bytes.

value, n, err := YourTypeMUS.Unmarshal(bs) // Returns the value, the number of 
// used bytes and any error encountered.

// Instead of unmarshalling, the value can be skipped:
n, err := YourTypeMUS.Skip(bs)
```

## Packages

`mus` offers several encoding options, each in a separate package.

| Package                                          | Use Case                                 | Key Strength                  | Trade-off                            |
| :----------------------------------------------- | :--------------------------------------- | :---------------------------- | :----------------------------------- |
| **[`varint`](#varint)**                          | Numbers                                  | Space efficient               | Slight CPU cost for encoding         |
| **[`raw`](#raw)**                                | Numbers, Time                            | Fast encoding                 | Higher space usage for small numbers |
| **[`ord`](#ord-ordinary)**                       | Pointers, Strings, Slices, Maps          | Variable-length types support | Standard allocations                 |
| **[`unsafe`](#unsafe)**                          | High-perf Numbers, Time, Strings, Arrays | Zero-allocation               | Uses unsafe type conversions         |
| **[`pm`](#pm-pointer-mapping)**                  | Pointers, Cyclic Graphs, Linked Lists    | Preserves pointer equality    | Slightly more complex than `ord`     |
| **[`typed`](#typed-data-type-metadata-support)** | Interface/Versioning                     | Typed serialization           | Requires DTM definition              |

### varint

This package provides Varint serializers for all `uint` (e.g., `uint64`,
`uint32`, ...), `int`, `float`, and `byte` data types. It also includes the 
`PositiveInt` serializer (Varint without ZigZag) for efficiently encoding 
positive `int` values (negative values are supported as well, though with 
reduced performance).

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

### typed (data type metadata support)

The `typed` package provides [DTM](https://medium.com/p/21d7be309e8d) 
support for the `mus` serializer. It wraps a type serializer and a DTM 
value, enabling [typed data serialization](https://ymz-ncnk.medium.com/mus-serialization-format-20f833df12d5)
to provide data versioning, the oneof feature, and [other capabilities](https://github.com/mus-format/examples-go/tree/main/typed).

## Structs Support

`mus` doesn’t support structs out of the box, which means you’ll need to 
implement the `mus.Serializer` interface yourself by choosing the specific
encoding for each field ([example](https://github.com/mus-format/examples-go/tree/main/types/struct)).
This approach provides greater flexibility and keeps `mus` simple, making it 
easy to implement in other programming languages.

## More Features

- Validation: Validate data during unmarshalling using custom functions:
  `func(v Type) error` ([examples](https://github.com/mus-format/examples-go/tree/main/validation)).
- Out-of-order deserialization: Decode fields partially or non-sequentially 
  ([example](https://github.com/mus-format/examples-go/tree/main/out_of_order)).
- Zero-allocation: Achieve maximum efficiency by using the `unsafe` package.

## Testing

To run all `mus` tests, use the following command:

```bash
go test ./...
```

### Fuzz Testing

`mus` also includes fuzz tests. To run them, you can use the `fuzz.sh` script:

```bash
./fuzz.sh 10s
```

Or you can run a specific fuzz test using the `go test` command:

```bash
go test -v -fuzz="^FuzzByte$" ./varint -fuzztime 10s
```

## Benchmarks

| NAME     | NS/OP   | B/SIZE | B/OP   | ALLOCS/OP |
| -------- | ------- | ------ | ------ | --------- |
| mus      | 102.90  | 58.00  | 48.00  | 1         |
| protobuf | 531.70  | 69.00  | 271.00 | 4         |
| json     | 2779.00 | 150.00 | 600.00 | 9         |

Data sourced from [ymz-ncnk/go-serialization-benchmarks](https://github.com/ymz-ncnk/go-serialization-benchmarks).

Why a separate benchmark suite? [Standard benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) 
sometimes produce inconsistent results across multiple runs. The 
`ymz-ncnk/go-serialization-benchmarks` suite was created to provide more 
consistent and reproducible measurements.

## Contributing & Security

We welcome contributions of all kinds! Please see [CONTRIBUTING.md](CONTRIBUTING.md) 
for details on how to get involved.

If you find a security vulnerability, please refer to our 
[Security Policy](SECURITY.md) for instructions on how to report it privately.

## Version Compatibility

Check [VERSIONS.md](VERSIONS.md) for the compatibility matrix of module 
versions.

## Used By

See who is building with MUS in [USERS.md](USERS.md).

Already using MUS and want to support the project's growth? **[Join the list!](USERS.md)**
