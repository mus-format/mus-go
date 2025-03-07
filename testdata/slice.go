package testdata

import (
	"testing"

	mock "github.com/mus-format/mus-go/testdata/mock"
)

func SliceSerData(t *testing.T) (sl []string, elemSer mock.Serializer[string]) {
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
