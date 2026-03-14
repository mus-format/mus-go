package ord

import (
	"testing"

	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestOrdIntegration_Pointer(t *testing.T) {
	ser := NewPtrSer(String)
	test.Test(ctest.PointerTestCases, ser, t)
	test.TestSkip(ctest.PointerTestCases, ser, t)
}

func TestOrdIntegration_Slice(t *testing.T) {
	ser := NewSliceSer(varint.Int)
	test.Test(ctest.SliceTestCases, ser, t)
	test.TestSkip(ctest.SliceTestCases, ser, t)
}

func TestOrdIntegration_ValidSlice(t *testing.T) {
	ser := NewValidSliceSer(varint.Int, nil, nil)
	test.Test(ctest.SliceTestCases, ser, t)
	test.TestSkip(ctest.SliceTestCases, ser, t)
}

func TestOrdIntegration_Map(t *testing.T) {
	ser := NewMapSer(varint.Float32, varint.Uint8)
	test.Test(ctest.MapTestCases, ser, t)
	test.TestSkip(ctest.MapTestCases, ser, t)
}

func TestOrdIntegration_ValidMap(t *testing.T) {
	ser := NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
	test.Test(ctest.MapTestCases, ser, t)
	test.TestSkip(ctest.MapTestCases, ser, t)
}
