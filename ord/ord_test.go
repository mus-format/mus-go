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
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/testdata/mock"
	"github.com/mus-format/mus-go/varint"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {

	t.Run("bool", func(t *testing.T) {

		t.Run("All MarshalBool, UnmarshalBool, SizeBool, SkipBool functions must work correctly",
			func(t *testing.T) {
				var (
					m  = mus.MarshallerFn[bool](MarshalBool)
					u  = mus.UnmarshallerFn[bool](UnmarshalBool)
					s  = mus.SizerFn[bool](SizeBool)
					sk = mus.SkipperFn(SkipBool)
				)
				testdata.Test[bool](com_testdata.BoolTestCases, m, u, s, t)
				testdata.TestSkip[bool](com_testdata.BoolTestCases, m, sk, s, t)
			})

		t.Run("UnmarshalBool should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = false
					wantN     = 0
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{}
					v, n, err = UnmarshalBool(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV     = false
					wantN     = 0
					wantErr   = com.ErrWrongFormat
					bs        = []byte{3}
					v, n, err = UnmarshalBool(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipBool should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = SkipBool(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipBool should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = com.ErrWrongFormat
					bs      = []byte{3}
					n, err  = SkipBool(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("string", func(t *testing.T) {

		t.Run("All MarshalStr, UnmarshalStr, SizeStr, SkipStr function must work correctly",
			func(t *testing.T) {
				var (
					m  mus.MarshallerFn[string]   = MarshalStr
					u  mus.UnmarshallerFn[string] = UnmarshalStr
					s  mus.SizerFn[string]        = SizeStr
					sk mus.SkipperFn              = SkipStr
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("All MarshalString, UnmarshalString, SizeString, SkipString function with default lenM, lenU, lenS must work correctly",
			func(t *testing.T) {
				var (
					m  mus.Marshaller[string]   = NewStringMarshallerFn(nil)
					u  mus.Unmarshaller[string] = NewStringUnmarshallerFn(nil)
					s  mus.Sizer[string]        = NewStringSizerFn(nil)
					sk mus.Skipper              = NewStringSkipperFn(nil)
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("All MarshalStringVarint, UnmarshalVarintString, SizeStringVarint, SkipStringVarint functions must work correctly",
			func(t *testing.T) {
				var (
					m mus.Marshaller[string] = NewStringMarshallerFn(
						mus.MarshallerFn[int](varint.MarshalInt))
					u mus.Unmarshaller[string] = NewStringUnmarshallerFn(
						mus.UnmarshallerFn[int](varint.UnmarshalInt))
					s mus.Sizer[string] = NewStringSizerFn(
						mus.SizerFn[int](varint.SizeInt))
					sk mus.Skipper = NewStringSkipperFn(mus.UnmarshallerFn[int](varint.UnmarshalInt))
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("MarshalString should panic with ErrTooSmallByteSlice if there is no space in bs",
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
				MarshalString("hello world", nil, make([]byte, 2))
			})

		t.Run("UnmarshalString should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN, bs = MakeBsWithNegativeLength()
					wantErr   = com.ErrNegativeLength
				)
				v, n, err := UnmarshalString(nil, bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalString should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{2, 2}
					v, n, err = UnmarshalString(nil, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalString should return ErrTooSmallByteSlice if meets invalid length", func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 2
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{200, 200}
				v, n, err = UnmarshalString(nil, bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("If skip == true and lenVl validator returns an error, UnmarshalValidString should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 3
					wantErr = errors.New("lenVl validator error")
					bs      = []byte{2, 2, 2}
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					v, n, err = UnmarshalValidString(nil, lenVl, true, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If skip == false and lenVl validator returns an error, UnmarshalStringValid should return an error",
			func(t *testing.T) {
				var (
					wantV   = ""
					wantN   = 1
					wantErr = errors.New("lenVl validator error")
					bs      = []byte{2, 2, 2}
					lenVl   = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					v, n, err = UnmarshalValidString(nil, lenVl, false, bs)
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
				v, n, err = UnmarshalValidString(nil, lenVl, false, bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("SkipString should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN, bs = MakeBsWithNegativeLength()
					wantErr   = com.ErrNegativeLength
				)
				n, err := SkipString(nil, bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipString should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{2, 2}
					n, err  = SkipString(nil, bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("If SkipString should return an error if it fails to unmarshal a length",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = SkipString(nil, bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("pointer", func(t *testing.T) {

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly for nil ptr",
			func(t *testing.T) {
				var (
					m  mus.MarshallerFn[*string]   = NewPtrMarshallerFn[string](nil)
					u  mus.UnmarshallerFn[*string] = NewPtrUnmarshallerFn[string](nil)
					s  mus.SizerFn[*string]        = NewPtrSizerFn[string](nil)
					sk mus.SkipperFn               = NewPtrSkipperFn(nil)
				)
				testdata.Test[*string]([]*string{nil}, m, u, s, t)
				testdata.TestSkip[*string]([]*string{nil}, m, sk, s, t)
			})

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly not nil ptr",
			func(t *testing.T) {
				var (
					str1                            = "one"
					str1Raw                         = append([]byte{6}, []byte(str1)...)
					ptr                             = &str1
					m1      mock.Marshaller[string] = mock.NewMarshaller[string]().RegisterNMarshal(2,
						func(v string, bs []byte) (n int) {
							switch v {
							case str1:
								return copy(bs, str1Raw)
							default:
								t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
								return
							}
						},
					)
					u1 mock.Unmarshaller[string] = mock.NewUnmarshaller[string]().RegisterNUnmarshal(1,
						func(bs []byte) (v string, n int, err error) {
							if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
								return str1, len(str1Raw), nil
							} else {
								t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
								return
							}
						},
					)
					s1 mock.Sizer[string] = mock.NewSizer[string]().RegisterNSize(2,
						func(v string) (size int) {
							switch v {
							case str1:
								return len(str1Raw)
							default:
								t.Fatalf("unexepcted string, want '%v' actual '%v'", str1, v)
								return
							}
						},
					)
					sk1 mock.Skipper = mock.NewSkipper().RegisterNSkip(1,
						func(bs []byte) (n int, err error) {
							if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
								return len(str1Raw), nil
							} else {
								t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
								return
							}
						},
					)
					m  mus.MarshallerFn[*string]   = NewPtrMarshallerFn[string](m1)
					u  mus.UnmarshallerFn[*string] = NewPtrUnmarshallerFn[string](u1)
					s  mus.SizerFn[*string]        = NewPtrSizerFn[string](s1)
					sk mus.SkipperFn               = NewPtrSkipperFn(sk1)
				)
				testdata.Test[*string]([]*string{ptr}, m, u, s, t)
				testdata.TestSkip[*string]([]*string{ptr}, m, sk, s, t)
			})

		t.Run("UnmarshalPtr should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     *string = nil
					wantN             = 0
					wantErr           = mus.ErrTooSmallByteSlice
					v, n, err         = UnmarshalPtr[string](nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalPtr should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantV     *string = nil
					wantN             = 0
					wantErr           = com.ErrWrongFormat
					v, n, err         = UnmarshalPtr[string](nil, []byte{2})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Unmarshaller fails with an error, UnmarshalPtr should return it",
			func(t *testing.T) {
				var (
					wantV   *string = nil
					wantN           = 5
					wantErr         = errors.New("Unmarshaller error")
					u               = mock.NewUnmarshaller[string]().RegisterUnmarshal(
						func(bs []byte) (v string, n int, err error) {
							return "", 4, wantErr
						},
					)
					v, n, err = UnmarshalPtr[string](u, []byte{byte(com.NotNil)})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipPtr should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = SkipPtr(nil, []byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipPtr should return ErrWrongFormat if meets wrong format",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = com.ErrWrongFormat
					n, err  = SkipPtr(nil, []byte{2})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Skipper fails with an error, SkipPtr should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("error")
					s       = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							return 2, wantErr
						},
					)
					n, err = SkipPtr(s, []byte{byte(com.NotNil)})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("byte_slice", func(t *testing.T) {

		t.Run("All MarshalByteSlice, UnmarshalByteSlice, SizeByteSlice, SkipByteSlice functions with default lenM, lenU, lenS and empty slice must work correctly",
			func(t *testing.T) {
				var (
					sl                            = []byte{}
					m  mus.MarshallerFn[[]byte]   = NewByteSliceMarshallerFn(nil)
					u  mus.UnmarshallerFn[[]byte] = NewByteSliceUnmarshallerFn(nil)
					s  mus.SizerFn[[]byte]        = NewByteSliceSizerFn(nil)
					sk mus.SkipperFn              = NewByteSliceSkipperFn(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, m, u, s, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, m, sk, s, t)
			})

		t.Run("All MarshalByteSlice, UnmarshalByteSlice, SizeByteSlice, SkipByteSlice functions with default lenM, lenU, lenS and not empty slice must work correctly",
			func(t *testing.T) {
				var (
					sl                            = []byte{0, 1, 1, 255, 100, 0, 1, 10}
					m  mus.MarshallerFn[[]byte]   = NewByteSliceMarshallerFn(nil)
					u  mus.UnmarshallerFn[[]byte] = NewByteSliceUnmarshallerFn(nil)
					s  mus.SizerFn[[]byte]        = NewByteSliceSizerFn(nil)
					sk mus.SkipperFn              = NewByteSliceSkipperFn(nil)
				)
				testdata.Test[[]byte]([][]byte{sl}, m, u, s, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, m, sk, s, t)
			})

		t.Run("All MarshalVarintByteSlice, UnmarshalVarintByteSlice, SizeVarintByteSlice, SkipVarintByteSlice functions must work correctly with empty slice",
			func(t *testing.T) {
				var (
					sl                            = []byte{}
					m  mus.MarshallerFn[[]byte]   = NewByteSliceMarshallerFn(mus.MarshallerFn[int](varint.MarshalInt))
					u  mus.UnmarshallerFn[[]byte] = NewByteSliceUnmarshallerFn(mus.UnmarshallerFn[int](varint.UnmarshalInt))
					s  mus.SizerFn[[]byte]        = NewByteSliceSizerFn(mus.SizerFn[int](varint.SizeInt))
					sk mus.SkipperFn              = NewByteSliceSkipperFn(mus.UnmarshallerFn[int](varint.UnmarshalInt))
				)
				testdata.Test[[]byte]([][]byte{sl}, m, u, s, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, m, sk, s, t)
			})

		t.Run("All MarshalVarintByteSlice, UnmarshalVarintByteSlice, SizeVarintByteSlice, SkipVarintByteSlice functions must work correctly with not empty slice",
			func(t *testing.T) {
				var (
					sl                            = []byte{0, 1, 1, 255, 100, 0, 1, 10}
					m  mus.MarshallerFn[[]byte]   = NewByteSliceMarshallerFn(mus.MarshallerFn[int](varint.MarshalInt))
					u  mus.UnmarshallerFn[[]byte] = NewByteSliceUnmarshallerFn(mus.UnmarshallerFn[int](varint.UnmarshalInt))
					s  mus.SizerFn[[]byte]        = NewByteSliceSizerFn(mus.SizerFn[int](varint.SizeInt))
					sk mus.SkipperFn              = NewByteSliceSkipperFn(mus.UnmarshallerFn[int](varint.UnmarshalInt))
				)

				testdata.Test[[]byte]([][]byte{sl}, m, u, s, t)
				testdata.TestSkip[[]byte]([][]byte{sl}, m, sk, s, t)
			})

		t.Run("MarshalByteSlice should panic with ErrTooSmallByteSlice if there is no space in bs",
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
				MarshalByteSlice([]byte{1, 2, 3, 4}, nil, make([]byte, 2))
			})

		t.Run("UnmarshalByteSlice should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 0
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = UnmarshalByteSlice(nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalByteSlice should return ErrTooSmallByteSlice if bs is too small",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN            = 1
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = UnmarshalByteSlice(nil, []byte{4, 1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalByteSlice should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []byte = nil
					wantN, bs        = MakeBsWithNegativeLength()
					wantErr          = com.ErrNegativeLength
				)
				v, n, err := UnmarshalByteSlice(nil, bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If lenVl validator returns an error, UnmarshalValidByteSlice should return it",
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
					v, n, err = UnmarshalValidByteSlice(nil, lenVl, false, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("If skip == true and lenVl validator returns an error, UnmarshalValidByteSlice should return it and skip the rest of the slice",
			func(t *testing.T) {
				var (
					wantV   []byte = nil
					wantN          = 3
					wantErr        = errors.New("too large slice")
					bs             = []byte{2, 4, 1}
					lenVl          = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					v, n, err = UnmarshalValidByteSlice(nil, lenVl, true, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil,
					t)
			})

		t.Run("SkipByteSlice should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = SkipByteSlice(nil, []byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipByteSlice should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = MakeBsWithNegativeLength()
					wantErr   = com.ErrNegativeLength
				)
				n, err := SkipByteSlice(nil, bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("slice", func(t *testing.T) {

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice function with default lenM, lenU, lenS and empty slice must work correctly",
			func(t *testing.T) {
				var (
					sl                            = []string{}
					m  mus.MarshallerFn[[]string] = NewSliceMarshallerFn[string](nil,
						mus.MarshallerFn[string](MarshalStr))
					u mus.UnmarshallerFn[[]string] = NewSliceUnmarshallerFn[string](nil,
						mus.UnmarshallerFn[string](UnmarshalStr))
					s mus.SizerFn[[]string] = NewSliceSizerFn[string](nil,
						mus.SizerFn[string](SizeStr))
					sk mus.SkipperFn = NewSliceSkipperFn(nil, mus.SkipperFn(SkipStr))
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice function with default lenM, lenU, lenS and not empty slice must work correctly",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 mus.Marshaller[string]) mus.MarshallerFn[[]string] {
						return NewSliceMarshallerFn(nil, m1)
					},
					func(u1 mus.Unmarshaller[string]) mus.UnmarshallerFn[[]string] {
						return NewSliceUnmarshallerFn(nil, u1)
					},
					func(s1 mus.Sizer[string]) mus.SizerFn[[]string] {
						return NewSliceSizerFn(nil, s1)
					},
					func(sk1 mus.Skipper) mus.SkipperFn {
						return NewSliceSkipperFn(nil, sk1)
					}, t)
			})

		t.Run("All MarshalVarintSlice, UnmarshalVarintSlice, SizeVarintSlice, SkipVarintSlice functions must work correctly with empty slice",
			func(t *testing.T) {
				var (
					sl                            = []string{}
					m  mus.MarshallerFn[[]string] = NewSliceMarshallerFn[string](
						mus.MarshallerFn[int](varint.MarshalInt),
						mus.MarshallerFn[string](MarshalStr),
					)
					u mus.UnmarshallerFn[[]string] = NewSliceUnmarshallerFn[string](
						mus.UnmarshallerFn[int](varint.UnmarshalInt),
						mus.UnmarshallerFn[string](UnmarshalStr),
					)
					s mus.SizerFn[[]string] = NewSliceSizerFn[string](
						mus.SizerFn[int](varint.SizeInt),
						mus.SizerFn[string](SizeStr),
					)
					sk mus.SkipperFn = NewSliceSkipperFn(
						mus.UnmarshallerFn[int](varint.UnmarshalInt),
						mus.SkipperFn(SkipStr),
					)
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalVarintSlice, UnmarshalVarintSlice, SizeVarintSlice, SkipVarintSlice functions must work correctly with not empty slice",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 mus.Marshaller[string]) mus.MarshallerFn[[]string] {
						return NewSliceMarshallerFn(mus.MarshallerFn[int](varint.MarshalInt), m1)
					},
					func(u1 mus.Unmarshaller[string]) mus.UnmarshallerFn[[]string] {
						return NewSliceUnmarshallerFn(mus.UnmarshallerFn[int](varint.UnmarshalInt), u1)
					},
					func(s1 mus.Sizer[string]) mus.SizerFn[[]string] {
						return NewSliceSizerFn(mus.SizerFn[int](varint.SizeInt), s1)
					},
					func(sk1 mus.Skipper) mus.SkipperFn {
						return NewSliceSkipperFn(mus.UnmarshallerFn[int](varint.UnmarshalInt), sk1)
					}, t)
			})

		t.Run("UnmarshalSlice should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN              = 0
					wantErr            = mus.ErrTooSmallByteSlice
					v, n, err          = UnmarshalSlice[string](nil, nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalSlice should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []string = nil
					wantN, bs          = MakeBsWithNegativeLength()
					wantErr            = com.ErrNegativeLength
				)
				v, n, err := UnmarshalSlice[string](nil, nil, bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Unmarshaller fails with an error, UnmarshalSlice should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = make([]uint, 1)
					wantN          = 3
					wantErr        = errors.New("Unmarshaller error")
					u              = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u.Mock}
					v, n, err = UnmarshalSlice[uint](nil, u, []byte{1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Skipper != nil and lenVl validator returns an error, UnmarshalValidSlice should return it",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 5
					wantErr          = errors.New("lenVl validator error")
					bs               = []byte{4, 4, 1, 1, 0}
					maxLength        = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 4 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					sk = mock.NewSkipper().RegisterNSkip(4,
						func(bs []byte) (n int, err error) { return 1, nil },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSlice[uint](nil, maxLength, nil, nil, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error and Skipper returns an error, UnmarshalValidSlice should return the last one",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 4
					wantErr          = errors.New("skip rest error")
					bs               = []byte{3, 4, 1, 1}
					maxLength        = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					sk = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) { return 3, wantErr },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSlice[uint](nil, maxLength, nil, nil, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Skipper == nil and lenVl validator returns an error, UnmarshalValidSlice should return it",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 1
					wantErr          = errors.New("lenVl Validator error")
					maxLength        = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					bs        = []byte{3, 10, 2, 3}
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidSlice[uint](nil, maxLength, nil, nil, nil, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Validator returns an error, UnmarshalVarintSlice should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = []uint{10, 0, 0}
					wantN          = 4
					wantErr        = errors.New("Validator error")
					bs             = []byte{3, 10, 2, 3}
					vl             = com_mock.NewValidator[uint]().RegisterValidate(
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
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					).RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 2, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3}, bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSlice[uint](nil, nil, u, vl, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Validator returns an error and Skipper returns an error, UnmarshalValidSlice should return the last one",
			func(t *testing.T) {
				var (
					wantV   = []uint{0, 0, 0}
					wantN   = 4
					wantErr = errors.New("skip rest error")
					bs      = []byte{3, 10, 2, 3}
					vl      = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return errors.New("validator error")
						},
					)
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{2, 3}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{2, 3},
									bs)
							}
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSlice[uint](nil, nil, u, vl, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("SkipSlice should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = SkipSlice(nil, nil, []byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipSlice should return ErrNegativeLength if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN, bs = MakeBsWithNegativeLength()
					wantErr   = com.ErrNegativeLength
				)
				n, err := SkipSlice(nil, nil, bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("map", func(t *testing.T) {

		t.Run("All MarshalMap, UnmarshalMap, SizeMap, SkipMap functions with default lenM, lenU, lenS, lenSk must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 mus.Marshaller[string], m2 mus.Marshaller[uint]) mus.MarshallerFn[map[string]uint] {
						return NewMapMarshallerFn(nil, m1, m2)
					},
					func(u1 mus.Unmarshaller[string], u2 mus.Unmarshaller[uint]) mus.UnmarshallerFn[map[string]uint] {
						return NewMapUnmarshallerFn(nil, u1, u2)
					},
					func(s1 mus.Sizer[string], s2 mus.Sizer[uint]) mus.SizerFn[map[string]uint] {
						return NewMapSizerFn(nil, s1, s2)
					},
					func(sk1, sk2 mus.Skipper) mus.SkipperFn {
						return NewMapSkipperFn[string, uint](nil, sk1, sk2)
					}, t)
			})

		t.Run("All MarshalVarintMap, UnmarshalVarintMap, SizeVarintMap, SkipVarintMap functions must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 mus.Marshaller[string], m2 mus.Marshaller[uint]) mus.MarshallerFn[map[string]uint] {
						return NewMapMarshallerFn(mus.MarshallerFn[int](varint.MarshalInt), m1, m2)
					},
					func(u1 mus.Unmarshaller[string], u2 mus.Unmarshaller[uint]) mus.UnmarshallerFn[map[string]uint] {
						return NewMapUnmarshallerFn(mus.UnmarshallerFn[int](varint.UnmarshalInt), u1, u2)
					},
					func(s1 mus.Sizer[string], s2 mus.Sizer[uint]) mus.SizerFn[map[string]uint] {
						return NewMapSizerFn(mus.SizerFn[int](varint.SizeInt), s1, s2)
					},
					func(sk1, sk2 mus.Skipper) mus.SkipperFn {
						return NewMapSkipperFn[string, uint](mus.UnmarshallerFn[int](varint.UnmarshalInt), sk1, sk2)
					}, t)
			})

		t.Run("UnmarshalMap should return ErrTooSmallByteSlice if there no space in bs",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 0
					wantErr                 = mus.ErrTooSmallByteSlice
					v, n, err               = UnmarshalMap[uint, uint](nil, nil, nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalMap should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN, bs               = MakeBsWithNegativeLength()
					wantErr                 = com.ErrNegativeLength
				)
				v, n, err := UnmarshalMap[uint, uint](nil, nil, nil, bs)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Key Unmarshaller fails with an error, UnmarshalMap should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 3
					wantErr = errors.New("Unmarshaller error")
					bs      = []byte{2, 100}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							if !reflect.DeepEqual(bs, []byte{100}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{100},
									bs)
							}
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock}
					v, n, err = UnmarshalMap[uint, uint](nil, u1, nil, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Value Unmarshaller fails with an error, UnmarshalMap should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 4
					wantErr = errors.New("Unmarshaller error")
					bs      = []byte{2, 1, 200, 200}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 1, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							if !reflect.DeepEqual(bs, []byte{200, 200}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{200, 200},
									bs)
							}
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock, u2.Mock}
					v, n, err = UnmarshalMap[uint, uint](nil, u1, u2, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error, UnmarshalValidMap should return it",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 5
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
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{199, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{199, 1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4},
									bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, lenVl, nil, nil,
						nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error and Key Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 2
					wantErr               = errors.New("skip key error")
					bs                    = []byte{4, 199, 1, 3, 4}
					lenVl                 = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 4 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 4, v)
							}
							return errors.New("lenVl validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{199, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{199, 1, 3, 4},
									bs)
							}
							return 1, wantErr
						},
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, lenVl, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If lenVl validator returns an error and Value Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 3
					wantErr               = errors.New("skip key error")
					bs                    = []byte{4, 199, 1, 3, 4}
					lenVl                 = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return errors.New("lenVl validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{199, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{199, 1, 3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{1, 3, 4},
									bs)
							}
							return 1, wantErr
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, lenVl, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Key Skipper == nil and lenVl validator returns an error, UnmarshalValidMap should return it immediately",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("lenVl Validator error")
					bs                    = []byte{4, 199, 1, 3, 4}
					sk2                   = mock.NewSkipper()
					lenVl                 = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, lenVl, nil, nil, nil,
						nil,
						nil,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Skipper == nil and lenVl validator returns an error, UnmarshalValidMap should return it immediately",
			func(t *testing.T) {
				var (
					wantV   map[uint]uint = nil
					wantN                 = 1
					wantErr               = errors.New("lenVl Validator error")
					bs                    = []byte{4, 199, 1, 3, 4}
					sk1                   = mock.NewSkipper()
					lenVl                 = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{lenVl.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, lenVl, nil, nil, nil,
						nil,
						sk1,
						nil,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Key Validator returns an error, UnmarshalValidMap should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					bs      = []byte{2, 10, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4},
									bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{u1.Mock, v1.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, nil, v1, nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Key Validator returns an error and Value Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("value Skipper error")
					bs      = []byte{2, 10, 100, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("key Validator error")
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{100, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{100, 1, 3, 4},
									bs)
							}
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock, v1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, nil, v1, nil,
						nil,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Skipper == nil and Key Validator returns an error, UnmarshalValidMap should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 2
					wantErr = errors.New("key Validator error")
					bs      = []byte{2, 10, 100, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock, v1.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, nil, v1, nil,
						nil,
						nil,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Key Validator retuns an error and Key Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					bs      = []byte{2, 10, 1, 200, 1, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("key Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{200, 1, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{200, 1, 4},
									bs)
							}
							return 2, wantErr
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 200, 1, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{1, 200, 1, 4},
									bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{u1.Mock, v1.Mock, sk1.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, nil, v1, nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Validator returns an error, UnmarshalValidMap should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("value Validator error")
					bs      = []byte{2, 10, 11, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 11 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 11, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4}, bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4}, bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{u1.Mock, u2.Mock, v2.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, u2, nil, v2,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If value Validator returns an error and Key Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("skip key error")
					bs      = []byte{2, 10, 11, 201, 4, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) { return 1, wantErr },
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{u1.Mock, u2.Mock, v2.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, u2, nil, v2,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Validator returns an error and Value Skipper returns an error, UnmarshalValidMap should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("skip key error")

					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshal(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) { return 1, nil },
					)
					sk2 = mock.NewSkipper().RegisterSkip(
						func(bs []byte) (n int, err error) { return 1, wantErr },
					)
					bs    = []byte{4, 10, 11, 3, 200, 1}
					mocks = []*mok.Mock{u1.Mock, u2.Mock, v1.Mock, v2.Mock, sk1.Mock,
						sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, nil, u1, u2, v1, v2,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("SkipMap should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN, bs = MakeBsWithNegativeLength()
					wantErr   = com.ErrNegativeLength
				)
				n, err := SkipMap(nil, nil, nil, bs)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipMap should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = SkipMap(nil, nil, nil, bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

}

func testAllSliceFunctions(
	m func(m1 mus.Marshaller[string]) mus.MarshallerFn[[]string],
	u func(u1 mus.Unmarshaller[string]) mus.UnmarshallerFn[[]string],
	s func(s1 mus.Sizer[string]) mus.SizerFn[[]string],
	sk func(sk1 mus.Skipper) mus.SkipperFn,
	t *testing.T,
) {
	var (
		str1    = "one"
		str1Raw = append([]byte{6}, []byte(str1)...)
		str2    = "two"
		str2Raw = append([]byte{6}, []byte(str2)...)
		sl      = []string{str1, str2}

		m1 mus.Marshaller[string] = mock.NewMarshaller[string]().RegisterNMarshal(4,
			func(v string, bs []byte) (n int) {
				switch v {
				case str1:
					return copy(bs, str1Raw)
				case str2:
					return copy(bs, str2Raw)
				default:
					t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
						str1, str2, v)
					return
				}
			},
		)
		u1 mus.Unmarshaller[string] = mock.NewUnmarshaller[string]().RegisterNUnmarshal(2,
			func(bs []byte) (v string, n int, err error) {
				if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
					return str1, len(str1Raw), nil
				} else if bytes.Equal(bs[:len(str2Raw)], str2Raw) {
					return str2, len(str2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						str1Raw, str2Raw, bs)
					return
				}
			},
		)
		s1 mus.Sizer[string] = mock.NewSizer[string]().RegisterNSize(4,
			func(v string) (size int) {
				switch v {
				case str1:
					return len(str1Raw)
				case str2:
					return len(str2Raw)
				default:
					t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
						str1, str2, v)
					return
				}
			},
		)
		sk1 mus.Skipper = mock.NewSkipper().RegisterNSkip(2,
			func(bs []byte) (n int, err error) {
				if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
					return len(str1Raw), nil
				} else if bytes.Equal(bs[:len(str2Raw)], str2Raw) {
					return len(str2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						str1Raw, str2Raw, bs)
					return
				}
			},
		)
	)
	testdata.Test[[]string]([][]string{sl},
		m(m1),
		u(u1),
		s(s1),
		t)
	testdata.TestSkip[[]string]([][]string{sl},
		m(m1),
		sk(sk1),
		s(s1),
		t)
}

func testAllMapFunctions(
	m func(m1 mus.Marshaller[string], m2 mus.Marshaller[uint]) mus.MarshallerFn[map[string]uint],
	u func(u1 mus.Unmarshaller[string], u2 mus.Unmarshaller[uint]) mus.UnmarshallerFn[map[string]uint],
	s func(s1 mus.Sizer[string], s2 mus.Sizer[uint]) mus.SizerFn[map[string]uint],
	sk func(sk1 mus.Skipper, sk2 mus.Skipper) mus.SkipperFn,
	t *testing.T,
) {
	var (
		str1                            = "one"
		str1Raw                         = append([]byte{6}, []byte(str1)...)
		str2                            = "two"
		str2Raw                         = append([]byte{6}, []byte(str2)...)
		int1    uint                    = 5
		int1Raw                         = []byte{5}
		int2    uint                    = 8
		int2Raw                         = []byte{8}
		mp                              = map[string]uint{str1: int1, str2: int2}
		m1      mock.Marshaller[string] = mock.NewMarshaller[string]().RegisterNMarshal(4,
			func(v string, bs []byte) (n int) {
				switch v {
				case str1:
					return copy(bs, str1Raw)
				case str2:
					return copy(bs, str2Raw)
				default:
					t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
						str1, str2, v)
					return
				}
			},
		)
		m2 mock.Marshaller[uint] = mock.NewMarshaller[uint]().RegisterNMarshal(4,
			func(v uint, bs []byte) (n int) {
				switch v {
				case int1:
					return copy(bs, int1Raw)
				case int2:
					return copy(bs, int2Raw)
				default:
					t.Fatalf("unexepcted uint, want '%v' or '%v' actual '%v'",
						int1, int2, v)
					return
				}
			},
		)
		u1 mock.Unmarshaller[string] = mock.NewUnmarshaller[string]().RegisterNUnmarshal(2,
			func(bs []byte) (v string, n int, err error) {
				if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
					return str1, len(str1Raw), nil
				} else if bytes.Equal(bs[:len(str2Raw)], str2Raw) {
					return str2, len(str2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						str1Raw, str2Raw, bs)
					return
				}
			},
		)
		u2 mock.Unmarshaller[uint] = mock.NewUnmarshaller[uint]().RegisterNUnmarshal(2,
			func(bs []byte) (v uint, n int, err error) {
				if bytes.Equal(bs[:len(int1Raw)], int1Raw) {
					return int1, len(int1Raw), nil
				} else if bytes.Equal(bs[:len(int2Raw)], int2Raw) {
					return int2, len(int2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						int1Raw, int2Raw, bs)
					return
				}
			},
		)
		s1 mock.Sizer[string] = mock.NewSizer[string]().RegisterNSize(4,
			func(v string) (size int) {
				switch v {
				case str1:
					return len(str1Raw)
				case str2:
					return len(str2Raw)
				default:
					t.Fatalf("unexepcted string, want '%v' or '%v' actual '%v'",
						str1, str2, v)
					return
				}
			},
		)
		s2 mock.Sizer[uint] = mock.NewSizer[uint]().RegisterNSize(4,
			func(v uint) (size int) {
				switch v {
				case int1:
					return len(int1Raw)
				case int2:
					return len(int2Raw)
				default:
					t.Fatalf("unexepcted uint, want '%v' or '%v' actual '%v'", int1,
						int2, v)
					return
				}
			},
		)
		sk1 mock.Skipper = mock.NewSkipper().RegisterNSkip(2,
			func(bs []byte) (n int, err error) {
				if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
					return len(str1Raw), nil
				} else if bytes.Equal(bs[:len(str2Raw)], str2Raw) {
					return len(str2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						str1Raw, str2Raw, bs)
					return
				}
			},
		)
		sk2 mock.Skipper = mock.NewSkipper().RegisterNSkip(2,
			func(bs []byte) (n int, err error) {
				if bytes.Equal(bs[:len(int1Raw)], int1Raw) {
					return len(int1Raw), nil
				} else if bytes.Equal(bs[:len(int2Raw)], int2Raw) {
					return len(int2Raw), nil
				} else {
					t.Fatalf("unexepcted bs, want '%v' or '%v' actual '%v'",
						int1Raw, int2Raw, bs)
					return
				}
			},
		)
		mocks = []*mok.Mock{m1.Mock, m2.Mock, u1.Mock, u2.Mock, s1.Mock,
			s2.Mock}
	)
	testdata.Test[map[string]uint]([]map[string]uint{mp},
		m(m1, m2),
		u(u1, u2),
		s(s1, s2),
		t)
	testdata.TestSkip[map[string]uint]([]map[string]uint{mp},
		m(m1, m2),
		sk(sk1, sk2),
		s(s1, s2),
		t)
	if info := mok.CheckCalls(mocks); len(info) > 0 {
		t.Error(info)
	}
}

func MakeBsWithNegativeLength() (n int, bs []byte) {
	n = varint.SizePositiveInt(-1)
	bs = make([]byte, n)
	varint.MarshalPositiveInt(-1, bs)
	return
}
