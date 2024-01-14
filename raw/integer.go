package raw

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

func marshalInteger64[T com.Integer64](t T, bs []byte) int {
	_ = bs[7]
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	bs[2] = byte(t >> 16)
	bs[3] = byte(t >> 24)
	bs[4] = byte(t >> 32)
	bs[5] = byte(t >> 40)
	bs[6] = byte(t >> 48)
	bs[7] = byte(t >> 56)
	return com.Num64RawSize
}

func marshalInteger32[T com.Integer32](t T, bs []byte) int {
	_ = bs[3]
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	bs[2] = byte(t >> 16)
	bs[3] = byte(t >> 24)
	return com.Num32RawSize
}

func marshalInteger16[T com.Integer16](t T, bs []byte) int {
	_ = bs[1]
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	return com.Num16RawSize
}

func marshalInteger8[T com.Integer8](t T, bs []byte) int {
	bs[0] = byte(t)
	return com.Num8RawSize
}

func unmarshalInteger64[T com.Integer64](bs []byte) (T, int, error) {
	var t T
	if len(bs) < com.Num64RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	t = T(bs[0])
	t |= T(bs[1]) << 8
	t |= T(bs[2]) << 16
	t |= T(bs[3]) << 24
	t |= T(bs[4]) << 32
	t |= T(bs[5]) << 40
	t |= T(bs[6]) << 48
	t |= T(bs[7]) << 56
	return t, com.Num64RawSize, nil
}

func unmarshalInteger32[T com.Integer32](bs []byte) (T, int, error) {
	var t T
	if len(bs) < com.Num32RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	t = T(bs[0])
	t |= T(bs[1]) << 8
	t |= T(bs[2]) << 16
	t |= T(bs[3]) << 24
	return t, com.Num32RawSize, nil
}

func unmarshalInteger16[T com.Integer16](bs []byte) (T, int, error) {
	var t T
	if len(bs) < com.Num16RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	t = T(bs[0])
	t |= T(bs[1]) << 8
	return t, com.Num16RawSize, nil
}

func unmarshalInteger8[T com.Integer8](bs []byte) (T, int, error) {
	var t T
	if len(bs) < com.Num8RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return T(bs[0]), com.Num8RawSize, nil
}

func sizeNum64[T com.Num64](t T) int {
	return com.Num64RawSize
}

func sizeNum32[T com.Num32](t T) int {
	return com.Num32RawSize
}

func sizeInteger16[T com.Integer16](t T) int {
	return com.Num16RawSize
}

func sizeInteger8[T com.Integer8](t T) int {
	return com.Num8RawSize
}

func skipInteger64(bs []byte) (int, error) {
	if len(bs) < com.Num64RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return com.Num64RawSize, nil
}

func skipInteger32(bs []byte) (int, error) {
	if len(bs) < com.Num32RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return com.Num32RawSize, nil
}

func skipInteger16(bs []byte) (int, error) {
	if len(bs) < com.Num16RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return com.Num16RawSize, nil
}

func skipInteger8(bs []byte) (int, error) {
	if len(bs) < com.Num8RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return com.Num8RawSize, nil
}
