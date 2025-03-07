package ord

import (
	"testing"

	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {

	t.Run("pointer", func(t *testing.T) {
		ser := NewPtrSer[string](String)
		testdata.Test[*string](com_testdata.PointerTestCases, ser, t)
		testdata.TestSkip[*string](com_testdata.PointerTestCases, ser, t)
	})

	t.Run("array", func(t *testing.T) {
		ser := NewArraySer[[3]int, int](3, varint.Int)
		testdata.Test[[3]int](com_testdata.ArrayTestCases, ser, t)
		testdata.TestSkip[[3]int](com_testdata.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := NewValidArraySer[[3]int, int](3, varint.Int, nil)
		testdata.Test[[3]int](com_testdata.ArrayTestCases, ser, t)
		testdata.TestSkip[[3]int](com_testdata.ArrayTestCases, ser, t)
	})

	t.Run("slice", func(t *testing.T) {
		ser := NewSliceSer[int](varint.Int)
		testdata.Test[[]int](com_testdata.SliceTestCases, ser, t)
		testdata.TestSkip[[]int](com_testdata.SliceTestCases, ser, t)
	})

	t.Run("valid slice", func(t *testing.T) {
		ser := NewValidSliceSer[int](varint.Int, nil, nil)
		testdata.Test[[]int](com_testdata.SliceTestCases, ser, t)
		testdata.TestSkip[[]int](com_testdata.SliceTestCases, ser, t)
	})

	t.Run("map", func(t *testing.T) {
		ser := NewMapSer[float32, uint8](varint.Float32, varint.Uint8)
		testdata.Test[map[float32]uint8](com_testdata.MapTestCases, ser, t)
		testdata.TestSkip[map[float32]uint8](com_testdata.MapTestCases, ser, t)
	})

	t.Run("valid map", func(t *testing.T) {
		ser := NewValidMapSer[float32, uint8](varint.Float32, varint.Uint8, nil, nil, nil)
		testdata.Test[map[float32]uint8](com_testdata.MapTestCases, ser, t)
		testdata.TestSkip[map[float32]uint8](com_testdata.MapTestCases, ser, t)
	})

}
