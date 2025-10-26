package ord

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	com_mock "github.com/mus-format/common-go/testdata/mock"
	"github.com/mus-format/mus-go"
	arrops "github.com/mus-format/mus-go/options/array"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	mapops "github.com/mus-format/mus-go/options/map"
	slops "github.com/mus-format/mus-go/options/slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/testdata"
	mock "github.com/mus-format/mus-go/testdata/mock"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		t.Run("Bool serializer should work correctly",
			func(t *testing.T) {
				ser := Bool
				testdata.Test[bool](com_testdata.BoolTestCases, ser, t)
				testdata.TestSkip[bool](com_testdata.BoolTestCases, ser, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = Bool.Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = com.ErrWrongFormat
					bs      = []byte{3}
					n, err  = Bool.Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("String serializer should work correctly",
			func(t *testing.T) {
				ser := String
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
			})

		t.Run("We should be able to set a length serializer",
			func(t *testing.T) {
				var (
					str, lenSer = testdata.StringLenSerData(t)
					ser         = NewStringSer(strops.WithLenSer(lenSer))
				)
				testdata.Test[string]([]string{str}, ser, t)
				testdata.TestSkip[string]([]string{str}, ser, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				v, n, err := String.Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					n, err = NewStringSer(strops.WithLenSer(lenSer)).Skip(nil)
					mocks  = []*mok.Mock{lenSer.Mock}
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				n, err := String.Skip(bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{2, 2}
					n, err  = String.Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid string serializer should work correctly",
			func(t *testing.T) {
				ser := NewValidStringSer(nil)
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
			})

		t.Run("Valid string serializer with varint length should work correctly",
			func(t *testing.T) {
				ser := NewValidStringSer(nil)
				testdata.Test[string](com_testdata.StringTestCases, ser, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, ser, t)
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
					v, n, err = NewValidStringSer(strops.WithLenSer(lenSer)).Unmarshal(nil)
					mocks     = []*mok.Mock{lenSer.Mock}
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				v, n, err := NewValidStringSer(nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("Valid Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{2, 2}
					v, n, err = NewValidStringSer(nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("If lenVl fails with an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("lenVl validator error")
					bs      = []byte{2, 2, 2}
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
				v, n, err = NewValidStringSer(strops.WithLenValidator(lenVl)).Unmarshal(bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})
	})

	t.Run("pointer", func(t *testing.T) {
		t.Run("Pointer seralizer should work correctly",
			func(t *testing.T) {
				var (
					ptr, baseSer = testdata.PtrSerData(t)
					ser          = NewPtrSer[int](baseSer)
				)
				testdata.Test[*int]([]*int{ptr}, ser, t)
				testdata.TestSkip[*int]([]*int{ptr}, ser, t)
			})

		t.Run("Pointer serializer should work correctly for nil ptr",
			func(t *testing.T) {
				ser := NewPtrSer[string](nil)
				testdata.Test[*string]([]*string{nil}, ser, t)
				testdata.TestSkip[*string]([]*string{nil}, ser, t)
			})

		t.Run("Pointer serializer should work correctly for not nil ptr",
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
					ser = NewPtrSer[string](strSer)
				)
				testdata.Test[*string]([]*string{ptr}, ser, t)
				testdata.TestSkip[*string]([]*string{ptr}, ser, t)
			})

		t.Run("Unmarshal should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     *string = nil
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					v, n, err         = NewPtrSer[string](nil).Unmarshal([]byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV     *string = nil
					wantN             = 0
					wantErr           = com.ErrWrongFormat
					v, n, err         = NewPtrSer[string](nil).Unmarshal([]byte{2})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					mocks     = []*mok.Mock{baseSer.Mock}
					v, n, err = NewPtrSer[string](baseSer).Unmarshal([]byte{byte(com.NotNil)})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = NewPtrSer[string](nil).Skip([]byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = com.ErrWrongFormat
					n, err  = NewPtrSer[string](nil).Skip([]byte{2})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
					mocks  = []*mok.Mock{baseSer.Mock}
					n, err = NewPtrSer[string](baseSer).Skip([]byte{byte(com.NotNil)})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})
	})

	t.Run("byte_slice", func(t *testing.T) {
		t.Run("ByteSlice serializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{}
					ser = ByteSlice
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("ByteSlice serializer should work correctly for non-empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{0, 1, 1, 255, 100, 0, 1, 10}
					ser = ByteSlice
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("We should be able to set a length serializer", func(t *testing.T) {
			var (
				sl, lenSer = testdata.ByteSliceLenSerData(t)
				ser        = NewByteSliceSer(bslops.WithLenSer(lenSer))
			)
			testdata.Test[[]byte]([][]byte{sl}, ser, t)
			testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If the length serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 0
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = ByteSlice.Unmarshal(nil)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
				)
				v, n, err := ByteSlice.Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Skip should return ErrTooSmallByteSlice if bs is too small",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{4, 1}
					n, err  = ByteSlice.Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
					n, err = NewByteSliceSer(bslops.WithLenSer(lenSer)).Skip(nil)
					mocks  = []*mok.Mock{lenSer.Mock}
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				n, err := ByteSlice.Skip(bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly for empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{}
					ser = NewValidByteSliceSer(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
			})

		t.Run("Valid ByteSlice serializer should work correctly for non-empty slice",
			func(t *testing.T) {
				var (
					sl  = []byte{0, 1, 1, 255, 100, 0, 1, 10}
					ser = NewValidByteSliceSer(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, ser, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, ser, t)
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
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					v, n, err = NewValidByteSliceSer(bslops.WithLenSer(lenSer)).Unmarshal(nil)
					mocks     = []*mok.Mock{lenSer.Mock}
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = NegativeLengthBs()
					wantErr          = com.ErrNegativeLength
				)
				v, n, err := NewValidByteSliceSer(nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 1
					wantErr        = errors.New("too large slice")
					bs             = []byte{2, 4, 1}
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					v, n, err = NewValidByteSliceSer(bslops.WithLenValidator(lenVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})
	})

	t.Run("array", func(t *testing.T) {
		t.Run("Array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testdata.ArraySerData(t)
				ser          = NewArraySer[[3]int, int](elemSer)
			)
			testdata.Test[[3]int]([][3]int{arr}, ser, t)
			testdata.TestSkip[[3]int]([][3]int{arr}, ser, t)
		})

		t.Run("Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV   [3]int = [3]int{0, 0, 0}
					wantN          = 1
					wantErr        = com.ErrTooLargeLength
					bs             = []byte{4, 0, 0}
				)
				v, n, err := NewArraySer[[3]int, int](nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("Valid array serializer should work correctly", func(t *testing.T) {
			var (
				arr, elemSer = testdata.ArraySerData(t)
				ser          = NewValidArraySer[[3]int, int](elemSer, nil)
			)
			testdata.Test[[3]int]([][3]int{arr}, ser, t)
			testdata.TestSkip[[3]int]([][3]int{arr}, ser, t)
		})

		t.Run("Valid Unmarshal of the too large array should return ErrTooLargeLength",
			func(t *testing.T) {
				var (
					wantV   [3]int = [3]int{0, 0, 0}
					wantN          = 1
					wantErr        = com.ErrTooLargeLength
					bs             = []byte{4, 0, 0}
				)
				v, n, err := NewValidArraySer[[3]int, int](nil, nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("If elemVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV    [3]int = [3]int{0, 0, 0}
					wantElem        = 11
					wantN           = 2
					wantErr         = errors.New("elemVl error")
					arr             = []byte{1, 2, 3}
					elemSer         = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (v int, n int, err error) {
							return wantElem, 1, nil
						},
					)
					elemVl = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != wantElem {
								t.Errorf("unexpected v, want '%v' actual '%v'", wantElem, v)
							}
							return wantErr
						},
					)
					// ser   = NewValidArraySer[[3]int, int](3, elemSer, elemVl)
					ser = NewValidArraySer[[3]int, int](elemSer,
						arrops.WithElemValidator[int](elemVl))
					mocks = []*mok.Mock{elemSer.Mock, elemVl.Mock}
				)
				v, n, err := ser.Unmarshal(arr)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err,
					mocks, t)
			})
	})

	t.Run("slice", func(t *testing.T) {
		t.Run("Slice serializer should work correctly with empty slice",
			func(t *testing.T) {
				var (
					sl  = []string{}
					ser = NewSliceSer[string](nil)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)
			})

		t.Run("Slice serializer should work correctly with not empty slice",
			func(t *testing.T) {
				var (
					sl, elemSer = testdata.SliceSerData(t)
					ser         = NewSliceSer[string](elemSer)
					mocks       = []*mok.Mock{elemSer.Mock}
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("We should be able to set a length serializer", func(t *testing.T) {
			var (
				sl, lenSer, elemSer = testdata.SliceLenSerData(t)
				ser                 = NewSliceSer[string](elemSer,
					slops.WithLenSer[string](lenSer))
				mocks = []*mok.Mock{lenSer.Mock, elemSer.Mock}
			)
			testdata.Test[[]string]([][]string{sl}, ser, t)
			testdata.TestSkip[[]string]([][]string{sl}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("If the length serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN              = 0
					wantErr            = mus.ErrTooSmallByteSlice
					v, n, err          = NewSliceSer[string](nil).Unmarshal([]byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN, bs          = NegativeLengthBs()
					wantErr            = com.ErrNegativeLength
				)
				v, n, err := NewSliceSer[string](nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					mocks     = []*mok.Mock{elemSer.Mock}
					v, n, err = NewSliceSer[uint](elemSer).Unmarshal([]byte{1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					mocks = []*mok.Mock{lenSer.Mock}
					// n, err = NewSliceSerWith[string](lenSer, nil).Skip([]byte{})
					n, err = NewSliceSer[string](nil, slops.WithLenSer[string](lenSer)).Skip([]byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				n, err := NewSliceSer[string](nil).Skip(bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
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
					mocks  = []*mok.Mock{elemSer.Mock}
					n, err = NewSliceSer[uint](elemSer).Skip([]byte{1})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Valid Slice serializer should work correctly with empty slice",
			func(t *testing.T) {
				var (
					sl  = []string{}
					ser = NewValidSliceSer[string](nil, nil, nil)
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)
			})

		t.Run("Valid Slice serializer should work correctly with not empty slice",
			func(t *testing.T) {
				var (
					sl, elemSer = testdata.SliceSerData(t)
					ser         = NewValidSliceSer[string](elemSer, nil, nil)
					mocks       = []*mok.Mock{elemSer.Mock}
				)
				testdata.Test[[]string]([][]string{sl}, ser, t)
				testdata.TestSkip[[]string]([][]string{sl}, ser, t)

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
					mocks = []*mok.Mock{lenSer.Mock}
					// v, n, err = NewValidSliceSerWith[string](lenSer, nil, nil, nil).Unmarshal([]byte{})
					v, n, err = NewValidSliceSer[string](nil, slops.WithLenSer[string](lenSer)).Unmarshal([]byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN, bs          = NegativeLengthBs()
					wantErr            = com.ErrNegativeLength
				)
				v, n, err := NewValidSliceSer[string](nil, nil, nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					mocks     = []*mok.Mock{elemSer.Mock}
					v, n, err = NewValidSliceSer[uint](elemSer, nil, nil).Unmarshal([]byte{1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = nil
					wantN          = 1
					wantErr        = errors.New("lenVl error")
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					bs        = []byte{3, 10, 2, 3}
					mocks     = []*mok.Mock{lenVl.Mock}
					v, n, err = NewValidSliceSer[uint](nil, slops.WithLenValidator[uint](lenVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If elemVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = []uint{10, 0, 0}
					wantN          = 3
					wantErr        = errors.New("elemVl error")
					bs             = []byte{3, 10, 2, 3}
					elemVl         = com_mock.NewValidator[uint]().RegisterValidate(
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
					mocks     = []*mok.Mock{elemSer.Mock, elemVl.Mock}
					v, n, err = NewValidSliceSer[uint](elemSer, slops.WithElemValidator[uint](elemVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})
	})

	t.Run("map", func(t *testing.T) {
		t.Run("Map should work correctly with empty map",
			func(t *testing.T) {
				var (
					mp  = map[string]int{}
					ser = NewMapSer[string, int](nil, nil)
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)
			})

		t.Run("Map should work correctly with not empty map",
			func(t *testing.T) {
				var (
					mp, keySer, elemSer = testdata.MapSerData(t)
					ser                 = NewMapSer[string, int](keySer, elemSer)
					mocks               = []*mok.Mock{keySer.Mock, elemSer.Mock}
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)

				if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
					t.Error(infomap)
				}
			})

		t.Run("We should be able to set a length serializer", func(t *testing.T) {
			var (
				mp, lenSer, keySer, valueSer = testdata.MapLenSerData(t)
				ser                          = NewMapSer[string, int](keySer, valueSer,
					mapops.WithLenSer[string, int](lenSer))
				mocks = []*mok.Mock{lenSer.Mock, keySer.Mock, valueSer.Mock}
			)
			testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
			testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)

			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})

		t.Run("If the length serializer fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 0
					wantErr                 = mus.ErrTooSmallByteSlice
					v, n, err               = NewMapSer[uint, uint](nil, nil).Unmarshal([]byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN, bs               = NegativeLengthBs()
					wantErr                 = com.ErrNegativeLength
				)
				v, n, err := NewMapSer[uint, uint](nil, nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					mocks     = []*mok.Mock{keySer.Mock}
					v, n, err = NewMapSer[uint, uint](keySer, nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					mocks     = []*mok.Mock{keySer.Mock, valueSer.Mock}
					v, n, err = NewMapSer[uint, uint](keySer, valueSer).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					mocks  = []*mok.Mock{keySer.Mock}
					n, err = NewMapSer[uint, uint](keySer, nil).Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
					mocks  = []*mok.Mock{keySer.Mock, valueSer.Mock}
					n, err = NewMapSer[uint, uint](keySer, valueSer).Skip(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
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
					mocks  = []*mok.Mock{lenSer.Mock}
					n, err = NewMapSer[uint, uint](nil, nil, mapops.WithLenSer[uint, uint](lenSer)).Skip([]byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = NegativeLengthBs()
					wantErr   = com.ErrNegativeLength
				)
				n, err := NewMapSer[string, string](nil, nil).Skip(bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("Valid Map serializer should work correctly with empty map",
			func(t *testing.T) {
				var (
					mp  = map[string]int{}
					ser = NewValidMapSer[string, int](nil, nil, nil, nil, nil)
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)
			})

		t.Run("Valid Map serializer should work correctly with not empty map",
			func(t *testing.T) {
				var (
					mp, keySer, elemSer = testdata.MapSerData(t)
					ser                 = NewValidMapSer[string, int](keySer, elemSer, nil, nil, nil)
					mocks               = []*mok.Mock{keySer.Mock, elemSer.Mock}
				)
				testdata.Test[map[string]int]([]map[string]int{mp}, ser, t)
				testdata.TestSkip[map[string]int]([]map[string]int{mp}, ser, t)

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
					mocks     = []*mok.Mock{lenSer.Mock}
					v, n, err = NewValidMapSer[uint, uint](nil, nil, mapops.WithLenSer[uint, uint](lenSer)).Unmarshal([]byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Valid Unmarshal should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN, bs               = NegativeLengthBs()
					wantErr                 = com.ErrNegativeLength
				)
				v, n, err := NewValidMapSer[uint, uint](nil, nil, nil, nil, nil).Unmarshal(bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
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
					mocks     = []*mok.Mock{keySer.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, nil, nil, nil, nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					mocks     = []*mok.Mock{keySer.Mock, valueSer.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, valueSer, nil, nil, nil).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl returns an error, valid Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("lenVl validator error")
					bs                    = []byte{2, 199, 1, 3, 4}
					lenVl                 = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](nil, nil, mapops.WithLenValidator[uint, uint](lenVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					keyVl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{keySer.Mock, keyVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, nil, mapops.WithKeyValidator[uint, uint](keyVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
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
					valueVl = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 11 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 11, v)
							}
							return wantErr
						},
					)
					mocks     = []*mok.Mock{keySer.Mock, valueSer.Mock, valueVl.Mock}
					v, n, err = NewValidMapSer[uint, uint](keySer, valueSer, mapops.WithValueValidator[uint, uint](valueVl)).Unmarshal(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})
	})
}

func NegativeLengthBs() (n int, bs []byte) {
	n = varint.PositiveInt.Size(-1)
	bs = make([]byte, n)
	varint.PositiveInt.Marshal(-1, bs)
	return
}
