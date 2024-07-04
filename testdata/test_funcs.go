package testdata

import (
	"reflect"
	"testing"

	"github.com/mus-format/mus-go"
)

func Test[T any](cases []T, m mus.Marshaller[T], u mus.Unmarshaller[T],
	s mus.Sizer[T],
	t *testing.T,
) {
	for i := 0; i < len(cases); i++ {
		var (
			size = s.Size(cases[i])
			bs   = make([]byte, size)
			n    int
			v    T
		)
		n = m.Marshal(cases[i], bs)
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		v, n, err := u.Unmarshal(bs)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		if !reflect.DeepEqual(v, cases[i]) {
			t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
		}
	}
}

func TestSkip[T any](cases []T, m mus.Marshaller[T], sk mus.Skipper,
	s mus.Sizer[T],
	t *testing.T,
) {
	for i := 0; i < len(cases); i++ {
		bs := make([]byte, s.Size(cases[i]))
		m.Marshal(cases[i], bs)
		n, err := sk.Skip(bs)
		if err != nil {
			t.Fatal(err)
		}
		if n != len(bs) {
			t.Fatal("skipped not enough")
		}
	}
}
