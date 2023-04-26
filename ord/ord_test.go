package ord

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	muscom "github.com/mus-format/mus-common-go"
	muscom_mock "github.com/mus-format/mus-common-go/testdata/mock"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/testdata/mock"
	"github.com/ymz-ncnk/mok"
)

func TestOrd(t *testing.T) {

	t.Run("Bool", func(t *testing.T) {

		t.Run("bool", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[bool](MarshalBool)
				u  = mus.UnmarshalerFn[bool](UnmarshalBool)
				s  = mus.SizerFn[bool](SizeBool)
				sk = mus.SkipperFn(SkipBool)
			)
			testdata.Test[bool](testdata.BoolTestCases, m, u, s, t)
			testdata.TestSkip[bool](testdata.BoolTestCases, m, sk, s, t)
		})

		t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     = false
				wantN     = 0
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{}
				v, n, err = UnmarshalBool(bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - ErrWrongFormat", func(t *testing.T) {
			var (
				wantV     = false
				wantN     = 0
				wantErr   = muscom.ErrWrongFormat
				bs        = []byte{3}
				v, n, err = UnmarshalBool(bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = SkipBool(bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrWrongFormat", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = muscom.ErrWrongFormat
				bs      = []byte{3}
				n, err  = SkipBool(bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("String", func(t *testing.T) {

		t.Run("string", func(t *testing.T) {
			var (
				m  = mus.MarshalerFn[string](MarshalString)
				u  = mus.UnmarshalerFn[string](UnmarshalString)
				s  = mus.SizerFn[string](SizeString)
				sk = mus.SkipperFn(SkipString)
			)
			testdata.Test[string](testdata.StringTestCases, m, u, s, t)
			testdata.TestSkip[string](testdata.StringTestCases, m, sk, s, t)
		})

		t.Run("Marshal - panic ErrTooSmallByteSlice", func(t *testing.T) {
			wantErr := mus.ErrTooSmallByteSlice
			defer func() {
				if r := recover(); r != nil {
					err := r.(error)
					if err != wantErr {
						t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
					}
				}
			}()
			MarshalString("hello world", make([]byte, 2))
		})

		t.Run("Unmarshal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 1
				wantErr   = muscom.ErrNegativeLength
				bs        = []byte{1}
				v, n, err = UnmarshalString(bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 1
				wantErr   = mus.ErrTooSmallByteSlice
				bs        = []byte{4, 2}
				v, n, err = UnmarshalString(bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("UnmarshalValid - MaxLength validator error", func(t *testing.T) {
			var (
				wantV     = ""
				wantN     = 3
				wantErr   = errors.New("MaxLength validator error")
				bs        = []byte{4, 2, 2}
				maxLength = muscom_mock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) { return wantErr },
				)
				v, n, err = UnmarshalValidString(maxLength, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				bs      = []byte{1}
				n, err  = SkipString(bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{4, 2}
				n, err  = SkipString(bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - unmarshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = SkipString(bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("Pointer", func(t *testing.T) {

		t.Run("Nil ptr", func(t *testing.T) {
			var (
				m = func() mus.MarshalerFn[*string] {
					return func(v *string, bs []byte) (n int) {
						return MarshalPtr(v, nil, bs)
					}
				}()
				u = func() mus.UnmarshalerFn[*string] {
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

		t.Run("Not nil ptr", func(t *testing.T) {
			var (
				str1    = "one"
				str1Raw = append([]byte{6}, []byte(str1)...)
				ptr     = &str1
				m1      = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(2,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterNUnmarshalMUS(1,
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
				m = func() mus.MarshalerFn[*string] {
					return func(v *string, bs []byte) (n int) {
						return MarshalPtr(v, mus.Marshaler[string](m1), bs)
					}
				}()
				u = func() mus.UnmarshalerFn[*string] {
					return func(bs []byte) (t *string, n int, err error) {
						return UnmarshalPtr(mus.Unmarshaler[string](u1), bs)
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

		t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     *string = nil
				wantN             = 0
				wantErr           = mus.ErrTooSmallByteSlice
				v, n, err         = UnmarshalPtr[string](nil, []byte{})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - ErrWrongFormat", func(t *testing.T) {
			var (
				wantV     *string = nil
				wantN             = 0
				wantErr           = muscom.ErrWrongFormat
				v, n, err         = UnmarshalPtr[string](nil, []byte{2})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - Unmarshaler error", func(t *testing.T) {
			var (
				wantV   *string = nil
				wantN           = 5
				wantErr         = errors.New("unmarshaler error")
				u               = mock.NewUnmarshaler[string]().RegisterUnmarshalMUS(
					func(bs []byte) (v string, n int, err error) {
						return "", 4, wantErr
					},
				)
				v, n, err = UnmarshalPtr[string](u, []byte{muscom.NotNilFlag})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				n, err  = SkipPtr(nil, []byte{})
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrWrongFormat", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = muscom.ErrWrongFormat
				n, err  = SkipPtr(nil, []byte{2})
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - Skipper error", func(t *testing.T) {
			var (
				wantN   = 3
				wantErr = errors.New("error")
				s       = mock.NewSkipper().RegisterSkipMUS(
					func(bs []byte) (n int, err error) {
						return 2, wantErr
					},
				)
				n, err = SkipPtr(s, []byte{muscom.NotNilFlag})
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("Slice", func(t *testing.T) {

		t.Run("Empty slice", func(t *testing.T) {
			var (
				sl = []string{}
				m  = func() mus.MarshalerFn[[]string] {
					return func(v []string, bs []byte) (n int) {
						return MarshalSlice(v, nil, bs)
					}
				}()
				u = func() mus.UnmarshalerFn[[]string] {
					return func(bs []byte) (v []string, n int, err error) {
						return UnmarshalSlice[string](nil, bs)
					}
				}()
				s = func() mus.SizerFn[[]string] {
					return func(v []string) (size int) {
						return SizeSlice(v, nil)
					}
				}()
				sk = func() mus.SkipperFn {
					return func(bs []byte) (n int, err error) {
						return SkipSlice(nil, bs)
					}
				}()
			)
			testdata.Test[[]string]([][]string{sl}, m, u, s, t)
			testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
		})

		t.Run("Not empty slice", func(t *testing.T) {
			var (
				str1    = "one"
				str1Raw = append([]byte{6}, []byte(str1)...)
				str2    = "two"
				str2Raw = append([]byte{6}, []byte(str2)...)
				sl      = []string{str1, str2}

				m1 = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(4,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterNUnmarshalMUS(2,
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
				m = func() mus.MarshalerFn[[]string] {
					return func(v []string, bs []byte) (n int) {
						return MarshalSlice(v, mus.Marshaler[string](m1), bs)
					}
				}()
				u = func() mus.UnmarshalerFn[[]string] {
					return func(bs []byte) (t []string, n int, err error) {
						return UnmarshalSlice(mus.Unmarshaler[string](u1), bs)
					}
				}()
				s = func() mus.SizerFn[[]string] {
					return func(t []string) (size int) {
						return SizeSlice(t, mus.Sizer[string](s1))
					}
				}()
				sk = func() mus.SkipperFn {
					return func(bs []byte) (n int, err error) {
						return SkipSlice(mus.Skipper(sk1), bs)
					}
				}()
			)
			testdata.Test[[]string]([][]string{sl}, m, u, s, t)
			testdata.TestSkip[[]string]([][]string{sl}, m, sk, s, t)
		})

		t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     []uint = nil
				wantN            = 0
				wantErr          = mus.ErrTooSmallByteSlice
				v, n, err        = UnmarshalSlice[uint](nil, []byte{})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV     []uint = nil
				wantN            = 1
				wantErr          = muscom.ErrNegativeLength
				v, n, err        = UnmarshalSlice[uint](nil, []byte{1})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - Unmarshaler error", func(t *testing.T) {
			var (
				wantV   []uint = make([]uint, 1)
				wantN          = 3
				wantErr        = errors.New("unmarshaler error")
				u              = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						return 0, 2, wantErr
					},
				)
				mocks     = []*mok.Mock{u.Mock}
				v, n, err = UnmarshalSlice[uint](u, []byte{2})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - MaxLength validator error", func(t *testing.T) {
			var (
				wantV     []uint = nil
				wantN            = 5
				wantErr          = errors.New("MaxLength validator error")
				bs               = []byte{4, 4, 1, 1, 0}
				maxLength        = muscom_mock.NewValidator[int]().RegisterValidate(
					func(v int) (err error) {
						if v != 2 {
							t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
						}
						return wantErr
					},
				)
				sk = mock.NewSkipper().RegisterSkipMUS(
					func(bs []byte) (n int, err error) { return 4, nil },
				)
				mocks     = []*mok.Mock{sk.Mock}
				v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, sk, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - MaxLength validator error, skip rest - error",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 4
					wantErr          = errors.New("skip rest error")
					bs               = []byte{4, 4, 1, 1}
					maxLength        = muscom_mock.NewValidator[int]().RegisterValidate(
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
					v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, sk, bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator error, Skipper == nil",
			func(t *testing.T) {
				var (
					wantV     []uint = nil
					wantN            = 1
					wantErr          = errors.New("MaxLength Validator error")
					maxLength        = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					bs        = []byte{6, 10, 2, 3}
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidSlice[uint](maxLength, nil, nil, nil, bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - Validator error", func(t *testing.T) {
			var (
				wantV   []uint = []uint{10, 0, 0}
				wantN          = 4
				wantErr        = errors.New("Validator error")
				bs             = []byte{6, 10, 2, 3}
				vl             = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				u = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
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
				v, n, err = UnmarshalValidSlice[uint](nil, u, vl, sk, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - Validator error, skip rest - error",
			func(t *testing.T) {
				var (
					wantV   = []uint{0, 0, 0}
					wantN   = 4
					wantErr = errors.New("skip rest error")
					bs      = []byte{6, 10, 2, 3}
					vl      = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							if v != 10 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 10, v)
							}
							return errors.New("validator error")
						},
					)
					u = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
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
					v, n, err = UnmarshalValidSlice[uint](nil, u, vl, sk, bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("Skip - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				n, err  = SkipSlice(nil, []byte{})
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				n, err  = SkipSlice(nil, []byte{1})
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

	t.Run("Map", func(t *testing.T) {

		t.Run("map", func(t *testing.T) {
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
				m1           = func() mock.Marshaler[string] {
					return mock.NewMarshaler[string]().RegisterNMarshalMUS(4,
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
				m2 = func() mock.Marshaler[uint] {
					return mock.NewMarshaler[uint]().RegisterNMarshalMUS(4,
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
				u1 = func() mock.Unmarshaler[string] {
					return mock.NewUnmarshaler[string]().RegisterNUnmarshalMUS(2,
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
				u2 = func() mock.Unmarshaler[uint] {
					return mock.NewUnmarshaler[uint]().RegisterNUnmarshalMUS(2,
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
				m = func() mus.MarshalerFn[map[string]uint] {
					return func(v map[string]uint, bs []byte) int {
						return MarshalMap(v,
							mus.Marshaler[string](m1),
							mus.Marshaler[uint](m2),
							bs)
					}
				}()
				u = func() mus.UnmarshalerFn[map[string]uint] {
					return func(bs []byte) (t map[string]uint, n int, err error) {
						return UnmarshalMap(
							mus.Unmarshaler[string](u1),
							mus.Unmarshaler[uint](u2),
							bs)
					}
				}()
				s = func() mus.SizerFn[map[string]uint] {
					return func(v map[string]uint) (size int) {
						return SizeMap(v,
							mus.Sizer[string](s1),
							mus.Sizer[uint](s2))
					}
				}()
				sk = func() mus.SkipperFn {
					return func(bs []byte) (n int, err error) {
						return SkipMap(mus.Skipper(sk1), mus.Skipper(sk2), bs)
					}
				}()
				mocks = []*mok.Mock{m1.Mock, m2.Mock, u1.Mock, u2.Mock, s1.Mock,
					s2.Mock}
			)
			testdata.Test[map[string]uint]([]map[string]uint{mp}, m, u, s, t)
			testdata.TestSkip[map[string]uint]([]map[string]uint{mp}, m, sk, s, t)
			if info := mok.CheckCalls(mocks); len(info) > 0 {
				t.Error(info)
			}
		})

		t.Run("Unmarshal - ErrTooSmallByteSlice", func(t *testing.T) {
			var (
				wantV     map[uint]uint = nil
				wantN                   = 0
				wantErr                 = mus.ErrTooSmallByteSlice
				v, n, err               = UnmarshalMap[uint, uint](nil, nil, []byte{})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - ErrNegativeLength", func(t *testing.T) {
			var (
				wantV     map[uint]uint = nil
				wantN                   = 1
				wantErr                 = muscom.ErrNegativeLength
				v, n, err               = UnmarshalMap[uint, uint](nil, nil, []byte{1})
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, nil, t)
		})

		t.Run("Unmarshal - key Unmarshaler error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 3
				wantErr = errors.New("unmarshaler error")
				bs      = []byte{2, 100}
				u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{100}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{100},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				mocks     = []*mok.Mock{u1.Mock}
				v, n, err = UnmarshalMap[uint, uint](u1, nil, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("Unmarshal - value Unmarshaler error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 1)
				wantN   = 4
				wantErr = errors.New("unmarshaler error")
				bs      = []byte{2, 1, 200, 200}
				u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						return 1, 1, nil
					},
				)
				u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						if !reflect.DeepEqual(bs, []byte{200, 200}) {
							t.Errorf("unexpected bs, want '%v' actual '%v'", []byte{200, 200},
								bs)
						}
						return 0, 2, wantErr
					},
				)
				mocks     = []*mok.Mock{u1.Mock, u2.Mock}
				v, n, err = UnmarshalMap[uint, uint](u1, u2, bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - MaxLength validator error", func(t *testing.T) {
			var (
				wantV     map[uint]uint = nil
				wantN                   = 5
				wantErr                 = errors.New("MaxLength validator error")
				bs                      = []byte{4, 199, 1, 3, 4}
				maxLength               = muscom_mock.NewValidator[int]().RegisterValidate(
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
				v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil, nil,
					sk1,
					sk2,
					bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - MaxLength validator error, skip key - error",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 2
					wantErr                 = errors.New("skip key error")
					bs                      = []byte{4, 199, 1, 3, 4}
					maxLength               = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							if v != 2 {
								t.Errorf("unexpected v, want '%v' actual '%v'", 2, v)
							}
							return errors.New("MaxLength validator error")
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
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength validator error, skip value - error",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 3
					wantErr                 = errors.New("skip key error")
					bs                      = []byte{4, 199, 1, 3, 4}
					maxLength               = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return errors.New("MaxLength validator error")
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
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
			})

		t.Run("UnmarshalValid - MaxLength Validator erorr, key Skipper = nil",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 1
					wantErr                 = errors.New("MaxLength Validator error")
					bs                      = []byte{4, 199, 1, 3, 4}
					sk2                     = mock.NewSkipper()
					maxLength               = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						nil,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - MaxLength Validator erorr, value Skipper == nil",
			func(t *testing.T) {
				var (
					wantV     map[uint]uint = nil
					wantN                   = 1
					wantErr                 = errors.New("MaxLength Validator error")
					bs                      = []byte{4, 199, 1, 3, 4}
					sk1                     = mock.NewSkipper()
					maxLength               = muscom_mock.NewValidator[int]().RegisterValidate(
						func(v int) (err error) {
							return wantErr
						},
					)
					mocks     = []*mok.Mock{maxLength.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](maxLength, nil, nil, nil,
						nil,
						sk1,
						nil,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - key Validator error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 5
				wantErr = errors.New("key Validator error")
				bs      = []byte{4, 10, 1, 3, 4}
				u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, sk1,
					sk2,
					bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - key Validator error, skip value - error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("skip value error")
					bs      = []byte{4, 10, 100, 1, 3, 4}
					u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, nil,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - key Validator error, skip key - error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("key Validator error")
					bs      = []byte{4, 10, 1, 200, 1, 4}
					u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, nil, v1, nil, sk1,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - value Validator error", func(t *testing.T) {
			var (
				wantV   = make(map[uint]uint, 2)
				wantN   = 5
				wantErr = errors.New("value Validator error")
				bs      = []byte{4, 10, 11, 3, 4}
				u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						return 10, 1, nil
					},
				)
				u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
					func(bs []byte) (v uint, n int, err error) {
						return 11, 1, nil
					},
				)
				v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
				v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, nil, v2, sk1,
					sk2,
					bs)
			)
			testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks, t)
		})

		t.Run("UnmarshalValid - value Validator error, skip key error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 4
					wantErr = errors.New("skip key error")
					bs      = []byte{4, 10, 11, 201, 4, 4}
					u1      = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return errors.New("value Validator error")
						},
					)
					sk1 = mock.NewSkipper().RegisterSkipMUS(
						func(bs []byte) (n int, err error) { return 1, wantErr },
					)
					sk2       = mock.NewSkipper()
					mocks     = []*mok.Mock{u1.Mock, u2.Mock, v2.Mock, sk1.Mock, sk2.Mock}
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, nil, v2,
						sk1,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("UnmarshalValid - value Validator error, skip value error",
			func(t *testing.T) {
				var (
					wantV   = make(map[uint]uint, 2)
					wantN   = 5
					wantErr = errors.New("skip key error")

					u1 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 10, 1, nil
						},
					)
					u2 = mock.NewUnmarshaler[uint]().RegisterUnmarshalMUS(
						func(bs []byte) (v uint, n int, err error) {
							return 11, 1, nil
						},
					)
					v1 = muscom_mock.NewValidator[uint]().RegisterValidate(
						func(v uint) (err error) {
							return nil
						},
					)
					v2 = muscom_mock.NewValidator[uint]().RegisterValidate(
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
					v, n, err = UnmarshalValidMap[uint, uint](nil, u1, u2, v1, v2, sk1,
						sk2,
						bs)
				)
				testdata.TestUnmarshalResults(wantV, v, wantN, n, wantErr, err, mocks,
					t)
			})

		t.Run("Skip - ErrNegativeLength", func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = muscom.ErrNegativeLength
				bs      = []byte{1}
				n, err  = SkipMap(nil, nil, bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

		t.Run("Skip - unmarshal length error", func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{}
				n, err  = SkipMap(nil, nil, bs)
			)
			testdata.TestSkipResults(wantN, n, wantErr, err, t)
		})

	})

}
