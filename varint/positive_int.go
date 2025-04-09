package varint

var (
	// PositiveInt64 is an int64 serializer, for positive values.
	PositiveInt64 = positiveInt64Ser{}
	// PositiveInt32 is an int32 serializer, for positive values.
	PositiveInt32 = positiveInt32Ser{}
	// PositiveInt16 is an int16 serializer, for positive values.
	PositiveInt16 = positiveInt16Ser{}
	// PositiveInt8 is an int8 serializer, for positive values.
	PositiveInt8 = positiveInt8Ser{}
	// PositiveInt is an int serializer, for positive values.
	PositiveInt = positiveIntSer{}
)

type positiveInt64Ser struct{}

// Marshal fills bs with an encoded (Varint) int64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s positiveInt64Ser) Marshal(v int64, bs []byte) (n int) {
	return marshalUint(uint64(v), bs)
}

// Unmarshal fills v with an decoded (Varint) int64 value.
//
// In addition to the int64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt64Ser) Unmarshal(bs []byte) (v int64, n int, err error) {
	var uv uint64
	uv, n, err = Uint64.Unmarshal(bs)
	v = int64(uv)
	return
}

// Size returns the size of an encoded (Varint) int64 value.
func (s positiveInt64Ser) Size(v int64) (size int) {
	return sizeUint(uint64(v))
}

// Skip skips the next encoded (Varint) int64 value.
//
// In addition to the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt64Ser) Skip(bs []byte) (n int, err error) {
	return Uint64.Skip(bs)
}

// -----------------------------------------------------------------------------

type positiveInt32Ser struct{}

// Marshal fills bs with an encoded (Varint) int32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s positiveInt32Ser) Marshal(v int32, bs []byte) (n int) {
	return marshalUint(uint32(v), bs)
}

// Unmarshal fills v with an decoded (Varint) int32 value.
//
// In addition to the int32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt32Ser) Unmarshal(bs []byte) (v int32, n int, err error) {
	var uv uint32
	uv, n, err = Uint32.Unmarshal(bs)
	v = int32(uv)
	return
}

// Size returns the size of an encoded (Varint) int32 value.
func (s positiveInt32Ser) Size(v int32) (size int) {
	return sizeUint(uint32(v))
}

// Skip skips the next encoded (Varint) int32 value.
//
// In addition to the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt32Ser) Skip(bs []byte) (n int, err error) {
	return Uint32.Skip(bs)
}

// -----------------------------------------------------------------------------

type positiveInt16Ser struct{}

// Marshal fills bs with an encoded (Varint) int16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s positiveInt16Ser) Marshal(v int16, bs []byte) (n int) {
	return marshalUint(uint16(v), bs)
}

// Unmarshal fills v with an decoded (Varint) int16 value.
//
// In addition to the int16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt16Ser) Unmarshal(bs []byte) (v int16, n int, err error) {
	var uv uint16
	uv, n, err = Uint16.Unmarshal(bs)
	v = int16(uv)
	return
}

// Size returns the size of an encoded (Varint) int16 value.
func (s positiveInt16Ser) Size(v int16) (size int) {
	return sizeUint(uint16(v))
}

// Skip skips the next encoded (Varint) int16 value.
//
// In addition to the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt16Ser) Skip(bs []byte) (n int, err error) {
	return Uint16.Skip(bs)
}

// -----------------------------------------------------------------------------

type positiveInt8Ser struct{}

// Marshal fills bs with an encoded (Varint) int8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s positiveInt8Ser) Marshal(v int8, bs []byte) (n int) {
	return marshalUint(uint8(v), bs)
}

// Unmarshal fills v with an decoded (Varint) int8 value.
//
// In addition to the int8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt8Ser) Unmarshal(bs []byte) (v int8, n int, err error) {
	var uv uint8
	uv, n, err = Uint8.Unmarshal(bs)
	v = int8(uv)
	return
}

// Size returns the size of an encoded (Varint) int8 value.
func (s positiveInt8Ser) Size(v int8) (size int) {
	return sizeUint(uint8(v))
}

// Skip skips the next encoded (Varint) int8 value.
//
// In addition to the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveInt8Ser) Skip(bs []byte) (n int, err error) {
	return Uint8.Skip(bs)
}

// -----------------------------------------------------------------------------

type positiveIntSer struct{}

// Marshal fills bs with an encoded (Varint) int value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s positiveIntSer) Marshal(v int, bs []byte) (n int) {
	return marshalUint(uint(v), bs)
}

// Unmarshal fills v with an decoded (Varint) int value.
//
// In addition to the int value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveIntSer) Unmarshal(bs []byte) (v int, n int, err error) {
	uv, n, err := Uint.Unmarshal(bs)
	if err != nil {
		return
	}
	return int(uv), n, nil
}

// Size returns the size of an encoded (Varint) int value.
func (s positiveIntSer) Size(v int) (size int) {
	return sizeUint(uint(v))
}

// Skip skips the next encoded (Varint) int value.
//
// In addition to the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s positiveIntSer) Skip(bs []byte) (n int, err error) {
	return Uint.Skip(bs)
}
