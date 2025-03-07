package testdata

import (
	"testing"

	mock "github.com/mus-format/mus-go/testdata/mock"
)

func PtrSerData(t *testing.T) (ptr *int, baseSer mock.Serializer[int]) {
	var (
		one = 1

		oneBs = []byte{byte(one)}
	)
	baseSer = mock.NewSerializer[int]().
		// unmarshal
		RegisterMarshal(m(one, oneBs, t)).
		RegisterUnmarshal(u(oneBs, one, t)).
		RegisterSize(s(one, 1, t)).

		// skip
		RegisterMarshal(m(one, oneBs, t)).
		RegisterSize(s(one, 1, t)).
		RegisterSkip(sk(oneBs, t))
	return
}
