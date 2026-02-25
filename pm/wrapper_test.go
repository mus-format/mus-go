package pm

import (
	"fmt"
	"testing"
	"unsafe"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	mus "github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testutil"
	mock "github.com/mus-format/mus-go/testutil/mock"
	"github.com/ymz-ncnk/mok"
)

func TestWrapper(t *testing.T) {
	t.Run("wrapper serializer should work correctly",
		func(t *testing.T) {
			var (
				st, baseSer = testutil.PtrStructTestData(t)
				ptrMap      = com.NewPtrMap()
				revPtrMap   = com.NewReversePtrMap()
				w           = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap, baseSer))
			)
			testutil.Test[ctestutil.PtrStruct]([]ctestutil.PtrStruct{st}, w, t)
			testutil.TestSkip[ctestutil.PtrStruct]([]ctestutil.PtrStruct{st}, w, t)
		})

	t.Run("Marshal should call ser.Marshal and empty the ptrMap",
		func(t *testing.T) {
			var (
				wantV  byte = 1
				wantN  int  = 1
				ptrMap      = com.NewPtrMap()
				ptrSer      = mock.NewSerializer[byte]().RegisterMarshal(
					func(v byte, bs []byte) (n int) {
						if v != wantV {
							t.Fatalf("unexpected v, want %v, actual %v", wantV, v)
						}
						ptrMap.Put(unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser   = Wrap[byte](ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			n := ser.Marshal(wantV, nil)
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if ptrMap.Len() != 0 {
				t.Fatal("ptrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
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
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						n = wantN
						revPtrMap.Put(1, unsafe.Pointer(&v))
						return
					},
				)
				ser   = Wrap[byte](nil, revPtrMap, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			v, n, err := ser.Unmarshal([]byte{wantV})
			if v != wantV {
				t.Fatalf("unexpected v, want %v, actual %v", wantV, v)
			}
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if revPtrMap.Len() != 0 {
				t.Fatal("revPtrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
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
				ser   = Wrap[byte](ptrMap, nil, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			size := ser.Size(wantV)
			if size != wantSize {
				t.Fatalf("unexpected size, want %v, actual %v", wantSize, size)
			}
			if ptrMap.Len() != 0 {
				t.Fatal("ptrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
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
						if v != wantV {
							err = fmt.Errorf("unexpected v, want %v, actual %v", wantV, v)
							return
						}
						revPtrMap.Put(1, unsafe.Pointer(&v))
						n = wantN
						return
					},
				)
				ser   = Wrap[byte](nil, revPtrMap, ptrSer)
				mocks = []*mok.Mock{ptrSer.Mock}
			)
			n, err := ser.Skip([]byte{wantV})
			if n != wantN {
				t.Fatalf("unexpected n, want %v, actual %v", wantN, n)
			}
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if revPtrMap.Len() != 0 {
				t.Fatal("revPtrMap should be empty")
			}
			if infomap := mok.CheckCalls(mocks); len(infomap) > 0 {
				t.Error(infomap)
			}
		})
}

func newPtrStructSer(ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer mus.Serializer[int],
) mus.Serializer[ctestutil.PtrStruct] {
	return ptrStructSer{NewPtrSer[int](ptrMap, revPtrMap, baseSer)}
}

type ptrStructSer struct {
	intPtrSer mus.Serializer[*int]
}

func (s ptrStructSer) Marshal(v ctestutil.PtrStruct, bs []byte) (n int) {
	n = s.intPtrSer.Marshal(v.A1, bs)
	n += s.intPtrSer.Marshal(v.A2, bs[n:])
	n += s.intPtrSer.Marshal(v.A3, bs[n:])
	return
}

func (s ptrStructSer) Unmarshal(bs []byte) (v ctestutil.PtrStruct, n int, err error) {
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

func (s ptrStructSer) Size(v ctestutil.PtrStruct) (size int) {
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
