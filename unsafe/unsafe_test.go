package unsafe

import (
	"errors"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	muscom_mock "github.com/mus-format/mus-common-go/testdata/mock"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/testdata"
)

func TestUnsafe(t *testing.T) {

	t.Run("setUpUintFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpUintFuncs should panic with ErrUnsupportedIntSize",
			func(t *testing.T) {
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

		t.Run("If the system int size is equal to 32, setUpUintFuncs should initialize the uint functions with 32-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(32)
				if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !muscom_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
					t.Error("unexpected sizeUint func")
				}
				if !muscom_testdata.ComparePtrs(skipUint, raw.SkipUint) {
					t.Error("unexpected skipUint func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(64)
				if !muscom_testdata.ComparePtrs(marshalUint, marshalInteger64[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !muscom_testdata.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if !muscom_testdata.ComparePtrs(sizeUint, raw.SizeUint) {
					t.Error("unexpected sizeUint func")
				}
				if !muscom_testdata.ComparePtrs(skipUint, raw.SkipUint) {
					t.Error("unexpected skipUint func")
				}
			})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize", func(t *testing.T) {
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

		t.Run("If the system int size is equal to 32, setUpIntFuncs should initialize the uint functions with 32-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(32)
				if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !muscom_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
					t.Error("unexpected sizeInt func")
				}
				if !muscom_testdata.ComparePtrs(skipInt, raw.SkipInt) {
					t.Error("unexpected skipInt func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(64)
				if !muscom_testdata.ComparePtrs(marshalInt, marshalInteger64[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !muscom_testdata.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if !muscom_testdata.ComparePtrs(sizeInt, raw.SizeInt) {
					t.Error("unexpected sizeInt func")
				}
				if !muscom_testdata.ComparePtrs(skipInt, raw.SkipInt) {
					t.Error("unexpected skipInt func")
				}
			})

	})

	t.Run("unmarshalInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint64 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1, 2, 3, 4, 5, 6, 7}
				v, n, err        = unmarshalInteger64[uint64](bs)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint32 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1, 2, 3}
				v, n, err        = unmarshalInteger32[uint32](bs)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1}
				v, n, err        = unmarshalInteger16[uint16](bs)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint8 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = unmarshalInteger8[uint8](bs)
			)
			muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("string", func(t *testing.T) {

		t.Run("All MarshalString, UnmarshalString, SizeString, SkipString functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[string](MarshalString)
					u  = mus.UnmarshalerFn[string](UnmarshalString)
					s  = mus.SizerFn[string](SizeString)
					sk = mus.SkipperFn(SkipString)
				)
				testdata.Test[string](muscom_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](muscom_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("Marshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					s       = "hello world"
					bs      = make([]byte, 4)
					wantErr = mus.ErrTooSmallByteSlice
				)
				defer func() {
					if r := recover(); r != nil {
						err := r.(error)
						if err != wantErr {
							t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
						}
					}
				}()

				MarshalString(s, bs)
			})

		t.Run("If UnmarshalValidString fails to unmarshal a length, it should return an error",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 0
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{}
					v, n, err = UnmarshalValidString(nil, false, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalValidString should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = muscom.ErrNegativeLength
					bs        = []byte{1}
					v, n, err = UnmarshalValidString(nil, false, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalValidString should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{6, 1, 1}
					v, n, err = UnmarshalValidString(nil, false, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If skip == true and MaxLength validator returns an error, UnmarshalValidString should return it",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 4
					wantErr   = errors.New("MaxLength validator error")
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					bs        = []byte{6, 1, 1, 1}
					v, n, err = UnmarshalValidString(maxLength, true, bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If skip == false and MaxLength validator returns an error, UnmarshalValidString should return it",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = errors.New("MaxLength validator error")
					maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					bs        = []byte{6, 1, 1, 1}
					v, n, err = UnmarshalValidString(maxLength, false, bs)
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

		t.Run("All MarshalUint16, UnmarshalUint16, SizeUint16, SkipUint16 functions must work correctly",
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

		t.Run("All MarshalUint, UnmarshalUint, SizeUint, SkipUint functions must work correctly",
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
					testdata.TestSkip[float64](muscom_testdata.Float64TestCases, m, sk, s,
						t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0.0
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

			t.Run("UnmarshalFloat32 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0.0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{}
						v, n, err         = UnmarshalFloat32(bs)
					)
					muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

		})

	})

	t.Run("bool", func(t *testing.T) {

		t.Run("All MarshalBool, UnmarshalBool, SizeBool, SkipBool functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshalerFn[bool](MarshalBool)
					u  = mus.UnmarshalerFn[bool](UnmarshalBool)
					s  = mus.SizerFn[bool](SizeBool)
					sk = mus.SkipperFn(SkipBool)
				)
				testdata.Test[bool](muscom_testdata.BoolTestCases, m, u, s, t)
				testdata.TestSkip[bool](muscom_testdata.BoolTestCases, m, sk, s, t)
			})

		t.Run("UnmarshalBool should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     bool = false
					wantN          = 0
					wantErr        = mus.ErrTooSmallByteSlice
					bs             = []byte{}
					v, n, err      = UnmarshalBool(bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV     bool = false
					wantN          = 0
					wantErr        = muscom.ErrWrongFormat
					bs             = []byte{2}
					v, n, err      = UnmarshalBool(bs)
				)
				muscom_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

	})

}
