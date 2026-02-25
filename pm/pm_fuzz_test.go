package pm

import (
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

// ptr -------------------------------------------------------------------------

func FuzzPtr(f *testing.F) {
	f.Add(true, 1, 1)  // nil
	f.Add(false, 1, 1) // single pointer
	f.Add(false, 2, 1) // two pointers to same value
	f.Add(false, 2, 2) // two pointers to different values

	f.Fuzz(func(t *testing.T, isNil bool, ptrsCount int, valuesCount int) {
		if ptrsCount < 0 || ptrsCount > 10 {
			return
		}
		if valuesCount < 1 || valuesCount > 10 {
			return
		}

		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			pSer      = NewPtrSer[int](ptrMap, revPtrMap, varint.Int)
			ser       = Wrap[*int](ptrMap, revPtrMap, pSer)
			values    = make([]int, valuesCount)
			ptrs      = make([]*int, ptrsCount)
		)
		for i := 0; i < valuesCount; i++ {
			values[i] = i
		}
		if isNil {
			for i := 0; i < ptrsCount; i++ {
				ptrs[i] = nil
			}
		} else {
			for i := 0; i < ptrsCount; i++ {
				ptrs[i] = &values[i%valuesCount]
			}
		}

		for _, p := range ptrs {
			testutil.Test[*int]([]*int{p}, ser, t)
			testutil.TestSkip[*int]([]*int{p}, ser, t)
		}
	})
}

func FuzzPtrUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			pSer      = NewPtrSer[int](ptrMap, revPtrMap, varint.Int)
			ser       = Wrap[*int](ptrMap, revPtrMap, pSer)
		)
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// wrap ------------------------------------------------------------------------

type node struct {
	Value int
	Next  *node
}

type nodeSer struct {
	ptrSer mus.Serializer[*node]
}

func (s *nodeSer) Marshal(v node, bs []byte) (n int) {
	n = varint.Int.Marshal(v.Value, bs)
	return n + s.ptrSer.Marshal(v.Next, bs[n:])
}

func (s *nodeSer) Unmarshal(bs []byte) (v node, n int, err error) {
	v.Value, n, err = varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	var next *node
	var n1 int
	next, n1, err = s.ptrSer.Unmarshal(bs[n:])
	n += n1
	v.Next = next
	return
}

func (s *nodeSer) Size(v node) (size int) {
	return varint.Int.Size(v.Value) + s.ptrSer.Size(v.Next)
}

func (s *nodeSer) Skip(bs []byte) (n int, err error) {
	n, err = varint.Int.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = s.ptrSer.Skip(bs[n:])
	return n + n1, err
}

func FuzzWrap(f *testing.F) {
	f.Add(0, false)
	f.Add(1, false)
	f.Add(5, false)
	f.Add(5, true) // Circular
	f.Fuzz(func(t *testing.T, length int, circular bool) {
		if length < 0 || length > 15 {
			return
		}
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			nSer      = &nodeSer{}
			pSer      = NewPtrSer[node](ptrMap, revPtrMap, nSer)
		)
		nSer.ptrSer = pSer
		ser := Wrap[*node](ptrMap, revPtrMap, pSer)

		var head *node
		if length > 0 {
			head = &node{Value: 0}
			curr := head
			for i := 1; i < length; i++ {
				curr.Next = &node{Value: i}
				curr = curr.Next
			}
			if circular {
				curr.Next = head
			}
		}

		testutil.Test[*node]([]*node{head}, ser, t)
		testutil.TestSkip[*node]([]*node{head}, ser, t)
	})
}

func FuzzWrapUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			nSer      = &nodeSer{}
			pSer      = NewPtrSer[node](ptrMap, revPtrMap, nSer)
		)
		nSer.ptrSer = pSer
		ser := Wrap[*node](ptrMap, revPtrMap, pSer)

		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}
