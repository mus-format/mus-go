package pm

import (
	"bytes"
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/varint"
)

func MarshalPointerMappingStruct(v com_testdata.PointerMappingStruct,
	bs []byte) (n int) {
	mp := com.NewPtrMap()
	n = MarshalPtr[int](v.A1, mus.MarshallerFn[int](varint.MarshalInt), mp,
		bs)
	n += MarshalPtr[int](v.A2, mus.MarshallerFn[int](varint.MarshalInt), mp,
		bs[n:])
	n += MarshalPtr[int](v.B1, mus.MarshallerFn[int](varint.MarshalInt), mp,
		bs[n:])
	n += MarshalPtr[int](v.B2, mus.MarshallerFn[int](varint.MarshalInt), mp,
		bs[n:])
	n += MarshalPtr[string](v.C1, mus.MarshallerFn[string](MarshalStringVarint),
		mp,
		bs[n:])
	n += MarshalPtr[string](v.C2, mus.MarshallerFn[string](MarshalStringVarint),
		mp,
		bs[n:])
	return
}

func UnmarshalPointerMappingStruct(bs []byte) (
	v com_testdata.PointerMappingStruct, n int, err error) {
	mp := com.NewReversePtrMap()
	v.A1, n, err = UnmarshalPtr[int](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		bs)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = UnmarshalPtr[int](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.B1, n1, err = UnmarshalPtr[int](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.B2, n1, err = UnmarshalPtr[int](mus.UnmarshallerFn[int](varint.UnmarshalInt),
		mp,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.C1, n1, err = UnmarshalPtr[string](
		mus.UnmarshallerFn[string](UnmarshalStringVarint),
		mp,
		bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.C2, n1, err = UnmarshalPtr[string](
		mus.UnmarshallerFn[string](UnmarshalStringVarint),
		mp,
		bs[n:])
	n += n1
	return
}

func SizePointerMappingStruct(v com_testdata.PointerMappingStruct) (size int) {
	mp := com.NewPtrMap()
	size = SizePtr[int](v.A1, mus.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.A2, mus.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.B1, mus.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[int](v.B2, mus.SizerFn[int](varint.SizeInt), mp)
	size += SizePtr[string](v.C1, mus.SizerFn[string](SizeStringVarint), mp)
	return size + SizePtr[string](v.C2, mus.SizerFn[string](SizeStringVarint),
		mp)
}

func SkipPointerMappingStruct(bs []byte) (n int, err error) {
	mp := com.NewReversePtrMap()
	n, err = SkipPtr(mus.SkipperFn(varint.SkipInt), mp, bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = SkipPtr(mus.SkipperFn(varint.SkipInt), mp, bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(mus.SkipperFn(varint.SkipInt), mp, bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(mus.SkipperFn(varint.SkipInt), mp, bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(mus.SkipperFn(SkipStringVarint), mp, bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = SkipPtr(mus.SkipperFn(SkipStringVarint), mp, bs[n:])
	n += n1
	return
}

func TestPM(t *testing.T) {

	t.Run("All MarshalPtr, UnmarshalPtr, SizePtr, SkipPtr functions must work correctly",
		func(t *testing.T) {
			var (
				m = mus.MarshallerFn[com_testdata.PointerMappingStruct](
					MarshalPointerMappingStruct)
				u = mus.UnmarshallerFn[com_testdata.PointerMappingStruct](
					UnmarshalPointerMappingStruct)
				s = mus.SizerFn[com_testdata.PointerMappingStruct](
					SizePointerMappingStruct)
				sk = mus.SkipperFn(SkipPointerMappingStruct)
			)
			testdata.Test[com_testdata.PointerMappingStruct](
				com_testdata.MakePointerMappingTestStruct(), m, u, s, t)
			testdata.TestSkip[com_testdata.PointerMappingStruct](
				com_testdata.MakePointerMappingTestStruct(), m, sk, s, t)
		})

	t.Run("MarshalPtr should be able to marshal nil pointer", func(t *testing.T) {
		var (
			wantN  = 1
			wantBS = []byte{1}
			mp     = com.NewPtrMap()
			bs     = make([]byte, SizePtr[int]((*int)(nil), nil, mp))
			n      = MarshalPtr[int](nil, nil, mp, bs)
		)
		if n != wantN {
			t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
		}
		if !bytes.Equal(bs, wantBS) {
			t.Errorf("unexpected bs, want '%v' actual '%v'", wantBS, bs)
		}
	})

	t.Run("UnmarshalPtr should return mus.ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				wantV     *int = nil
				wantN          = 0
				wantErr        = mus.ErrTooSmallByteSlice
				v, n, err      = UnmarshalPtr[int](nil, com.ReversePtrMap{}, []byte{})
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

	t.Run("UnmarshalPtr should be able to unmarshal nil pointer",
		func(t *testing.T) {
			var (
				wantV     *int  = nil
				wantN           = 1
				wantErr   error = nil
				bs              = []byte{1}
				v, n, err       = UnmarshalPtr[int](nil, com.ReversePtrMap{}, bs)
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

	t.Run("If unmarshal pointer id fails with an error, UnmarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantV     *int = nil
				wantN          = 1
				wantErr        = mus.ErrTooSmallByteSlice
				bs             = []byte{byte(com.Mapping)}
				v, n, err      = UnmarshalPtr[int](nil, com.ReversePtrMap{}, bs)
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

	t.Run("If unmarshal data fails with an error, UnmarshalPtr should return it",
		func(t *testing.T) {
			var (
				wantV   *int                    = nil
				wantN                           = 2
				wantErr                         = errors.New("unmarshal data error")
				u       mus.UnmarshallerFn[int] = func(bs []byte) (t int, n int, err error) {
					err = wantErr
					return
				}
				bs        = []byte{byte(com.Mapping), 2, 1}
				mp        = com.NewReversePtrMap()
				v, n, err = UnmarshalPtr[int](u, mp, bs)
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

	t.Run("UnmarshalPtr should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantV     *int = nil
				wantN          = 0
				wantErr        = com.ErrWrongFormat
				bs             = []byte{byte(com.Mapping) + 100}
				v, n, err      = UnmarshalPtr[int](nil, com.ReversePtrMap{}, bs)
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

	t.Run("SizePtr should return 1 for nil pointer", func(t *testing.T) {
		var (
			wantSize = 1
			size     = SizePtr[int](nil, nil, nil)
		)
		if size != wantSize {
			t.Errorf("unexpected size, want '%v' actual '%v'", wantSize, size)
		}
	})

	t.Run("If unmarshal id fails with an error, SkipPtr should return it",
		func(t *testing.T) {
			var (
				wantN   = 1
				wantErr = mus.ErrTooSmallByteSlice
				bs      = []byte{byte(com.Mapping)}
				n, err  = SkipPtr(nil, com.ReversePtrMap{}, bs)
			)
			if n != wantN {
				t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
			}
			if err != wantErr {
				t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
			}
		})

	t.Run("If Skipper fails, SkipPtr should return it", func(t *testing.T) {
		var (
			wantErr               = errors.New("Skipper error")
			wantN                 = 4
			mp                    = com.NewReversePtrMap()
			sk      mus.SkipperFn = func(bs []byte) (n int, err error) {
				return 2, wantErr
			}
			bs     = []byte{byte(com.Mapping), 1, 1}
			n, err = SkipPtr(sk, mp, bs)
		)
		if n != wantN {
			t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
		}
		if err != wantErr {
			t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
		}
	})

	t.Run("SkipPtr should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = com.ErrWrongFormat
				bs      = []byte{byte(com.Mapping) + 100}
				n, err  = SkipPtr(nil, com.ReversePtrMap{}, bs)
			)
			if n != wantN {
				t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
			}
			if err != wantErr {
				t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
			}
		})

	t.Run("SkipPtr should return mus.ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice
				n, err  = SkipPtr(nil, com.ReversePtrMap{}, []byte{})
			)
			if n != wantN {
				t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
			}
			if err != wantErr {
				t.Errorf("unexpected err, want '%v' actual '%v'", wantErr, err)
			}
		})

}

// StringVarint

func MarshalStringVarint(v string, bs []byte) (n int) {
	return ord.MarshalString(v, nil, bs)
}

func UnmarshalStringVarint(bs []byte) (v string,
	n int, err error) {
	return ord.UnmarshalString(nil, bs)
}

func UnmarshalValidStringVarint(lenVl com.Validator[int], skip bool, bs []byte) (
	v string, n int, err error) {
	return ord.UnmarshalValidString(nil, lenVl, skip, bs)
}

func SizeStringVarint(v string) (n int) {
	return ord.SizeString(v, nil)
}

func SkipStringVarint(bs []byte) (n int, err error) {
	return ord.SkipString(nil, bs)
}
