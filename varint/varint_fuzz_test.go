package varint

import (
	"math"
	"testing"

	"github.com/mus-format/mus-go/testutil"
)

// byte ------------------------------------------------------------------------

func FuzzByte(f *testing.F) {
	seeds := []byte{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		testutil.Test[byte]([]byte{v}, Byte, t)
		testutil.TestSkip[byte]([]byte{v}, Byte, t)
	})
}

func FuzzByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Byte.Unmarshal(bs)
		Byte.Skip(bs)
	})
}

// uint ------------------------------------------------------------------------

func FuzzUint64(f *testing.F) {
	seeds := []uint64{0, 1, 127, 128, 255, 256, math.MaxUint64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		testutil.Test[uint64]([]uint64{v}, Uint64, t)
		testutil.TestSkip[uint64]([]uint64{v}, Uint64, t)
	})
}

func FuzzUint32(f *testing.F) {
	seeds := []uint32{0, 1, 127, 128, 255, 256, math.MaxUint32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		testutil.Test[uint32]([]uint32{v}, Uint32, t)
		testutil.TestSkip[uint32]([]uint32{v}, Uint32, t)
	})
}

func FuzzUint16(f *testing.F) {
	seeds := []uint16{0, 1, 127, 128, 255, 256, math.MaxUint16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		testutil.Test[uint16]([]uint16{v}, Uint16, t)
		testutil.TestSkip[uint16]([]uint16{v}, Uint16, t)
	})
}

func FuzzUint8(f *testing.F) {
	seeds := []uint8{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		testutil.Test[uint8]([]uint8{v}, Uint8, t)
		testutil.TestSkip[uint8]([]uint8{v}, Uint8, t)
	})
}

func FuzzUint(f *testing.F) {
	seeds := []uint{0, 1, 127, 128, 255, 256, math.MaxUint}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		testutil.Test[uint]([]uint{v}, Uint, t)
		testutil.TestSkip[uint]([]uint{v}, Uint, t)
	})
}

func FuzzUint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint64.Unmarshal(bs)
		Uint64.Skip(bs)
	})
}

func FuzzUint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint32.Unmarshal(bs)
		Uint32.Skip(bs)
	})
}

func FuzzUint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint16.Unmarshal(bs)
		Uint16.Skip(bs)
	})
}

func FuzzUint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint8.Unmarshal(bs)
		Uint8.Skip(bs)
	})
}

func FuzzUintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint.Unmarshal(bs)
		Uint.Skip(bs)
	})
}

// int -------------------------------------------------------------------------

func FuzzInt64(f *testing.F) {
	seeds := []int64{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt64, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		testutil.Test[int64]([]int64{v}, Int64, t)
		testutil.TestSkip[int64]([]int64{v}, Int64, t)
	})
}

func FuzzInt32(f *testing.F) {
	seeds := []int32{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt32, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		testutil.Test[int32]([]int32{v}, Int32, t)
		testutil.TestSkip[int32]([]int32{v}, Int32, t)
	})
}

func FuzzInt16(f *testing.F) {
	seeds := []int16{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt16, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		testutil.Test[int16]([]int16{v}, Int16, t)
		testutil.TestSkip[int16]([]int16{v}, Int16, t)
	})
}

func FuzzInt8(f *testing.F) {
	seeds := []int8{0, 1, -1, 127, -127, math.MinInt8, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		testutil.Test[int8]([]int8{v}, Int8, t)
		testutil.TestSkip[int8]([]int8{v}, Int8, t)
	})
}

func FuzzInt(f *testing.F) {
	seeds := []int{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt, math.MaxInt}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		testutil.Test[int]([]int{v}, Int, t)
		testutil.TestSkip[int]([]int{v}, Int, t)
	})
}

func FuzzInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int64.Unmarshal(bs)
		Int64.Skip(bs)
	})
}

func FuzzInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int32.Unmarshal(bs)
		Int32.Skip(bs)
	})
}

func FuzzInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int16.Unmarshal(bs)
		Int16.Skip(bs)
	})
}

func FuzzInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int8.Unmarshal(bs)
		Int8.Skip(bs)
	})
}

func FuzzIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int.Unmarshal(bs)
		Int.Skip(bs)
	})
}

// float -----------------------------------------------------------------------

func FuzzFloat64(f *testing.F) {
	seeds := []float64{0, 1, -1, 0.1, -0.1, math.Pi, math.E, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		testutil.Test[float64]([]float64{v}, Float64, t)
		testutil.TestSkip[float64]([]float64{v}, Float64, t)
	})
}

func FuzzFloat32(f *testing.F) {
	seeds := []float32{0, 1, -1, 0.1, -0.1, math.Pi, math.E, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		testutil.Test[float32]([]float32{v}, Float32, t)
		testutil.TestSkip[float32]([]float32{v}, Float32, t)
	})
}

func FuzzFloat64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float64.Unmarshal(bs)
		Float64.Skip(bs)
	})
}

func FuzzFloat32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float32.Unmarshal(bs)
		Float32.Skip(bs)
	})
}

// positive_int ----------------------------------------------------------------

func FuzzPositiveInt64(f *testing.F) {
	seeds := []int64{0, 1, 127, 128, 255, 256, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		if v < 0 {
			return
		}
		testutil.Test[int64]([]int64{v}, PositiveInt64, t)
		testutil.TestSkip[int64]([]int64{v}, PositiveInt64, t)
	})
}

func FuzzPositiveInt32(f *testing.F) {
	seeds := []int32{0, 1, 127, 128, 255, 256, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		if v < 0 {
			return
		}
		testutil.Test[int32]([]int32{v}, PositiveInt32, t)
		testutil.TestSkip[int32]([]int32{v}, PositiveInt32, t)
	})
}

func FuzzPositiveInt16(f *testing.F) {
	seeds := []int16{0, 1, 127, 128, 255, 256, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		if v < 0 {
			return
		}
		testutil.Test[int16]([]int16{v}, PositiveInt16, t)
		testutil.TestSkip[int16]([]int16{v}, PositiveInt16, t)
	})
}

func FuzzPositiveInt8(f *testing.F) {
	seeds := []int8{0, 1, 127, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		if v < 0 {
			return
		}
		testutil.Test[int8]([]int8{v}, PositiveInt8, t)
		testutil.TestSkip[int8]([]int8{v}, PositiveInt8, t)
	})
}

func FuzzPositiveInt(f *testing.F) {
	seeds := []int{0, 1, 127, 128, 255, 256, math.MaxInt}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		if v < 0 {
			return
		}
		testutil.Test[int]([]int{v}, PositiveInt, t)
		testutil.TestSkip[int]([]int{v}, PositiveInt, t)
	})
}

func FuzzPositiveInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt64.Unmarshal(bs)
		PositiveInt64.Skip(bs)
	})
}

func FuzzPositiveInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt32.Unmarshal(bs)
		PositiveInt32.Skip(bs)
	})
}

func FuzzPositiveInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt16.Unmarshal(bs)
		PositiveInt16.Skip(bs)
	})
}

func FuzzPositiveInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt8.Unmarshal(bs)
		PositiveInt8.Skip(bs)
	})
}

func FuzzPositiveIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt.Unmarshal(bs)
		PositiveInt.Skip(bs)
	})
}
