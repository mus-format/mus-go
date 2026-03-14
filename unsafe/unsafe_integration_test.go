package unsafe

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationUnsafe(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		ser := NewArraySer[[3]int](varint.Int)
		test.Test(ctestutil.ArrayTestCases, ser, t)
		test.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := NewValidArraySer[[3]int](varint.Int, nil)
		test.Test(ctestutil.ArrayTestCases, ser, t)
		test.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})
}
