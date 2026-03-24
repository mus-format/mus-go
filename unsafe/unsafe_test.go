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
	asserterror "github.com/ymz-ncnk/assert/error"
	"github.com/ymz-ncnk/mok"
)

func TestUnsafe_setUpUintFuncs(t *testing.T) {
	t.Run("If the system int size is not 32 or 64, setUpUintFuncs should panic with ErrUnsupportedIntSize",
		func(t *testing.T) {
			wantErr := com.ErrUnsupportedIntSize
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					asserterror.EqualError(t, err, wantErr)
				}
			}()
			setUpUintFuncs(16)
		})

	t.Run("If the system int size is equal to 32, setUpUintFuncs should initialize the uint functions with 32-bit versions",
		func(t *testing.T) {
			setUpUintFuncs(32)
			asserterror.Equal(t, ctest.ComparePtrs(marshalUint, marshalInteger32[uint]),
				true, "unexpected marshalUint func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalUint, unmarshalInteger32[uint]),
				true, "unexpected unmarshalUint func")
			asserterror.Equal(t, sizeUint, com.Num32RawSize, "unexpected sizeUint value")
			asserterror.Equal(t, ctest.ComparePtrs(skipUint, raw.SkipInteger32),
				true, "unexpected skipUint func")
		})

	t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpUintFuncs(64)
			asserterror.Equal(t, ctest.ComparePtrs(marshalUint, marshalInteger64[uint]),
				true, "unexpected marshalUint func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]),
				true, "unexpected unmarshalUint func")
			asserterror.Equal(t, sizeUint, com.Num64RawSize, "unexpected sizeUint value")
			asserterror.Equal(t, ctest.ComparePtrs(skipUint, raw.SkipInteger64),
				true, "unexpected skipUint func")
		})
}

func TestUnsafe_setUpIntFuncs(t *testing.T) {
	t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize", func(t *testing.T) {
		wantErr := com.ErrUnsupportedIntSize
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				asserterror.EqualError(t, err, wantErr)
			}
		}()
		setUpIntFuncs(16)
	})

	t.Run("If the system int size is equal to 32, setUpIntFuncs should initialize the uint functions with 32-bit versions",
		func(t *testing.T) {
			setUpIntFuncs(32)
			asserterror.Equal(t, ctest.ComparePtrs(marshalInt, marshalInteger32[int]),
				true, "unexpected marshalInt func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalInt, unmarshalInteger32[int]),
				true, "unexpected unmarshalInt func")
			asserterror.Equal(t, sizeInt, com.Num32RawSize, "unexpected sizeInt value")
			asserterror.Equal(t, ctest.ComparePtrs(skipInt, raw.SkipInteger32),
				true, "unexpected skipInt func")
		})

	t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpIntFuncs(64)
			asserterror.Equal(t, ctest.ComparePtrs(marshalInt, marshalInteger64[int]),
				true, "unexpected marshalInt func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalInt, unmarshalInteger64[int]),
				true, "unexpected unmarshalInt func")
			asserterror.Equal(t, sizeInt, com.Num64RawSize, "unexpected sizeInt value")
			asserterror.Equal(t, ctest.ComparePtrs(skipInt, raw.SkipInteger64),
				true, "unexpected skipInt func")
		})
}

func TestUnsafe_unmarshalInteger64(t *testing.T) {
	var (
		want = test.UnmarshalResult[uint64]{
			V:   0,
			N:   0,
			Err: mus.ErrTooSmallByteSlice,
		}
		bs = []byte{1, 2, 3, 4, 5, 6, 7}
	)
	test.TestUnmarshalOnly(bs, Uint64, want, nil, t)
}

func TestUnsafe_unmarshalInteger32(t *testing.T) {
	var (
		want = test.UnmarshalResult[uint32]{
			V:   0,
			N:   0,
			Err: mus.ErrTooSmallByteSlice,
		}
		bs = []byte{1, 2, 3}
	)
	test.TestUnmarshalOnly(bs, Uint32, want, nil, t)
}

func TestUnsafe_unmarshalInteger16(t *testing.T) {
	var (
		want = test.UnmarshalResult[uint16]{
			V:   0,
			N:   0,
			Err: mus.ErrTooSmallByteSlice,
		}
		bs = []byte{1}
	)
	test.TestUnmarshalOnly(bs, Uint16, want, nil, t)
}

func TestUnsafe_unmarshalInteger8(t *testing.T) {
	var (
		want = test.UnmarshalResult[uint8]{
			V:   0,
			N:   0,
			Err: mus.ErrTooSmallByteSlice,
		}
		bs = []byte{}
	)
	test.TestUnmarshalOnly(bs, Uint8, want, nil, t)
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
					asserterror.EqualError(t, err, wantErr)
				}
			}()

			String.Marshal(s, bs)
		})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			want := test.UnmarshalResult[string]{
				V:   "",
				N:   0,
				Err: mus.ErrTooSmallByteSlice,
			}
			test.TestUnmarshalOnly(nil, String, want, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				want      = test.UnmarshalResult[string]{
					V:   "",
					N:   wantN,
					Err: com.ErrNegativeLength,
				}
			)
			test.TestUnmarshalOnly(bs, String, want, nil, t)
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[string]{
					V:   "",
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{3, 1, 1}
			)
			test.TestUnmarshalOnly(bs, String, want, nil, t)
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
				wantErr = mus.ErrTooSmallByteSlice
				want    = test.UnmarshalResult[string]{
					V:   "",
					N:   0,
					Err: wantErr,
				}
				lenSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser = NewValidStringSer(stropts.WithLenSer(lenSer))
			)
			test.TestUnmarshalOnly(nil, ser, want, []*mok.Mock{lenSer.Mock}, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				want      = test.UnmarshalResult[string]{
					V:   "",
					N:   wantN,
					Err: com.ErrNegativeLength,
				}
			)
			test.TestUnmarshalOnly(bs, NewValidStringSer(nil), want, nil, t)
		})

	t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[string]{
					V:   "",
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{3, 1, 1}
			)
			test.TestUnmarshalOnly(bs, NewValidStringSer(nil), want, nil, t)
		})

	t.Run("If lenVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("lenVl validator error")
				want    = test.UnmarshalResult[string]{
					V:   "",
					N:   1,
					Err: wantErr,
				}
				lenVl = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						wantV := 3
						asserterror.Equal(t, v, wantV)
						return wantErr
					},
				)
				bs  = []byte{3, 1, 1, 1}
				ser = NewValidStringSer(stropts.WithLenValidator(lenVl))
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
		})

	t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
		var (
			wantErr = errors.New("empty string")
			want    = test.UnmarshalResult[string]{
				V:   "",
				N:   1,
				Err: wantErr,
			}
			bs                         = []byte{0}
			lenVl com.ValidatorFn[int] = func(t int) (err error) {
				return wantErr
			}
			ser = NewValidStringSer(stropts.WithLenValidator(lenVl))
		)
		test.TestUnmarshalOnly(bs, ser, want, nil, t)
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
				want = test.UnmarshalResult[float64]{
					V:   0.0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Float64, want, nil, t)
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
				want = test.UnmarshalResult[float32]{
					V:   0.0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Float32, want, nil, t)
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
				want = test.UnmarshalResult[bool]{
					V:   false,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Bool, want, nil, t)
		})

	t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[bool]{
					V:   false,
					N:   0,
					Err: com.ErrWrongFormat,
				}
				bs = []byte{2}
			)
			test.TestUnmarshalOnly(bs, Bool, want, nil, t)
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
					asserterror.EqualError(t, err, wantErr)
				}
			}()
			ByteSlice.Marshal(s, bs)
		})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			want := test.UnmarshalResult[[]byte]{
				V:   nil,
				N:   0,
				Err: mus.ErrTooSmallByteSlice,
			}
			test.TestUnmarshalOnly(nil, ByteSlice, want, nil, t)
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if bs is too small for slice content",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{2, 1}
			)
			test.TestUnmarshalOnly(bs, ByteSlice, want, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				want      = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   wantN,
					Err: com.ErrNegativeLength,
				}
			)
			test.TestUnmarshalOnly(bs, ByteSlice, want, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestSkipOnly(bs, ByteSlice, want, nil, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				want      = test.SkipResult{
					N:   wantN,
					Err: com.ErrNegativeLength,
				}
			)
			test.TestSkipOnly(bs, ByteSlice, want, nil, t)
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
				wantErr = errors.New("lenSer error")
				want    = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   0,
					Err: wantErr,
				}
				lenSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser = NewValidByteSliceSer(bslopts.WithLenSer(lenSer))
			)
			test.TestUnmarshalOnly(nil, ser, want, []*mok.Mock{lenSer.Mock}, t)
		})

	t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if bs is too small for slice content",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{2, 1}
			)
			test.TestUnmarshalOnly(bs, NewValidByteSliceSer(nil), want, nil, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				want      = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   wantN,
					Err: com.ErrNegativeLength,
				}
			)
			test.TestUnmarshalOnly(bs, NewValidByteSliceSer(nil), want, nil, t)
		})

	t.Run("If lenVl fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantErr = errors.New("too large slice")
				want    = test.UnmarshalResult[[]byte]{
					V:   nil,
					N:   1,
					Err: wantErr,
				}
				bs    = []byte{3, 4, 1, 1}
				lenVl = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						return wantErr
					},
				)
				ser = NewValidByteSliceSer(bslopts.WithLenValidator(lenVl))
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
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
					want = test.UnmarshalResult[time.Time]{
						V:   time.Time{},
						N:   0,
						Err: mus.ErrTooSmallByteSlice,
					}
					bs = []byte{}
				)
				test.TestUnmarshalOnly(bs, TimeUnixUTC, want, nil, t)
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
					want = test.UnmarshalResult[time.Time]{
						V:   time.Time{},
						N:   0,
						Err: mus.ErrTooSmallByteSlice,
					}
					bs = []byte{}
				)
				test.TestUnmarshalOnly(bs, TimeUnixMilliUTC, want, nil, t)
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
					want = test.UnmarshalResult[time.Time]{
						V:   time.Time{},
						N:   0,
						Err: mus.ErrTooSmallByteSlice,
					}
					bs = []byte{}
				)
				test.TestUnmarshalOnly(bs, TimeUnixMicroUTC, want, nil, t)
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
					want = test.UnmarshalResult[time.Time]{
						V:   time.Time{},
						N:   0,
						Err: mus.ErrTooSmallByteSlice,
					}
					bs = []byte{}
				)
				test.TestUnmarshalOnly(bs, TimeUnixNanoUTC, want, nil, t)
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
				want = test.UnmarshalResult[[3]int]{
					V:   [3]int{0, 0, 0},
					N:   1,
					Err: com.ErrTooLargeLength,
				}
				bs = []byte{4}
			)
			test.TestUnmarshalOnly(bs, NewArraySer[[3]int, int](nil), want, nil, t)
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
				want = test.UnmarshalResult[[3]int]{
					V:   [3]int{0, 0, 0},
					N:   1,
					Err: com.ErrTooLargeLength,
				}
				bs = []byte{4}
			)
			test.TestUnmarshalOnly(bs, NewValidArraySer[[3]int, int](nil, nil), want, nil, t)
		})

	t.Run("If elemVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantElem = 11
				wantErr  = errors.New("elemVl error")
				want     = test.UnmarshalResult[[3]int]{
					V:   [3]int{0, 0, 0},
					N:   2,
					Err: wantErr,
				}
				bs      = []byte{3, 11}
				elemSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						return 11, 1, nil
					},
				)
				elemVl = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						asserterror.Equal(t, v, wantElem)
						return wantErr
					},
				)
				ser   = NewValidArraySer[[3]int](elemSer, arropts.WithElemValidator(elemVl))
				mocks = []*mok.Mock{elemSer.Mock, elemVl.Mock}
			)
			test.TestUnmarshalOnly(bs, ser, want, mocks, t)
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
