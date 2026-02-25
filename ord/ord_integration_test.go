package ord

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		ser := NewPtrSer[string](String)
		testutil.Test[*string](ctestutil.PointerTestCases, ser, t)
		testutil.TestSkip[*string](ctestutil.PointerTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := NewSliceSer[int](varint.Int)
		testutil.Test[[]int](ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip[[]int](ctestutil.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := NewValidSliceSer[int](varint.Int, nil, nil)
		testutil.Test[[]int](ctestutil.SliceTestCases, ser, t)
		testutil.TestSkip[[]int](ctestutil.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := NewMapSer[float32, uint8](varint.Float32, varint.Uint8)
		testutil.Test[map[float32]uint8](ctestutil.MapTestCases, ser, t)
		testutil.TestSkip[map[float32]uint8](ctestutil.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := NewValidMapSer[float32, uint8](varint.Float32, varint.Uint8, nil, nil, nil)
		testutil.Test[map[float32]uint8](ctestutil.MapTestCases, ser, t)
		testutil.TestSkip[map[float32]uint8](ctestutil.MapTestCases, ser, t)
	})
}
