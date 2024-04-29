package raw

import (
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
)

func TestRaw(t *testing.T) {

	t.Run("setUpUintFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpUintFuncs should panic with ErrUnsupportedIntSize",
			func(t *testing.T) {
				wantErr := com.ErrUnsupportedIntSize
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
				if !com_testdata.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !com_testdata.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if sizeUint != com.Num32RawSize {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, skipInteger32) {
					t.Error("unexpected skipUint func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(64)
				if !com_testdata.ComparePtrs(marshalUint, marshalInteger64[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !com_testdata.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if sizeUint != com.Num64RawSize {
					t.Error("unexpected sizeUint func")
				}
				if !com_testdata.ComparePtrs(skipUint, skipInteger64) {
					t.Error("unexpected skipUint func")
				}
			})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize",
			func(t *testing.T) {
				wantErr := com.ErrUnsupportedIntSize
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
				if !com_testdata.ComparePtrs(marshalInt, marshalInteger32[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !com_testdata.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if sizeInt != com.Num32RawSize {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, skipInteger32) {
					t.Error("unexpected skipInt func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(64)
				if !com_testdata.ComparePtrs(marshalInt, marshalInteger64[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !com_testdata.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if sizeInt != com.Num64RawSize {
					t.Error("unexpected sizeInt func")
				}
				if !com_testdata.ComparePtrs(skipInt, skipInteger64) {
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
				bs               = []byte{1, 2, 3, 4, 5}
				v, n, err        = unmarshalInteger64[uint64](bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3, 4, 5, 6, 7}
				n, err  = skipInteger64(bs)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3}
				n, err  = skipInteger32(bs)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1}
				n, err  = skipInteger16(bs)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = skipInteger8(bs)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("All MarshalByte, UnmarshalByte, SizeByte, SkipByte functions must work correctly", func(t *testing.T) {
		var (
			m  = mus.MarshallerFn[byte](MarshalByte)
			u  = mus.UnmarshallerFn[byte](UnmarshalByte)
			s  = mus.SizerFn[byte](SizeByte)
			sk = mus.SkipperFn(SkipByte)
		)
		testdata.Test[byte](com_testdata.ByteTestCases, m, u, s, t)
		testdata.TestSkip[byte](com_testdata.ByteTestCases, m, sk, s, t)
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

		t.Run("All MarshalUint16, UnmarshalUint16, SizeUint16, SkipUint16 functions must work correctly",
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

		t.Run("All MarshalUint, UnmarshalUint, SizeUint, SkipUint functions must work correctly",
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

			t.Run("UnmarshalFloat64 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2, 3, 4, 5}
						v, n, err         = UnmarshalFloat64(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("SkipFloat64 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2, 3, 4, 5}
						n, err  = SkipFloat64(bs)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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

			t.Run("UnmarshalFloat32 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2}
						v, n, err         = UnmarshalFloat32(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("SkipFloat32 should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2}
						n, err  = SkipFloat32(bs)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
				})

		})

	})

}
