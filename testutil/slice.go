package testutil

import (
	"testing"

	mock "github.com/mus-format/mus-go/testutil/mock"
)

func SliceTestData(t *testing.T) (sl []string, elemSer mock.Serializer[string]) {
	var (
		aBs = append([]byte{1}, []byte("a")...)
		bBs = append([]byte{1}, []byte("b")...)
		cBs = append([]byte{1}, []byte("c")...)
	)
	sl = []string{"a", "b", "c"}
	elemSer = mock.NewSerializer[string]().
		// unmarshal
		RegisterMarshal(m("a", aBs, t)).
		RegisterMarshal(m("b", bBs, t)).
		RegisterMarshal(m("c", cBs, t)).
		RegisterUnmarshal(u(aBs, "a", t)).
		RegisterUnmarshal(u(bBs, "b", t)).
		RegisterUnmarshal(u(cBs, "c", t)).
		RegisterSize(s("a", 2, t)).
		RegisterSize(s("b", 2, t)).
		RegisterSize(s("c", 2, t)).
		// skip
		RegisterMarshal(m("a", aBs, t)).
		RegisterMarshal(m("b", bBs, t)).
		RegisterMarshal(m("c", cBs, t)).
		RegisterSize(s("a", 2, t)).
		RegisterSize(s("b", 2, t)).
		RegisterSize(s("c", 2, t)).
		RegisterSkip(sk(aBs, t)).
		RegisterSkip(sk(bBs, t)).
		RegisterSkip(sk(cBs, t))
	return
}

func SliceLenTestData(t *testing.T) (sl []string, lenSer mock.Serializer[int],
	elemSer mock.Serializer[string],
) {
	sl, elemSer = SliceTestData(t)
	var (
		l    = len(sl)
		lBs  = []byte{byte(l * 2)}
		size = 1
	)
	lenSer = mock.NewSerializer[int]().
		// unmarshal
		RegisterMarshal(m(l, lBs, t)).
		RegisterUnmarshal(u(lBs, l, t)).
		RegisterSize(s(l, size, t)).
		// skip
		RegisterMarshal(m(l, lBs, t)).
		RegisterUnmarshal(u(lBs, l, t)).
		RegisterSize(s(l, size, t))
	return
}
