package varint

import (
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
)

func TestVarint(t *testing.T) {

	t.Run("unmarshalUint", func(t *testing.T) {

		t.Run("ErrTooSmallByteSlice, empty buf", func(t *testing.T) {
			var (
				wantV     uint64 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{}
				v, n, err        = unmarshalUint[uint64](0, 0, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("ErrOverflow", func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 3
				wantErr          = muscom.ErrOverflow
				bs               = []byte{200, 200, 200}
				v, n, err        = unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
					muscom.Uint16MaxLastByte, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 2
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{200, 200}
				v, n, err        = unmarshalUint[uint16](muscom.Uint16MaxVarintLen,
					muscom.Uint16MaxLastByte, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	})

	t.Run("skipUint", func(t *testing.T) {

		t.Run("ErrTooSmallByteSlice, empty buf", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = skipUint(0, 0, bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("ErrOverflow", func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = muscom.ErrOverflow
				bs      = []byte{200, 200, 200, 200, 200}
				n, err  = skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte,
					bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantN   = 2
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{200, 200}
				n, err  = skipUint(muscom.Uint16MaxVarintLen, muscom.Uint16MaxLastByte,
					bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)

		})

	})

	t.Run("Unsigned", func(t *testing.T) {

		t.Run("uint64", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint64](MarshalUint64)
				u  = mus.UnmarshalerFn[uint64](UnmarshalUint64)
				s  = mus.SizerFn[uint64](SizeUint64)
				sk = mus.SkipperFn(SkipUint64)
			)
			testdata.Test[uint64](testdata.Uint64TestCases, m, u, s, t)
			testdata.TestSkip[uint64](testdata.Uint64TestCases, m, sk, s, t)
		})

		t.Run("uint32", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint32](MarshalUint32)
				u  = mus.UnmarshalerFn[uint32](UnmarshalUint32)
				s  = mus.SizerFn[uint32](SizeUint32)
				sk = mus.SkipperFn(SkipUint32)
			)
			testdata.Test[uint32](testdata.Uint32TestCases, m, u, s, t)
			testdata.TestSkip[uint32](testdata.Uint32TestCases, m, sk, s, t)
		})

		t.Run("uint16", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint16](MarshalUint16)
				u  = mus.UnmarshalerFn[uint16](UnmarshalUint16)
				s  = mus.SizerFn[uint16](SizeUint16)
				sk = mus.SkipperFn(SkipUint16)
			)
			testdata.Test[uint16](testdata.Uint16TestCases, m, u, s, t)
			testdata.TestSkip[uint16](testdata.Uint16TestCases, m, sk, s, t)
		})

		t.Run("uint8", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint8](MarshalUint8)
				u  = mus.UnmarshalerFn[uint8](UnmarshalUint8)
				s  = mus.SizerFn[uint8](SizeUint8)
				sk = mus.SkipperFn(SkipUint8)
			)
			testdata.Test[uint8](testdata.Uint8TestCases, m, u, s, t)
			testdata.TestSkip[uint8](testdata.Uint8TestCases, m, sk, s, t)
		})

		t.Run("uint", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint](MarshalUint)
				u  = mus.UnmarshalerFn[uint](UnmarshalUint)
				s  = mus.SizerFn[uint](SizeUint)
				sk = mus.SkipperFn(SkipUint)
			)
			testdata.Test[uint](testdata.UintTestCases, m, u, s, t)
			testdata.TestSkip[uint](testdata.UintTestCases, m, sk, s, t)
		})

	})

	t.Run("Signed", func(t *testing.T) {

		t.Run("Int64", func(t *testing.T) {

			t.Run("int64", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int64](MarshalInt64)
					u  = mus.UnmarshalerFn[int64](UnmarshalInt64)
					s  = mus.SizerFn[int64](SizeInt64)
					sk = mus.SkipperFn(SkipInt64)
				)
				testdata.Test[int64](testdata.Int64TestCases, m, u, s, t)
				testdata.TestSkip[int64](testdata.Int64TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     int64 = 0
					wantN           = 0
					wantErr         = mus.ErrTooSmallByteSlice
					bs              = []byte{}
					v, n, err       = UnmarshalInt64(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

		t.Run("Int32", func(t *testing.T) {

			t.Run("int32", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int32](MarshalInt32)
					u  = mus.UnmarshalerFn[int32](UnmarshalInt32)
					s  = mus.SizerFn[int32](SizeInt32)
					sk = mus.SkipperFn(SkipInt32)
				)
				testdata.Test[int32](testdata.Int32TestCases, m, u, s, t)
				testdata.TestSkip[int32](testdata.Int32TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     int32 = 0
					wantN           = 0
					wantErr         = mus.ErrTooSmallByteSlice
					bs              = []byte{}
					v, n, err       = UnmarshalInt32(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

		t.Run("Int16", func(t *testing.T) {

			t.Run("int16", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int16](MarshalInt16)
					u  = mus.UnmarshalerFn[int16](UnmarshalInt16)
					s  = mus.SizerFn[int16](SizeInt16)
					sk = mus.SkipperFn(SkipInt16)
				)
				testdata.Test[int16](testdata.Int16TestCases, m, u, s, t)
				testdata.TestSkip[int16](testdata.Int16TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     int16 = 0
					wantN           = 0
					wantErr         = mus.ErrTooSmallByteSlice
					bs              = []byte{}
					v, n, err       = UnmarshalInt16(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

		t.Run("Int8", func(t *testing.T) {

			t.Run("int8", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int8](MarshalInt8)
					u  = mus.UnmarshalerFn[int8](UnmarshalInt8)
					s  = mus.SizerFn[int8](SizeInt8)
					sk = mus.SkipperFn(SkipInt8)
				)
				testdata.Test[int8](testdata.Int8TestCases, m, u, s, t)
				testdata.TestSkip[int8](testdata.Int8TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     int8 = 0
					wantN          = 0
					wantErr        = mus.ErrTooSmallByteSlice
					bs             = []byte{}
					v, n, err      = UnmarshalInt8(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

		t.Run("Int", func(t *testing.T) {

			t.Run("int", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[int](MarshalInt)
					u  = mus.UnmarshalerFn[int](UnmarshalInt)
					s  = mus.SizerFn[int](SizeInt)
					sk = mus.SkipperFn(SkipInt)
				)
				testdata.Test[int](testdata.IntTestCases, m, u, s, t)
				testdata.TestSkip[int](testdata.IntTestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     int = 0
					wantN         = 0
					wantErr       = mus.ErrTooSmallByteSlice
					bs            = []byte{}
					v, n, err     = UnmarshalInt(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

	})

	t.Run("byte", func(t *testing.T) {
		var (
			m  = mus.MarshalerFn[byte](MarshalByte)
			u  = mus.UnmarshalerFn[byte](UnmarshalByte)
			s  = mus.SizerFn[byte](SizeByte)
			sk = mus.SkipperFn(SkipByte)
		)
		testdata.Test[byte](testdata.ByteTestCases, m, u, s, t)
		testdata.TestSkip[byte](testdata.ByteTestCases, m, sk, s, t)
	})

	t.Run("Float", func(t *testing.T) {

		t.Run("Float64", func(t *testing.T) {

			t.Run("float64", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[float64](MarshalFloat64)
					u  = mus.UnmarshalerFn[float64](UnmarshalFloat64)
					s  = mus.SizerFn[float64](SizeFloat64)
					sk = mus.SkipperFn(SkipFloat64)
				)
				testdata.Test[float64](testdata.Float64TestCases, m, u, s, t)
				testdata.TestSkip[float64](testdata.Float64TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     float64 = 0
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					bs                = []byte{}
					v, n, err         = UnmarshalFloat64(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

		t.Run("Float32", func(t *testing.T) {

			t.Run("float32", func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[float32](MarshalFloat32)
					u  = mus.UnmarshalerFn[float32](UnmarshalFloat32)
					s  = mus.SizerFn[float32](SizeFloat32)
					sk = mus.SkipperFn(SkipFloat32)
				)
				testdata.Test[float32](testdata.Float32TestCases, m, u, s, t)
				testdata.TestSkip[float32](testdata.Float32TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - error", func(t *testing.T) {
				var (
					wantV     float32 = 0
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					bs                = []byte{}
					v, n, err         = UnmarshalFloat32(bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		})

	})

}
