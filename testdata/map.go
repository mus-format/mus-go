package testdata

import (
	"bytes"
	"testing"

	mock "github.com/mus-format/mus-go/testdata/mock"
)

func MapSerData(t *testing.T) (mp map[string]int, keySer mock.Serializer[string],
	valueSer mock.Serializer[int],
) {
	var (
		aBs = append([]byte{1}, []byte("a")...)
		bBs = append([]byte{1}, []byte("b")...)
		cBs = append([]byte{1}, []byte("c")...)

		oneBs   = []byte{1}
		twoBs   = []byte{2}
		threeBs = []byte{3}
	)
	mp = map[string]int{"a": 1, "b": 2, "c": 3}
	keySer = mock.NewSerializer[string]().RegisterMarshalN(6, func(v string, bs []byte) (n int) {
		switch v {
		case "a":
			return m("a", aBs, t)(v, bs)
		case "b":
			return m("b", bBs, t)(v, bs)
		case "c":
			return m("c", cBs, t)(v, bs)
		default:
			t.Fatalf("unexepcted string %v", v)
			return
		}
	}).RegisterUnmarshalN(3, func(bs []byte) (v string, n int, err error) {
		switch {
		case bytes.Equal(bs[:len(aBs)], aBs):
			return u(aBs, "a", t)(bs)
		case bytes.Equal(bs[:len(bBs)], bBs):
			return u(bBs, "b", t)(bs)
		case bytes.Equal(bs[:len(cBs)], cBs):
			return u(cBs, "c", t)(bs)
		default:
			t.Fatalf("unexepcted bs '%v'", bs)
			return
		}
	}).RegisterSizeN(6, func(v string) (size int) {
		switch v {
		case "a":
			return s("a", 2, t)(v)
		case "b":
			return s("b", 2, t)(v)
		case "c":
			return s("c", 2, t)(v)
		default:
			t.Fatalf("unexepcted string %v", v)
		}
		return
	}).RegisterSkipN(3, func(bs []byte) (n int, err error) {
		switch {
		case bytes.Equal(bs[:len(aBs)], aBs):
			return sk(aBs, t)(bs)
		case bytes.Equal(bs[:len(bBs)], bBs):
			return sk(bBs, t)(bs)
		case bytes.Equal(bs[:len(cBs)], cBs):
			return sk(cBs, t)(bs)
		default:
			t.Fatalf("unexepcted bs '%v'", bs)
		}
		return
	})

	valueSer = mock.NewSerializer[int]().RegisterMarshalN(6, func(v int, bs []byte) (n int) {
		switch v {
		case 1:
			return m(1, oneBs, t)(v, bs)
		case 2:
			return m(2, twoBs, t)(v, bs)
		case 3:
			return m(3, threeBs, t)(v, bs)
		default:
			t.Fatalf("unexepcted int %v", v)
		}
		return
	}).RegisterUnmarshalN(3, func(bs []byte) (v int, n int, err error) {
		switch {
		case bytes.Equal(bs[:len(oneBs)], oneBs):
			return u(oneBs, 1, t)(bs)
		case bytes.Equal(bs[:len(twoBs)], twoBs):
			return u(twoBs, 2, t)(bs)
		case bytes.Equal(bs[:len(threeBs)], threeBs):
			return u(threeBs, 3, t)(bs)
		default:
			t.Fatalf("unexepcted bs '%v'", bs)
			return
		}
	}).RegisterSizeN(6, func(v int) (size int) {
		switch v {
		case 1:
			return s(1, 1, t)(v)
		case 2:
			return s(2, 1, t)(v)
		case 3:
			return s(3, 1, t)(v)
		default:
			t.Fatalf("unexepcted int %v", v)
		}
		return
	}).RegisterSkipN(3, func(bs []byte) (n int, err error) {
		switch {
		case bytes.Equal(bs[:len(oneBs)], oneBs):
			return sk(oneBs, t)(bs)
		case bytes.Equal(bs[:len(twoBs)], twoBs):
			return sk(twoBs, t)(bs)
		case bytes.Equal(bs[:len(threeBs)], threeBs):
			return sk(threeBs, t)(bs)
		default:
			t.Fatalf("unexepcted bs '%v'", bs)
		}
		return
	})
	return
}

func MapLenSerData(t *testing.T) (mp map[string]int, lenSer mock.Serializer[int],
	keySer mock.Serializer[string], valueSer mock.Serializer[int],
) {
	mp, keySer, valueSer = MapSerData(t)
	var (
		l    = len(mp)
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
