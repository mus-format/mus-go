package mus

import (
	"testing"

	"github.com/mus-format/mus-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestMarshal(t *testing.T) {
	var (
		size = 10
		n    = 5
		m    = mock.NewMarshaller().RegisterSizeMUS(
			func() int { return size },
		).RegisterMarshalMUS(
			func(bs []byte) int { return n },
		)
		bs = Marshal(m)
	)
	asserterror.Equal(t, size, len(bs))
}

func TestMarshalTyped(t *testing.T) {
	var (
		size = 10
		n    = 5
		m    = mock.NewMarshallerTyped().RegisterSizeTypedMUS(
			func() int { return size },
		).RegisterMarshalTypedMUS(
			func(bs []byte) int { return n },
		)
		bs = MarshalTyped(m)
	)
	asserterror.Equal(t, size, len(bs))
}
