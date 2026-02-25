package pm

import (
	"testing"

	com "github.com/mus-format/common-go"
	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationWrapper(t *testing.T) {
	t.Run("wrapper", func(t *testing.T) {
		var (
			ptrMap    = com.NewPtrMap()
			revPtrMap = com.NewReversePtrMap()
			ser       = Wrap(ptrMap, revPtrMap, newPtrStructSer(ptrMap, revPtrMap,
				varint.Int))
		)
		testutil.Test[ctestutil.PtrStruct](
			ctestutil.PointerMappingTestCases(), ser, t)
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
		testutil.Test[ctestutil.PtrStruct](
			[]ctestutil.PtrStruct{{A1: &a, A2: &b, A3: &c}},
			ser, t)

		testutil.TestSkip[ctestutil.PtrStruct](
			[]ctestutil.PtrStruct{{A1: &d, A2: &e, A3: &f}},
			ser, t)
	})
}
