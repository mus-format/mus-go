package ord

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestOrdIntegration_Pointer(t *testing.T) {
	ser := NewPtrSer(String)
	test.Test(ctestutil.PointerTestCases, ser, t)
	test.TestSkip(ctestutil.PointerTestCases, ser, t)
}

func TestOrdIntegration_Slice(t *testing.T) {
	ser := NewSliceSer(varint.Int)
	test.Test(ctestutil.SliceTestCases, ser, t)
	test.TestSkip(ctestutil.SliceTestCases, ser, t)
}

func TestOrdIntegration_ValidSlice(t *testing.T) {
	ser := NewValidSliceSer(varint.Int, nil, nil)
	test.Test(ctestutil.SliceTestCases, ser, t)
	test.TestSkip(ctestutil.SliceTestCases, ser, t)
}

func TestOrdIntegration_Map(t *testing.T) {
	ser := NewMapSer(varint.Float32, varint.Uint8)
	test.Test(ctestutil.MapTestCases, ser, t)
	test.TestSkip(ctestutil.MapTestCases, ser, t)
}

func TestOrdIntegration_ValidMap(t *testing.T) {
	ser := NewValidMapSer(varint.Float32, varint.Uint8, nil, nil, nil)
	test.Test(ctestutil.MapTestCases, ser, t)
	test.TestSkip(ctestutil.MapTestCases, ser, t)
}
