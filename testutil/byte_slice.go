package testutil

import (
	"testing"

	mock "github.com/mus-format/mus-go/testutil/mock"
)

func ByteSliceLenTestData(t *testing.T) (sl []byte, lenSer mock.Serializer[int]) {
	sl = []byte{1, 2, 45, 255, 123, 70, 0, 0}
	var (
		l    = 8
		lBs  = []byte{16}
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
