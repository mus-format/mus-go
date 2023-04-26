package testdata

import (
	"reflect"
	"testing"

	"github.com/mus-format/mus-go"
	"github.com/ymz-ncnk/mok"
)

func Test[T any](cases []T, m mus.Marshaler[T], u mus.Unmarshaler[T],
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

func TestSkip[T any](cases []T, m mus.Marshaler[T], sk mus.Skipper,
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

func TestUnmarshalResults[T any](wantV, v T, wantN, n int, wantErr, err error,
	mocks []*mok.Mock, t *testing.T) {
	if !reflect.DeepEqual(v, wantV) {
		t.Errorf("unexpected v, want '%v' actual '%v'", wantV, v)
	}
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
	if info := mok.CheckCalls(mocks); len(info) > 0 {
		t.Error(info)
	}
}

func TestSkipResults(wantN, n int, wantErr, err error, t *testing.T) {
	if n != wantN {
		t.Errorf("unexpected n, want '%v' actual '%v'", wantN, n)
	}
	if err != wantErr {
		t.Errorf("unexpected error, want '%v' actual '%v'", wantErr, err)
	}
}

func ComparePtrs(t, v any) bool {
	p1 := reflect.ValueOf(t).Pointer()
	p2 := reflect.ValueOf(v).Pointer()
	return p1 == p2
}
