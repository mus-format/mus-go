package ivarint

import (
	"testing"
)

func TestIVarint(t *testing.T) {
	var (
		num uint64 = 7205759403792111111
		bs         = make([]byte, 9)
	)
	marshalUint64(num, bs)
	anum, _, err := unmarshalUint64(bs)
	if err != nil {
		t.Errorf("unexpected err, want %v actual %v", nil, err)
	}
	if num != anum {
		t.Errorf("unexpected num, want %v actual %v", num, anum)
	}
}
