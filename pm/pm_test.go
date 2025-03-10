package pm

import (
	"bytes"
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	mock "github.com/mus-format/mus-go/testdata/mock"
)

func TestPM(t *testing.T) {

	t.Run("pointer", func(t *testing.T) {

		t.Run("Marshal should be able to marshal nil pointer", func(t *testing.T) {
			var (
				wantN  = 1
				wantBS = []byte{byte(com.Nil)}
				ser    = NewPtrSer[int](com.NewPtrMap(), nil, nil)
				size   = ser.Size(nil)
				bs     = make([]byte, size)
				n      = ser.Marshal(nil, bs)
			)
			if n != wantN {
				t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
			}
			if !bytes.Equal(bs, wantBS) {
				t.Errorf("unexpected bs, want '%v' actual '%v'", wantBS, bs)
			}
		})

		t.Run("Unmarshal should return mus.ErrTooSmallByteSlice if bs is too small",
			func(t *testing.T) {
				var (
					wantV     *int = nil
					wantN          = 0
					wantErr        = mus.ErrTooSmallByteSlice
					ser            = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
					v, n, err      = ser.Unmarshal([]byte{})
				)
				if v != wantV {
					t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
				}
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Unmarshal should be able to unmarshal nil pointer",
			func(t *testing.T) {
				var (
					wantV     *int  = nil
					wantN           = 1
					wantErr   error = nil
					bs              = []byte{byte(com.Nil)}
					ser             = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
					v, n, err       = ser.Unmarshal(bs)
				)
				if v != wantV {
					t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
				}
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("If unmarshaling pointer id fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV     *int = nil
					wantN          = 1
					wantErr        = mus.ErrTooSmallByteSlice
					bs             = []byte{byte(com.Mapping)}
					ser            = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
					v, n, err      = ser.Unmarshal(bs)
				)
				if v != wantV {
					t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
				}
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("If unmarshaling data fails with an error, Unmarshal should return it",
			func(t *testing.T) {
				var (
					wantV   *int = nil
					wantN        = 2
					wantErr      = errors.New("unmarshal data error")
					baseSer      = mock.NewSerializer[int]().RegisterUnmarshal(
						func(bs []byte) (t int, n int, err error) {
							err = wantErr
							return
						},
					)
					ptrMap    = com.NewReversePtrMap()
					ser       = NewPtrSer[int](nil, ptrMap, baseSer)
					bs        = []byte{byte(com.Mapping), 2, 1}
					v, n, err = ser.Unmarshal(bs)
				)
				if v != wantV {
					t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
				}
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Unmarshal should fail with com.ErrWrongFormat if meets unknown pointer flag",
			func(t *testing.T) {
				var (
					wantV     *int = nil
					wantN          = 0
					wantErr        = com.ErrWrongFormat
					bs             = []byte{byte(com.Mapping) + 100}
					ser            = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
					v, n, err      = ser.Unmarshal(bs)
				)
				if v != wantV {
					t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
				}
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Size should return 1 for nil pointer", func(t *testing.T) {
			var (
				wantSize = 1
				size     = NewPtrSer[int](nil, nil, nil).Size(nil)
			)
			if size != wantSize {
				t.Errorf("unexpected size, want '%v' actual '%v'", wantSize, size)
			}
		})

		t.Run("If id unmarshaling fails with an error, Skip should return it",
			func(t *testing.T) {
				var (
					wantN   = 1
					wantErr = mus.ErrTooSmallByteSlice
					bs      = []byte{byte(com.Mapping)}
					n, err  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip(bs)
				)
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Skip should fail with com.ErrWrongFormat if meets unknown pointer flag",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = com.ErrWrongFormat
					bs      = []byte{byte(com.Mapping) + 100}
					n, err  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip(bs)
				)
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Skip should return mus.ErrTooSmallByteSlice if bs is too small",
			func(t *testing.T) {
				var (
					wantN   = 0
					wantErr = mus.ErrTooSmallByteSlice
					n, err  = NewPtrSer[int](nil, com.NewReversePtrMap(), nil).Skip([]byte{})
				)
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("Skip should be able to skip nil pointer",
			func(t *testing.T) {
				var (
					wantN         = 1
					wantErr error = nil
					bs            = []byte{byte(com.Nil)}
					ser           = NewPtrSer[int](nil, com.NewReversePtrMap(), nil)
					n, err        = ser.Skip(bs)
				)
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

		t.Run("If unmarshaling data fails with an error, SKip should return it",
			func(t *testing.T) {
				var (
					wantN   = 2
					wantErr = errors.New("unmarshal data error")
					baseSer = mock.NewSerializer[int]().RegisterSkip(
						func(bs []byte) (n int, err error) {
							err = wantErr
							return
						},
					)
					ptrMap = com.NewReversePtrMap()
					ser    = NewPtrSer[int](nil, ptrMap, baseSer)
					bs     = []byte{byte(com.Mapping), 2, 1}
					n, err = ser.Skip(bs)
				)
				if n != wantN {
					t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
				}
				if err != wantErr {
					t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
				}
			})

	})

}
