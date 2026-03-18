package typed

import (
	"testing"

	"github.com/mus-format/mus-go"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestDTMSer(t *testing.T) {
	t.Run("Marshal should panic if receives too small byte slice",
		func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("should panic")
				}
			}()
			DTMSer.Marshal(1, make([]byte, DTMSer.Size(1)-1))
		})

	t.Run("Unmarshal and Skip should return error if receives too small byte slice",
		func(t *testing.T) {
			var (
				bs  = make([]byte, DTMSer.Size(1)-1)
				err error
			)
			_, _, err = DTMSer.Unmarshal(bs)
			asserterror.EqualError(t, mus.ErrTooSmallByteSlice, err)

			_, err = DTMSer.Skip(bs)
			asserterror.EqualError(t, mus.ErrTooSmallByteSlice, err)
		})
}
