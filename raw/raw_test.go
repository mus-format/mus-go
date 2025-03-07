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
				if !com_testdata.ComparePtrs(skipUint, SkipInteger32) {
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
				if !com_testdata.ComparePtrs(skipUint, SkipInteger64) {
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
				if !com_testdata.ComparePtrs(skipInt, SkipInteger32) {
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
				if !com_testdata.ComparePtrs(skipInt, SkipInteger64) {
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
				n, err  = SkipInteger64(bs)
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
				n, err  = SkipInteger32(bs)
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
				n, err  = SkipInteger16(bs)
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
				n, err  = SkipInteger8(bs)
			)
			com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("byte", func(t *testing.T) {

		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testdata.Test[byte](com_testdata.ByteTestCases, ser, t)
				testdata.TestSkip[byte](com_testdata.ByteTestCases, ser, t)
			})

	})

	t.Run("unsigned", func(t *testing.T) {

		t.Run("Uint64 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint64
				testdata.Test[uint64](com_testdata.Uint64TestCases, ser, t)
				testdata.TestSkip[uint64](com_testdata.Uint64TestCases, ser, t)
			})

		t.Run("Uint32 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint32
				testdata.Test[uint32](com_testdata.Uint32TestCases, ser, t)
				testdata.TestSkip[uint32](com_testdata.Uint32TestCases, ser, t)
			})

		t.Run("Uint16 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint16
				testdata.Test[uint16](com_testdata.Uint16TestCases, ser, t)
				testdata.TestSkip[uint16](com_testdata.Uint16TestCases, ser, t)
			})

		t.Run("Uint8 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint8
				testdata.Test[uint8](com_testdata.Uint8TestCases, ser, t)
				testdata.TestSkip[uint8](com_testdata.Uint8TestCases, ser, t)
			})

		t.Run("Uint serializer should work correctly",
			func(t *testing.T) {
				ser := Uint
				testdata.Test[uint](com_testdata.UintTestCases, ser, t)
				testdata.TestSkip[uint](com_testdata.UintTestCases, ser, t)
			})

	})

	t.Run("signed", func(t *testing.T) {

		t.Run("Int64 serializer should work correctly",
			func(t *testing.T) {
				ser := Int64
				testdata.Test[int64](com_testdata.Int64TestCases, ser, t)
				testdata.TestSkip[int64](com_testdata.Int64TestCases, ser, t)
			})

		t.Run("Int32 serializer should work correctly",
			func(t *testing.T) {
				ser := Int32
				testdata.Test[int32](com_testdata.Int32TestCases, ser, t)
				testdata.TestSkip[int32](com_testdata.Int32TestCases, ser, t)
			})

		t.Run("Int16 serializer should work correctly",
			func(t *testing.T) {
				ser := Int16
				testdata.Test[int16](com_testdata.Int16TestCases, ser, t)
				testdata.TestSkip[int16](com_testdata.Int16TestCases, ser, t)
			})

		t.Run("Int8 serializer should work correctly",
			func(t *testing.T) {
				ser := Int8
				testdata.Test[int8](com_testdata.Int8TestCases, ser, t)
				testdata.TestSkip[int8](com_testdata.Int8TestCases, ser, t)
			})

		t.Run("Int serializer should work correctly",
			func(t *testing.T) {
				ser := Int
				testdata.Test[int](com_testdata.IntTestCases, ser, t)
				testdata.TestSkip[int](com_testdata.IntTestCases, ser, t)
			})

	})

	t.Run("float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("Float64 serializer should work correctly",
				func(t *testing.T) {
					ser := Float64
					testdata.Test[float64](com_testdata.Float64TestCases, ser, t)
					testdata.TestSkip[float64](com_testdata.Float64TestCases, ser, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2, 3, 4, 5}
						v, n, err         = Float64.Unmarshal(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2, 3, 4, 5}
						n, err  = Float64.Skip(bs)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
				})

		})

		t.Run("float32", func(t *testing.T) {

			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testdata.Test[float32](com_testdata.Float32TestCases, ser, t)
					testdata.TestSkip[float32](com_testdata.Float32TestCases, ser, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2}
						v, n, err         = Float32.Unmarshal(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2}
						n, err  = Float32.Skip(bs)
					)
					com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
				})

		})

	})

}
