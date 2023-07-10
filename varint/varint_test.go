package varint

import (
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
			func(t *testing.T) {
				var (
					wantV     uint16 = 0
					wantN            = 3
					wantErr          = muscom.ErrOverflow
					bs               = []byte{200, 200, 200}
					v, n, err        = unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
						muscom.Uint16MaxLastByte, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("unmarshalUint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     uint16 = 0
					wantN            = 2
					wantErr          = mus.ErrTooSmallByteSlice
					bs               = []byte{200, 200}
					v, n, err        = unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
						muscom.Uint16MaxLastByte, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
			})

		t.Run("skipUint should return ErrOverflow if there is no varint end",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = muscom.ErrOverflow
					bs      = []byte{200, 200, 200, 200, 200}
					n, err  = skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte,
						bs)
				)
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
			})

		t.Run("skipUint shold return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{200, 200}
					n, err  = skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte,
						bs)
				)
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)

			})

	})

	t.Run("unsigned", func(t *testing.T) {

		t.Run("All MarshalUint64, UnmarshalUint64, SizeUint64, SkipUint64 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[uint64](MarshalUint64)
					u  = mus.UnmarshalerFn[uint64](UnmarshalUint64)
					s  = mus.SizerFn[uint64](SizeUint64)
					sk = mus.SkipperFn(SkipUint64)
				)
				testdata.Test[uint64](muscom_testdata.Uint64TestCases, m, u, s, t)
				testdata.TestSkip[uint64](muscom_testdata.Uint64TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint32, UnmarshalUint32, SizeUint32, SkipUint32 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[uint32](MarshalUint32)
					u  = mus.UnmarshalerFn[uint32](UnmarshalUint32)
					s  = mus.SizerFn[uint32](SizeUint32)
					sk = mus.SkipperFn(SkipUint32)
				)
				testdata.Test[uint32](muscom_testdata.Uint32TestCases, m, u, s, t)
				testdata.TestSkip[uint32](muscom_testdata.Uint32TestCases, m, sk, s, t)
			})

		t.Run("All Marshal16, Unmarshal16, Size16, Skip16 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[uint16](MarshalUint16)
					u  = mus.UnmarshalerFn[uint16](UnmarshalUint16)
					s  = mus.SizerFn[uint16](SizeUint16)
					sk = mus.SkipperFn(SkipUint16)
				)
				testdata.Test[uint16](muscom_testdata.Uint16TestCases, m, u, s, t)
				testdata.TestSkip[uint16](muscom_testdata.Uint16TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint8, UnmarshalUint8, SizeUint8, SkipUint8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[uint8](MarshalUint8)
					u  = mus.UnmarshalerFn[uint8](UnmarshalUint8)
					s  = mus.SizerFn[uint8](SizeUint8)
					sk = mus.SkipperFn(SkipUint8)
				)
				testdata.Test[uint8](muscom_testdata.Uint8TestCases, m, u, s, t)
				testdata.TestSkip[uint8](muscom_testdata.Uint8TestCases, m, sk, s, t)
			})

		t.Run("All MarshalUint8, UnmarshalUint8, SizeUint8, SkipUint8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[uint](MarshalUint)
					u  = mus.UnmarshalerFn[uint](UnmarshalUint)
					s  = mus.SizerFn[uint](SizeUint)
					sk = mus.SkipperFn(SkipUint)
				)
				testdata.Test[uint](muscom_testdata.UintTestCases, m, u, s, t)
				testdata.TestSkip[uint](muscom_testdata.UintTestCases, m, sk, s, t)
			})

	})

	t.Run("signed", func(t *testing.T) {

		t.Run("All MarshalInt64, UnmarshalInt64, SizeInt64, SkipInt64 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int64](MarshalInt64)
					u  = mus.UnmarshalerFn[int64](UnmarshalInt64)
					s  = mus.SizerFn[int64](SizeInt64)
					sk = mus.SkipperFn(SkipInt64)
				)
				testdata.Test[int64](muscom_testdata.Int64TestCases, m, u, s, t)
				testdata.TestSkip[int64](muscom_testdata.Int64TestCases, m, sk, s, t)
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("All MarshalInt32, UnmarshalInt32, SizeInt32, SkipInt32 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int32](MarshalInt32)
					u  = mus.UnmarshalerFn[int32](UnmarshalInt32)
					s  = mus.SizerFn[int32](SizeInt32)
					sk = mus.SkipperFn(SkipInt32)
				)
				testdata.Test[int32](muscom_testdata.Int32TestCases, m, u, s, t)
				testdata.TestSkip[int32](muscom_testdata.Int32TestCases, m, sk, s, t)
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Int16", func(t *testing.T) {

			t.Run("All MarshalInt16, UnmarshalInt16, SizeInt16, SkipInt16 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshalerFn[int16](MarshalInt16)
						u  = mus.UnmarshalerFn[int16](UnmarshalInt16)
						s  = mus.SizerFn[int16](SizeInt16)
						sk = mus.SkipperFn(SkipInt16)
					)
					testdata.Test[int16](muscom_testdata.Int16TestCases, m, u, s, t)
					testdata.TestSkip[int16](muscom_testdata.Int16TestCases, m, sk, s, t)
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
					muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("All MarshalInt8, UnmarshalInt8, SizeInt8, SkipInt8 functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int8](MarshalInt8)
					u  = mus.UnmarshalerFn[int8](UnmarshalInt8)
					s  = mus.SizerFn[int8](SizeInt8)
					sk = mus.SkipperFn(SkipInt8)
				)
				testdata.Test[int8](muscom_testdata.Int8TestCases, m, u, s, t)
				testdata.TestSkip[int8](muscom_testdata.Int8TestCases, m, sk, s, t)
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("All MarshalInt, UnmarshalInt, SizeInt, SkipInt functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int](MarshalInt)
					u  = mus.UnmarshalerFn[int](UnmarshalInt)
					s  = mus.SizerFn[int](SizeInt)
					sk = mus.SkipperFn(SkipInt)
				)
				testdata.Test[int](muscom_testdata.IntTestCases, m, u, s, t)
				testdata.TestSkip[int](muscom_testdata.IntTestCases, m, sk, s, t)
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
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("All MarshalByte, UnmarshalByte, SizeByte, SkipByte functions must work correctly",
		func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[byte](MarshalByte)
				u  = mus.UnmarshalerFn[byte](UnmarshalByte)
				s  = mus.SizerFn[byte](SizeByte)
				sk = mus.SkipperFn(SkipByte)
			)
			testdata.Test[byte](muscom_testdata.ByteTestCases, m, u, s, t)
			testdata.TestSkip[byte](muscom_testdata.ByteTestCases, m, sk, s, t)
		})

	t.Run("float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("All MarshalFloat64, UnmarshalFloat64, SizeFloat64, SkipFloat64 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshalerFn[float64](MarshalFloat64)
						u  = mus.UnmarshalerFn[float64](UnmarshalFloat64)
						s  = mus.SizerFn[float64](SizeFloat64)
						sk = mus.SkipperFn(SkipFloat64)
					)
					testdata.Test[float64](muscom_testdata.Float64TestCases, m, u, s, t)
					testdata.TestSkip[float64](muscom_testdata.Float64TestCases, m, sk, s, t)
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
					muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("All MarshalFloat32, UnmarshalFloat32, SizeFloat32, SkipFloat32 functions must work correctly",
				func(t *testing.T) {
					var (
						m  = mus.MarshalerFn[float32](MarshalFloat32)
						u  = mus.UnmarshalerFn[float32](UnmarshalFloat32)
						s  = mus.SizerFn[float32](SizeFloat32)
						sk = mus.SkipperFn(SkipFloat32)
					)
					testdata.Test[float32](muscom_testdata.Float32TestCases, m, u, s, t)
					testdata.TestSkip[float32](muscom_testdata.Float32TestCases, m, sk, s, t)
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
					muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})

		})

	})

}
