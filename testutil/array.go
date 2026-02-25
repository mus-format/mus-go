package testutil

import (
	"testing"

	mock "github.com/mus-format/mus-go/testutil/mock"
)

func ArrayTestData(t *testing.T) (ar [3]int, elemSer mock.Serializer[int]) {
	var (
		oneBs   = []byte{1}
		twoBs   = []byte{2}
		threeBs = []byte{3}
	)
	ar = [3]int{1, 2, 3}
	elemSer = mock.NewSerializer[int]().
		// unmarshal
		RegisterMarshal(m(1, oneBs, t)).
		RegisterMarshal(m(2, twoBs, t)).
		RegisterMarshal(m(3, threeBs, t)).
		RegisterUnmarshal(u(oneBs, 1, t)).
		RegisterUnmarshal(u(twoBs, 2, t)).
		RegisterUnmarshal(u(threeBs, 3, t)).
		RegisterSize(s(1, 1, t)).
		RegisterSize(s(2, 1, t)).
		RegisterSize(s(3, 1, t)).
		// skip
		RegisterMarshal(m(1, oneBs, t)).
		RegisterMarshal(m(2, twoBs, t)).
		RegisterMarshal(m(3, threeBs, t)).
		RegisterSize(s(1, 1, t)).
		RegisterSize(s(2, 1, t)).
		RegisterSize(s(3, 1, t)).
		RegisterSkip(sk(oneBs, t)).
		RegisterSkip(sk(twoBs, t)).
		RegisterSkip(sk(threeBs, t))
	return
}

func ArrayLenTestData(t *testing.T) (arr [3]int, lenSer mock.Serializer[int]) {
	arr = [3]int{1, 2, 3}
	var (
		l    = 3
		lBs  = []byte{6}
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
