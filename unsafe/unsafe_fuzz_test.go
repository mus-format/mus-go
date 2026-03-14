package unsafe

import (
	"errors"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

const maxLen = 1000

// bool ------------------------------------------------------------------------

func FuzzBool(f *testing.F) {
	f.Add(true)
	f.Add(false)
	f.Fuzz(func(t *testing.T, v bool) {
		test.Test([]bool{v}, Bool, t)
		test.TestSkip([]bool{v}, Bool, t)
	})
}

func FuzzBoolUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Bool.Unmarshal(bs)
		Bool.Skip(bs)
	})
}

// byte ------------------------------------------------------------------------

func FuzzByte(f *testing.F) {
	seeds := []byte{0, 1, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		test.Test([]byte{v}, Byte, t)
		test.TestSkip([]byte{v}, Byte, t)
	})
}

func FuzzByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Byte.Unmarshal(bs)
		Byte.Skip(bs)
	})
}

// uint64 ----------------------------------------------------------------------

func FuzzUint64(f *testing.F) {
	seeds := []uint64{0, 1, 1 << 63}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		test.Test([]uint64{v}, Uint64, t)
		test.TestSkip([]uint64{v}, Uint64, t)
	})
}

func FuzzUint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint64.Unmarshal(bs)
		Uint64.Skip(bs)
	})
}

// uint32 ----------------------------------------------------------------------

func FuzzUint32(f *testing.F) {
	seeds := []uint32{0, 1, 1 << 31}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		test.Test([]uint32{v}, Uint32, t)
		test.TestSkip([]uint32{v}, Uint32, t)
	})
}

func FuzzUint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint32.Unmarshal(bs)
		Uint32.Skip(bs)
	})
}

// uint16 ----------------------------------------------------------------------

func FuzzUint16(f *testing.F) {
	seeds := []uint16{0, 1, 1 << 15}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		test.Test([]uint16{v}, Uint16, t)
		test.TestSkip([]uint16{v}, Uint16, t)
	})
}

func FuzzUint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint16.Unmarshal(bs)
		Uint16.Skip(bs)
	})
}

// uint8 -----------------------------------------------------------------------

func FuzzUint8(f *testing.F) {
	seeds := []uint8{0, 1, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		test.Test([]uint8{v}, Uint8, t)
		test.TestSkip([]uint8{v}, Uint8, t)
	})
}

func FuzzUint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint8.Unmarshal(bs)
		Uint8.Skip(bs)
	})
}

// uint ------------------------------------------------------------------------

func FuzzUint(f *testing.F) {
	seeds := []uint{0, 1}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		test.Test([]uint{v}, Uint, t)
		test.TestSkip([]uint{v}, Uint, t)
	})
}

func FuzzUintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint.Unmarshal(bs)
		Uint.Skip(bs)
	})
}

// int64 -----------------------------------------------------------------------

func FuzzInt64(f *testing.F) {
	seeds := []int64{0, 1, -1, 1 << 62, -(1 << 62)}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		test.Test([]int64{v}, Int64, t)
		test.TestSkip([]int64{v}, Int64, t)
	})
}

func FuzzInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int64.Unmarshal(bs)
		Int64.Skip(bs)
	})
}

// int32 -----------------------------------------------------------------------

func FuzzInt32(f *testing.F) {
	seeds := []int32{0, 1, -1, 1 << 30, -(1 << 30)}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		test.Test([]int32{v}, Int32, t)
		test.TestSkip([]int32{v}, Int32, t)
	})
}

func FuzzInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int32.Unmarshal(bs)
		Int32.Skip(bs)
	})
}

// int16 -----------------------------------------------------------------------

func FuzzInt16(f *testing.F) {
	seeds := []int16{0, 1, -1, 1 << 14, -(1 << 14)}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		test.Test([]int16{v}, Int16, t)
		test.TestSkip([]int16{v}, Int16, t)
	})
}

func FuzzInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int16.Unmarshal(bs)
		Int16.Skip(bs)
	})
}

// int8 ------------------------------------------------------------------------

func FuzzInt8(f *testing.F) {
	seeds := []int8{0, 1, -1, 127, -128}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		test.Test([]int8{v}, Int8, t)
		test.TestSkip([]int8{v}, Int8, t)
	})
}

func FuzzInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int8.Unmarshal(bs)
		Int8.Skip(bs)
	})
}

// int -------------------------------------------------------------------------

func FuzzInt(f *testing.F) {
	seeds := []int{0, 1, -1}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		test.Test([]int{v}, Int, t)
		test.TestSkip([]int{v}, Int, t)
	})
}

func FuzzIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int.Unmarshal(bs)
		Int.Skip(bs)
	})
}

// float64 ---------------------------------------------------------------------

func FuzzFloat64(f *testing.F) {
	seeds := []float64{0, 1, -1, 3.14}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		test.Test([]float64{v}, Float64, t)
		test.TestSkip([]float64{v}, Float64, t)
	})
}

func FuzzFloat64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float64.Unmarshal(bs)
		Float64.Skip(bs)
	})
}

// float32 ---------------------------------------------------------------------

func FuzzFloat32(f *testing.F) {
	seeds := []float32{0, 1, -1, 3.14}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		test.Test([]float32{v}, Float32, t)
		test.TestSkip([]float32{v}, Float32, t)
	})
}

func FuzzFloat32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float32.Unmarshal(bs)
		Float32.Skip(bs)
	})
}

// string ----------------------------------------------------------------------

func FuzzString(f *testing.F) {
	seeds := []string{"", "hello", "world", "mus-format"}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) > maxLen {
			v = v[:maxLen]
		}
		test.Test([]string{v}, String, t)
		test.TestSkip([]string{v}, String, t)
	})
}

func FuzzStringUnmarshal(f *testing.F) {
	ser := NewValidStringSer(strops.WithLenValidator(
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// byte slice ------------------------------------------------------------------

func FuzzByteSlice(f *testing.F) {
	seeds := [][]byte{{}, {1, 2, 3}, {255, 0, 255}}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v []byte) {
		if len(v) > maxLen {
			v = v[:maxLen]
		}
		test.Test([][]byte{v}, ByteSlice, t)
		test.TestSkip([][]byte{v}, ByteSlice, t)
	})
}

func FuzzByteSliceUnmarshal(f *testing.F) {
	ser := NewValidByteSliceSer(bslops.WithLenValidator(
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// time ------------------------------------------------------------------------

func FuzzTimeUnixUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0).UTC()
		test.Test([]time.Time{v}, TimeUnixUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixUTC, t)
	})
}

func FuzzTimeUnixMilliUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli).UTC()
		test.Test([]time.Time{v}, TimeUnixMilliUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixMilliUTC, t)
	})
}

func FuzzTimeUnixMicroUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro).UTC()
		test.Test([]time.Time{v}, TimeUnixMicroUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixMicroUTC, t)
	})
}

func FuzzTimeUnixNanoUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano).UTC()
		test.Test([]time.Time{v}, TimeUnixNanoUTC, t)
		test.TestSkip([]time.Time{v}, TimeUnixNanoUTC, t)
	})
}

func FuzzTimeUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		TimeUnixUTC.Unmarshal(bs)
		TimeUnixUTC.Skip(bs)
		TimeUnixMilliUTC.Unmarshal(bs)
		TimeUnixMilliUTC.Skip(bs)
		TimeUnixMicroUTC.Unmarshal(bs)
		TimeUnixMicroUTC.Skip(bs)
		TimeUnixNanoUTC.Unmarshal(bs)
		TimeUnixNanoUTC.Skip(bs)
	})
}

// array -----------------------------------------------------------------------

func FuzzArray(f *testing.F) {
	seeds := [][3]byte{
		{1, 2, 3},
		{0, 0, 0},
		{255, 255, 255},
	}
	for _, seed := range seeds {
		f.Add(seed[0], seed[1], seed[2])
	}
	f.Fuzz(func(t *testing.T, b1, b2, b3 byte) {
		v := [3]int{int(b1), int(b2), int(b3)}
		ser := NewArraySer[[3]int](varint.Int)
		test.Test([][3]int{v}, ser, t)
		test.TestSkip([][3]int{v}, ser, t)
	})
}

func FuzzArrayUnmarshal(f *testing.F) {
	ser := NewArraySer[[3]int](varint.Int)
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}
