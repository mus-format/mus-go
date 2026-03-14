package raw

import (
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test"
)

func TestRaw_setUpUintFuncs(t *testing.T) {
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
				t.Error("unexpected sizeUint func")
			}
			if !ctestutil.ComparePtrs(skipUint, SkipInteger32) {
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
				t.Error("unexpected sizeUint func")
			}
			if !ctestutil.ComparePtrs(skipUint, SkipInteger64) {
				t.Error("unexpected skipUint func")
			}
		})
}

func TestRaw_setUpIntFuncs(t *testing.T) {
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
			if !ctestutil.ComparePtrs(marshalInt, marshalInteger32[int]) {
				t.Error("unexpected marshalInt func")
			}
			if !ctestutil.ComparePtrs(unmarshalInt, unmarshalInteger32[int]) {
				t.Error("unexpected unmarshalInt func")
			}
			if sizeInt != com.Num32RawSize {
				t.Error("unexpected sizeInt func")
			}
			if !ctestutil.ComparePtrs(skipInt, SkipInteger32) {
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
				t.Error("unexpected sizeInt func")
			}
			if !ctestutil.ComparePtrs(skipInt, SkipInteger64) {
				t.Error("unexpected skipInt func")
			}
		})
}

func TestRaw_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctestutil.ByteTestCases, ser, t)
	test.TestSkip(ctestutil.ByteTestCases, ser, t)
}

func TestRaw_Uint64(t *testing.T) {
	t.Run("Uint64 serializer should succeed", func(t *testing.T) {
		ser := Uint64
		test.Test(ctestutil.Uint64TestCases, ser, t)
		test.TestSkip(ctestutil.Uint64TestCases, ser, t)
	})

	t.Run("unmarshalInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   uint64 = 0
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{1, 2, 3, 4, 5}
			)
			v, n, err := unmarshalInteger64[uint64](bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3, 4, 5, 6, 7}
			)
			n, err := SkipInteger64(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestRaw_Uint32(t *testing.T) {
	t.Run("Uint32 serializer should succeed", func(t *testing.T) {
		ser := Uint32
		test.Test(ctestutil.Uint32TestCases, ser, t)
		test.TestSkip(ctestutil.Uint32TestCases, ser, t)
	})

	t.Run("unmarshalInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   uint32 = 0
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{1, 2, 3}
			)
			v, n, err := unmarshalInteger32[uint32](bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3}
			)
			n, err := SkipInteger32(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestRaw_Uint16(t *testing.T) {
	t.Run("Uint16 serializer should succeed", func(t *testing.T) {
		ser := Uint16
		test.Test(ctestutil.Uint16TestCases, ser, t)
		test.TestSkip(ctestutil.Uint16TestCases, ser, t)
	})

	t.Run("unmarshalInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   uint16 = 0
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{1}
			)
			v, n, err := unmarshalInteger16[uint16](bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1}
			)
			n, err := SkipInteger16(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestRaw_Uint8(t *testing.T) {
	t.Run("Uint8 serializer should succeed", func(t *testing.T) {
		ser := Uint8
		test.Test(ctestutil.Uint8TestCases, ser, t)
		test.TestSkip(ctestutil.Uint8TestCases, ser, t)
	})

	t.Run("unmarshalInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   uint8 = 0
				wantN         = 0
				wantErr       = mus.ErrTooSmallByteSlice
				bs            = []byte{}
			)
			v, n, err := unmarshalInteger8[uint8](bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
			)
			n, err := SkipInteger8(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestRaw_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctestutil.UintTestCases, ser, t)
	test.TestSkip(ctestutil.UintTestCases, ser, t)
}

func TestRaw_Int64(t *testing.T) {
	ser := Int64
	test.Test(ctestutil.Int64TestCases, ser, t)
	test.TestSkip(ctestutil.Int64TestCases, ser, t)
}

func TestRaw_Int32(t *testing.T) {
	ser := Int32
	test.Test(ctestutil.Int32TestCases, ser, t)
	test.TestSkip(ctestutil.Int32TestCases, ser, t)
}

func TestRaw_Int16(t *testing.T) {
	ser := Int16
	test.Test(ctestutil.Int16TestCases, ser, t)
	test.TestSkip(ctestutil.Int16TestCases, ser, t)
}

func TestRaw_Int8(t *testing.T) {
	ser := Int8
	test.Test(ctestutil.Int8TestCases, ser, t)
	test.TestSkip(ctestutil.Int8TestCases, ser, t)
}

func TestRaw_Int(t *testing.T) {
	ser := Int
	test.Test(ctestutil.IntTestCases, ser, t)
	test.TestSkip(ctestutil.IntTestCases, ser, t)
}

func TestRaw_Float64(t *testing.T) {
	t.Run("Float64 serializer should succeed", func(t *testing.T) {
		ser := Float64
		test.Test(ctestutil.Float64TestCases, ser, t)
		test.TestSkip(ctestutil.Float64TestCases, ser, t)
	})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   float64 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{1, 2, 3, 4, 5}
			)
			v, n, err := Float64.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3, 4, 5}
			)
			n, err := Float64.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestRaw_Float32(t *testing.T) {
	t.Run("Float32 serializer should succeed", func(t *testing.T) {
		ser := Float32
		test.Test(ctestutil.Float32TestCases, ser, t)
		test.TestSkip(ctestutil.Float32TestCases, ser, t)
	})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   float32 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{1, 2}
			)
			v, n, err := Float32.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2}
			)
			n, err := Float32.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
				wantV   = time.Time{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
			)
			v, n, err := TimeUnixUTC.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				wantV   = time.Time{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
			)
			v, n, err := TimeUnixMilliUTC.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				wantV   = time.Time{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
			)
			v, n, err := TimeUnixMicroUTC.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				wantV   = time.Time{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
			)
			v, n, err := TimeUnixNanoUTC.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}
