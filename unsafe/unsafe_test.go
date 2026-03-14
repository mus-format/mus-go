package unsafe

import (
	"errors"
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	cmock "github.com/mus-format/common-go/test/mock"
	"github.com/mus-format/mus-go"
	arropts "github.com/mus-format/mus-go/options/array"
	bslopts "github.com/mus-format/mus-go/options/byte_slice"
	stropts "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/test/mock"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe_setUpUintFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(marshalUint, marshalInteger32[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !ctest.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if sizeUint != com.Num32RawSize {
				t.Error("unexpected sizeUint value")
			}
			if !ctest.ComparePtrs(skipUint, raw.SkipInteger32) {
				t.Error("unexpected skipUint func")
			}
		})

	t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpUintFuncs(64)
			if !ctest.ComparePtrs(marshalUint, marshalInteger64[uint]) {
				t.Error("unexpected marshalUint func")
			}
			if !ctest.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]) {
				t.Error("unexpected unmarshalUint func")
			}
			if sizeUint != com.Num64RawSize {
				t.Error("unexpected sizeUint value")
			}
			if !ctest.ComparePtrs(skipUint, raw.SkipInteger64) {
				t.Error("unexpected skipUint func")
			}
		})
}

func TestUnsafe_setUpIntFuncs(t *testing.T) {
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
			if !ctest.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !ctest.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if sizeInt != com.Num32RawSize {
				t.Error("unexpected sizeInt value")
			}
			if !ctest.ComparePtrs(skipInt, raw.SkipInteger32) {
				t.Error("unexpected skipInt func")
			}
		})

	t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpIntFuncs(64)
			if !ctest.ComparePtrs(marshalInt, marshalInteger64[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !ctest.ComparePtrs(unmarshalInt, unmarshalInteger64[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if sizeInt != com.Num64RawSize {
				t.Error("unexpected sizeInt value")
			}
			if !ctest.ComparePtrs(skipInt, raw.SkipInteger64) {
				t.Error("unexpected skipInt func")
			}
		})
}

func TestUnsafe_unmarshalInteger64(t *testing.T) {
	var (
		wantV     uint64 = 0
		wantN            = 0
		wantErr          = mus.ErrTooSmallByteSlice
		bs               = []byte{1, 2, 3, 4, 5, 6, 7}
		v, n, err        = unmarshalInteger64[uint64](bs)
	)
	ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
}

func TestUnsafe_unmarshalInteger32(t *testing.T) {
	var (
		wantV     uint32 = 0
		wantN            = 0
		wantErr          = mus.ErrTooSmallByteSlice
		bs               = []byte{1, 2, 3}
		v, n, err        = unmarshalInteger32[uint32](bs)
	)
	ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
}

func TestUnsafe_unmarshalInteger16(t *testing.T) {
	var (
		wantV     uint16 = 0
		wantN            = 0
		wantErr          = mus.ErrTooSmallByteSlice
		bs               = []byte{1}
		v, n, err        = unmarshalInteger16[uint16](bs)
	)
	ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
}

func TestUnsafe_unmarshalInteger8(t *testing.T) {
	var (
		wantV     uint8 = 0
		wantN           = 0
		wantErr         = mus.ErrTooSmallByteSlice
		bs              = []byte{}
		v, n, err       = unmarshalInteger8[uint8](bs)
	)
	ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
}

func TestUnsafe_String(t *testing.T) {
	t.Run("String serializer should succeed",
		func(t *testing.T) {
			ser := String
			test.Test(ctest.StringTestCases, ser, t)
			test.TestSkip(ctest.StringTestCases, ser, t)
		})

	t.Run("We should be able to set a length serializer",
		func(t *testing.T) {
			var (
				str, lenSer = test.StringLenTestData(t)
				ser         = NewStringSer(stropts.WithLenSer(lenSer))
			)
			test.Test([]string{str}, ser, t)
			test.TestSkip([]string{str}, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid String serializer should succeed",
		func(t *testing.T) {
			ser := NewValidStringSer()
			test.Test(ctest.StringTestCases, ser, t)
			test.TestSkip(ctest.StringTestCases, ser, t)
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
				v, n, err = NewValidStringSer(stropts.WithLenSer(lenSer)).Unmarshal(nil)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     = ""
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				v, n, err = NewValidStringSer(nil).Unmarshal(bs)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				v, n, err = NewValidStringSer(stropts.WithLenValidator(lenVl)).Unmarshal(bs)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			v, n, err = NewValidStringSer(stropts.WithLenValidator(lenVl)).Unmarshal(bs)
		)
		ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})
}

func TestUnsafe_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctest.ByteTestCases, ser, t)
	test.TestSkip(ctest.ByteTestCases, ser, t)
}

func TestUnsafe_Uint64(t *testing.T) {
	ser := Uint64
	test.Test(ctest.Uint64TestCases, ser, t)
	test.TestSkip(ctest.Uint64TestCases, ser, t)
}

func TestUnsafe_Uint32(t *testing.T) {
	ser := Uint32
	test.Test(ctest.Uint32TestCases, ser, t)
	test.TestSkip(ctest.Uint32TestCases, ser, t)
}

func TestUnsafe_Uint16(t *testing.T) {
	ser := Uint16
	test.Test(ctest.Uint16TestCases, ser, t)
	test.TestSkip(ctest.Uint16TestCases, ser, t)
}

func TestUnsafe_Uint8(t *testing.T) {
	ser := Uint8
	test.Test(ctest.Uint8TestCases, ser, t)
	test.TestSkip(ctest.Uint8TestCases, ser, t)
}

func TestUnsafe_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctest.UintTestCases, ser, t)
	test.TestSkip(ctest.UintTestCases, ser, t)
}

func TestUnsafe_Int64(t *testing.T) {
	ser := Int64
	test.Test(ctest.Int64TestCases, ser, t)
	test.TestSkip(ctest.Int64TestCases, ser, t)
}

func TestUnsafe_Int32(t *testing.T) {
	ser := Int32
	test.Test(ctest.Int32TestCases, ser, t)
	test.TestSkip(ctest.Int32TestCases, ser, t)
}

func TestUnsafe_Int16(t *testing.T) {
	ser := Int16
	test.Test(ctest.Int16TestCases, ser, t)
	test.TestSkip(ctest.Int16TestCases, ser, t)
}

func TestUnsafe_Int8(t *testing.T) {
	ser := Int8
	test.Test(ctest.Int8TestCases, ser, t)
	test.TestSkip(ctest.Int8TestCases, ser, t)
}

func TestUnsafe_Int(t *testing.T) {
	ser := Int
	test.Test(ctest.IntTestCases, ser, t)
	test.TestSkip(ctest.IntTestCases, ser, t)
}

func TestUnsafe_Float64(t *testing.T) {
	t.Run("Float64 serializer should succeed", func(t *testing.T) {
		ser := Float64
		test.Test(ctest.Float64TestCases, ser, t)
		test.TestSkip(ctest.Float64TestCases, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestUnsafe_Float32(t *testing.T) {
	t.Run("Float32 serializer should succeed", func(t *testing.T) {
		ser := Float32
		test.Test(ctest.Float32TestCases, ser, t)
		test.TestSkip(ctest.Float32TestCases, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestUnsafe_Bool(t *testing.T) {
	t.Run("Bool serializer should succeed", func(t *testing.T) {
		ser := Bool
		test.Test(ctest.BoolTestCases, ser, t)
		test.TestSkip(ctest.BoolTestCases, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestUnsafe_ByteSlice(t *testing.T) {
	t.Run("ByteSlice serializer should succeed with empty slice",
		func(t *testing.T) {
			ser := ByteSlice
			test.Test([][]byte{{}}, ser, t)
			test.TestSkip([][]byte{{}}, ser, t)
		})

	t.Run("ByteSlice serializer should succeed with non-empty slice",
		func(t *testing.T) {
			ser := ByteSlice
			test.Test([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
			test.TestSkip([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
		})

	t.Run("We should be able to set a length serializer", func(t *testing.T) {
		var (
			sl, lenSer = test.ByteSliceLenTestData(t)
			ser        = NewByteSliceSer(bslopts.WithLenSer(lenSer))
		)
		test.Test([][]byte{sl}, ser, t)
		test.TestSkip([][]byte{sl}, ser, t)
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN, bs        = NegativeLengthBs()
				wantErr          = com.ErrNegativeLength
				v, n, err        = ByteSlice.Unmarshal(bs)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				n, err  = ByteSlice.Skip([]byte{})
			)
			ctest.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				n, err    = ByteSlice.Skip(bs)
			)
			ctest.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid ByteSlice serializer should succeed with empty slice",
		func(t *testing.T) {
			ser := NewValidByteSliceSer(nil)
			test.Test([][]byte{{}}, ser, t)
			test.TestSkip([][]byte{{}}, ser, t)
		})

	t.Run("Valid ByteSlice serializer should succeed with non-empty slice",
		func(t *testing.T) {
			ser := NewValidByteSliceSer(nil)
			test.Test([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
			test.TestSkip([][]byte{{0, 1, 1, 255, 100, 0, 1, 10}}, ser, t)
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
				v, n, err = NewValidByteSliceSer(bslopts.WithLenSer(lenSer)).Unmarshal(nil)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN, bs        = NegativeLengthBs()
				wantErr          = com.ErrNegativeLength
				v, n, err        = NewValidByteSliceSer(nil).Unmarshal(bs)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				v, n, err = NewValidByteSliceSer(bslopts.WithLenValidator(lenVl)).Unmarshal(bs)
			)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestUnsafe_Time(t *testing.T) {
	os.Setenv("TZ", "")

	t.Run("time_unix_utc", func(t *testing.T) {
		t.Run("TimeUnixUTC serializer should succeed",
			func(t *testing.T) {
				var (
					sec = time.Now().Unix()
					tm  = time.Unix(sec, 0)
				)
				test.Test([]time.Time{tm}, TimeUnixUTC, t)
				test.TestSkip([]time.Time{tm}, TimeUnix, t)
			})

		t.Run("We should be able to serializer the zero Time",
			func(t *testing.T) {
				test.Test([]time.Time{{}}, TimeUnixUTC, t)
				test.TestSkip([]time.Time{{}}, TimeUnixUTC, t)
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
				ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})
	})

	t.Run("time_unix_milli_UTC", func(t *testing.T) {
		t.Run("TimeUnixMilliUTC serializer should succeed",
			func(t *testing.T) {
				var (
					milli = time.Now().UnixMilli()
					tm    = time.UnixMilli(milli)
				)
				test.Test([]time.Time{tm}, TimeUnixMilliUTC, t)
				test.TestSkip([]time.Time{tm}, TimeUnixMilliUTC, t)
			})

		t.Run("We should be able to serializer the zero Time",
			func(t *testing.T) {
				test.Test([]time.Time{{}}, TimeUnixMilliUTC, t)
				test.TestSkip([]time.Time{{}}, TimeUnixMilliUTC, t)
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
				ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})
	})

	t.Run("time_unix_micro_utc", func(t *testing.T) {
		t.Run("TimeUnixMicroUTC serializer should succeed",
			func(t *testing.T) {
				var (
					milli = time.Now().UnixMicro()
					tm    = time.UnixMicro(milli)
				)
				test.Test([]time.Time{tm}, TimeUnixMicroUTC, t)
				test.TestSkip([]time.Time{tm}, TimeUnixMicroUTC, t)
			})

		t.Run("We should be able to serializer the zero Time",
			func(t *testing.T) {
				test.Test([]time.Time{{}}, TimeUnixMicroUTC, t)
				test.TestSkip([]time.Time{{}}, TimeUnixMicroUTC, t)
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
				ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})
	})

	t.Run("time_unix_nano_utc", func(t *testing.T) {
		t.Run("TimeUnixNanoUTC serializer should succeed",
			func(t *testing.T) {
				var (
					nano = time.Now().UnixNano()
					tm   = time.Unix(0, nano)
				)
				test.Test([]time.Time{tm}, TimeUnixNanoUTC, t)
				test.TestSkip([]time.Time{tm}, TimeUnixNanoUTC, t)
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
				ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					nil, t)
			})
	})
}

func TestUnsafe_Array(t *testing.T) {
	t.Run("Array serializer should succeed", func(t *testing.T) {
		var (
			arr, elemSer = test.ArrayTestData(t)
			mocks        = []*mok.Mock{elemSer.Mock}
			ser          = NewArraySer[[3]int](elemSer)
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)

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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				nil, t)
		})

	t.Run("Valid array serializer should succeed", func(t *testing.T) {
		var (
			arr, elemSer = test.ArrayTestData(t)
			mocks        = []*mok.Mock{elemSer.Mock}
			ser          = NewValidArraySer[[3]int](elemSer, nil)
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)

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
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
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
				ser   = NewValidArraySer[[3]int](elemSer, arropts.WithElemValidator[int](elemVl))
				mocks = []*mok.Mock{elemSer.Mock, elemVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctest.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				mocks, t)
		})

	t.Run("We should be able to set a length serializer", func(t *testing.T) {
		var (
			arr, elemSer = test.ArrayTestData(t)
			_, lenSer    = test.ArrayLenTestData(t)
			ser          = NewArraySer[[3]int](elemSer, arropts.WithLenSer[int](lenSer))
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)
	})

	t.Run("Valid array: We should be able to set a length serializer", func(t *testing.T) {
		var (
			arr, elemSer = test.ArrayTestData(t)
			_, lenSer    = test.ArrayLenTestData(t)
			ser          = NewValidArraySer[[3]int](elemSer, arropts.WithLenSer[int](lenSer))
		)
		test.Test([][3]int{arr}, ser, t)
		test.TestSkip([][3]int{arr}, ser, t)
	})
}

func NegativeLengthBs() (n int, bs []byte) {
	n = varint.PositiveInt.Size(-1)
	bs = make([]byte, n)
	varint.PositiveInt.Marshal(-1, bs)
	return
}
