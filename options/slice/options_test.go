package slops

import (
	"testing"

	cmock "github.com/mus-format/common-go/testutil/mock"
	"github.com/mus-format/mus-go/testutil/mock"
)

func TestOptions(t *testing.T) {
	var (
		o          = Options[any]{}
		wantLenSer = mock.NewSerializer[int]()
		wantLenVl  = cmock.NewValidator[int]()
		wantElemVl = cmock.NewValidator[any]()
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
