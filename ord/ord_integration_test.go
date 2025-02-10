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
			m  = NewPtrMarshallerFn[string](mus.MarshallerFn[string](MarshalStr))
			u  = NewPtrUnmarshallerFn[string](mus.UnmarshallerFn[string](UnmarshalStr))
			s  = NewPtrSizerFn[string](mus.SizerFn[string](SizeStr))
			sk = NewPtrSkipperFn(mus.SkipperFn(SkipStr))
		)
		testdata.Test[*string](com_testdata.PointerTestCases, m, u, s, t)
		testdata.TestSkip[*string](com_testdata.PointerTestCases, m, sk, s, t)
	})

	t.Run("slice", func(t *testing.T) {
		var (
			m = NewSliceMarshallerFn[int](nil,
				mus.MarshallerFn[int](varint.MarshalInt))
			u = NewSliceUnmarshallerFn[int](nil,
				mus.UnmarshallerFn[int](varint.UnmarshalInt))
			s  = NewSliceSizerFn[int](nil, mus.SizerFn[int](varint.SizeInt))
			sk = NewSliceSkipperFn(nil, mus.SkipperFn(varint.SkipInt))
		)
		testdata.Test[[]int](com_testdata.SliceTestCases, m, u, s, t)
		testdata.TestSkip[[]int](com_testdata.SliceTestCases, m, sk, s, t)
	})

	t.Run("map", func(t *testing.T) {
		var (
			m = NewMapMarshallerFn[float32, uint8](nil,
				mus.MarshallerFn[float32](varint.MarshalFloat32),
				mus.MarshallerFn[uint8](varint.MarshalUint8))

			u = NewMapUnmarshallerFn[float32, uint8](nil,
				mus.UnmarshallerFn[float32](varint.UnmarshalFloat32),
				mus.UnmarshallerFn[uint8](varint.UnmarshalUint8))

			s = NewMapSizerFn[float32, uint8](nil,
				mus.SizerFn[float32](varint.SizeFloat32),
				mus.SizerFn[uint8](varint.SizeUint8))

			sk = NewMapSkipperFn[float32, uint8](nil,
				mus.SkipperFn(varint.SkipFloat32),
				mus.SkipperFn(varint.SkipUint8))
		)
		testdata.Test[map[float32]uint8](com_testdata.MapTestCases, m, u, s, t)
		testdata.TestSkip[map[float32]uint8](com_testdata.MapTestCases, m, sk, s, t)
	})

}
