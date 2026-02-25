package raw

import (
	"testing"
	"time"

	"github.com/mus-format/mus-go/testutil"
)

// byte ------------------------------------------------------------------------

func FuzzByte(f *testing.F) {
	seeds := []byte{0, 1, 255}
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

// uint64 ----------------------------------------------------------------------

func FuzzUint64(f *testing.F) {
	seeds := []uint64{0, 1, 1 << 63}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v uint64) {
		testutil.Test[uint64]([]uint64{v}, Uint64, t)
		testutil.TestSkip[uint64]([]uint64{v}, Uint64, t)
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
		testutil.Test[uint32]([]uint32{v}, Uint32, t)
		testutil.TestSkip[uint32]([]uint32{v}, Uint32, t)
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
		testutil.Test[uint16]([]uint16{v}, Uint16, t)
		testutil.TestSkip[uint16]([]uint16{v}, Uint16, t)
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
		testutil.Test[uint8]([]uint8{v}, Uint8, t)
		testutil.TestSkip[uint8]([]uint8{v}, Uint8, t)
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
		testutil.Test[uint]([]uint{v}, Uint, t)
		testutil.TestSkip[uint]([]uint{v}, Uint, t)
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
		testutil.Test[int64]([]int64{v}, Int64, t)
		testutil.TestSkip[int64]([]int64{v}, Int64, t)
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
		testutil.Test[int32]([]int32{v}, Int32, t)
		testutil.TestSkip[int32]([]int32{v}, Int32, t)
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
		testutil.Test[int16]([]int16{v}, Int16, t)
		testutil.TestSkip[int16]([]int16{v}, Int16, t)
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
		testutil.Test[int8]([]int8{v}, Int8, t)
		testutil.TestSkip[int8]([]int8{v}, Int8, t)
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
		testutil.Test[int]([]int{v}, Int, t)
		testutil.TestSkip[int]([]int{v}, Int, t)
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
		testutil.Test[float64]([]float64{v}, Float64, t)
		testutil.TestSkip[float64]([]float64{v}, Float64, t)
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
		testutil.Test[float32]([]float32{v}, Float32, t)
		testutil.TestSkip[float32]([]float32{v}, Float32, t)
	})
}

func FuzzFloat32Unmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Float32.Unmarshal(bs)
		Float32.Skip(bs)
	})
}

// time ------------------------------------------------------------------------

func FuzzTimeUnixUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, sec int64) {
		v := time.Unix(sec, 0).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixUTC, t)
	})
}

func FuzzTimeUnixMilliUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, milli int64) {
		v := time.UnixMilli(milli).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMilliUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMilliUTC, t)
	})
}

func FuzzTimeUnixMicroUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, micro int64) {
		v := time.UnixMicro(micro).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixMicroUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixMicroUTC, t)
	})
}

func FuzzTimeUnixNanoUTC(f *testing.F) {
	f.Fuzz(func(t *testing.T, nano int64) {
		v := time.Unix(0, nano).UTC()
		testutil.Test[time.Time]([]time.Time{v}, TimeUnixNanoUTC, t)
		testutil.TestSkip[time.Time]([]time.Time{v}, TimeUnixNanoUTC, t)
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
