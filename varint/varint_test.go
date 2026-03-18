package varint

import (
	"testing"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestVarint_unmarshalUint(t *testing.T) {
	t.Run("unmarshalUint should return ErrTooSmallByteSlice if bs is empty",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			v, n, err := unmarshalUint[uint64](0, 0, bs)
			asserterror.Equal(t, want.V, v)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})

	t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint16]{
					V:   0,
					N:   3,
					Err: com.ErrOverflow,
				}
				bs = []byte{200, 200, 200}
			)
			v, n, err := unmarshalUint[uint16](com.Uint16MaxVarintLen,
				com.Uint16MaxLastByte, bs)
			asserterror.Equal(t, want.V, v)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})

	t.Run("unmarshalUint should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[uint16]{
					V:   0,
					N:   2,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{200, 200}
			)
			v, n, err := unmarshalUint[uint16](com.Uint16MaxVarintLen,
				com.Uint16MaxLastByte, bs)
			asserterror.Equal(t, want.V, v)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})
}

func TestVarint_skipUint(t *testing.T) {
	t.Run("skipUint should return ErrTooSmallByteSlice if bs is empty",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			n, err := skipUint(0, 0, bs)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})

	t.Run("skipUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   3,
					Err: com.ErrOverflow,
				}
				bs = []byte{200, 200, 200, 200, 200}
			)
			n, err := skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
				bs)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})

	t.Run("skipUint shold return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   2,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{200, 200}
			)
			n, err := skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
				bs)
			asserterror.Equal(t, want.N, n)
			asserterror.EqualError(t, err, want.Err)
		})
}

func TestVarint_Uint64(t *testing.T) {
	ser := Uint64
	test.Test(ctest.Uint64TestCases, ser, t)
	test.TestSkip(ctest.Uint64TestCases, ser, t)
}

func TestVarint_Uint32(t *testing.T) {
	ser := Uint32
	test.Test(ctest.Uint32TestCases, ser, t)
	test.TestSkip(ctest.Uint32TestCases, ser, t)
}

func TestVarint_Uint16(t *testing.T) {
	ser := Uint16
	test.Test(ctest.Uint16TestCases, ser, t)
	test.TestSkip(ctest.Uint16TestCases, ser, t)
}

func TestVarint_Uint8(t *testing.T) {
	ser := Uint8
	test.Test(ctest.Uint8TestCases, ser, t)
	test.TestSkip(ctest.Uint8TestCases, ser, t)
}

func TestVarint_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctest.UintTestCases, ser, t)
	test.TestSkip(ctest.UintTestCases, ser, t)
}

func TestVarint_Int64(t *testing.T) {
	t.Run("Int64 serializer should succeed",
		func(t *testing.T) {
			ser := Int64
			test.Test(ctest.Int64TestCases, ser, t)
			test.TestSkip(ctest.Int64TestCases, ser, t)
		})

	t.Run("UnmarshalInt64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Int64, want, nil, t)
		})
}

func TestVarint_Int32(t *testing.T) {
	t.Run("Int32 serializer should succeed",
		func(t *testing.T) {
			ser := Int32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})

	t.Run("UnmarshalInt32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int32]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Int32, want, nil, t)
		})
}

func TestVarint_Int16(t *testing.T) {
	t.Run("Int16 serializer should succeed",
		func(t *testing.T) {
			ser := Int16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})

	t.Run("UnmarshalInt16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int16]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Int16, want, nil, t)
		})
}

func TestVarint_Int8(t *testing.T) {
	t.Run("Int8 serializer should succeed",
		func(t *testing.T) {
			ser := Int8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})

	t.Run("UnmarshalInt8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int8]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Int8, want, nil, t)
		})
}

func TestVarint_Int(t *testing.T) {
	t.Run("Int serializer should succeed",
		func(t *testing.T) {
			ser := Int
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})

	t.Run("UnmarshaPositivelInt should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Int, want, nil, t)
		})
}

func TestVarint_PositiveInt64(t *testing.T) {
	t.Run("PositiveInt64 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt64
			test.Test(ctest.Int64TestCases, ser, t)
			test.TestSkip(ctest.Int64TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, PositiveInt64, want, nil, t)
		})
}

func TestVarint_PositiveInt32(t *testing.T) {
	t.Run("PositiveInt32 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt32
			test.Test(ctest.Int32TestCases, ser, t)
			test.TestSkip(ctest.Int32TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int32]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, PositiveInt32, want, nil, t)
		})
}

func TestVarint_PositiveInt16(t *testing.T) {
	t.Run("PositiveInt16 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt16
			test.Test(ctest.Int16TestCases, ser, t)
			test.TestSkip(ctest.Int16TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int16]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, PositiveInt16, want, nil, t)
		})
}

func TestVarint_PositiveInt8(t *testing.T) {
	t.Run("PositiveInt8 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt8
			test.Test(ctest.Int8TestCases, ser, t)
			test.TestSkip(ctest.Int8TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int8]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, PositiveInt8, want, nil, t)
		})
}

func TestVarint_PositiveInt(t *testing.T) {
	t.Run("PositiveInt serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt
			test.Test(ctest.IntTestCases, ser, t)
			test.TestSkip(ctest.IntTestCases, ser, t)
		})

	t.Run("UnmarshaPositivelInt should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[int]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, PositiveInt, want, nil, t)
		})
}

func TestVarint_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctest.ByteTestCases, ser, t)
	test.TestSkip(ctest.ByteTestCases, ser, t)
}

func TestVarint_Float64(t *testing.T) {
	t.Run("Float64 serializer should succeed",
		func(t *testing.T) {
			ser := Float64
			test.Test(ctest.Float64TestCases, ser, t)
			test.TestSkip(ctest.Float64TestCases, ser, t)
		})

	t.Run("UnmarshalFloat64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[float64]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Float64, want, nil, t)
		})
}

func TestVarint_Float32(t *testing.T) {
	t.Run("Float32 serializer should succeed",
		func(t *testing.T) {
			ser := Float32
			test.Test(ctest.Float32TestCases, ser, t)
			test.TestSkip(ctest.Float32TestCases, ser, t)
		})

	t.Run("UnmarshalFloat32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[float32]{
					V:   0,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs = []byte{}
			)
			test.TestUnmarshalOnly(bs, Float32, want, nil, t)
		})
}
