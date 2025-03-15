package mapops

import (
	"testing"

	com_mock "github.com/mus-format/common-go/testdata/mock"
	"github.com/mus-format/mus-go/testdata/mock"
)

func TestOptions(t *testing.T) {
	var (
		o           = Options[any, any]{}
		wantLenSer  = mock.NewSerializer[int]()
		wantLenVl   = com_mock.NewValidator[int]()
		wantKeyVl   = com_mock.NewValidator[any]()
		wantValueVl = com_mock.NewValidator[any]()
	)
	Apply([]SetOption[any, any]{
		WithLenSer[any, any](wantLenSer),
		WithLenValidator[any, any](wantLenVl),
		WithKeyValidator[any, any](wantKeyVl),
		WithValueValidator[any, any](wantValueVl),
	}, &o)

	if o.LenSer != wantLenSer {
		t.Errorf("unexpected LenSer, want %v actual %v", wantLenSer, o.LenSer)
	}

	if o.LenVl != wantLenVl {
		t.Errorf("unexpected LenVl, want %v actual %v", wantLenVl, o.LenVl)
	}

	if o.KeyVl != wantKeyVl {
		t.Errorf("unexpected KeyVl, want %v actual %v", wantKeyVl, o.KeyVl)
	}

	if o.ValueVl != wantValueVl {
		t.Errorf("unexpected ValueVl, want %v actual %v", wantValueVl, o.ValueVl)
	}
}
