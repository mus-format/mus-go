package pm

import (
	"errors"
	"fmt"
	"testing"
	"unsafe"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test"
	mock "github.com/mus-format/mus-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
	"github.com/ymz-ncnk/mok"
)

func TestPM_Pointer(t *testing.T) {
	t.Run("Marshal should be able to marshal nil pointer", func(t *testing.T) {
		var (
			wantN  = 1
			wantBS = []byte{byte(com.Nil)}
			ser    = NewPtrSer(com.NewPtrMap(), nil, mus.Serializer[int](nil))
			size   = ser.Size(nil)
			bs     = make([]byte, size)
			n      = ser.Marshal(nil, bs)
		)
		asserterror.Equal(t, n, wantN,
			fmt.Sprintf("unexpected n, want '%v' actual '%v'", wantN, n))
		asserterror.EqualDeep(t, bs, wantBS,
			fmt.Sprintf("unexpected bs, want '%v' actual '%v'", wantBS, bs))
	})

	t.Run("Unmarshal should return mus.ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[*int]{
					V:   nil,
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestUnmarshalOnly([]byte{}, ser, want, nil, t)
		})

	t.Run("Unmarshal should be able to unmarshal nil pointer",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[*int]{
					V:   nil,
					N:   1,
					Err: nil,
				}
				bs  = []byte{byte(com.Nil)}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
		})

	t.Run("If unmarshaling pointer id fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[*int]{
					V:   nil,
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs  = []byte{byte(com.Mapping)}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
		})

	t.Run("If unmarshaling data fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[*int]{
					V:   nil,
					N:   2,
					Err: errors.New("unmarshal data error"),
				}
				ptrMap  = com.NewReversePtrMap()
				baseSer = mock.NewSerializer[int]().RegisterUnmarshal(
					func(bs []byte) (t int, n int, err error) {
						err = want.Err
						return
					},
				)
				ser = NewPtrSer(nil, ptrMap, baseSer)
				bs  = []byte{byte(com.Mapping), 2, 1}
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
		})

	t.Run("Unmarshal should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				want = test.UnmarshalResult[*int]{
					V:   nil,
					N:   0,
					Err: com.ErrWrongFormat,
				}
				bs  = []byte{byte(com.Mapping) + 100}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestUnmarshalOnly(bs, ser, want, nil, t)
		})

	t.Run("Size should return 1 for nil pointer", func(t *testing.T) {
		var (
			wantSize = 1
			ser      = NewPtrSer(nil, nil, mus.Serializer[int](nil))
		)
		size := ser.Size(nil)
		asserterror.Equal(t, size, wantSize,
			fmt.Sprintf("unexpected size, want '%v' actual '%v'", wantSize, size))
	})

	t.Run("If id unmarshaling fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   1,
					Err: mus.ErrTooSmallByteSlice,
				}
				bs  = []byte{byte(com.Mapping)}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestSkipOnly(bs, ser, want, nil, t)
		})

	t.Run("Skip should fail with com.ErrWrongFormat if meets unknown pointer flag",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: com.ErrWrongFormat,
				}
				bs  = []byte{byte(com.Mapping) + 100}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestSkipOnly(bs, ser, want, nil, t)
		})

	t.Run("Skip should return mus.ErrTooSmallByteSlice if bs is too small",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   0,
					Err: mus.ErrTooSmallByteSlice,
				}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestSkipOnly([]byte{}, ser, want, nil, t)
		})

	t.Run("Skip should be able to skip nil pointer",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   1,
					Err: nil,
				}
				bs  = []byte{byte(com.Nil)}
				ser = NewPtrSer(nil, com.NewReversePtrMap(), mus.Serializer[int](nil))
			)
			test.TestSkipOnly(bs, ser, want, nil, t)
		})

	t.Run("If unmarshaling data fails with an error, SKip should return it",
		func(t *testing.T) {
			var (
				want = test.SkipResult{
					N:   2,
					Err: errors.New("unmarshal data error"),
				}
				ptrMap  = com.NewReversePtrMap()
				baseSer = mock.NewSerializer[int]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						err = want.Err
						return
					},
				)
				ser = NewPtrSer(nil, ptrMap, baseSer)
				bs  = []byte{byte(com.Mapping), 2, 1}
			)
			test.TestSkipOnly(bs, ser, want, nil, t)
		})
}

func TestPM_Wrapper(t *testing.T) {
	t.Run("Wrapped serializer should succeed",
		func(t *testing.T) {
			var (
				st, baseSer = test.PtrStructTestData(t)
				ptrMap      = com.NewPtrMap()
				revPtrMap   = com.NewReversePtrMap()
				ser         = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap, baseSer))
			)
			test.Test([]ctest.PtrStruct{st}, ser, t)
			test.TestSkip([]ctest.PtrStruct{st}, ser, t)
		})

	t.Run("Marshal should call ser.Marshal and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV  byte = 1
				wantN  int  = 1
				ptrMap      = com.NewPtrMap()
				ptrSer      = mock.NewSerializer[byte]().RegisterMarshal(
					func(v byte, bs []byte) (n int) {
						assertfatal.Equal(t, v, wantV,
							fmt.Sprintf("unexpected v, want %v, actual %v", wantV, v))
						ptrMap.Put(unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser   = Wrap(ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			n := ser.Marshal(wantV, nil)
			assertfatal.Equal(t, n, wantN,
				fmt.Sprintf("unexpected n, want %v, actual %v", wantN, n))
			assertfatal.Equal(t, ptrMap.Len(), 0, "ptrMap should be empty")
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap,
				"unexpected mocks")
		})

	t.Run("Unmarshal should call ser.Unmarshal and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN     int  = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterUnmarshal(
					func(bs []byte) (v byte, n int, err error) {
						v = bs[0]
						assertfatal.Equal(t, v, wantV,
							fmt.Sprintf("unexpected v, want %v, actual %v", wantV, v))
						n = wantN
						revPtrMap.Put(1, unsafe.Pointer(&v))
						return
					},
				)
				ser   = Wrap(nil, revPtrMap, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{wantV})
			assertfatal.Equal(t, v, wantV,
				fmt.Sprintf("unexpected v, want %v, actual %v", wantV, v))
			assertfatal.Equal(t, n, wantN,
				fmt.Sprintf("unexpected n, want %v, actual %v", wantN, n))
			assertfatal.EqualError(t, err, nil, "unexpected error")
			assertfatal.Equal(t, revPtrMap.Len(), 0, "revPtrMap should be empty")
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap,
				"unexpected mocks")
		})

	t.Run("Size should call ser.Size and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV    byte = 1
				wantSize int  = 1
				ptrMap        = com.NewPtrMap()
				ptrSer        = mock.NewSerializer[byte]().RegisterSize(
					func(v byte) (size int) {
						ptrMap.Put(unsafe.Pointer(&v))
						return wantSize
					},
				)
				ser   = Wrap(ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			size := ser.Size(wantV)
			assertfatal.Equal(t, size, wantSize,
				fmt.Sprintf("unexpected size, want %v, actual %v", wantSize, size))
			assertfatal.Equal(t, ptrMap.Len(), 0, "ptrMap should be empty")
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap,
				"unexpected mocks")
		})

	t.Run("Skip should call ser.Skip and empty the revPtrMap",
		func(t *testing.T) {
			var (
				wantV     byte = 1
				wantN     int  = 1
				revPtrMap      = com.NewReversePtrMap()
				ptrSer         = mock.NewSerializer[byte]().RegisterSkip(
					func(bs []byte) (n int, err error) {
						v := bs[0]
						assertfatal.Equal(t, v, wantV,
							fmt.Sprintf("unexpected v, want %v, actual %v", wantV, v))
						revPtrMap.Put(1, unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser   = Wrap(nil, revPtrMap, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			n, err := ser.Skip([]byte{wantV})
			assertfatal.Equal(t, n, wantN,
				fmt.Sprintf("unexpected n, want %v, actual %v", wantN, n))
			assertfatal.EqualError(t, err, nil, "unexpected error")
			assertfatal.Equal(t, revPtrMap.Len(), 0, "revPtrMap should be empty")
			asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap,
				"unexpected mocks")
		})
}

func newPtrStructSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer mus.Serializer[int],
) mus.Serializer[ctest.PtrStruct] {
	return ptrStructSer{NewPtrSer(ptrMap, revPtrMap, baseSer)}
}

type ptrStructSer struct {
	intPtrSer mus.Serializer[*int]
}

func (s ptrStructSer) Marshal(v ctest.PtrStruct, bs []byte) (n int) {
	n = s.intPtrSer.Marshal(v.A1, bs)
	n += s.intPtrSer.Marshal(v.A2, bs[n:])
	n += s.intPtrSer.Marshal(v.A3, bs[n:])
	return
}

func (s ptrStructSer) Unmarshal(bs []byte) (v ctest.PtrStruct, n int, err error) {
	v.A1, n, err = s.intPtrSer.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	v.A2, n1, err = s.intPtrSer.Unmarshal(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.A3, n1, err = s.intPtrSer.Unmarshal(bs[n:])
	n += n1
	return
}

func (s ptrStructSer) Size(v ctest.PtrStruct) (size int) {
	size = s.intPtrSer.Size(v.A1)
	size += s.intPtrSer.Size(v.A2)
	size += s.intPtrSer.Size(v.A3)

	return
}

func (s ptrStructSer) Skip(bs []byte) (n int, err error) {
	n, err = s.intPtrSer.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.intPtrSer.Skip(bs[n:])
	n += n1
	if err != nil {
		return
	}
	n1, err = s.intPtrSer.Skip(bs[n:])
	n += n1
	return
}
