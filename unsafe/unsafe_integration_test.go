package unsafe

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationUnsafe(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		ser := NewArraySer[[3]int](varint.Int)
		testutil.Test(ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})

	t.Run("valid array", func(t *testing.T) {
		ser := NewValidArraySer[[3]int](varint.Int, nil)
		testutil.Test(ctestutil.ArrayTestCases, ser, t)
		testutil.TestSkip(ctestutil.ArrayTestCases, ser, t)
	})
}
