package unsafe

import (
	"testing"

	ctestutil "github.com/mus-format/common-go/testutil"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestUnsafeIntegration_Array(t *testing.T) {
	ser := NewArraySer[[3]int](varint.Int)
	test.Test(ctestutil.ArrayTestCases, ser, t)
	test.TestSkip(ctestutil.ArrayTestCases, ser, t)
}

func TestUnsafeIntegration_ValidArray(t *testing.T) {
	ser := NewValidArraySer[[3]int](varint.Int, nil)
	test.Test(ctestutil.ArrayTestCases, ser, t)
	test.TestSkip(ctestutil.ArrayTestCases, ser, t)
}
