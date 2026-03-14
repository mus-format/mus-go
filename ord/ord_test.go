package ord

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	cmock "github.com/mus-format/common-go/testutil/mock"
	"github.com/mus-format/mus-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	mapops "github.com/mus-format/mus-go/options/map"
	slops "github.com/mus-format/mus-go/options/slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/test"
	mock "github.com/mus-format/mus-go/test/mock"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestOrd_Bool(t *testing.T) {
	t.Run("Bool serializer should succeed",
		func(t *testing.T) {
			ser := Bool
			test.Test(ctestutil.BoolTestCases, ser, t)
			test.TestSkip(ctestutil.BoolTestCases, ser, t)
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     = false
				wantN     = 0
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{}
				v, n, err = Bool.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				wantV     = false
				wantN     = 0
				wantErr   = com.ErrWrongFormat
				bs        = []byte{3}
				v, n, err = Bool.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = Bool.Skip(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = com.ErrWrongFormat
				bs      = []byte{3}
				n, err  = Bool.Skip(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})
}

func TestOrd_String(t *testing.T) {
	t.Run("String serializer should succeed",
		func(t *testing.T) {
			ser := String
			test.Test(ctestutil.StringTestCases, ser, t)
			test.TestSkip(ctestutil.StringTestCases, ser, t)
		})

	t.Run("We should be able to set a length serializer",
		func(t *testing.T) {
			var (
				str, lenSer = test.StringLenTestData(t)
				ser         = NewStringSer(strops.WithLenSer(lenSer))
			)
			test.Test([]string{str}, ser, t)
			test.TestSkip([]string{str}, ser, t)
		})

	t.Run("Marshal should panic with ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			wantErr := mus.ErrTooSmallByteSlice
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					if err != wantErr {
						t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
					}
				}
			}()
			String.Marshal("hello world", make([]byte, 2))
		})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 0
				wantErr   = mus.ErrTooSmallByteSlice
				v, n, err = String.Unmarshal(nil)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantV     = ""
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				v, n, err = String.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 1
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{2, 2}
				v, n, err = String.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If lenSer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("lenSer error")
				lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser    = NewStringSer(strops.WithLenSer(lenSer))
				n, err = ser.Skip(nil)
				mocks  = []*mok.Mock{lenSer.Mock}
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				n, err    = String.Skip(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{2, 2}
				n, err  = String.Skip(bs)
			)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid string serializer should succeed",
		func(t *testing.T) {
			ser := NewValidStringSer(nil)
			test.Test(ctestutil.StringTestCases, ser, t)
			test.TestSkip(ctestutil.StringTestCases, ser, t)
		})

	t.Run("Valid string serializer with varint length should succeed",
		func(t *testing.T) {
			ser := NewValidStringSer(nil)
			test.Test(ctestutil.StringTestCases, ser, t)
			test.TestSkip(ctestutil.StringTestCases, ser, t)
		})

	t.Run("If lenSer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = ""
				wantN   = 0
				wantErr = errors.New("lenSer error")
				lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser   = NewValidStringSer(strops.WithLenSer(lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			v, n, err := ser.Unmarshal(nil)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantV     = ""
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				ser       = NewValidStringSer(nil)
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   = ""
				wantN   = 1
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{2, 2}
				ser     = NewValidStringSer(nil)
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If lenVl fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = ""
				wantN   = 1
				wantErr = errors.New("lenVl validator error")
				bs      = []byte{2, 2, 2}
				lenVl   = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) { return wantErr },
				)
				ser = NewValidStringSer(strops.WithLenValidator(lenVl))
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If string length == 0 lenVl should work", func(t *testing.T) {
		var (
			wantV                        = ""
			wantN                        = 1
			wantErr                      = errors.New("empty string")
			bs                           = []byte{0}
			lenVl   com.ValidatorFn[int] = func(t int) (err error) {
				return wantErr
			}
			ser = NewValidStringSer(strops.WithLenValidator(lenVl))
		)

		v, n, err := ser.Unmarshal(bs)
		ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
	})
}

func TestOrd_Pointer(t *testing.T) {
	t.Run("Pointer seralizer should succeed",
		func(t *testing.T) {
			var (
				ptr, baseSer = test.PtrTestData(t)
				ser          = NewPtrSer(baseSer)
			)
			test.Test([]*int{ptr}, ser, t)
			test.TestSkip([]*int{ptr}, ser, t)
		})

	t.Run("Pointer serializer should succeed for nil ptr",
		func(t *testing.T) {
			ser := NewPtrSer(mus.Serializer[string](nil))
			test.Test([]*string{nil}, ser, t)
			test.TestSkip([]*string{nil}, ser, t)
		})

	t.Run("Pointer serializer should succeed for not nil ptr",
		func(t *testing.T) {
			var (
				str1                            = "one"
				str1Raw                         = append([]byte{6}, []byte(str1)...)
				ptr                             = &str1
				strSer  mock.Serializer[string] = mock.NewSerializer[string]().RegisterMarshalN(2,
					func(v string, bs []byte) (n int) {
						switch v {
						case str1:
							return copy(bs, str1Raw)
						default:
							t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
							return
						}
					},
				).RegisterUnmarshal(
					func(bs []byte) (v string, n int, err error) {
						if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
							return str1, len(str1Raw), nil
						} else {
							t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
							return
						}
					},
				).RegisterSizeN(2,
					func(v string) (size int) {
						switch v {
						case str1:
							return len(str1Raw)
						default:
							t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
							return
						}
					},
				).RegisterSkip(
					func(bs []byte) (n int, err error) {
						if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
							return len(str1Raw), nil
						} else {
							t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
							return
						}
					},
				)
				ser = NewPtrSer(strSer)
			)
			test.Test([]*string{ptr}, ser, t)
			test.TestSkip([]*string{ptr}, ser, t)
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				ser             = NewPtrSer(mus.Serializer[string](nil))
			)
			v, n, err := ser.Unmarshal([]byte{})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 0
				wantErr         = com.ErrWrongFormat
				ser             = NewPtrSer(mus.Serializer[string](nil))
			)
			v, n, err := ser.Unmarshal([]byte{2})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If base serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 5
				wantErr         = errors.New("base serializer error")
				baseSer         = mock.NewSerializer[string]().RegisterUnmarshal(
					func(bs []byte) (v string, n int, err error) {
						return "", 4, wantErr
					},
				)
				ser   = NewPtrSer(baseSer)
				mocks = []*mok.Mock{baseSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{byte(com.NotNil)})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				ser     = NewPtrSer(mus.Serializer[string](nil))
			)
			n, err := ser.Skip([]byte{})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrWrongFormat if meets wrong format",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = com.ErrWrongFormat
				ser     = NewPtrSer(mus.Serializer[string](nil))
			)
			n, err := ser.Skip([]byte{2})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("If base serializer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = errors.New("error")
				baseSer = mock.NewSerializer[string]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						return 2, wantErr
					},
				)
				ser   = NewPtrSer(baseSer)
				mocks = []*mok.Mock{baseSer.Mock}
			)
			n, err := ser.Skip([]byte{byte(com.NotNil)})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})
}

func TestOrd_ByteSlice(t *testing.T) {
	t.Run("ByteSlice serializer should succeed for empty slice",
		func(t *testing.T) {
			var (
				sl  = []byte{}
				ser = ByteSlice
			)
			test.Test([][]byte{sl}, ser, t)
			test.TestSkip([][]byte{sl}, ser, t)
		})

	t.Run("ByteSlice serializer should succeed for non-empty slice",
		func(t *testing.T) {
			var (
				sl  = []byte{0, 1, 1, 255, 100, 0, 1, 10}
				ser = ByteSlice
			)
			test.Test([][]byte{sl}, ser, t)
			test.TestSkip([][]byte{sl}, ser, t)
		})

	t.Run("We should be able to set a length serializer", func(t *testing.T) {
		var (
			sl, lenSer = test.ByteSliceLenTestData(t)
			ser        = NewByteSliceSer(bslops.WithLenSer(lenSer))
		)
		test.Test([][]byte{sl}, ser, t)
		test.TestSkip([][]byte{sl}, ser, t)
	})

	t.Run("Marshal should panic with ErrTooSmallByteSlice if there is no space in bs",
		func(t *testing.T) {
			wantErr := mus.ErrTooSmallByteSlice
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					if err != wantErr {
						t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
					}
				}
			}()
			ByteSlice.Marshal([]byte{1, 2, 3, 4}, make([]byte, 2))
		})

	t.Run("Unmarshal should return ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN            = 1
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{4, 1}
				v, n, err        = ByteSlice.Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []byte = nil
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
			)
			v, n, err := ByteSlice.Unmarshal(nil)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN, bs        = NegativeLengthBs()
				wantErr          = com.ErrNegativeLength
			)
			v, n, err := ByteSlice.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Skip should return ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{4, 1}
			)
			n, err := ByteSlice.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("If lenSer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = errors.New("lenSer error")
				lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						return 0, wantN, wantErr
					},
				)
				ser   = NewByteSliceSer(bslops.WithLenSer(lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			n, err := ser.Skip(nil)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
			)
			n, err := ByteSlice.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid ByteSlice serializer should succeed for empty slice",
		func(t *testing.T) {
			var (
				sl  = []byte{}
				ser = NewValidByteSliceSer(nil)
			)
			test.Test([][]byte{sl}, ser, t)
			test.TestSkip([][]byte{sl}, ser, t)
		})

	t.Run("Valid ByteSlice serializer should succeed for non-empty slice",
		func(t *testing.T) {
			var (
				sl  = []byte{0, 1, 1, 255, 100, 0, 1, 10}
				ser = NewValidByteSliceSer(nil)
			)
			test.Test([][]byte{sl}, ser, t)
			test.TestSkip([][]byte{sl}, ser, t)
		})

	t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN            = 1
				wantErr          = mus.ErrTooSmallByteSlice
				bs               = []byte{4, 1}
				v, n, err        = NewValidByteSliceSer(nil).Unmarshal(bs)
			)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If lenSer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []byte = nil
				wantN          = 1
				wantErr        = errors.New("lenSer error")
				lenSer         = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						return 0, wantN, wantErr
					},
				)
				ser   = NewValidByteSliceSer(bslops.WithLenSer(lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			v, n, err := ser.Unmarshal(nil)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []byte = nil
				wantN, bs        = NegativeLengthBs()
				wantErr          = com.ErrNegativeLength
				ser              = NewValidByteSliceSer(nil)
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If lenVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []byte = nil
				wantN          = 1
				wantErr        = errors.New("too large slice")
				bs             = []byte{2, 4, 1}
				lenVl          = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						return wantErr
					},
				)
				ser = NewValidByteSliceSer(bslops.WithLenValidator(lenVl))
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
}

func TestOrd_Slice(t *testing.T) {
	t.Run("Slice serializer should succeed with empty slice",
		func(t *testing.T) {
			var (
				sl  = []string{}
				ser = NewSliceSer(mus.Serializer[string](nil))
			)
			test.Test([][]string{sl}, ser, t)
			test.TestSkip([][]string{sl}, ser, t)
		})

	t.Run("Slice serializer should succeed with not empty slice",
		func(t *testing.T) {
			var (
				sl, elemSer = test.SliceTestData(t)
				ser         = NewSliceSer(elemSer)
				mocks       = []*mok.Mock{elemSer.Mock}
			)
			test.Test([][]string{sl}, ser, t)
			test.TestSkip([][]string{sl}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("We should be able to set a length serializer", func(t *testing.T) {
		var (
			sl, lenSer, elemSer = test.SliceLenTestData(t)
			ser                 = NewSliceSer(elemSer, slops.WithLenSer[string](lenSer))
			mocks               = []*mok.Mock{lenSer.Mock, elemSer.Mock}
		)
		test.Test([][]string{sl}, ser, t)
		test.TestSkip([][]string{sl}, ser, t)

		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []string = nil
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				ser              = NewSliceSer(mus.Serializer[string](nil))
			)
			v, n, err := ser.Unmarshal([]byte{})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []string = nil
				wantN, bs          = NegativeLengthBs()
				wantErr            = com.ErrNegativeLength
				ser                = NewSliceSer(mus.Serializer[string](nil))
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If elemSer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []uint = make([]uint, 1)
				wantN          = 3
				wantErr        = errors.New("Unmarshaller error")
				elemSer        = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 0, 2, wantErr
					},
				)
				ser   = NewSliceSer(elemSer)
				mocks = []*mok.Mock{elemSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{1})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If lenSer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = errors.New("lenSer error")
				lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser   = NewSliceSer(mus.Serializer[string](nil), slops.WithLenSer[string](lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			n, err := ser.Skip([]byte{})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				ser       = NewSliceSer(mus.Serializer[string](nil))
			)
			n, err := ser.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("If elemSer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = errors.New("Unmarshaller error")
				elemSer = mock.NewSerializer[uint]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser   = NewSliceSer(elemSer)
				mocks = []*mok.Mock{elemSer.Mock}
			)
			n, err := ser.Skip([]byte{1})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Valid Slice serializer should succeed with empty slice",
		func(t *testing.T) {
			var (
				sl  = []string{}
				ser = NewValidSliceSer(mus.Serializer[string](nil), nil, nil)
			)
			test.Test([][]string{sl}, ser, t)
			test.TestSkip([][]string{sl}, ser, t)
		})

	t.Run("Valid Slice serializer should succeed with not empty slice",
		func(t *testing.T) {
			var (
				sl, elemSer = test.SliceTestData(t)
				ser         = NewValidSliceSer(elemSer, nil, nil)
				mocks       = []*mok.Mock{elemSer.Mock}
			)
			test.Test([][]string{sl}, ser, t)
			test.TestSkip([][]string{sl}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("If lenSer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []string = nil
				wantN            = 0
				wantErr          = errors.New("lenSer error")
				lenSer           = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser   = NewValidSliceSer(mus.Serializer[string](nil), slops.WithLenSer[string](lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantV     []string = nil
				wantN, bs          = NegativeLengthBs()
				wantErr            = com.ErrNegativeLength
				ser                = NewValidSliceSer(mus.Serializer[string](nil), nil, nil)
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If elemSer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []uint = make([]uint, 1)
				wantN          = 3
				wantErr        = errors.New("Unmarshaller error")
				elemSer        = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 0, 2, wantErr
					},
				)
				ser   = NewValidSliceSer(elemSer, nil, nil)
				mocks = []*mok.Mock{elemSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{1})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If lenVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []uint = nil
				wantN          = 1
				wantErr        = errors.New("lenVl error")
				lenVl          = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						return wantErr
					},
				)
				bs    = []byte{3, 10, 2, 3}
				ser   = NewValidSliceSer(mus.Serializer[uint](nil), slops.WithLenValidator[uint](lenVl))
				mocks = []*mok.Mock{lenVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If elemVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   []uint = []uint{10, 0, 0}
				wantN          = 3
				wantErr        = errors.New("elemVl error")
				bs             = []byte{3, 10, 2, 3}
				elemVl         = cmock.NewValidator[uint]().RegisterValidate(
					func(v uint) (err error) {
						if v != 10 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
						}
						return nil
					},
				).RegisterValidate(
					func(v uint) (err error) {
						if v != 2 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
						}
						return wantErr
					},
				)
				elemSer = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 10, 1, nil
					},
				).RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 2, 1, nil
					},
				)
				ser   = NewValidSliceSer(elemSer, nil, slops.WithElemValidator(elemVl))
				mocks = []*mok.Mock{elemSer.Mock, elemVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func TestOrd_Map(t *testing.T) {
	t.Run("Map should succeed with empty map",
		func(t *testing.T) {
			var (
				mp  = map[string]int{}
				ser = NewMapSer(mus.Serializer[string](nil), mus.Serializer[int](nil))
			)
			test.Test([]map[string]int{mp}, ser, t)
			test.TestSkip([]map[string]int{mp}, ser, t)
		})

	t.Run("Map should succeed with not empty map",
		func(t *testing.T) {
			var (
				mp, keySer, elemSer = test.MapTestData(t)
				ser                 = NewMapSer(keySer, elemSer)
				mocks               = []*mok.Mock{keySer.Mock, elemSer.Mock}
			)
			test.Test([]map[string]int{mp}, ser, t)
			test.TestSkip([]map[string]int{mp}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("We should be able to set a length serializer", func(t *testing.T) {
		var (
			mp, lenSer, keySer, valueSer = test.MapLenTestData(t)
			ser                          = NewMapSer(keySer, valueSer,
				mapops.WithLenSer[string, int](lenSer))
			mocks = []*mok.Mock{lenSer.Mock, keySer.Mock, valueSer.Mock}
		)
		test.Test([]map[string]int{mp}, ser, t)
		test.TestSkip([]map[string]int{mp}, ser, t)

		if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
			t.Error(infomap)
		}
	})

	t.Run("If the length serializer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   map[uint]uint = nil
				wantN                 = 0
				wantErr               = mus.ErrTooSmallByteSlice
				ser                   = NewMapSer(mus.Serializer[uint](nil), mus.Serializer[uint](nil))
			)
			v, n, err := ser.Unmarshal([]byte{})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantV     map[uint]uint = nil
				wantN, bs               = NegativeLengthBs()
				wantErr                 = com.ErrNegativeLength
				ser                     = NewMapSer(mus.Serializer[uint](nil), mus.Serializer[uint](nil))
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If keySer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 3
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 100}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{100}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{100},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				ser   = NewMapSer(keySer, mus.Serializer[uint](nil))
				mocks = []*mok.Mock{keySer.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If valueSer fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 4
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 1, 200, 200}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 1, 1, nil
					},
				)
				valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{200, 200}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{200, 200},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				ser   = NewMapSer(keySer, valueSer)
				mocks = []*mok.Mock{keySer.Mock, valueSer.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If keySer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 100}
				keySer  = mock.NewSerializer[uint]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser   = NewMapSer(keySer, mus.Serializer[uint](nil))
				mocks = []*mok.Mock{keySer.Mock}
			)
			n, err := ser.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If valueSer fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 2
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 1, 200}
				keySer  = mock.NewSerializer[uint]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						return 1, nil
					},
				)
				valueSer = mock.NewSerializer[uint]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						return 0, wantErr
					},
				)
				ser   = NewMapSer(keySer, valueSer)
				mocks = []*mok.Mock{keySer.Mock, valueSer.Mock}
			)
			n, err := ser.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If lenSer fails with an error, SKip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				lenSer  = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser = NewMapSer(mus.Serializer[string](nil), mus.Serializer[int](nil),
					mapops.WithLenSer[string, int](lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			n, err := ser.Skip([]byte{})
			ctestutil.TestSkipResults(wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Skip should return ErrNegativeLength if meets a negative length",
		func(t *testing.T) {
			var (
				wantN, bs = NegativeLengthBs()
				wantErr   = com.ErrNegativeLength
				ser       = NewMapSer(mus.Serializer[uint](nil), mus.Serializer[uint](nil))
			)
			n, err := ser.Skip(bs)
			ctestutil.TestSkipResults(wantN, n, wantErr, err, nil, t)
		})

	t.Run("Valid Map serializer should succeed with empty map",
		func(t *testing.T) {
			var (
				mp  = map[string]int{}
				ser = NewValidMapSer(mus.Serializer[string](nil), mus.Serializer[int](nil),
					nil, nil, nil)
			)
			test.Test([]map[string]int{mp}, ser, t)
			test.TestSkip([]map[string]int{mp}, ser, t)
		})

	t.Run("Valid Map serializer should succeed with not empty map",
		func(t *testing.T) {
			var (
				mp, keySer, elemSer = test.MapTestData(t)
				ser                 = NewValidMapSer(keySer, elemSer, nil, nil, nil)
				mocks               = []*mok.Mock{keySer.Mock, elemSer.Mock}
			)
			test.Test([]map[string]int{mp}, ser, t)
			test.TestSkip([]map[string]int{mp}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

	t.Run("If lenSer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   map[uint]uint = nil
				wantN                 = 0
				wantErr               = mus.ErrTooSmallByteSlice
				lenSer                = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (v int, n int, err error) {
						return 0, 0, wantErr
					},
				)
				ser = NewValidMapSer(mus.Serializer[uint](nil), mus.Serializer[uint](nil),
					mapops.WithLenSer[uint, uint](lenSer))
				mocks = []*mok.Mock{lenSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{})
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("Valid Unmarshal should return ErrNegativeLength if meets negative length",
		func(t *testing.T) {
			var (
				wantV     map[uint]uint = nil
				wantN, bs               = NegativeLengthBs()
				wantErr                 = com.ErrNegativeLength
				ser                     = NewValidMapSer(mus.Serializer[uint](nil),
					mus.Serializer[uint](nil), nil, nil, nil)
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

	t.Run("If keySer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 3
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 100}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{100}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{100},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				ser   = NewValidMapSer(keySer, mus.Serializer[uint](nil), nil, nil, nil)
				mocks = []*mok.Mock{keySer.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If valueSer fails with an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 4
				wantErr = errors.New("Unmarshaller error")
				bs      = []byte{2, 1, 200, 200}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 1, 1, nil
					},
				)
				valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{200, 200}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{200, 200},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				ser   = NewValidMapSer(keySer, valueSer, nil, nil, nil)
				mocks = []*mok.Mock{keySer.Mock, valueSer.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If lenVl returns an error, valid Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   map[uint]uint = nil
				wantN                 = 1
				wantErr               = errors.New("lenVl validator error")
				bs                    = []byte{2, 199, 1, 3, 4}
				lenVl                 = cmock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						if v != 2 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
						}
						return wantErr
					},
				)
				ser = NewValidMapSer(mus.Serializer[uint](nil), mus.Serializer[uint](nil),
					mapops.WithLenValidator[uint, uint](lenVl))
				mocks = []*mok.Mock{lenVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If keyVl returns an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 2
				wantErr = errors.New("key Validator error")
				bs      = []byte{2, 10, 1, 3, 4}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				keyVl = cmock.NewValidator[uint]().RegisterValidate(
					func(v uint) (err error) {
						if v != 10 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
						}
						return wantErr
					},
				)
				ser = NewValidMapSer(keySer, mus.Serializer[uint](nil),
					mapops.WithKeyValidator[uint, uint](keyVl))
				mocks = []*mok.Mock{keySer.Mock, keyVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

	t.Run("If valueVl returns an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 3
				wantErr = errors.New("value Validator error")
				bs      = []byte{2, 10, 11, 3, 4}
				keySer  = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				valueSer = mock.NewSerializer[uint]().RegisterUnmarshal(
					func(bs []byte) (v uint, n int, err error) {
						return 11, 1, nil
					},
				)
				valueVl = cmock.NewValidator[uint]().RegisterValidate(
					func(v uint) (err error) {
						if v != 11 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 11, v)
						}
						return wantErr
					},
				)
				ser = NewValidMapSer(keySer, valueSer,
					mapops.WithValueValidator[uint](valueVl))
				mocks = []*mok.Mock{keySer.Mock, valueSer.Mock, valueVl.Mock}
			)
			v, n, err := ser.Unmarshal(bs)
			ctestutil.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})
}

func NegativeLengthBs() (n int, bs []byte) {
	n = varint.PositiveInt.Size(-1)
	bs = make([]byte, n)
	varint.PositiveInt.Marshal(-1, bs)
	return
}
