package testdata

import (
	"testing"

	mock "github.com/mus-format/mus-go/testdata/mock"
)

func ArraySerData(t *testing.T) (ar [3]int, elemSer mock.Serializer[int]) {
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
