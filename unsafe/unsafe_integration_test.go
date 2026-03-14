package unsafe

import (
	"testing"

	ctest "github.com/mus-format/common-go/test"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

func TestUnsafeIntegration_Array(t *testing.T) {
	ser := NewArraySer[[3]int](varint.Int)
	test.Test(ctest.ArrayTestCases, ser, t)
	test.TestSkip(ctest.ArrayTestCases, ser, t)
}

func TestUnsafeIntegration_ValidArray(t *testing.T) {
	ser := NewValidArraySer[[3]int](varint.Int, nil)
	test.Test(ctest.ArrayTestCases, ser, t)
	test.TestSkip(ctest.ArrayTestCases, ser, t)
}
