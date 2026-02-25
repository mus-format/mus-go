package raw

import (
	"os"
	"testing"
	"time"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testutil"
)

func TestRaw(t *testing.T) {
	t.Run("setUpUintFuncs", func(t *testing.T) {
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
	})

	t.Run("setUpIntFuncs", func(t *testing.T) {
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
	})

	t.Run("unmarshalInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint64 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1, 2, 3, 4, 5}
				v, n, err        = unmarshalInteger64[uint64](bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger64 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3, 4, 5, 6, 7}
				n, err  = SkipInteger64(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint32 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1, 2, 3}
				v, n, err        = unmarshalInteger32[uint32](bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger32 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1, 2, 3}
				n, err  = SkipInteger32(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint16 = 0
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{1}
				v, n, err        = unmarshalInteger16[uint16](bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger16 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{1}
				n, err  = SkipInteger16(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("unmarshalInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     uint8 = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
				v, n, err       = unmarshalInteger8[uint8](bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("skipInteger8 should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = SkipInteger8(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("byte", func(t *testing.T) {
		t.Run("Byte serializer should work correctly",
			func(t *testing.T) {
				ser := Byte
				testutil.Test[byte](ctestutil.ByteTestCases, ser, t)
				testutil.TestSkip[byte](ctestutil.ByteTestCases, ser, t)
			})
	})

	t.Run("unsigned", func(t *testing.T) {
		t.Run("Uint64 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint64
				testutil.Test[uint64](ctestutil.Uint64TestCases, ser, t)
				testutil.TestSkip[uint64](ctestutil.Uint64TestCases, ser, t)
			})

		t.Run("Uint32 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint32
				testutil.Test[uint32](ctestutil.Uint32TestCases, ser, t)
				testutil.TestSkip[uint32](ctestutil.Uint32TestCases, ser, t)
			})

		t.Run("Uint16 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint16
				testutil.Test[uint16](ctestutil.Uint16TestCases, ser, t)
				testutil.TestSkip[uint16](ctestutil.Uint16TestCases, ser, t)
			})

		t.Run("Uint8 serializer should work correctly",
			func(t *testing.T) {
				ser := Uint8
				testutil.Test[uint8](ctestutil.Uint8TestCases, ser, t)
				testutil.TestSkip[uint8](ctestutil.Uint8TestCases, ser, t)
			})

		t.Run("Uint serializer should work correctly",
			func(t *testing.T) {
				ser := Uint
				testutil.Test[uint](ctestutil.UintTestCases, ser, t)
				testutil.TestSkip[uint](ctestutil.UintTestCases, ser, t)
			})
	})

	t.Run("signed", func(t *testing.T) {
		t.Run("Int64 serializer should work correctly",
			func(t *testing.T) {
				ser := Int64
				testutil.Test[int64](ctestutil.Int64TestCases, ser, t)
				testutil.TestSkip[int64](ctestutil.Int64TestCases, ser, t)
			})

		t.Run("Int32 serializer should work correctly",
			func(t *testing.T) {
				ser := Int32
				testutil.Test[int32](ctestutil.Int32TestCases, ser, t)
				testutil.TestSkip[int32](ctestutil.Int32TestCases, ser, t)
			})

		t.Run("Int16 serializer should work correctly",
			func(t *testing.T) {
				ser := Int16
				testutil.Test[int16](ctestutil.Int16TestCases, ser, t)
				testutil.TestSkip[int16](ctestutil.Int16TestCases, ser, t)
			})

		t.Run("Int8 serializer should work correctly",
			func(t *testing.T) {
				ser := Int8
				testutil.Test[int8](ctestutil.Int8TestCases, ser, t)
				testutil.TestSkip[int8](ctestutil.Int8TestCases, ser, t)
			})

		t.Run("Int serializer should work correctly",
			func(t *testing.T) {
				ser := Int
				testutil.Test[int](ctestutil.IntTestCases, ser, t)
				testutil.TestSkip[int](ctestutil.IntTestCases, ser, t)
			})
	})

	t.Run("float", func(t *testing.T) {
		t.Run("float64", func(t *testing.T) {
			t.Run("Float64 serializer should work correctly",
				func(t *testing.T) {
					ser := Float64
					testutil.Test[float64](ctestutil.Float64TestCases, ser, t)
					testutil.TestSkip[float64](ctestutil.Float64TestCases, ser, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float64 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2, 3, 4, 5}
						v, n, err         = Float64.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2, 3, 4, 5}
						n, err  = Float64.Skip(bs)
					)
					ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
				})
		})

		t.Run("float32", func(t *testing.T) {
			t.Run("Float32 serializer should work correctly",
				func(t *testing.T) {
					ser := Float32
					testutil.Test[float32](ctestutil.Float32TestCases, ser, t)
					testutil.TestSkip[float32](ctestutil.Float32TestCases, ser, t)
				})

			t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantV     float32 = 0
						wantN             = 0
						wantErr           = mus.ErrTooSmallByteSlice
						bs                = []byte{1, 2}
						v, n, err         = Float32.Unmarshal(bs)
					)
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
				})

			t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
				func(t *testing.T) {
					var (
						wantN   = 0
						wantErr = mus.ErrTooSmallByteSlice
						bs      = []byte{1, 2}
						n, err  = Float32.Skip(bs)
					)
					ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
				})
		})
	})

	t.Run("time", func(t *testing.T) {
		os.Setenv("TZ", "")

		t.Run("time_unix_utc", func(t *testing.T) {
			t.Run("TimeUnixUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						sec = time.Now().Unix()
						tm  = time.Unix(sec, 0)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixUTC, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_milli_utc", func(t *testing.T) {
			t.Run("TimeUnixMilliUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMilli()
						tm    = time.UnixMilli(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMilliUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixMilliUTC, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_micro_utc", func(t *testing.T) {
			t.Run("TimeUnixMicroUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						milli = time.Now().UnixMicro()
						tm    = time.UnixMicro(milli)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixMicroUTC, t)
				})

			t.Run("We should be able to serializer the zero Time",
				func(t *testing.T) {
					testutil.Test[time.Time]([]time.Time{{}}, TimeUnixMicroUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{{}}, TimeUnixMicroUTC, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})

		t.Run("time_unix_nano_utc", func(t *testing.T) {
			t.Run("TimeUnixNanoUTC serializer should work correctly",
				func(t *testing.T) {
					var (
						nano = time.Now().UnixNano()
						tm   = time.Unix(0, nano)
					)
					testutil.Test[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
					testutil.TestSkip[time.Time]([]time.Time{tm}, TimeUnixNanoUTC, t)
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
					ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
						nil, t)
				})
		})
	})
}
