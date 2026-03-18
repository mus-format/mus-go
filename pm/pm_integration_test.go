package pm

import (
	"testing"

	com "github.com/mus-format/common-go"
	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestPMIntegration_Wrapper(t *testing.T) {
	t.Run("Wrapped serializer should succeed", func(t *testing.T) {
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			ser       = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap,
				varint.Int))
		)
		test.Test(ctest.PointerMappingTestCases(), ser, t)
	})

	t.Run("We should be able to use same serializer several times", func(t *testing.T) {
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			ser       = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap,
				varint.Int))
			a = 1
			b = 2
			c = 3
			d = 4
			e = 5
			f = 6
		)
		test.Test([]ctest.PtrStruct{{A1: &a, A2: &b, A3: &c}}, ser, t)
		test.TestSkip([]ctest.PtrStruct{{A1: &d, A2: &e, A3: &f}}, ser, t)
	})
}
