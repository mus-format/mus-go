package raw

import (
	"time"

	com "github.com/mus-format/common-go"
)

var (
	// TimeUnix is a time.Time serializer that encodes a value as a Unix
	// timestamp in seconds.
	TimeUnix = timeUnixSer{}
	// TimeUnixMilli is a time.Time serializer that encodes a value as a Unix
	// timestamp in milliseconds.
	TimeUnixMilli = timeUnixMilliSer{}
	// TimeUnixMicro is a time.Time serializer that encodes a value as a Unix
	// timestamp in microseconds.
	TimeUnixMicro = timeUnixMicroSer{}
	// TimeUnixNano is a time.Time serializer that encodes a value as a Unix
	// timestamp in nanoseconds.
	TimeUnixNano = timeUnixNanoSer{}

	// TimeUnixUTC is a time.Time serializer that encodes a value as a Unix
	// timestamp in seconds. The deserialized value is always in UTC.
	TimeUnixUTC = timeUnixUTCSer{}
	// TimeUnixUTCMilli is a time.Time serializer that encodes a value as a Unix
	// timestamp in milliseconds. The deserialized value is always in UTC.
	TimeUnixMilliUTC = timeUnixMilliUTCSer{}
	// TimeUnixUTCMicro is a time.Time serializer that encodes a value as a Unix
	// timestamp in microseconds. The deserialized value is always in UTC.
	TimeUnixMicroUTC = timeUnixMicroUTCSer{}
	// TimeUnixUTCNano is a time.Time serializer that encodes a value as a Unix
	// timestamp in nanoseconds. The deserialized value is always in UTC.
	TimeUnixNanoUTC = timeUnixNanoUTCSer{}
)

// -----------------------------------------------------------------------------

type timeUnixSer struct{}

// Marshal fills bs with an encoded time.Time value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s timeUnixSer) Marshal(v time.Time, bs []byte) (n int) {
	return Int64.Marshal(v.Unix(), bs)
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	sec, n, err := Int64.Unmarshal(bs)
	if err != nil {
		return
	}
	v = time.Unix(sec, 0)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s timeUnixSer) Skip(bs []byte) (n int, err error) {
	return Int64.Skip(bs)
}

// -----------------------------------------------------------------------------

type timeUnixMilliSer struct{}

// Marshal fills bs with an encoded time.Time value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s timeUnixMilliSer) Marshal(v time.Time, bs []byte) (n int) {
	return Int64.Marshal(v.UnixMilli(), bs)
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixMilliSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	milli, n, err := Int64.Unmarshal(bs)
	if err != nil {
		return
	}
	v = time.UnixMilli(milli)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixMilliSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s timeUnixMilliSer) Skip(bs []byte) (n int, err error) {
	return Int64.Skip(bs)
}

// -----------------------------------------------------------------------------

type timeUnixMicroSer struct{}

// Marshal fills bs with an encoded time.Time value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s timeUnixMicroSer) Marshal(v time.Time, bs []byte) (n int) {
	return Int64.Marshal(v.UnixMicro(), bs)
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixMicroSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	micro, n, err := Int64.Unmarshal(bs)
	if err != nil {
		return
	}
	v = time.UnixMicro(micro)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixMicroSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s timeUnixMicroSer) Skip(bs []byte) (n int, err error) {
	return Int64.Skip(bs)
}

// -----------------------------------------------------------------------------

type timeUnixNanoSer struct{}

// Marshal fills bs with an encoded time.Time value.
//
// Returns the number of used bytes. It will panic if bs is too small. The
// result will be unpredictable if v is the zero Time.
func (s timeUnixNanoSer) Marshal(v time.Time, bs []byte) (n int) {
	return Int64.Marshal(v.UnixNano(), bs)
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixNanoSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	nano, n, err := Int64.Unmarshal(bs)
	if err != nil {
		return
	}
	v = time.Unix(0, nano)
	return
}

// Size returns the size of an encoded time.Time value.
func (s timeUnixNanoSer) Size(v time.Time) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded time.Time value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s timeUnixNanoSer) Skip(bs []byte) (n int, err error) {
	return Int64.Skip(bs)
}

// -----------------------------------------------------------------------------

type timeUnixUTCSer struct {
	timeUnixSer
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixUTCSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixSer.Unmarshal(bs)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixMilliUTCSer struct {
	timeUnixMilliSer
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixMilliUTCSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixMilliSer.Unmarshal(bs)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixMicroUTCSer struct {
	timeUnixMicroSer
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixMicroUTCSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixMicroSer.Unmarshal(bs)
	if err == nil {
		v = v.UTC()
	}
	return
}

// -----------------------------------------------------------------------------

type timeUnixNanoUTCSer struct {
	timeUnixNanoSer
}

// Unmarshal parses an encoded time.Time value from bs.
//
// In addition to the time.Time value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s timeUnixNanoUTCSer) Unmarshal(bs []byte) (v time.Time, n int, err error) {
	v, n, err = s.timeUnixNanoSer.Unmarshal(bs)
	if err == nil {
		v = v.UTC()
	}
	return
}
