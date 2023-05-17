package raw

import (
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
)

func TestRaw(t *testing.T) {

	t.Run("setUpUintFuncs", func(t *testing.T) {

		t.Run("ErrUnsupportedIntSize", func(t *testing.T) {
			wantErr := muscom.ErrUnsupportedIntSize
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					if err != wantErr {
						t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
					}
				}
			}()
			setUpUintFuncs(16)
		})

		t.Run("intSize == 32", func(t *testing.T) {
			setUpUintFuncs(32)
			if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !muscom_testdata.ComparePtrs(sizeUint, sizeNum32[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, skipInteger32) {
				t.Error("unexpected skipUint func")
			}
		})

		t.Run("intSize == 64", func(t *testing.T) {
			setUpUintFuncs(64)
			if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger64[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if !muscom_testdata.ComparePtrs(sizeUint, sizeNum64[uint]) {
				t.Error("unexpected sizeUint func")
			}
			if !muscom_testdata.ComparePtrs(skipUint, skipInteger64) {
				t.Error("unexpected skipUint func")
			}
		})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("ErrUnsupportedIntSize", func(t *testing.T) {
			wantErr := muscom.ErrUnsupportedIntSize
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					if err != wantErr {
						t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
					}
				}
			}()
			setUpIntFuncs(16)
		})

		t.Run("intSize == 32", func(t *testing.T) {
			setUpIntFuncs(32)
			if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !muscom_testdata.ComparePtrs(sizeInt, sizeNum32[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, skipInteger32) {
				t.Error("unexpected skipInt func")
			}
		})

		t.Run("intSize == 64", func(t *testing.T) {
			setUpIntFuncs(64)
			if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger64[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if !muscom_testdata.ComparePtrs(sizeInt, sizeNum64[int]) {
				t.Error("unexpected sizeInt func")
			}
			if !muscom_testdata.ComparePtrs(skipInt, skipInteger64) {
				t.Error("unexpected skipInt func")
			}
		})

	})

	t.Run("unmarshalInteger64 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantV     uint64 = 0
			wantN            = 0
			wantErr          = mus.ErrTooSmallByteSlice
			bs               = []byte{1, 2, 3, 4, 5}
			v, n, err        = unmarshalInteger64[uint64](bs)
		)
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})

	t.Run("skipInteger64 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantN   = 0
			wantErr = mus.ErrTooSmallByteSlice
			bs      = []byte{1, 2, 3, 4, 5, 6, 7}
			n, err  = skipInteger64(bs)
		)
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("unmarshalInteger32 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantV     uint32 = 0
			wantN            = 0
			wantErr          = mus.ErrTooSmallByteSlice
			bs               = []byte{1, 2, 3}
			v, n, err        = unmarshalInteger32[uint32](bs)
		)
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})

	t.Run("skipInteger32 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantN   = 0
			wantErr = mus.ErrTooSmallByteSlice
			bs      = []byte{1, 2, 3}
			n, err  = skipInteger32(bs)
		)
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("unmarshalInteger16 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantV     uint16 = 0
			wantN            = 0
			wantErr          = mus.ErrTooSmallByteSlice
			bs               = []byte{1}
			v, n, err        = unmarshalInteger16[uint16](bs)
		)
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})

	t.Run("skipInteger16 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantN   = 0
			wantErr = mus.ErrTooSmallByteSlice
			bs      = []byte{1}
			n, err  = skipInteger16(bs)
		)
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("unmarshalInteger8 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantV     uint8 = 0
			wantN           = 0
			wantErr         = mus.ErrTooSmallByteSlice
			bs              = []byte{}
			v, n, err       = unmarshalInteger8[uint8](bs)
		)
		muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})

	t.Run("skipInteger8 - ErrTooSmallByteSlice", func(t *testing.T) {
		var (
			wantN   = 0
			wantErr = mus.ErrTooSmallByteSlice
			bs      = []byte{}
			n, err  = skipInteger8(bs)
		)
		muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
	})

	t.Run("byte", func(t *testing.T) {
		var (
			m  = mus.MarshalerFn[byte](MarshalByte)
			u  = mus.UnmarshalerFn[byte](UnmarshalByte)
			s  = mus.SizerFn[byte](SizeByte)
			sk = mus.SkipperFn(SkipByte)
		)
		testdata.Test[byte](muscom_testdata.ByteTestCases, m, u, s, t)
		testdata.TestSkip[byte](muscom_testdata.ByteTestCases, m, sk, s, t)
	})

	t.Run("Unsigned", func(t *testing.T) {

		t.Run("uint64", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint64](MarshalUint64)
				u  = mus.UnmarshalerFn[uint64](UnmarshalUint64)
				s  = mus.SizerFn[uint64](SizeUint64)
				sk = mus.SkipperFn(SkipUint64)
			)
			testdata.Test[uint64](muscom_testdata.Uint64TestCases, m, u, s, t)
			testdata.TestSkip[uint64](muscom_testdata.Uint64TestCases, m, sk, s, t)
		})

		t.Run("uint32", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint32](MarshalUint32)
				u  = mus.UnmarshalerFn[uint32](UnmarshalUint32)
				s  = mus.SizerFn[uint32](SizeUint32)
				sk = mus.SkipperFn(SkipUint32)
			)
			testdata.Test[uint32](muscom_testdata.Uint32TestCases, m, u, s, t)
			testdata.TestSkip[uint32](muscom_testdata.Uint32TestCases, m, sk, s, t)
		})

		t.Run("uint16", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint16](MarshalUint16)
				u  = mus.UnmarshalerFn[uint16](UnmarshalUint16)
				s  = mus.SizerFn[uint16](SizeUint16)
				sk = mus.SkipperFn(SkipUint16)
			)
			testdata.Test[uint16](muscom_testdata.Uint16TestCases, m, u, s, t)
			testdata.TestSkip[uint16](muscom_testdata.Uint16TestCases, m, sk, s, t)
		})

		t.Run("uint8", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[uint8](MarshalUint8)
				u  = mus.UnmarshalerFn[uint8](UnmarshalUint8)
				s  = mus.SizerFn[uint8](SizeUint8)
				sk = mus.SkipperFn(SkipUint8)
			)
			testdata.Test[uint8](muscom_testdata.Uint8TestCases, m, u, s, t)
			testdata.TestSkip[uint8](muscom_testdata.Uint8TestCases, m, sk, s, t)
		})

		t.Run("uint", func(t *testing.T) {
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

	t.Run("Signed", func(t *testing.T) {

		t.Run("int64", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[int64](MarshalInt64)
				u  = mus.UnmarshalerFn[int64](UnmarshalInt64)
				s  = mus.SizerFn[int64](SizeInt64)
				sk = mus.SkipperFn(SkipInt64)
			)
			testdata.Test[int64](muscom_testdata.Int64TestCases, m, u, s, t)
			testdata.TestSkip[int64](muscom_testdata.Int64TestCases, m, sk, s, t)
		})

		t.Run("int32", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[int32](MarshalInt32)
				u  = mus.UnmarshalerFn[int32](UnmarshalInt32)
				s  = mus.SizerFn[int32](SizeInt32)
				sk = mus.SkipperFn(SkipInt32)
			)
			testdata.Test[int32](muscom_testdata.Int32TestCases, m, u, s, t)
			testdata.TestSkip[int32](muscom_testdata.Int32TestCases, m, sk, s, t)
		})

		t.Run("int16", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[int16](MarshalInt16)
				u  = mus.UnmarshalerFn[int16](UnmarshalInt16)
				s  = mus.SizerFn[int16](SizeInt16)
				sk = mus.SkipperFn(SkipInt16)
			)
			testdata.Test[int16](muscom_testdata.Int16TestCases, m, u, s, t)
			testdata.TestSkip[int16](muscom_testdata.Int16TestCases, m, sk, s, t)
		})

		t.Run("int8", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[int8](MarshalInt8)
				u  = mus.UnmarshalerFn[int8](UnmarshalInt8)
				s  = mus.SizerFn[int8](SizeInt8)
				sk = mus.SkipperFn(SkipInt8)
			)
			testdata.Test[int8](muscom_testdata.Int8TestCases, m, u, s, t)
			testdata.TestSkip[int8](muscom_testdata.Int8TestCases, m, sk, s, t)
		})

		t.Run("int", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[int](MarshalInt)
				u  = mus.UnmarshalerFn[int](UnmarshalInt)
				s  = mus.SizerFn[int](SizeInt)
				sk = mus.SkipperFn(SkipInt)
			)
			testdata.Test[int](muscom_testdata.IntTestCases, m, u, s, t)
			testdata.TestSkip[int](muscom_testdata.IntTestCases, m, sk, s, t)
		})

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
				testdata.Test[float64](muscom_testdata.Float64TestCases, m, u, s, t)
				testdata.TestSkip[float64](muscom_testdata.Float64TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
				var (
					wantV     float64 = 0
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					bs                = []byte{1, 2, 3, 4, 5}
					v, n, err         = UnmarshalFloat64(bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

			t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{1, 2, 3, 4, 5}
					n, err  = SkipFloat64(bs)
				)
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
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
				testdata.Test[float32](muscom_testdata.Float32TestCases, m, u, s, t)
				testdata.TestSkip[float32](muscom_testdata.Float32TestCases, m, sk, s, t)
			})

			t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
				var (
					wantV     float32 = 0
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					bs                = []byte{1, 2}
					v, n, err         = UnmarshalFloat32(bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

			t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{1, 2}
					n, err  = SkipFloat32(bs)
				)
				muscom_testdata.TestSkipResults(wantN, n, wantErr, err, t)
			})

		})

	})

}
