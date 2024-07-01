package varint

// MarshalPositiveInt64 fills bs with the MUS encoding (Varint) of an int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPositiveInt64(v int64, bs []byte) (n int) {
	return marshalUint(uint64(v), bs)
}

// MarshalPositiveInt32 fills bs with the MUS encoding (Varint) of an int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPositiveInt32(v int32, bs []byte) (n int) {
	return marshalUint(uint32(v), bs)
}

// MarshalPositiveInt16 fills bs with the MUS encoding (Varint) of an int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPositiveInt16(v int16, bs []byte) (n int) {
	return marshalUint(uint16(v), bs)
}

// MarshalPositiveInt8 fills bs with the MUS encoding (Varint) of an int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPositiveInt8(v int8, bs []byte) (n int) {
	return marshalUint(uint8(v), bs)
}

// MarshalPositiveInt fills bs with the MUS encoding (Varint) of an int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPositiveInt(v int, bs []byte) (n int) {
	return marshalUint(uint(v), bs)
}

// UnmarshalPositiveInt64 parses a MUS-encoded (Varint) int64 value from bs.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int64 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalPositiveInt64(bs []byte) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(bs)
	if err != nil {
		return
	}
	return int64(uv), n, nil
}

// UnmarshalPositiveInt32 parses a MUS-encoded (Varint) int32 value from bs.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int32 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalPositiveInt32(bs []byte) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(bs)
	if err != nil {
		return
	}
	return int32(uv), n, nil
}

// UnmarshalPositiveInt16 parses a MUS-encoded (Varint) int16 value from bs.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int16 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalPositiveInt16(bs []byte) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(bs)
	if err != nil {
		return
	}
	return int16(uv), n, nil
}

// UnmarshalPositiveInt8 parses a MUS-encoded (Varint) int8 value from bs.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int8 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalPositiveInt8(bs []byte) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(bs)
	if err != nil {
		return
	}
	return int8(uv), n, nil
}

// UnmarshalPositiveInt parses a MUS-encoded (Varint) int value from bs.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the int value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalPositiveInt(bs []byte) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(bs)
	if err != nil {
		return
	}
	return int(uv), n, nil
}

// SizePositiveInt64 returns the size of a MUS-encoded (Varint) int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt64(v int64) int {
	return sizeUint(uint64(v))
}

// SizePositiveInt32 returns the size of a MUS-encoded (Varint) int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt32(v int32) int {
	return SizeUint32(uint32(v))
}

// SizePositiveInt16 returns the size of a MUS-encoded (Varint) int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt16(v int16) (size int) {
	return SizeUint16(uint16(v))
}

// SizePositiveInt8 returns the size of a MUS-encoded (Varint) int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt8(v int8) (size int) {
	return SizeUint8(uint8(v))
}

// SizePositiveInt returns the size of a MUS-encoded (Varint) int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
func SizePositiveInt(v int) (size int) {
	return SizeUint(uint(v))
}

// SkipPositiveInt64 skips a MUS-encoded (Varint) int64 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipPositiveInt64(bs []byte) (n int, err error) {
	return SkipUint64(bs)
}

// SkipPositiveInt32 skips a MUS-encoded (Varint) int32 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipPositiveInt32(bs []byte) (n int, err error) {
	return SkipUint32(bs)
}

// SkipPositiveInt16 skips a MUS-encoded (Varint) int16 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipPositiveInt16(bs []byte) (n int, err error) {
	return SkipUint16(bs)
}

// SkipPositiveInt8 skips a MUS-encoded (Varint) int8 value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipPositiveInt8(bs []byte) (n int, err error) {
	return SkipUint8(bs)
}

// SkipPositiveInt skips a MUS-encoded (Varint) int value.
// It should be used with positive values, like string length (does not use
// ZigZag).
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipPositiveInt(bs []byte) (n int, err error) {
	return SkipUint(bs)
}
