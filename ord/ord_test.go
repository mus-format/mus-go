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

		t.Run("All MarshalString, UnmarshalString, SizeString, SkipString function with default lenM, lenU, lenS must work correctly",
			func(t *testing.T) {
				var (
					m mus.MarshallerFn[string] = func(t string, bs []byte) (n int) {
						return MarshalString(t, nil, bs)
					}
					u mus.UnmarshallerFn[string] = func(bs []byte) (t string, n int, err error) {
						return UnmarshalString(nil, bs)
					}
					s mus.SizerFn[string] = func(t string) (size int) {
						return SizeString(t, nil)
					}
					sk mus.SkipperFn = func(bs []byte) (n int, err error) {
						return SkipString(nil, bs)
					}
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("All MarshalStringVarint, UnmarshalStringVarint, SizeStringVarint, SkipStringVarint functions must work correctly",
			func(t *testing.T) {
				var (
					m  mus.MarshallerFn[string]   = MarshalStringVarint
					u  mus.UnmarshallerFn[string] = UnmarshalStringVarint
					s  mus.SizerFn[string]        = SizeStringVarint
					sk mus.SkipperFn              = SkipStringVarint
				)
				testdata.Test[string](com_testdata.StringTestCases, m, u, s, t)
				testdata.TestSkip[string](com_testdata.StringTestCases, m, sk, s, t)
			})

		t.Run("MarshalStringVarint should panic with ErrTooSmallByteSlice if there is no space in bs",
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
				MarshalStringVarint("hello world", make([]byte, 2))
			})

		t.Run("UnmarshalStringVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = com.ErrNegativeLength
					bs        = []byte{1}
					v, n, err = UnmarshalStringVarint(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalStringVarint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = mus.ErrTooSmallByteSlice
					bs        = []byte{4, 2}
					v, n, err = UnmarshalStringVarint(bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If skip == true and lenVl validator returns an error, UnmarshalStringValid should return an error",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 3
					wantErr   = errors.New("lenVl validator error")
					bs        = []byte{4, 2, 2}
					maxLength = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					v, n, err = UnmarshalValidStringVarint(maxLength, true, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If skip == false and lenVl validator returns an error, UnmarshalStringValid should return an error",
			func(t *testing.T) {
				var (
					wantV     = ""
					wantN     = 1
					wantErr   = errors.New("lenVl validator error")
					bs        = []byte{4, 2, 2}
					maxLength = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) { return wantErr },
					)
					v, n, err = UnmarshalValidStringVarint(maxLength, false, bs)
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
				v, n, err = UnmarshalValidStringVarint(lenVl, false, bs)
			)
			com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("SkipStringVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					bs      = []byte{1}
					n, err  = SkipStringVarint(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipStringVarint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{4, 2}
					n, err  = SkipStringVarint(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("If SkipStringVarint should return an error if it fails to unmarshal a length",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = SkipStringVarint(bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("pointer", func(t *testing.T) {

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly for nil ptr",
			func(t *testing.T) {
				var (
					m = func() mus.MarshallerFn[*string] {
						return func(v *string, bs []byte) (n int) {
							return MarshalPtr(v, nil, bs)
						}
					}()
					u = func() mus.UnmarshallerFn[*string] {
						return func(bs []byte) (t *string, n int, err error) {
							return UnmarshalPtr[string](nil, bs)
						}
					}()
					s = func() mus.SizerFn[*string] {
						return func(v *string) (size int) {
							return SizePtr(v, nil)
						}
					}()
					sk = func() mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipPtr(nil, bs)
						}
					}()
				)
				testdata.Test[*string]([]*string{nil}, m, u, s, t)
				testdata.TestSkip[*string]([]*string{nil}, m, sk, s, t)
			})

		t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly not nil ptr",
			func(t *testing.T) {
				var (
					str1    = "one"
					str1Raw = append([]byte{6}, []byte(str1)...)
					ptr     = &str1
					m1      = func() mock.Marshaller[string] {
						return mock.NewMarshaller[string]().RegisterNMarshalMUS(2,
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
					}()
					u1 = func() mock.Unmarshaller[string] {
						return mock.NewUnmarshaller[string]().RegisterNUnmarshalMUS(1,
							func(bs []byte) (v string, n int, err error) {
								if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
									return str1, len(str1Raw), nil
								} else {
									t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
									return
								}
							},
						)
					}()
					s1 = func() mock.Sizer[string] {
						return mock.NewSizer[string]().RegisterNSizeMUS(2,
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
					}()
					sk1 = func() mock.Skipper {
						return mock.NewSkipper().RegisterNSkipMUS(1,
							func(bs []byte) (n int, err error) {
								if bytes.Equal(bs[:len(str1Raw)], str1Raw) {
									return len(str1Raw), nil
								} else {
									t.Fatalf("unexepcted bs, want '%v' actual '%v'", str1Raw, bs)
									return
								}
							},
						)
					}()
					m = func() mus.MarshallerFn[*string] {
						return func(v *string, bs []byte) (n int) {
							return MarshalPtr(v, mus.Marshaller[string](m1), bs)
						}
					}()
					u = func() mus.UnmarshallerFn[*string] {
						return func(bs []byte) (t *string, n int, err error) {
							return UnmarshalPtr(mus.Unmarshaller[string](u1), bs)
						}
					}()
					s = func() mus.SizerFn[*string] {
						return func(v *string) (size int) {
							return SizePtr(v, mus.Sizer[string](s1))
						}
					}()
					sk = func() mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipPtr(mus.Skipper(sk1), bs)
						}
					}()
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

		t.Run("Unmarshal should return ErrWrongFormat if meets wrong format",
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
					u               = mock.NewUnmarshaller[string]().RegisterUnmarshalMUS(
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

		t.Run("If Skipper fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 3
					wantErr = errors.New("error")
					s       = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							return 2, wantErr
						},
					)
					n, err = SkipPtr(s, []byte{byte(com.NotNil)})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("slice", func(t *testing.T) {

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice function with default lenM, lenU, lenS and empty slice must work correctly",
			func(t *testing.T) {
				var (
					sl = []string{}
					m  = mus.MarshallerFn[[]string](func(v []string, bs []byte) (n int) {
						return MarshalSlice(v, nil, nil, bs)
					})
					u = mus.UnmarshallerFn[[]string](func(bs []byte) (v []string, n int, err error) {
						return UnmarshalSlice[string](nil, nil, bs)
					})
					s = mus.SizerFn[[]string](func(v []string) (size int) {
						return SizeSlice(v, nil, nil)
					})
					sk = mus.SkipperFn(func(bs []byte) (n int, err error) {
						return SkipSlice(nil, nil, bs)
					})
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalSlice, UnmarshalSlice, SizeSlice, SkipSlice function with default lenM, lenU, lenS and not empty slice must work correctly",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 mus.Marshaller[string]) mus.MarshallerFn[[]string] {
						return func(v []string, bs []byte) (n int) {
							return MarshalSlice(v, nil, m1, bs)
						}
					},
					func(u1 mus.Unmarshaller[string]) mus.UnmarshallerFn[[]string] {
						return func(bs []byte) (v []string, n int, err error) {
							return UnmarshalSlice[string](nil, u1, bs)
						}
					},
					func(s1 mus.Sizer[string]) mus.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSlice(v, nil, s1)
						}
					},
					func(sk1 mus.Skipper) mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipSlice(nil, sk1, bs)
						}
					},
					t,
				)
			})

		t.Run("All MarshalSliceVarint, UnmarshalSliceVarint, SizeSliceVarint, SkipSliceVarint functions must work correctly with empty slice",
			func(t *testing.T) {
				var (
					sl = []string{}
					m  = func() mus.MarshallerFn[[]string] {
						return func(v []string, bs []byte) (n int) {
							return MarshalSliceVarint(v, nil, bs)
						}
					}()
					u = func() mus.UnmarshallerFn[[]string] {
						return func(bs []byte) (v []string, n int, err error) {
							return UnmarshalSliceVarint[string](nil, bs)
						}
					}()
					s = func() mus.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSliceVarint(v, nil)
						}
					}()
					sk = func() mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipSliceVarint(nil, bs)
						}
					}()
				)
				testdata.Test[[]string]([][]string{sl}, m, u, s, t)
				testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
			})

		t.Run("All MarshalSliceVarint, UnmarshalSliceVarint, SizeSliceVarint, SkipSliceVarint functions must work correctly with not empty slice",
			func(t *testing.T) {
				testAllSliceFunctions(
					func(m1 mus.Marshaller[string]) mus.MarshallerFn[[]string] {
						return func(v []string, bs []byte) (n int) {
							return MarshalSliceVarint(v, m1, bs)
						}
					},
					func(u1 mus.Unmarshaller[string]) mus.UnmarshallerFn[[]string] {
						return func(bs []byte) (v []string, n int, err error) {
							return UnmarshalSliceVarint[string](u1, bs)
						}
					},
					func(s1 mus.Sizer[string]) mus.SizerFn[[]string] {
						return func(v []string) (size int) {
							return SizeSliceVarint(v, s1)
						}
					},
					func(sk1 mus.Skipper) mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipSliceVarint(sk1, bs)
						}
					},
					t,
				)
			})

		t.Run("UnmarshalSliceVarint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 0
					wantErr          = mus.ErrTooSmallByteSlice
					v, n, err        = UnmarshalSliceVarint[uint](nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalSliceVarint should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 1
					wantErr          = com.ErrNegativeLength
					v, n, err        = UnmarshalSliceVarint[uint](nil, []byte{1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Unmarshaller fails with an error, UnmarshalSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = make([]uint, 1)
					wantN          = 3
					wantErr        = errors.New("Unmarshaller error")
					u              = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u.Mock}
					v, n, err = UnmarshalSliceVarint[uint](u, []byte{2})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Skipper != nil and lenVl validator returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 5
					wantErr          = errors.New("lenVl validator error")
					bs               = []byte{8, 4, 1, 1, 0}
					maxLength        = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 4 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					sk = mock.NewSkipper().RegisterNSkipMUS(4,
						func(bs []byte) (n int, err error) { return 1, nil },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](maxLength, nil, nil, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error and Skipper returns an error, UnmarshalValidSliceVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 4
					wantErr          = errors.New("skip rest error")
					bs               = []byte{4, 4, 1, 1}
					maxLength        = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 5, v)
							}
							return wantErr
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) { return 3, wantErr },
					)
					mocks     = []*mok.Mock{sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](maxLength, nil, nil, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Skipper == nil and lenVl validator returns an error, UnmarshalValidSliceVarint should return it",
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
					bs        = []byte{6, 10, 2, 3}
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](maxLength, nil, nil, nil, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Validator returns an error, UnmarshalValidSliceVarint should return it",
			func(t *testing.T) {
				var (
					wantV   []uint = []uint{10, 0, 0}
					wantN          = 4
					wantErr        = errors.New("Validator error")
					bs             = []byte{6, 10, 2, 3}
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
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					).RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 2, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3}, bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](nil, u, vl, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Validator returns an error and Skipper returns an error, UnmarshalValidSliceVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV   = []uint{0, 0, 0}
					wantN   = 4
					wantErr = errors.New("skip rest error")
					bs      = []byte{6, 10, 2, 3}
					vl      = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return errors.New("validator error")
						},
					)
					u = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					sk = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{2, 3}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{2, 3},
									bs)
							}
							return 2, wantErr
						},
					)
					mocks     = []*mok.Mock{vl.Mock, u.Mock, sk.Mock}
					v, n, err = UnmarshalValidSliceVarint[uint](nil, u, vl, sk, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("SkipSliceVarint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = SkipSliceVarint(nil, []byte{})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipSliceVarint should return ErrNegativeLength if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					n, err  = SkipSliceVarint(nil, []byte{1})
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

	})

	t.Run("map", func(t *testing.T) {

		t.Run("All MarshalMap, UnmarshalMap, SizeMap, SkipMap functions with default lenM, lenU, lenS, lenSk must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 mus.Marshaller[string], m2 mus.Marshaller[uint]) mus.MarshallerFn[map[string]uint] {
						return func(v map[string]uint, bs []byte) (n int) {
							return MarshalMap(v, nil, m1, m2, bs)
						}
					},
					func(u1 mus.Unmarshaller[string], u2 mus.Unmarshaller[uint]) mus.UnmarshallerFn[map[string]uint] {
						return func(bs []byte) (v map[string]uint, n int, err error) {
							return UnmarshalMap[string, uint](nil, u1, u2, bs)
						}
					},
					func(s1 mus.Sizer[string], s2 mus.Sizer[uint]) mus.SizerFn[map[string]uint] {
						return func(v map[string]uint) (size int) {
							return SizeMap(v, nil, s1, s2)
						}
					},
					func(sk1, sk2 mus.Skipper) mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipMap(nil, sk1, sk2, bs)
						}
					},
					t,
				)
			})

		t.Run("All MarshalMapVarint, UnmarshalMapVarint, SizeMapVarint, SkipMapVarint functions must work correctly",
			func(t *testing.T) {
				testAllMapFunctions(
					func(m1 mus.Marshaller[string], m2 mus.Marshaller[uint]) mus.MarshallerFn[map[string]uint] {
						return func(v map[string]uint, bs []byte) (n int) {
							return MarshalMapVarint(v, m1, m2, bs)
						}
					},
					func(u1 mus.Unmarshaller[string], u2 mus.Unmarshaller[uint]) mus.UnmarshallerFn[map[string]uint] {
						return func(bs []byte) (v map[string]uint, n int, err error) {
							return UnmarshalMapVarint[string, uint](u1, u2, bs)
						}
					},
					func(s1 mus.Sizer[string], s2 mus.Sizer[uint]) mus.SizerFn[map[string]uint] {
						return func(v map[string]uint) (size int) {
							return SizeMapVarint(v, s1, s2)
						}
					},
					func(sk1, sk2 mus.Skipper) mus.SkipperFn {
						return func(bs []byte) (n int, err error) {
							return SkipMapVarint(sk1, sk2, bs)
						}
					},
					t,
				)
			})

		t.Run("UnmarshalMapVarint should return ErrTooSmallByteSlice if there no space in bs",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 0
					wantErr                 = mus.ErrTooSmallByteSlice
					v, n, err               = UnmarshalMapVarint[uint, uint](nil, nil, []byte{})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("UnmarshalMapVarint should return ErrNegativeLength if meets negative length",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 1
					wantErr                 = com.ErrNegativeLength
					v, n, err               = UnmarshalMapVarint[uint, uint](nil, nil, []byte{1})
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
			})

		t.Run("If Key Unmarshaller fails with an error, UnmarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 3
					wantErr = errors.New("Unmarshaller error")
					bs      = []byte{2, 100}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							if !reflect.DeepEqual(bs, []byte{100}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{100},
									bs)
							}
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](u1, nil, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Value Unmarshaller fails with an error, UnmarshalMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 1)
					wantN   = 4
					wantErr = errors.New("Unmarshaller error")
					bs      = []byte{2, 1, 200, 200}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 1, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							if !reflect.DeepEqual(bs, []byte{200, 200}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{200, 200},
									bs)
							}
							return 0, 2, wantErr
						},
					)
					mocks     = []*mok.Mock{u1.Mock, u2.Mock}
					v, n, err = UnmarshalMapVarint[uint, uint](u1, u2, bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 5
					wantErr                 = errors.New("lenVl validator error")
					bs                      = []byte{4, 199, 1, 3, 4}
					maxLength               = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return wantErr
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{199, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{199, 1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4},
									bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](maxLength, nil, nil,
						nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If lenVl validator returns an error and Key Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 2
					wantErr                 = errors.New("skip key error")
					bs                      = []byte{4, 199, 1, 3, 4}
					maxLength               = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return errors.New("lenVl validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
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
					mocks     = []*mok.Mock{maxLength.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If lenVl validator returns an error and Value Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 3
					wantErr                 = errors.New("skip key error")
					bs                      = []byte{4, 199, 1, 3, 4}
					maxLength               = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return errors.New("lenVl validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{199, 1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{199, 1, 3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{1, 3, 4},
									bs)
							}
							return 1, wantErr
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Key Skipper == nil and lenVl validator returns an error, UnmarshalValidMapVarint should return it immediately",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 1
					wantErr                 = errors.New("lenVl Validator error")
					bs                      = []byte{4, 199, 1, 3, 4}
					sk2                     = mock.NewSkipper()
					maxLength               = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](maxLength, nil, nil, nil,
						nil,
						nil,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Skipper == nil and lenVl validator returns an error, UnmarshalValidMapVarint should return it immediately",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 1
					wantErr                 = errors.New("lenVl Validator error")
					bs                      = []byte{4, 199, 1, 3, 4}
					sk1                     = mock.NewSkipper()
					maxLength               = com_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						nil,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Key Validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					bs      = []byte{4, 10, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
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
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4},
									bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{1, 3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{1, 3, 4},
									bs)
							}
							return 1, nil
						},
					).RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4},
									bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{u1.Mock, v1.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If Key Validator returns an error and Value Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("value Skipper error")
					bs      = []byte{4, 10, 100, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("key Validator error")
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
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
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, nil,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Skipper == nil and Key Validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 2
					wantErr = errors.New("key Validator error")
					bs      = []byte{4, 10, 100, 1, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
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
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, nil,
						nil,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Key Validator retuns an error and Key Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					bs      = []byte{4, 10, 1, 200, 1, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("key Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{200, 1, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'",
									[]byte{200, 1, 4},
									bs)
							}
							return 2, wantErr
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
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
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Validator returns an error, UnmarshalValidMapVarint should return it",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("value Validator error")
					bs      = []byte{4, 10, 11, 3, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
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
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{3, 4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{3, 4}, bs)
							}
							return 1, nil
						},
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) {
							if !reflect.DeepEqual(bs, []byte{4}) {
								t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{4}, bs)
							}
							return 1, nil
						},
					)
					mocks     = []*mok.Mock{u1.Mock, u2.Mock, v2.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, nil, v2, sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("If value Validator returns an error and Key Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("skip key error")
					bs      = []byte{4, 10, 11, 201, 4, 4}
					u1      = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = com_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) { return 1, wantErr },
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{u1.Mock, u2.Mock, v2.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, nil, v2,
						sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("If Value Validator returns an error and Value Skipper returns an error, UnmarshalValidMapVarint should return the last one",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("skip key error")

					u1 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaller[uint]().RegisterUnmarshalMUS(
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
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) { return 1, nil },
					)
					sk2 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) { return 1, wantErr },
					)
					bs    = []byte{4, 10, 11, 3, 200, 1}
					mocks = []*mok.Mock{u1.Mock, u2.Mock, v1.Mock, v2.Mock, sk1.Mock,
						sk2.Mock}
					v, n, err = UnmarshalValidMapVarint[uint, uint](nil, u1, u2, v1, v2, sk1,
						sk2,
						bs)
				)
				com_testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("SkipMapVarint should return ErrNegativeLength if meets a negative length",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = com.ErrNegativeLength
					bs      = []byte{1}
					n, err  = SkipMapVarint(nil, nil, bs)
				)
				com_testdata.TestSkipResults(wantN, n, wantErr, err, nil, t)
			})

		t.Run("SkipMapVarint should return ErrTooSmallByteSlice if there is no space in bs",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{}
					n, err  = SkipMapVarint(nil, nil, bs)
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

		m1 mus.Marshaller[string] = mock.NewMarshaller[string]().RegisterNMarshalMUS(4,
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
		u1 mus.Unmarshaller[string] = mock.NewUnmarshaller[string]().RegisterNUnmarshalMUS(2,
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
		s1 mus.Sizer[string] = mock.NewSizer[string]().RegisterNSizeMUS(4,
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
		sk1 mus.Skipper = mock.NewSkipper().RegisterNSkipMUS(2,
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
		// m mus.MarshallerFn[[]string] = func(v []string, bs []byte) (n int) {
		// 	return MarshalSlice(v, nil, m1, bs)
		// }
		// u mus.UnmarshallerFn[[]string] = func(bs []byte) (t []string, n int, err error) {
		// 	return UnmarshalSlice(nil, u1, bs)
		// }
		// s mus.SizerFn[[]string] = func(v []string) (size int) {
		// 	return SizeSlice(v, nil, s1)
		// }
		// sk mus.SkipperFn = func(bs []byte) (n int, err error) {
		// 	return SkipSlice(nil, sk1, bs)
		// }

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
		str1         = "one"
		str1Raw      = append([]byte{6}, []byte(str1)...)
		str2         = "two"
		str2Raw      = append([]byte{6}, []byte(str2)...)
		int1    uint = 5
		int1Raw      = []byte{5}
		int2    uint = 8
		int2Raw      = []byte{8}
		mp           = map[string]uint{str1: int1, str2: int2}
		m1           = func() mock.Marshaller[string] {
			return mock.NewMarshaller[string]().RegisterNMarshalMUS(4,
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
		}()
		m2 = func() mock.Marshaller[uint] {
			return mock.NewMarshaller[uint]().RegisterNMarshalMUS(4,
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
		}()
		u1 = func() mock.Unmarshaller[string] {
			return mock.NewUnmarshaller[string]().RegisterNUnmarshalMUS(2,
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
		}()
		u2 = func() mock.Unmarshaller[uint] {
			return mock.NewUnmarshaller[uint]().RegisterNUnmarshalMUS(2,
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
		}()
		s1 = func() mock.Sizer[string] {
			return mock.NewSizer[string]().RegisterNSizeMUS(4,
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
		}()
		s2 = func() mock.Sizer[uint] {
			return mock.NewSizer[uint]().RegisterNSizeMUS(4,
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
		}()
		sk1 = func() mock.Skipper {
			return mock.NewSkipper().RegisterNSkipMUS(2,
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
		}()
		sk2 = func() mock.Skipper {
			return mock.NewSkipper().RegisterNSkipMUS(2,
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
		}()
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

// StringVarint

func MarshalStringVarint(v string, bs []byte) (n int) {
	return MarshalString(v, mus.MarshallerFn[int](varint.MarshalInt), bs)
}

func UnmarshalStringVarint(bs []byte) (v string,
	n int, err error) {
	return UnmarshalValidStringVarint(nil, false, bs)
}

func UnmarshalValidStringVarint(lenVl com.Validator[int], skip bool, bs []byte) (
	v string, n int, err error) {
	return UnmarshalValidString(mus.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, skip, bs)
}

func SizeStringVarint(v string) (n int) {
	return SizeString(v, mus.SizerFn[int](varint.SizeInt))
}

func SkipStringVarint(bs []byte) (n int, err error) {
	return SkipString(mus.UnmarshallerFn[int](varint.UnmarshalInt), bs)
}

// SliceVarint

func MarshalSliceVarint[T any](v []T, m mus.Marshaller[T], bs []byte) (n int) {
	return MarshalSlice(v, mus.MarshallerFn[int](varint.MarshalInt), m, bs)
}

func UnmarshalSliceVarint[T any](u mus.Unmarshaller[T], bs []byte) (v []T,
	n int, err error) {
	return UnmarshalValidSliceVarint[T](nil, u, nil, nil, bs)
}

func UnmarshalValidSliceVarint[T any](lenVl com.Validator[int],
	u mus.Unmarshaller[T],
	vl com.Validator[T],
	sk mus.Skipper,
	bs []byte,
) (v []T, n int, err error) {
	return UnmarshalValidSlice[T](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, u, vl, sk, bs)
}

func SizeSliceVarint[T any](v []T, s mus.Sizer[T]) (size int) {
	return SizeSlice(v, mus.SizerFn[int](varint.SizeInt), s)
}

func SkipSliceVarint(sk mus.Skipper, bs []byte) (n int,
	err error) {
	return SkipSlice(mus.UnmarshallerFn[int](varint.UnmarshalInt), sk, bs)
}

// MapVarint

func MarshalMapVarint[T comparable, V any](v map[T]V, m1 mus.Marshaller[T],
	m2 mus.Marshaller[V],
	bs []byte,
) (n int) {
	return MarshalMap(v, mus.MarshallerFn[int](varint.MarshalInt), m1, m2, bs)
}

func UnmarshalMapVarint[T comparable, V any](u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	bs []byte,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMapVarint[T, V](nil, u1, u2, nil, nil, nil, nil, bs)
}

func UnmarshalValidMapVarint[T comparable, V any](lenVl com.Validator[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 mus.Skipper,
	bs []byte,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMap[T, V](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		lenVl, u1, u2, vl1, vl2, sk1, sk2, bs)
}

func SizeMapVarint[T comparable, V any](v map[T]V, s1 mus.Sizer[T],
	s2 mus.Sizer[V]) (size int) {
	return SizeMap(v, mus.SizerFn[int](varint.SizeInt), s1, s2)
}

func SkipMapVarint(sk1, sk2 mus.Skipper, bs []byte) (n int, err error) {
	return SkipMap(mus.UnmarshallerFn[int](varint.UnmarshalInt), sk1, sk2, bs)
}
