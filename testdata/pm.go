package testdata

import (
	"testing"

	com_testdata "github.com/mus-format/common-go/testdata"
	mock "github.com/mus-format/mus-go/testdata/mock"
)

type PointerMappingStruct struct {
	A1 *int
	A2 *int
	A3 *int
}

func PtrStructSerData(t *testing.T) (st com_testdata.PtrStruct, baseSer mock.Serializer[int]) {
	var (
		one = 1
		two = 2

		oneBs   = []byte{2}
		threeBs = []byte{2}
	)
	st = com_testdata.PtrStruct{A1: &one, A2: &one, A3: &two}
	baseSer = mock.NewSerializer[int]().
		// unmarshal
		RegisterMarshal(m(one, oneBs, t)).
		RegisterMarshal(m(two, threeBs, t)).
		RegisterUnmarshal(u(oneBs, one, t)).
		RegisterUnmarshal(u(threeBs, two, t)).
		RegisterSize(s(one, len(oneBs), t)).
		RegisterSize(s(two, len(threeBs), t)).
		// skip
		RegisterMarshal(m(one, oneBs, t)).
		RegisterMarshal(m(two, threeBs, t)).
		RegisterSize(s(one, len(oneBs), t)).
		RegisterSize(s(two, len(threeBs), t)).
		RegisterSkip(sk(oneBs, t)).
		RegisterSkip(sk(threeBs, t))
	return
}
