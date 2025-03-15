package strops

import (
	"testing"

	com_mock "github.com/mus-format/common-go/testdata/mock"
	"github.com/mus-format/mus-go/testdata/mock"
)

func TestOptions(t *testing.T) {
	var (
		o          = Options{}
		wantLenSer = mock.NewSerializer[int]()
		wantLenVl  = com_mock.NewValidator[int]()
	)
	Apply([]SetOption{
		WithLenSer(wantLenSer),
		WithLenValidator(wantLenVl),
	}, &o)

	if o.LenSer != wantLenSer {
		t.Errorf("unexpected LenSer, want %v actual %v", wantLenSer, o.LenSer)
	}

	if o.LenVl != wantLenVl {
		t.Errorf("unexpected LenVl, want %v actual %v", wantLenVl, o.LenVl)
	}
}
