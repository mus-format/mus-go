package raw

import (
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestRaw_setUpUintFuncs(t *testing.T) {
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
			asserterror.Equal(t, sizeUint, com.Num32RawSize, "unexpected sizeUint func")
			asserterror.Equal(t, ctest.ComparePtrs(skipUint, SkipInteger32),
				true, "unexpected skipUint func")
		})

	t.Run("If the system int size is equal to 64, setUpUintFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpUintFuncs(64)
			asserterror.Equal(t, ctest.ComparePtrs(marshalUint, marshalInteger64[uint]),
				true, "unexpected marshalUint func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalUint, unmarshalInteger64[uint]),
				true, "unexpected unmarshalUint func")
			asserterror.Equal(t, sizeUint, com.Num64RawSize, "unexpected sizeUint func")
			asserterror.Equal(t, ctest.ComparePtrs(skipUint, SkipInteger64),
				true, "unexpected skipUint func")
		})
}

func TestRaw_setUpIntFuncs(t *testing.T) {
	t.Run("If the system int size is not 32 or 64, setUpIntFuncs should panic with ErrUnsupportedIntSize",
		func(t *testing.T) {
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
			asserterror.Equal(t, sizeInt, com.Num32RawSize, "unexpected sizeInt func")
			asserterror.Equal(t, ctest.ComparePtrs(skipInt, SkipInteger32),
				true, "unexpected skipInt func")
		})

	t.Run("If the system int size is equal to 64, setUpIntFuncs should initialize the uint functions with 64-bit versions",
		func(t *testing.T) {
			setUpIntFuncs(64)
			asserterror.Equal(t, ctest.ComparePtrs(marshalInt, marshalInteger64[int]),
				true, "unexpected marshalInt func")
			asserterror.Equal(t, ctest.ComparePtrs(unmarshalInt, unmarshalInteger64[int]),
				true, "unexpected unmarshalInt func")
			asserterror.Equal(t, sizeInt, com.Num64RawSize, "unexpected sizeInt func")
			asserterror.Equal(t, ctest.ComparePtrs(skipInt, SkipInteger64),
				true, "unexpected skipInt func")
		})
}

func TestRaw_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctest.ByteTestCases, ser, t)
	test.TestSkip(ctest.ByteTestCases, ser, t)
}

func TestRaw_Uint64(t *testing.T) {
	t.Run("Uint64 serializer should succeed", func(t *testing.T) {
		ser := Uint64
		test.Test(ctest.Uint64TestCases, ser, t)
		test.TestSkip(ctest.Uint64TestCases, ser, t)
	})

	t.Run("unmarshalInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3, 4, 5}
			)
			test.TestUnmarshalOnly(bs, Uint64, want, nil, t)
		})

	t.Run("skipInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3, 4, 5, 6, 7}
			)
			test.TestSkipOnly(bs, Uint64, want, nil, t)
		})
}

func TestRaw_Uint32(t *testing.T) {
	t.Run("Uint32 serializer should succeed", func(t *testing.T) {
		ser := Uint32
		test.Test(ctest.Uint32TestCases, ser, t)
		test.TestSkip(ctest.Uint32TestCases, ser, t)
	})

	t.Run("unmarshalInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint32]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3}
			)
			test.TestUnmarshalOnly(bs, Uint32, want, nil, t)
		})

	t.Run("skipInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3}
			)
			test.TestSkipOnly(bs, Uint32, want, nil, t)
		})
}

func TestRaw_Uint16(t *testing.T) {
	t.Run("Uint16 serializer should succeed", func(t *testing.T) {
		ser := Uint16
		test.Test(ctest.Uint16TestCases, ser, t)
		test.TestSkip(ctest.Uint16TestCases, ser, t)
	})

	t.Run("unmarshalInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint16]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1}
			)
			test.TestUnmarshalOnly(bs, Uint16, want, nil, t)
		})

	t.Run("skipInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1}
			)
			test.TestSkipOnly(bs, Uint16, want, nil, t)
		})
}

func TestRaw_Uint8(t *testing.T) {
	t.Run("Uint8 serializer should succeed", func(t *testing.T) {
		ser := Uint8
		test.Test(ctest.Uint8TestCases, ser, t)
		test.TestSkip(ctest.Uint8TestCases, ser, t)
	})

	t.Run("unmarshalInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint8]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Uint8, want, nil, t)
		})

	t.Run("skipInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestSkipOnly(bs, Uint8, want, nil, t)
		})
}

func TestRaw_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctest.UintTestCases, ser, t)
	test.TestSkip(ctest.UintTestCases, ser, t)
}

func TestRaw_Int64(t *testing.T) {
	ser := Int64
	test.Test(ctest.Int64TestCases, ser, t)
	test.TestSkip(ctest.Int64TestCases, ser, t)
}

func TestRaw_Int32(t *testing.T) {
	ser := Int32
	test.Test(ctest.Int32TestCases, ser, t)
	test.TestSkip(ctest.Int32TestCases, ser, t)
}

func TestRaw_Int16(t *testing.T) {
	ser := Int16
	test.Test(ctest.Int16TestCases, ser, t)
	test.TestSkip(ctest.Int16TestCases, ser, t)
}

func TestRaw_Int8(t *testing.T) {
	ser := Int8
	test.Test(ctest.Int8TestCases, ser, t)
	test.TestSkip(ctest.Int8TestCases, ser, t)
}

func TestRaw_Int(t *testing.T) {
	ser := Int
	test.Test(ctest.IntTestCases, ser, t)
	test.TestSkip(ctest.IntTestCases, ser, t)
}

func TestRaw_Float64(t *testing.T) {
	t.Run("Float64 serializer should succeed", func(t *testing.T) {
		ser := Float64
		test.Test(ctest.Float64TestCases, ser, t)
		test.TestSkip(ctest.Float64TestCases, ser, t)
	})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[float64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3, 4, 5}
			)
			test.TestUnmarshalOnly(bs, Float64, want, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2, 3, 4, 5}
			)
			test.TestSkipOnly(bs, Float64, want, nil, t)
		})
}

func TestRaw_Float32(t *testing.T) {
	t.Run("Float32 serializer should succeed", func(t *testing.T) {
		ser := Float32
		test.Test(ctest.Float32TestCases, ser, t)
		test.TestSkip(ctest.Float32TestCases, ser, t)
	})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[float32]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2}
			)
			test.TestUnmarshalOnly(bs, Float32, want, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{1, 2}
			)
			test.TestSkipOnly(bs, Float32, want, nil, t)
		})
}

func TestRaw_TimeUnixUTC(t *testing.T) {
	os.Setenv("TZ", "")

	t.Run("TimeUnixUTC serializer should succeed", func(t *testing.T) {
		var (
			sec = time.Now().Unix()
			tm  = time.Unix(sec, 0)
		)
		test.Test([]time.Time{tm}, TimeUnixUTC, t)
		test.TestSkip([]time.Time{tm}, TimeUnixUTC, t)
	})

	t.Run("We should be able to serializer the zero Time", func(t *testing.T) {
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
}

func TestRaw_TimeUnixMilliUTC(t *testing.T) {
	os.Setenv("TZ", "")

	t.Run("TimeUnixMilliUTC serializer should succeed", func(t *testing.T) {
		var (
			milli = time.Now().UnixMilli()
			tm    = time.UnixMilli(milli)
		)
		test.Test([]time.Time{tm}, TimeUnixMilliUTC, t)
		test.TestSkip([]time.Time{tm}, TimeUnixMilliUTC, t)
	})

	t.Run("We should be able to serializer the zero Time", func(t *testing.T) {
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
}

func TestRaw_TimeUnixMicroUTC(t *testing.T) {
	os.Setenv("TZ", "")

	t.Run("TimeUnixMicroUTC serializer should succeed", func(t *testing.T) {
		var (
			milli = time.Now().UnixMicro()
			tm    = time.UnixMicro(milli)
		)
		test.Test([]time.Time{tm}, TimeUnixMicroUTC, t)
		test.TestSkip([]time.Time{tm}, TimeUnixMicroUTC, t)
	})

	t.Run("We should be able to serializer the zero Time", func(t *testing.T) {
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
}

func TestRaw_TimeUnixNanoUTC(t *testing.T) {
	os.Setenv("TZ", "")

	t.Run("TimeUnixNanoUTC serializer should succeed", func(t *testing.T) {
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
}
