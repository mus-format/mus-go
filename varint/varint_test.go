package varint

import (
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
)

func TestVarint(t *testing.T) {

	t.Run("unmarshalUint", func(t *testing.T) {

		t.Run("unmarshalUint should return ErrTooSmallByteSlice if bs is empty",
			func(t *testing.T) {
				var (
					wantV     uint64 = 0
					wantN            = 0
					wantErr          = mus.ErrTooSmallByteSlice
					bs               = []byte{}
					v, n, err        = unmarshalUint[uint64](0, 0, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
			func(t *testing.T) {
				var (
					wantV     uint16 = 0
					wantN            = 3
					wantErr          = com.ErrOverflow
					bs               = []byte{200, 200, 200}
					v, n, err        = unmarshalUint[uint16](com.Uint16MaxVarintLen,
						com.Uint16MaxLastByte, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("unmarshalUint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     uint16 = 0
					wantN            = 2
					wantErr          = mus.ErrTooSmallByteSlice
					bs               = []byte{200, 200}
					v, n, err        = unmarshalUint[uint16](com.Uint16MaxVarintLen,
						com.Uint16MaxLastByte, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("skipUint", func(t *testing.T) {

		t.Run("skipUint should return ErrTooSmallByteSlice if bs is empty",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = skipUint(0, 0, bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("skipUint should return ErrOverflow if there is no varint end",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = com.ErrOverflow
					bs      = []byte{200, 200, 200, 200, 200}
					n, err  = skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
						bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("skipUint shold return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{200, 200}
					n, err  = skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
						bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)

			})

	})

	t.Run("unsigned", func(t *testing.T) {

		t.Run("All MarshalUint64, UnmarshalUint64, SizeUint64, SkipUint64 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[uint64](MarshalUint64)
					u  = mus.UnmarshallerFn[uint64](UnmarshalUint64)
					s  = mus.SizerFn[uint64](SizeUint64)
					sk = mus.SkipperFn(SkipUint64)
				)
				testdata.Test[uint64](com_testdata.Uint64TestCases, m, u, s, t)
				testdata.TestSkip[uint64](com_testdata.Uint64TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint32, UnmarshalUint32, SizeUint32, SkipUint32 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[uint32](MarshalUint32)
					u  = mus.UnmarshallerFn[uint32](UnmarshalUint32)
					s  = mus.SizerFn[uint32](SizeUint32)
					sk = mus.SkipperFn(SkipUint32)
				)
				testdata.Test[uint32](com_testdata.Uint32TestCases, m, u, s, t)
				testdata.TestSkip[uint32](com_testdata.Uint32TestCases, m, sk, s, t)
			})

		t.Run("All Marshal16, Unmarshal16, Size16, Skip16 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[uint16](MarshalUint16)
					u  = mus.UnmarshallerFn[uint16](UnmarshalUint16)
					s  = mus.SizerFn[uint16](SizeUint16)
					sk = mus.SkipperFn(SkipUint16)
				)
				testdata.Test[uint16](com_testdata.Uint16TestCases, m, u, s, t)
				testdata.TestSkip[uint16](com_testdata.Uint16TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint8, UnmarshalUint8, SizeUint8, SkipUint8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[uint8](MarshalUint8)
					u  = mus.UnmarshallerFn[uint8](UnmarshalUint8)
					s  = mus.SizerFn[uint8](SizeUint8)
					sk = mus.SkipperFn(SkipUint8)
				)
				testdata.Test[uint8](com_testdata.Uint8TestCases, m, u, s, t)
				testdata.TestSkip[uint8](com_testdata.Uint8TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint8, UnmarshalUint8, SizeUint8, SkipUint8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[uint](MarshalUint)
					u  = mus.UnmarshallerFn[uint](UnmarshalUint)
					s  = mus.SizerFn[uint](SizeUint)
					sk = mus.SkipperFn(SkipUint)
				)
				testdata.Test[uint](com_testdata.UintTestCases, m, u, s, t)
				testdata.TestSkip[uint](com_testdata.UintTestCases, m, sk, s, t)
			})

	})

	t.Run("signed", func(t *testing.T) {

		t.Run("int64", func(t *testing.T) {

			t.Run("All MarshalInt64, UnmarshalInt64, SizeInt64, SkipInt64 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[int64](MarshalInt64)
						u  = mus.UnmarshallerFn[int64](UnmarshalInt64)
						s  = mus.SizerFn[int64](SizeInt64)
						sk = mus.SkipperFn(SkipInt64)
					)
					testdata.Test[int64](com_testdata.Int64TestCases, m, u, s, t)
					testdata.TestSkip[int64](com_testdata.Int64TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalInt64 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int64 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalInt64(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("int32", func(t *testing.T) {

			t.Run("All MarshalInt32, UnmarshalInt32, SizeInt32, SkipInt32 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[int32](MarshalInt32)
						u  = mus.UnmarshallerFn[int32](UnmarshalInt32)
						s  = mus.SizerFn[int32](SizeInt32)
						sk = mus.SkipperFn(SkipInt32)
					)
					testdata.Test[int32](com_testdata.Int32TestCases, m, u, s, t)
					testdata.TestSkip[int32](com_testdata.Int32TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalInt32 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int32 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalInt32(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("int16", func(t *testing.T) {

			t.Run("All MarshalInt16, UnmarshalInt16, SizeInt16, SkipInt16 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[int16](MarshalInt16)
						u  = mus.UnmarshallerFn[int16](UnmarshalInt16)
						s  = mus.SizerFn[int16](SizeInt16)
						sk = mus.SkipperFn(SkipInt16)
					)
					testdata.Test[int16](com_testdata.Int16TestCases, m, u, s, t)
					testdata.TestSkip[int16](com_testdata.Int16TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalInt16 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int16 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalInt16(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("int8", func(t *testing.T) {

			t.Run("All MarshalInt8, UnmarshalInt8, SizeInt8, SkipInt8 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[int8](MarshalInt8)
						u  = mus.UnmarshallerFn[int8](UnmarshalInt8)
						s  = mus.SizerFn[int8](SizeInt8)
						sk = mus.SkipperFn(SkipInt8)
					)
					testdata.Test[int8](com_testdata.Int8TestCases, m, u, s, t)
					testdata.TestSkip[int8](com_testdata.Int8TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalInt8 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int8 = 0
						wantN          = 0
						wantErr        = mus.ErrTooSmallByteSlice
						bs             = []byte{}
						v, n, err      = UnmarshalInt8(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("int", func(t *testing.T) {

			t.Run("All MarshalInt, UnmarshalInt, SizeInt, SkipInt functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[int](MarshalInt)
						u  = mus.UnmarshallerFn[int](UnmarshalInt)
						s  = mus.SizerFn[int](SizeInt)
						sk = mus.SkipperFn(SkipInt)
					)
					testdata.Test[int](com_testdata.IntTestCases, m, u, s, t)
					testdata.TestSkip[int](com_testdata.IntTestCases, m, sk, s, t)
				})

			t.Run("UnmarshalInt should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int = 0
						wantN         = 0
						wantErr       = mus.ErrTooSmallByteSlice
						bs            = []byte{}
						v, n, err     = UnmarshalInt(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("Positive int64", func(t *testing.T) {

			t.Run("All MarshalPositiveInt64, UnmarshalPositiveInt64, SizePositiveInt64, SkipPositiveInt64 functions must work correctly",
				func(t *testing.T) {
					var (
						m  mus.MarshallerFn[int64]   = MarshalPositiveInt64
						u  mus.UnmarshallerFn[int64] = UnmarshalPositiveInt64
						s  mus.SizerFn[int64]        = SizePositiveInt64
						sk mus.SkipperFn             = SkipPositiveInt64
					)
					testdata.Test[int64](com_testdata.Int64TestCases, m, u, s, t)
					testdata.TestSkip[int64](com_testdata.Int64TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalPositiveInt64 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int64 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalPositiveInt64(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("Positive int32", func(t *testing.T) {

			t.Run("All MarshalPositiveInt32, UnmarshalPositiveInt32, SizePositiveInt32, SkipPositiveInt32 functions must work correctly",
				func(t *testing.T) {
					var (
						m  mus.MarshallerFn[int32]   = MarshalPositiveInt32
						u  mus.UnmarshallerFn[int32] = UnmarshalPositiveInt32
						s  mus.SizerFn[int32]        = SizePositiveInt32
						sk mus.SkipperFn             = SkipPositiveInt32
					)
					testdata.Test[int32](com_testdata.Int32TestCases, m, u, s, t)
					testdata.TestSkip[int32](com_testdata.Int32TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalPositiveInt32 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int32 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalPositiveInt32(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("Positive int16", func(t *testing.T) {

			t.Run("All MarshalPositiveInt16, UnmarshalPositiveInt16, SizePositiveInt16, SkipPositiveInt16 functions must work correctly",
				func(t *testing.T) {
					var (
						m  mus.MarshallerFn[int16]   = MarshalPositiveInt16
						u  mus.UnmarshallerFn[int16] = UnmarshalPositiveInt16
						s  mus.SizerFn[int16]        = SizePositiveInt16
						sk mus.SkipperFn             = SkipPositiveInt16
					)
					testdata.Test[int16](com_testdata.Int16TestCases, m, u, s, t)
					testdata.TestSkip[int16](com_testdata.Int16TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalPositiveInt16 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int16 = 0
						wantN           = 0
						wantErr         = mus.ErrTooSmallByteSlice
						bs              = []byte{}
						v, n, err       = UnmarshalPositiveInt16(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("Positive int8", func(t *testing.T) {

			t.Run("All MarshalPositiveInt8, UnmarshalPositiveInt8, SizePositiveInt8, SkipPositiveInt8 functions must work correctly",
				func(t *testing.T) {
					var (
						m  mus.MarshallerFn[int8]   = MarshalPositiveInt8
						u  mus.UnmarshallerFn[int8] = UnmarshalPositiveInt8
						s  mus.SizerFn[int8]        = SizePositiveInt8
						sk mus.SkipperFn            = SkipPositiveInt8
					)
					testdata.Test[int8](com_testdata.Int8TestCases, m, u, s, t)
					testdata.TestSkip[int8](com_testdata.Int8TestCases, m, sk, s, t)
				})

			t.Run("UnmarshalPositiveInt8 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int8 = 0
						wantN          = 0
						wantErr        = mus.ErrTooSmallByteSlice
						bs             = []byte{}
						v, n, err      = UnmarshalPositiveInt8(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("Positive int", func(t *testing.T) {

			t.Run("All MarshalPositiveInt, UnmarshalPositiveInt, SizePositiveInt, SkipPositiveInt functions must work correctly",
				func(t *testing.T) {
					var (
						m  mus.MarshallerFn[int]   = MarshalPositiveInt
						u  mus.UnmarshallerFn[int] = UnmarshalPositiveInt
						s  mus.SizerFn[int]        = SizePositiveInt
						sk mus.SkipperFn           = SkipPositiveInt
					)
					testdata.Test[int](com_testdata.IntTestCases, m, u, s, t)
					testdata.TestSkip[int](com_testdata.IntTestCases, m, sk, s, t)
				})

			t.Run("UnmarshaPositivelInt should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     int = 0
						wantN         = 0
						wantErr       = mus.ErrTooSmallByteSlice
						bs            = []byte{}
						v, n, err     = UnmarshalPositiveInt(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("All MarshalByte, UnmarshalByte, SizeByte, SkipByte functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[byte](MarshalByte)
					u  = mus.UnmarshallerFn[byte](UnmarshalByte)
					s  = mus.SizerFn[byte](SizeByte)
					sk = mus.SkipperFn(SkipByte)
				)
				testdata.Test[byte](com_testdata.ByteTestCases, m, u, s, t)
				testdata.TestSkip[byte](com_testdata.ByteTestCases, m, sk, s, t)
			})

	})

	t.Run("float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("All MarshalFloat64, UnmarshalFloat64, SizeFloat64, SkipFloat64 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[float64](MarshalFloat64)
						u  = mus.UnmarshallerFn[float64](UnmarshalFloat64)
						s  = mus.SizerFn[float64](SizeFloat64)
						sk = mus.SkipperFn(SkipFloat64)
					)
					testdata.Test[float64](com_testdata.Float64TestCases, m, u, s, t)
					testdata.TestSkip[float64](com_testdata.Float64TestCases, m, sk, s, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{}
						v, n, err         = UnmarshalFloat64(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("All MarshalFloat32, UnmarshalFloat32, SizeFloat32, SkipFloat32 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshallerFn[float32](MarshalFloat32)
						u  = mus.UnmarshallerFn[float32](UnmarshalFloat32)
						s  = mus.SizerFn[float32](SizeFloat32)
						sk = mus.SkipperFn(SkipFloat32)
					)
					testdata.Test[float32](com_testdata.Float32TestCases, m, u, s, t)
					testdata.TestSkip[float32](com_testdata.Float32TestCases, m, sk, s, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{}
						v, n, err         = UnmarshalFloat32(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})

		})

	})

}
