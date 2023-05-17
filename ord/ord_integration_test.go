package ord

import (
	"testing"

	muscom_testdata "github.com/mus-format/mus-common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {

	t.Run("Pointer", func(t *testing.T) {
		var (
			m = func() mus.MarshalerFn[*string] {
				return func(t *string, bs []byte) (n int) {
					return MarshalPtr[string](t, mus.MarshalerFn[string](MarshalString), bs)
				}
			}()
			u = func() mus.UnmarshalerFn[*string] {
				return func(bs []byte) (t *string, n int, err error) {
					return UnmarshalPtr[string](mus.UnmarshalerFn[string](UnmarshalString), bs)
				}
			}()
			s = func() mus.SizerFn[*string] {
				return func(t *string) (size int) {
					return SizePtr[string](t, mus.SizerFn[string](SizeString))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipPtr(mus.SkipperFn(SkipString), bs)
				}
			}()
		)
		testdata.Test[*string](muscom_testdata.PointerTestCases, m, u, s, t)
		testdata.TestSkip[*string](muscom_testdata.PointerTestCases, m, sk, s, t)
	})

	t.Run("Slice", func(t *testing.T) {
		var (
			m = func() mus.MarshalerFn[[]int] {
				return func(t []int, bs []byte) (n int) {
					return MarshalSlice[int](t, mus.MarshalerFn[int](varint.MarshalInt), bs)
				}
			}()
			u = func() mus.UnmarshalerFn[[]int] {
				return func(bs []byte) (t []int, n int, err error) {
					return UnmarshalSlice[int](mus.UnmarshalerFn[int](varint.UnmarshalInt), bs)
				}
			}()
			s = func() mus.SizerFn[[]int] {
				return func(t []int) (size int) {
					return SizeSlice[int](t, mus.SizerFn[int](varint.SizeInt))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipSlice(mus.SkipperFn(varint.SkipInt), bs)
				}
			}()
		)
		testdata.Test[[]int](muscom_testdata.SliceTestCases, m, u, s, t)
		testdata.TestSkip[[]int](muscom_testdata.SliceTestCases, m, sk, s, t)
	})

	t.Run("Map", func(t *testing.T) {
		var (
			m = func() mus.MarshalerFn[map[float32]uint8] {
				return func(t map[float32]uint8, bs []byte) int {
					return MarshalMap[float32, uint8](t,
						mus.MarshalerFn[float32](varint.MarshalFloat32),
						mus.MarshalerFn[uint8](varint.MarshalUint8),
						bs)
				}
			}()
			u = func() mus.UnmarshalerFn[map[float32]uint8] {
				return func(bs []byte) (t map[float32]uint8, n int, err error) {
					return UnmarshalMap[float32, uint8](
						mus.UnmarshalerFn[float32](varint.UnmarshalFloat32),
						mus.UnmarshalerFn[uint8](varint.UnmarshalUint8),
						bs)
				}
			}()
			s = func() mus.SizerFn[map[float32]uint8] {
				return func(t map[float32]uint8) (size int) {
					return SizeMap[float32, uint8](t,
						mus.SizerFn[float32](varint.SizeFloat32),
						mus.SizerFn[uint8](varint.SizeUint8))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipMap(mus.SkipperFn(varint.SkipFloat32),
						mus.SkipperFn(varint.SkipUint8),
						bs)
				}
			}()
		)
		testdata.Test[map[float32]uint8](muscom_testdata.MapTestCases, m, u, s, t)
		testdata.TestSkip[map[float32]uint8](muscom_testdata.MapTestCases, m, sk, s, t)
	})

}
