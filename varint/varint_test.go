package varint

import (
	"testing"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test"
)

func TestVarint_unmarshalUint(t *testing.T) {
	t.Run("unmarshalUint should return ErrTooSmallByteSlice if bs is empty",
		func(t *testing.T) {
			var (
				wantV     uint64 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{}
				v, n, err        = unmarshalUint[uint64](0, 0, bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 3
				wantErr          = com.ErrOverflow
				bs               = []byte{200, 200, 200}
				v, n, err        = unmarshalUint[uint16](com.Uint16MaxVarintLen,
					com.Uint16MaxLastByte, bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalUint should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 2
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{200, 200}
				v, n, err        = unmarshalUint[uint16](com.Uint16MaxVarintLen,
					com.Uint16MaxLastByte, bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_skipUint(t *testing.T) {
	t.Run("skipUint should return ErrTooSmallByteSlice if bs is empty",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = skipUint(0, 0, bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipUint should return ErrOverflow if there is no varint end",
		func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = com.ErrOverflow
				bs      = []byte{200, 200, 200, 200, 200}
				n, err  = skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
					bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipUint shold return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 2
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{200, 200}
				n, err  = skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
					bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Uint64(t *testing.T) {
	ser := Uint64
	test.Test(ctestutil.Uint64TestCases, ser, t)
	test.TestSkip(ctestutil.Uint64TestCases, ser, t)
}

func TestVarint_Uint32(t *testing.T) {
	ser := Uint32
	test.Test(ctestutil.Uint32TestCases, ser, t)
	test.TestSkip(ctestutil.Uint32TestCases, ser, t)
}

func TestVarint_Uint16(t *testing.T) {
	ser := Uint16
	test.Test(ctestutil.Uint16TestCases, ser, t)
	test.TestSkip(ctestutil.Uint16TestCases, ser, t)
}

func TestVarint_Uint8(t *testing.T) {
	ser := Uint8
	test.Test(ctestutil.Uint8TestCases, ser, t)
	test.TestSkip(ctestutil.Uint8TestCases, ser, t)
}

func TestVarint_Uint(t *testing.T) {
	ser := Uint
	test.Test(ctestutil.UintTestCases, ser, t)
	test.TestSkip(ctestutil.UintTestCases, ser, t)
}

func TestVarint_Int64(t *testing.T) {
	t.Run("Int64 serializer should succeed",
		func(t *testing.T) {
			ser := Int64
			test.Test(ctestutil.Int64TestCases, ser, t)
			test.TestSkip(ctestutil.Int64TestCases, ser, t)
		})

	t.Run("UnmarshalInt64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int64 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = Int64.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Int32(t *testing.T) {
	t.Run("Int32 serializer should succeed",
		func(t *testing.T) {
			ser := Int32
			test.Test(ctestutil.Int32TestCases, ser, t)
			test.TestSkip(ctestutil.Int32TestCases, ser, t)
		})

	t.Run("UnmarshalInt32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int32 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = Int32.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Int16(t *testing.T) {
	t.Run("Int16 serializer should succeed",
		func(t *testing.T) {
			ser := Int16
			test.Test(ctestutil.Int16TestCases, ser, t)
			test.TestSkip(ctestutil.Int16TestCases, ser, t)
		})

	t.Run("UnmarshalInt16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int16 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = Int16.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Int8(t *testing.T) {
	t.Run("Int8 serializer should succeed",
		func(t *testing.T) {
			ser := Int8
			test.Test(ctestutil.Int8TestCases, ser, t)
			test.TestSkip(ctestutil.Int8TestCases, ser, t)
		})

	t.Run("UnmarshalInt8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int8 = 0
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{}
				v, n, err      = Int8.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Int(t *testing.T) {
	t.Run("Int serializer should succeed",
		func(t *testing.T) {
			ser := Int
			test.Test(ctestutil.IntTestCases, ser, t)
			test.TestSkip(ctestutil.IntTestCases, ser, t)
		})

	t.Run("UnmarshaPositivelInt should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     = 0
				wantN     = 0
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{}
				v, n, err = Int.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_PositiveInt64(t *testing.T) {
	t.Run("PositiveInt64 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt64
			test.Test(ctestutil.Int64TestCases, ser, t)
			test.TestSkip(ctestutil.Int64TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int64 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = PositiveInt64.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_PositiveInt32(t *testing.T) {
	t.Run("PositiveInt32 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt32
			test.Test(ctestutil.Int32TestCases, ser, t)
			test.TestSkip(ctestutil.Int32TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int32 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = PositiveInt32.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_PositiveInt16(t *testing.T) {
	t.Run("PositiveInt16 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt16
			test.Test(ctestutil.Int16TestCases, ser, t)
			test.TestSkip(ctestutil.Int16TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int16 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = PositiveInt16.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_PositiveInt8(t *testing.T) {
	t.Run("PositiveInt8 serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt8
			test.Test(ctestutil.Int8TestCases, ser, t)
			test.TestSkip(ctestutil.Int8TestCases, ser, t)
		})

	t.Run("UnmarshalPositiveInt8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     int8 = 0
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{}
				v, n, err      = PositiveInt8.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_PositiveInt(t *testing.T) {
	t.Run("PositiveInt serializer should succeed",
		func(t *testing.T) {
			ser := PositiveInt
			test.Test(ctestutil.IntTestCases, ser, t)
			test.TestSkip(ctestutil.IntTestCases, ser, t)
		})

	t.Run("UnmarshaPositivelInt should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     = 0
				wantN     = 0
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{}
				v, n, err = PositiveInt.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestVarint_Byte(t *testing.T) {
	ser := Byte
	test.Test(ctestutil.ByteTestCases, ser, t)
	test.TestSkip(ctestutil.ByteTestCases, ser, t)
}

func TestVarint_Float64(t *testing.T) {
	t.Run("Float64 serializer should succeed",
		func(t *testing.T) {
			ser := Float64
			test.Test(ctestutil.Float64TestCases, ser, t)
			test.TestSkip(ctestutil.Float64TestCases, ser, t)
		})

	t.Run("UnmarshalFloat64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     float64 = 0
				wantN             = 0
				wantErr           = mus.ErrTooSmallByteSlice
				bs                = []byte{}
				v, n, err         = Float64.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				nil, t)
		})
}

func TestVarint_Float32(t *testing.T) {
	t.Run("Float32 serializer should succeed",
		func(t *testing.T) {
			ser := Float32
			test.Test(ctestutil.Float32TestCases, ser, t)
			test.TestSkip(ctestutil.Float32TestCases, ser, t)
		})

	t.Run("UnmarshalFloat32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     float32 = 0
				wantN             = 0
				wantErr           = mus.ErrTooSmallByteSlice
				bs                = []byte{}
				v, n, err         = Float32.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
				nil, t)
		})
}
