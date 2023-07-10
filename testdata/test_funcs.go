package testdata

import (
	"reflect"
	"testing"

	"github.com/mus-format/mus-go"
)

func Test[T any](cases []T, m mus.Marshaller[T], u mus.UnMarshaller[T],
	s mus.Sizer[T],
	t *testing.T,
) {
	for i := 0; i < len(cases); i++ {
		var (
			size = s.SizeMUS(cases[i])
			bs   = make([]byte, size)
			n    int
			v    T
		)
		n = m.MarshalMUS(cases[i], bs)
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		v, n, err := u.UnmarshalMUS(bs)
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
		bs := make([]byte, s.SizeMUS(cases[i]))
		m.MarshalMUS(cases[i], bs)
		n, err := sk.SkipMUS(bs)
		if err != nil {
			t.Fatal(err)
		}
		if n != len(bs) {
			t.Fatal("skipped not enough")
		}
	}
}
