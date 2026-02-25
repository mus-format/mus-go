package unsafe

import (
	"errors"
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	cmock "github.com/mus-format/common-go/testutil/mock"
	"github.com/mus-format/mus-go"
	arrops "github.com/mus-format/mus-go/options/array"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/testutil/mock"
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
				if !ctestutil.ComparePtrs(marshalUint, marshalInteger32[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !ctestutil.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if sizeUint != com.Num32RawSize {
					t.Error("unexpected sizeUint value")
				}
				if !ctestutil.ComparePtrs(skipUint, raw.SkipInteger32) {
					t.Error("unexpected skipUint func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpUintFuncs(64)
				if !ctestutil.ComparePtrs(marshalUint, marshalInteger64[uint]) {
					t.Error("unexpected marshalUint func")
				}
				if !ctestutil.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
					t.Error("unexpected unmarshalUint func")
				}
				if sizeUint != com.Num64RawSize {
					t.Error("unexpected sizeUint value")
				}
				if !ctestutil.ComparePtrs(skipUint, raw.SkipInteger64) {
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
				if !ctestutil.ComparePtrs(marshalInt, marshalInteger32[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !ctestutil.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if sizeInt != com.Num32RawSize {
					t.Error("unexpected sizeInt value")
				}
				if !ctestutil.ComparePtrs(skipInt, raw.SkipInteger32) {
					t.Error("unexpected skipInt func")
				}
			})

		t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
			func(t *testing.T) {
				setUpIntFuncs(64)
				if !ctestutil.ComparePtrs(marshalInt, marshalInteger64[int]) {
					t.Error("unexpected marshalInt func")
				}
				if !ctestutil.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
					t.Error("unexpected unmarshalInt func")
				}
				if sizeInt != com.Num64RawSize {
					t.Error("unexpected sizeInt value")
				}
				if !ctestutil.ComparePtrs(skipInt, raw.SkipInteger64) {
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("string", func(t *testing.T) {
		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := String
				testutil.Test[string](ctestutil.StringTestCases, ser, t)
				testutil.TestSkip[string](ctestutil.StringTestCases, ser, t)
			})

		t.Run("We should be able to set a length serializer",
			func(t *testing.T) {
				var (
					str, lenSer = testutil.StringLenTestData(t)
					ser         = NewStringSer(strops.WithLenSer(lenSer))
				)
				testutil.Test[string]([]string{str}, ser, t)
				testutil.TestSkip[string]([]string{str}, ser, t)
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

		t.Run("If the length serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 0
					wantErr   = mus.ErrTooSmallByteSlice
					v, n, err = String.Unmarshal(nil)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					v, n, err = String.Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := NewValidStringSer(nil)
				testutil.Test[string](ctestutil.StringTestCases, ser, t)
				testutil.TestSkip[string](ctestutil.StringTestCases, ser, t)
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
					v, n, err = NewValidStringSer(strops.WithLenSer(lenSer)).Unmarshal(nil)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					v, n, err = NewValidStringSer(nil).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("lenVl validator error")
					lenVl   = cmock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							wantV := 3
							if v != wantV {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
							}
							return wantErr
						},
					)
					bs        = []byte{3, 1, 1, 1}
					v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
	})

	t.Run("byte", func(t *testing.T) {
		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testutil.Test[byte](ctestutil.ByteTestCases, ser, t)
				testutil.TestSkip[byte](ctestutil.ByteTestCases, ser, t)
			})
	})

	t.Run("unsigned", func(t *testing.T) {
		t.Run("Uint64 serializer should work correctly", func(t *testing.T) {
			ser := Uint64
			testutil.Test[uint64](ctestutil.Uint64TestCases, ser, t)
			testutil.TestSkip[uint64](ctestutil.Uint64TestCases, ser, t)
		})

		t.Run("Uint32 serializer should work correctly", func(t *testing.T) {
			ser := Uint32
			testutil.Test[uint32](ctestutil.Uint32TestCases, ser, t)
			testutil.TestSkip[uint32](ctestutil.Uint32TestCases, ser, t)
		})

		t.Run("Uint16 serializer should work correctly", func(t *testing.T) {
			ser := Uint16
			testutil.Test[uint16](ctestutil.Uint16TestCases, ser, t)
			testutil.TestSkip[uint16](ctestutil.Uint16TestCases, ser, t)
		})

		t.Run("Uint8 serializer should work correctly", func(t *testing.T) {
			ser := Uint8
			testutil.Test[uint8](ctestutil.Uint8TestCases, ser, t)
			testutil.TestSkip[uint8](ctestutil.Uint8TestCases, ser, t)
		})

		t.Run("Uint serializer should work correctly", func(t *testing.T) {
			ser := Uint
			testutil.Test[uint](ctestutil.UintTestCases, ser, t)
			testutil.TestSkip[uint](ctestutil.UintTestCases, ser, t)
		})
	})

	t.Run("signed", func(t *testing.T) {
		t.Run("Int64 serializer should work correctly", func(t *testing.T) {
			ser := Int64
			testutil.Test[int64](ctestutil.Int64TestCases, ser, t)
			testutil.TestSkip[int64](ctestutil.Int64TestCases, ser, t)
		})

		t.Run("Int32 serializer should work correctly", func(t *testing.T) {
			ser := Int32
			testutil.Test[int32](ctestutil.Int32TestCases, ser, t)
			testutil.TestSkip[int32](ctestutil.Int32TestCases, ser, t)
		})

		t.Run("Int16 serializer should work correctly", func(t *testing.T) {
			ser := Int16
			testutil.Test[int16](ctestutil.Int16TestCases, ser, t)
			testutil.TestSkip[int16](ctestutil.Int16TestCases, ser, t)
		})

		t.Run("Int8 serializer should work correctly", func(t *testing.T) {
			ser := Int8
			testutil.Test[int8](ctestutil.Int8TestCases, ser, t)
			testutil.TestSkip[int8](ctestutil.Int8TestCases, ser, t)
		})

		t.Run("Int serializer should work correctly", func(t *testing.T) {
			ser := Int
			testutil.Test[int](ctestutil.IntTestCases, ser, t)
			testutil.TestSkip[int](ctestutil.IntTestCases, ser, t)
		})
	})

	t.Run("float", func(t *testing.T) {
		t.Run("float64", func(t *testing.T) {
			t.Run("Float64 serializer should work correctly", func(t *testing.T) {
				ser := Float64
				testutil.Test[float64](ctestutil.Float64TestCases, ser, t)
				testutil.TestSkip[float64](ctestutil.Float64TestCases, ser, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})
		})

		t.Run("float32", func(t *testing.T) {
			t.Run("Float32 serializer should work correctly", func(t *testing.T) {
				ser := Float32
				testutil.Test[float32](ctestutil.Float32TestCases, ser, t)
				testutil.TestSkip[float32](ctestutil.Float32TestCases, ser, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})
		})
	})

	t.Run("bool", func(t *testing.T) {
		t.Run("Bool serializer should work correctly", func(t *testing.T) {
			ser := Bool
			testutil.Test[bool](ctestutil.BoolTestCases, ser, t)
			testutil.TestSkip[bool](ctestutil.BoolTestCases, ser, t)
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})
	})

	t.Run("byte_slice", func(t *testing.T) {
		t.Run("ByteSlice serializer should work correctly with empty slice",
			func(t *testing.T) {
				ser := ByteSlice
				testutil.Test[[]byte]([][]byte{{}}, ser, t)
				testutil.TestSkip[[]byte]([][]byte{{}}, ser, t)
			})

		t.Run("ByteSlice serializer should work correctly with non-empty slice",
			func(t *testing.T) {
				ser := ByteSlice
				testutil.Test[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
				testutil.TestSkip[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
			})

		t.Run("We should be able to set a length serializer", func(t *testing.T) {
			var (
				sl, lenSer = testutil.ByteSliceLenTestData(t)
				ser        = NewByteSliceSer(bslops.WithLenSer(lenSer))
			)
			testutil.Test[[]byte]([][]byte{sl}, ser, t)
			testutil.TestSkip[[]byte]([][]byte{sl}, ser, t)
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

		t.Run("If the length serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 0
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = ByteSlice.Unmarshal(nil)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("Unmarshal should return ErrTooSmallByteSlice if bs is too small for slice content",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 1
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = ByteSlice.Unmarshal([]byte{2, 1}) // Length 2 but only 1 byte available
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
					v, n, err        = ByteSlice.Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = ByteSlice.Skip([]byte{})
				)
				ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
					n, err    = ByteSlice.Skip(bs)
				)
				ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly with empty slice",
			func(t *testing.T) {
				ser := NewValidByteSliceSer(nil)
				testutil.Test[[]byte]([][]byte{{}}, ser, t)
				testutil.TestSkip[[]byte]([][]byte{{}}, ser, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly with non-empty slice",
			func(t *testing.T) {
				ser := NewValidByteSliceSer(nil)
				testutil.Test[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
				testutil.TestSkip[[]byte]([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
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
					v, n, err = NewValidByteSliceSer(bslops.WithLenSer(lenSer)).Unmarshal(nil)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
					v, n, err        = NewValidByteSliceSer(nil).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl fails with an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 1
					wantErr        = errors.New("too large slice")
					bs             = []byte{3, 4, 1, 1}
					lenVl          = cmock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					v, n, err = NewValidByteSliceSer(bslops.WithLenValidator(lenVl)).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})
	})

	t.Run("time", func(t *testing.T) {
		os.Setenv("TZ", "")

		t.Run("time_unix_utc", func(t *testing.T) {
			t.Run("TimeUnixUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						sec = time.Now().Unix()
						tm  = time.Unix(sec, 0)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnix, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     = time.Time{}
						wantN     = 0
						wantErr   = mus.ErrTooSmallByteSlice
						bs        = []byte{}
						v, n, err = TimeUnixUTC.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_milli_UTC", func(t *testing.T) {
			t.Run("TimeUnixMilliUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMilli()
						tm    = time.UnixMilli(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     = time.Time{}
						wantN     = 0
						wantErr   = mus.ErrTooSmallByteSlice
						bs        = []byte{}
						v, n, err = TimeUnixMilliUTC.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_micro_utc", func(t *testing.T) {
			t.Run("TimeUnixMicroUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMicro()
						tm    = time.UnixMicro(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixMicroUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixMicroUTC, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     = time.Time{}
						wantN     = 0
						wantErr   = mus.ErrTooSmallByteSlice
						bs        = []byte{}
						v, n, err = TimeUnixMicroUTC.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_nano_utc", func(t *testing.T) {
			t.Run("TimeUnixNanoUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						nano = time.Now().UnixNano()
						tm   = time.Unix(0, nano)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     = time.Time{}
						wantN     = 0
						wantErr   = mus.ErrTooSmallByteSlice
						bs        = []byte{}
						v, n, err = TimeUnixNanoUTC.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})
	})

	t.Run("array", func(t *testing.T) {
		t.Run("Array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArrayTestData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewArraySer[[3]int, int](elemSer)
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV     [3]int = [3]int{0, 0, 0}
					wantN            = 1
					wantErr          = com.ErrTooLargeLength
					bs               = []byte{4}
					v, n, err        = NewArraySer[[3]int, int](nil).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})

		t.Run("Valid array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArrayTestData(t)
				mocks        = []*mok.Mock{elemSer.Mock}
				ser          = NewValidArraySer[[3]int, int](elemSer, nil)
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("Valid Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV     [3]int = [3]int{0, 0, 0}
					wantN            = 1
					wantErr          = com.ErrTooLargeLength
					bs               = []byte{4}
					v, n, err        = NewValidArraySer[[3]int, int](nil, nil).Unmarshal(bs)
				)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})

		t.Run("If elemVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV    [3]int = [3]int{0, 0, 0}
					wantElem        = 11
					wantN           = 2
					wantErr         = errors.New("elemVl error")
					bs              = []byte{3, 11}
					elemSer         = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							return 11, 1, nil
						},
					)
					elemVl = cmock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantElem {
								t.Errorf("unexpected v, want %v actual %v", wantElem, v)
							}
							return wantErr
						},
					)
					ser   = NewValidArraySer[[3]int, int](elemSer, arrops.WithElemValidator[int](elemVl))
					mocks = []*mok.Mock{elemSer.Mock, elemVl.Mock}
				)
				v, n, err := ser.Unmarshal(bs)
				ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("We should be able to set a length serializer", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArrayTestData(t)
				_, lenSer    = testutil.ArrayLenTestData(t)
				ser          = NewArraySer[[3]int, int](elemSer, arrops.WithLenSer[int](lenSer))
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)
		})

		t.Run("Valid array: We should be able to set a length serializer", func(t *testing.T) {
			var (
				arr, elemSer = testutil.ArrayTestData(t)
				_, lenSer    = testutil.ArrayLenTestData(t)
				ser          = NewValidArraySer[[3]int, int](elemSer, arrops.WithLenSer[int](lenSer))
			)
			testutil.Test[[3]int]([][3]int{arr}, ser, t)
			testutil.TestSkip[[3]int]([][3]int{arr}, ser, t)
		})
	})
}

func NegativeLengthBs() (n int, bs []byte) {
	n = varint.PositiveInt.Size(-1)
	bs = make([]byte, n)
	varint.PositiveInt.Marshal(-1, bs)
	return
}
