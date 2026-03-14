package ord

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		ser := NewPtrSer(String)
		test.Test(ctestutil.PointerTestCases, ser, t)
		test.TestSkip(ctestutil.PointerTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := NewSliceSer(varint.Int)
		test.Test(ctestutil.SliceTestCases, ser, t)
		test.Test(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := NewValidSliceSer(varint.Int, nil, nil)
		test.Test(ctestutil.SliceTestCases, ser, t)
		test.Test(ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := NewMapSer(varint.Float32, varint.Uint8)
		test.Test(ctestutil.MapTestCases, ser, t)
		test.Test(ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
		test.Test(ctestutil.MapTestCases, ser, t)
		test.Test(ctestutil.MapTestCases, ser, t)
	})
}
