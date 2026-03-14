package ord

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		ser := NewPtrSer(String)
		testutil.Test(ctestutil.PointerTestCases, ser, t)
		testutil.TestSkip(ctestutil.PointerTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := NewSliceSer(varint.Int)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := NewValidSliceSer(varint.Int, nil, nil)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
		testutil.Test(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := NewMapSer(varint.Float32, varint.Uint8)
		testutil.Test(ctestutil.MapTestCases, ser, t)
		testutil.Test(ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
		testutil.Test(ctestutil.MapTestCases, ser, t)
		testutil.Test(ctestutil.MapTestCases, ser, t)
	})
}
