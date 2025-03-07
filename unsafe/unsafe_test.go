package unsafe

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	com_mock "github.com/mus-format/common-go/testdata/mock"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/testdata/mock"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe(t *testing.T) {

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
					t.Error("unexpected sizeUint value")
				}
				if !com_testdata.ComparePtrs(skipUint, raw.SkipInteger32) {
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
					t.Error("unexpected sizeUint value")
				}
				if !com_testdata.ComparePtrs(skipUint, raw.SkipInteger64) {
					t.Error("unexpected skipUint func")
				}
			})

	})

	t.Run("setUpIntFuncs", func(t *testing.T) {

		t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize", func(t *testing.T) {
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
					t.Error("unexpected sizeInt value")
				}
				if !com_testdata.ComparePtrs(skipInt, raw.SkipInteger32) {
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
					t.Error("unexpected sizeInt value")
				}
				if !com_testdata.ComparePtrs(skipInt, raw.SkipInteger64) {
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
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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

	t.Run("string", func(t *testing.T) {

		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := String
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
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

				String.Marshal(s, bs)
			})

		t.Run("If lenSer fails to unmarshal length, Unmarshal should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							return 0, 0, wantErr
						},
					)
					v, n, err = NewStringSerWith(lenSer).Unmarshal(nil)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					v, n, err = String.Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{3, 1, 1}
					v, n, err = String.Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := NewValidStringSer(nil)
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
			})

		t.Run("If lenSer fails to unmarshal length, Unmarshal should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							return 0, 0, wantErr
						},
					)
					v, n, err = NewValidStringSerWith(lenSer, nil).Unmarshal(nil)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					v, n, err = NewValidStringSer(nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{3, 1, 1}
					v, n, err = NewValidStringSer(nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("lenVl validator error")
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							var wantV = 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					bs        = []byte{3, 1, 1, 1}
					v, n, err = NewValidStringSer(lenVl).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
			var (
				wantV                        = ""
				wantN                        = 1
				wantErr                      = errors.New("empty string")
				bs                           = []byte{0}
				lenVl   com.ValidatorFn[int] = func(t int) (err error) {
					return wantErr
				}
				v, n, err = NewValidStringSer(lenVl).Unmarshal(bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

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

		t.Run("Uint64 serializer should work correctly", func(t *testing.T) {
			ser := Uint64
			testdata.Test[uint64](com_testdata.Uint64TestCases, ser, t)
			testdata.TestSkip[uint64](com_testdata.Uint64TestCases, ser, t)
		})

		t.Run("Uint32 serializer should work correctly", func(t *testing.T) {
			ser := Uint32
			testdata.Test[uint32](com_testdata.Uint32TestCases, ser, t)
			testdata.TestSkip[uint32](com_testdata.Uint32TestCases, ser, t)
		})

		t.Run("Uint16 serializer should work correctly", func(t *testing.T) {
			ser := Uint16
			testdata.Test[uint16](com_testdata.Uint16TestCases, ser, t)
			testdata.TestSkip[uint16](com_testdata.Uint16TestCases, ser, t)
		})

		t.Run("Uint8 serializer should work correctly", func(t *testing.T) {
			ser := Uint8
			testdata.Test[uint8](com_testdata.Uint8TestCases, ser, t)
			testdata.TestSkip[uint8](com_testdata.Uint8TestCases, ser, t)
		})

		t.Run("Uint serializer should work correctly", func(t *testing.T) {
			ser := Uint
			testdata.Test[uint](com_testdata.UintTestCases, ser, t)
			testdata.TestSkip[uint](com_testdata.UintTestCases, ser, t)
		})

	})

	t.Run("signed", func(t *testing.T) {

		t.Run("Int64 serializer should work correctly", func(t *testing.T) {
			ser := Int64
			testdata.Test[int64](com_testdata.Int64TestCases, ser, t)
			testdata.TestSkip[int64](com_testdata.Int64TestCases, ser, t)
		})

		t.Run("Int32 serializer should work correctly", func(t *testing.T) {
			ser := Int32
			testdata.Test[int32](com_testdata.Int32TestCases, ser, t)
			testdata.TestSkip[int32](com_testdata.Int32TestCases, ser, t)
		})

		t.Run("Int16 serializer should work correctly", func(t *testing.T) {
			ser := Int16
			testdata.Test[int16](com_testdata.Int16TestCases, ser, t)
			testdata.TestSkip[int16](com_testdata.Int16TestCases, ser, t)
		})

		t.Run("Int8 serializer should work correctly", func(t *testing.T) {
			ser := Int8
			testdata.Test[int8](com_testdata.Int8TestCases, ser, t)
			testdata.TestSkip[int8](com_testdata.Int8TestCases, ser, t)
		})

		t.Run("Int serializer should work correctly", func(t *testing.T) {
			ser := Int
			testdata.Test[int](com_testdata.IntTestCases, ser, t)
			testdata.TestSkip[int](com_testdata.IntTestCases, ser, t)
		})

	})

	t.Run("float", func(t *testing.T) {

		t.Run("float64", func(t *testing.T) {

			t.Run("Float64 serializer should work correctly", func(t *testing.T) {
				ser := Float64
				testdata.Test[float64](com_testdata.Float64TestCases, ser, t)
				testdata.TestSkip[float64](com_testdata.Float64TestCases, ser, t)
			})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0.0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{}
						v, n, err         = Float64.Unmarshal(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})
		})

		t.Run("float32", func(t *testing.T) {

			t.Run("Float32 serializer should work correctly", func(t *testing.T) {
				ser := Float32
				testdata.Test[float32](com_testdata.Float32TestCases, ser, t)
				testdata.TestSkip[float32](com_testdata.Float32TestCases, ser, t)
			})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0.0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{}
						v, n, err         = Float32.Unmarshal(bs)
					)
					com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})
		})

	})

	t.Run("bool", func(t *testing.T) {

		t.Run("Bool serializer should work correctly", func(t *testing.T) {
			ser := Bool
			testdata.Test[bool](com_testdata.BoolTestCases, ser, t)
			testdata.TestSkip[bool](com_testdata.BoolTestCases, ser, t)
		})

		t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     bool = false
					wantN          = 0
					wantErr        = mus.ErrTooSmallByteSlice
					bs             = []byte{}
					v, n, err      = Bool.Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV     bool = false
					wantN          = 0
					wantErr        = com.ErrWrongFormat
					bs             = []byte{2}
					v, n, err      = Bool.Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("byte_slice", func(t *testing.T) {

		t.Run("ByteSlice serializer should work correctly with empty slice",
			func(t *testing.T) {
				ser := ByteSlice
				testdata.Test[[]byte]([][]byte{{}}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{{}}, ser, t)
			})

		t.Run("ByteSlice serializer should work correctly with non-empty slice",
			func(t *testing.T) {
				ser := ByteSlice
				testdata.Test[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
			})

		t.Run("Marshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					s       = []byte{1, 2, 3, 4}
					bs      = make([]byte, 2)
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
				ByteSlice.Marshal(s, bs)
			})

		t.Run("If lenSer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 0
					wantErr        = errors.New("lenSer error")
					lenSer         = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							return 0, 0, wantErr
						},
					)
					mocks     = []*mok.Mock{lenSer.Mock}
					v, n, err = NewByteSliceSerWith(lenSer).Unmarshal(nil)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Unmarshal should return ErrTooSmallByteSlice if bs is too small for slice content",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 1
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = ByteSlice.Unmarshal([]byte{2, 1}) // Length 2 but only 1 byte available
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
					v, n, err        = ByteSlice.Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = ByteSlice.Skip([]byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					n, err    = ByteSlice.Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly with empty slice",
			func(t *testing.T) {
				ser := NewValidByteSliceSer(nil)
				testdata.Test[[]byte]([][]byte{{}}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{{}}, ser, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly with non-empty slice",
			func(t *testing.T) {
				ser := NewValidByteSliceSer(nil)
				testdata.Test[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
			})

		t.Run("If lenSer fails with an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 0
					wantErr        = errors.New("lenSer error")
					lenSer         = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							return 0, 0, wantErr
						},
					)
					mocks     = []*mok.Mock{lenSer.Mock}
					v, n, err = NewValidByteSliceSerWith(lenSer, nil).Unmarshal(nil)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if bs is too small for slice content",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 1
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = NewValidByteSliceSer(nil).Unmarshal([]byte{2, 1}) // Length 2 but only 1 byte available
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
					v, n, err        = NewValidByteSliceSer(nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl fails with an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 1
					wantErr        = errors.New("too large slice")
					bs             = []byte{3, 4, 1, 1}
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					v, n, err = NewValidByteSliceSer(lenVl).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

	})

}

func NegativeLengthBs() (n int, bs []byte) {
	n = varint.PositiveInt.Size(-1)
	bs = make([]byte, n)
	varint.PositiveInt.Marshal(-1, bs)
	return
}
