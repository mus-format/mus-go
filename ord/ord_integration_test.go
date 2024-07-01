package ord

import (
	"testing"

	com_testdata "github.com/mus-format/common-go/testdata"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/testdata"
	"github.com/mus-format/mus-go/varint"
)

func TestIntegrationOrd(t *testing.T) {

	t.Run("pointer", func(t *testing.T) {
		var (
			m = func() mus.MarshallerFn[*string] {
				return func(t *string, bs []byte) (n int) {
					return MarshalPtr[string](t, mus.MarshallerFn[string](MarshalStringVarint), bs)
				}
			}()
			u = func() mus.UnmarshallerFn[*string] {
				return func(bs []byte) (t *string, n int, err error) {
					return UnmarshalPtr[string](mus.UnmarshallerFn[string](UnmarshalStringVarint), bs)
				}
			}()
			s = func() mus.SizerFn[*string] {
				return func(t *string) (size int) {
					return SizePtr[string](t, mus.SizerFn[string](SizeStringVarint))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipPtr(mus.SkipperFn(SkipStringVarint), bs)
				}
			}()
		)
		testdata.Test[*string](com_testdata.PointerTestCases, m, u, s, t)
		testdata.TestSkip[*string](com_testdata.PointerTestCases, m, sk, s, t)
	})

	t.Run("slice", func(t *testing.T) {
		var (
			m = func() mus.MarshallerFn[[]int] {
				return func(t []int, bs []byte) (n int) {
					return MarshalSliceVarint[int](t,
						mus.MarshallerFn[int](varint.MarshalInt), bs)
				}
			}()
			u = func() mus.UnmarshallerFn[[]int] {
				return func(bs []byte) (t []int, n int, err error) {
					return UnmarshalSliceVarint[int](
						mus.UnmarshallerFn[int](varint.UnmarshalInt), bs)
				}
			}()
			s = func() mus.SizerFn[[]int] {
				return func(t []int) (size int) {
					return SizeSliceVarint[int](t, mus.SizerFn[int](varint.SizeInt))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipSliceVarint(mus.SkipperFn(varint.SkipInt), bs)
				}
			}()
		)
		testdata.Test[[]int](com_testdata.SliceTestCases, m, u, s, t)
		testdata.TestSkip[[]int](com_testdata.SliceTestCases, m, sk, s, t)
	})

	t.Run("map", func(t *testing.T) {
		var (
			m = func() mus.MarshallerFn[map[float32]uint8] {
				return func(t map[float32]uint8, bs []byte) int {
					return MarshalMapVarint[float32, uint8](t,
						mus.MarshallerFn[float32](varint.MarshalFloat32),
						mus.MarshallerFn[uint8](varint.MarshalUint8),
						bs)
				}
			}()
			u = func() mus.UnmarshallerFn[map[float32]uint8] {
				return func(bs []byte) (t map[float32]uint8, n int, err error) {
					return UnmarshalMapVarint[float32, uint8](
						mus.UnmarshallerFn[float32](varint.UnmarshalFloat32),
						mus.UnmarshallerFn[uint8](varint.UnmarshalUint8),
						bs)
				}
			}()
			s = func() mus.SizerFn[map[float32]uint8] {
				return func(t map[float32]uint8) (size int) {
					return SizeMapVarint[float32, uint8](t,
						mus.SizerFn[float32](varint.SizeFloat32),
						mus.SizerFn[uint8](varint.SizeUint8))
				}
			}()
			sk = func() mus.SkipperFn {
				return func(bs []byte) (n int, err error) {
					return SkipMapVarint(mus.SkipperFn(varint.SkipFloat32),
						mus.SkipperFn(varint.SkipUint8),
						bs)
				}
			}()
		)
		testdata.Test[map[float32]uint8](com_testdata.MapTestCases, m, u, s, t)
		testdata.TestSkip[map[float32]uint8](com_testdata.MapTestCases, m, sk, s, t)
	})

}
