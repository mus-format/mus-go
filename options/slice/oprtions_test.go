package slops

import (
	"testing"

	com_mock "github.com/mus-format/common-go/testdata/mock"
	"github.com/mus-format/mus-go/testdata/mock"
)

func TestOptions(t *testing.T) {
	var (
		o          = Options[any]{}
		wantLenSer = mock.NewSerializer[int]()
		wantLenVl  = com_mock.NewValidator[int]()
		wantElemVl = com_mock.NewValidator[any]()
	)
	Apply([]SetOption[any]{
		WithLenSer[any](wantLenSer),
		WithLenValidator[any](wantLenVl),
		WithElemValidator[any](wantElemVl),
	}, &o)

	if o.LenSer != wantLenSer {
		t.Errorf("unexpected LenSer, want %v actual %v", wantLenSer, o.LenSer)
	}

	if o.LenVl != wantLenVl {
		t.Errorf("unexpected LenVl, want %v actual %v", wantLenVl, o.LenVl)
	}

	if o.ElemVl != wantElemVl {
		t.Errorf("unexpected ElemVl, want %v actual %v", wantElemVl, o.ElemVl)
	}
}
