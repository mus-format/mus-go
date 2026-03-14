package varint

import (
	"math"
	"testing"

	"github.com/mus-format/mus-go/test"
)

// byte ------------------------------------------------------------------------

func FuzzVarint_Byte(f *testing.F) {
	seeds := []byte{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v byte) {
		test.Test([]byte{v}, Byte, t)
		test.TestSkip([]byte{v}, Byte, t)
	})
}

func FuzzVarint_ByteUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Byte.Unmarshal(bs)
		Byte.Skip(bs)
	})
}

// uint ------------------------------------------------------------------------

func FuzzVarint_Uint64(f *testing.F) {
	seeds := []uint64{0, 1, 127, 128, 255, 256, math.MaxUint64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		test.Test([]uint64{v}, Uint64, t)
		test.TestSkip([]uint64{v}, Uint64, t)
	})
}

func FuzzVarint_Uint32(f *testing.F) {
	seeds := []uint32{0, 1, 127, 128, 255, 256, math.MaxUint32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint32) {
		test.Test([]uint32{v}, Uint32, t)
		test.TestSkip([]uint32{v}, Uint32, t)
	})
}

func FuzzVarint_Uint16(f *testing.F) {
	seeds := []uint16{0, 1, 127, 128, 255, 256, math.MaxUint16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint16) {
		test.Test([]uint16{v}, Uint16, t)
		test.TestSkip([]uint16{v}, Uint16, t)
	})
}

func FuzzVarint_Uint8(f *testing.F) {
	seeds := []uint8{0, 1, 127, 128, 255}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint8) {
		test.Test([]uint8{v}, Uint8, t)
		test.TestSkip([]uint8{v}, Uint8, t)
	})
}

func FuzzVarint_Uint(f *testing.F) {
	seeds := []uint{0, 1, 127, 128, 255, 256, math.MaxUint}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint) {
		test.Test([]uint{v}, Uint, t)
		test.TestSkip([]uint{v}, Uint, t)
	})
}

func FuzzVarint_Uint64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint64.Unmarshal(bs)
		Uint64.Skip(bs)
	})
}

func FuzzVarint_Uint32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint32.Unmarshal(bs)
		Uint32.Skip(bs)
	})
}

func FuzzVarint_Uint16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint16.Unmarshal(bs)
		Uint16.Skip(bs)
	})
}

func FuzzVarint_Uint8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint8.Unmarshal(bs)
		Uint8.Skip(bs)
	})
}

func FuzzVarint_UintUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Uint.Unmarshal(bs)
		Uint.Skip(bs)
	})
}

// int -------------------------------------------------------------------------

func FuzzVarint_Int64(f *testing.F) {
	seeds := []int64{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt64, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		test.Test([]int64{v}, Int64, t)
		test.TestSkip([]int64{v}, Int64, t)
	})
}

func FuzzVarint_Int32(f *testing.F) {
	seeds := []int32{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt32, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		test.Test([]int32{v}, Int32, t)
		test.TestSkip([]int32{v}, Int32, t)
	})
}

func FuzzVarint_Int16(f *testing.F) {
	seeds := []int16{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt16, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		test.Test([]int16{v}, Int16, t)
		test.TestSkip([]int16{v}, Int16, t)
	})
}

func FuzzVarint_Int8(f *testing.F) {
	seeds := []int8{0, 1, -1, 127, -127, math.MinInt8, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		test.Test([]int8{v}, Int8, t)
		test.TestSkip([]int8{v}, Int8, t)
	})
}

func FuzzVarint_Int(f *testing.F) {
	seeds := []int{0, 1, -1, 127, -127, 128, -128, 255, -255, 256, -256, math.MinInt, math.MaxInt}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		test.Test([]int{v}, Int, t)
		test.TestSkip([]int{v}, Int, t)
	})
}

func FuzzVarint_Int64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int64.Unmarshal(bs)
		Int64.Skip(bs)
	})
}

func FuzzVarint_Int32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int32.Unmarshal(bs)
		Int32.Skip(bs)
	})
}

func FuzzVarint_Int16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int16.Unmarshal(bs)
		Int16.Skip(bs)
	})
}

func FuzzVarint_Int8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int8.Unmarshal(bs)
		Int8.Skip(bs)
	})
}

func FuzzVarint_IntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Int.Unmarshal(bs)
		Int.Skip(bs)
	})
}

// float -----------------------------------------------------------------------

func FuzzVarint_Float64(f *testing.F) {
	seeds := []float64{0, 1, -1, 0.1, -0.1, math.Pi, math.E, math.Inf(1), math.Inf(-1), math.NaN()}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float64) {
		test.Test([]float64{v}, Float64, t)
		test.TestSkip([]float64{v}, Float64, t)
	})
}

func FuzzVarint_Float32(f *testing.F) {
	seeds := []float32{0, 1, -1, 0.1, -0.1, math.Pi, math.E, float32(math.Inf(1)), float32(math.Inf(-1)), float32(math.NaN())}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v float32) {
		test.Test([]float32{v}, Float32, t)
		test.TestSkip([]float32{v}, Float32, t)
	})
}

func FuzzVarint_Float64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float64.Unmarshal(bs)
		Float64.Skip(bs)
	})
}

func FuzzVarint_Float32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float32.Unmarshal(bs)
		Float32.Skip(bs)
	})
}

// positive_int ----------------------------------------------------------------

func FuzzVarint_PositiveInt64(f *testing.F) {
	seeds := []int64{0, 1, 127, 128, 255, 256, math.MaxInt64}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int64) {
		if v < 0 {
			return
		}
		test.Test([]int64{v}, PositiveInt64, t)
		test.TestSkip([]int64{v}, PositiveInt64, t)
	})
}

func FuzzVarint_PositiveInt32(f *testing.F) {
	seeds := []int32{0, 1, 127, 128, 255, 256, math.MaxInt32}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int32) {
		if v < 0 {
			return
		}
		test.Test([]int32{v}, PositiveInt32, t)
		test.TestSkip([]int32{v}, PositiveInt32, t)
	})
}

func FuzzVarint_PositiveInt16(f *testing.F) {
	seeds := []int16{0, 1, 127, 128, 255, 256, math.MaxInt16}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int16) {
		if v < 0 {
			return
		}
		test.Test([]int16{v}, PositiveInt16, t)
		test.TestSkip([]int16{v}, PositiveInt16, t)
	})
}

func FuzzVarint_PositiveInt8(f *testing.F) {
	seeds := []int8{0, 1, 127, math.MaxInt8}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int8) {
		if v < 0 {
			return
		}
		test.Test([]int8{v}, PositiveInt8, t)
		test.TestSkip([]int8{v}, PositiveInt8, t)
	})
}

func FuzzVarint_PositiveInt(f *testing.F) {
	seeds := []int{0, 1, 127, 128, 255, 256, math.MaxInt}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v int) {
		if v < 0 {
			return
		}
		test.Test([]int{v}, PositiveInt, t)
		test.TestSkip([]int{v}, PositiveInt, t)
	})
}

func FuzzVarint_PositiveInt64Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt64.Unmarshal(bs)
		PositiveInt64.Skip(bs)
	})
}

func FuzzVarint_PositiveInt32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt32.Unmarshal(bs)
		PositiveInt32.Skip(bs)
	})
}

func FuzzVarint_PositiveInt16Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt16.Unmarshal(bs)
		PositiveInt16.Skip(bs)
	})
}

func FuzzVarint_PositiveInt8Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt8.Unmarshal(bs)
		PositiveInt8.Skip(bs)
	})
}

func FuzzVarint_PositiveIntUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		PositiveInt.Unmarshal(bs)
		PositiveInt.Skip(bs)
	})
}
