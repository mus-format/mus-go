package strops

import (
	"testing"

	cmock "github.com/mus-format/common-go/testutil/mock"
	"github.com/mus-format/mus-go/testutil/mock"
)

func TestOptions(t *testing.T) {
	var (
		o          = Options{}
		wantLenSer = mock.NewSerializer[int]()
		wantLenVl  = cmock.NewValidator[int]()
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
