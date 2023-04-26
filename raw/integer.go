package raw

import (
	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
)

func marshalInteger64[T muscom.Integer64](t T, bs []byte) int {
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	bs[2] = byte(t >> 16)
	bs[3] = byte(t >> 24)
	bs[4] = byte(t >> 32)
	bs[5] = byte(t >> 40)
	bs[6] = byte(t >> 48)
	bs[7] = byte(t >> 56)
	return muscom.Num64RawSize
}

func marshalInteger32[T muscom.Integer32](t T, bs []byte) int {
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	bs[2] = byte(t >> 16)
	bs[3] = byte(t >> 24)
	return muscom.Num32RawSize
}

func marshalInteger16[T muscom.Integer16](t T, bs []byte) int {
	bs[0] = byte(t)
	bs[1] = byte(t >> 8)
	return muscom.Num16RawSize
}

func marshalInteger8[T muscom.Integer8](t T, bs []byte) int {
	bs[0] = byte(t)
	return muscom.Num8RawSize
}

// -----------------------------------------------------------------------------
func unmarshalInteger64[T muscom.Integer64](bs []byte) (T, int, error) {
	var t T
	if len(bs) < muscom.Num64RawSize {
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
	return t, muscom.Num64RawSize, nil
}

func unmarshalInteger32[T muscom.Integer32](bs []byte) (T, int, error) {
	var t T
	if len(bs) < muscom.Num32RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	t = T(bs[0])
	t |= T(bs[1]) << 8
	t |= T(bs[2]) << 16
	t |= T(bs[3]) << 24
	return t, muscom.Num32RawSize, nil
}

func unmarshalInteger16[T muscom.Integer16](bs []byte) (T, int, error) {
	var t T
	if len(bs) < muscom.Num16RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	t = T(bs[0])
	t |= T(bs[1]) << 8
	return t, muscom.Num16RawSize, nil
}

func unmarshalInteger8[T muscom.Integer8](bs []byte) (T, int, error) {
	var t T
	if len(bs) < muscom.Num8RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return T(bs[0]), muscom.Num8RawSize, nil
}

// -----------------------------------------------------------------------------
func sizeNum64[T muscom.Num64](t T) int {
	return muscom.Num64RawSize
}

func sizeNum32[T muscom.Num32](t T) int {
	return muscom.Num32RawSize
}

func sizeInteger16[T muscom.Integer16](t T) int {
	return muscom.Num16RawSize
}

func sizeInteger8[T muscom.Integer8](t T) int {
	return muscom.Num8RawSize
}

// -----------------------------------------------------------------------------
func skipInteger64(bs []byte) (int, error) {
	if len(bs) < muscom.Num64RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return muscom.Num64RawSize, nil
}

func skipInteger32(bs []byte) (int, error) {
	if len(bs) < muscom.Num32RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return muscom.Num32RawSize, nil
}

func skipInteger16(bs []byte) (int, error) {
	if len(bs) < muscom.Num16RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return muscom.Num16RawSize, nil
}

func skipInteger8(bs []byte) (int, error) {
	if len(bs) < muscom.Num8RawSize {
		return 0, mus.ErrTooSmallByteSlice
	}
	return muscom.Num8RawSize, nil
}
