package testutil

import (
	"testing"

	mock "github.com/mus-format/mus-go/testutil/mock"
)

func StringLenTestData(t *testing.T) (str string, lenSer mock.Serializer[int]) {
	str = "abc"
	var (
		l    = len(str)
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
