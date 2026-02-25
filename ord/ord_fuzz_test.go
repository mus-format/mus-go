package ord

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	mapops "github.com/mus-format/mus-go/options/map"
	slops "github.com/mus-format/mus-go/options/slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/testutil"
	"github.com/mus-format/mus-go/varint"
)

const maxLen = 1000

// byte_slice ------------------------------------------------------------------

func FuzzByteSlice(f *testing.F) {
	seeds := [][]byte{
		{},
		{1, 2, 3},
		{0, 255},
	}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v []byte) {
		if len(v) > maxLen {
			v = v[:maxLen]
		}
		testutil.Test[[]byte]([][]byte{v}, ByteSlice, t)
		testutil.TestSkip[[]byte]([][]byte{v}, ByteSlice, t)
	})
}

func FuzzByteSliceUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidByteSliceSer(bslops.WithLenValidator(
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// slice -----------------------------------------------------------------------

func FuzzSlice(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		if len(bs) > maxLen {
			bs = bs[:maxLen]
		}
		v := make([]int, len(bs))
		for i, b := range bs {
			v[i] = int(b)
		}
		ser := NewSliceSer[int](varint.Int)
		testutil.Test[[]int]([][]int{v}, ser, t)
		testutil.TestSkip[[]int]([][]int{v}, ser, t)
	})
}

func FuzzSliceUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidSliceSer[int](varint.Int, slops.WithLenValidator[int](
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// bool ------------------------------------------------------------------------

func FuzzBool(f *testing.F) {
	seeds := []bool{true, false}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v bool) {
		testutil.Test[bool]([]bool{v}, Bool, t)
		testutil.TestSkip[bool]([]bool{v}, Bool, t)
	})
}

func FuzzBoolUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Bool.Unmarshal(bs)
		Bool.Skip(bs)
	})
}

// map -------------------------------------------------------------------------

func FuzzMap(f *testing.F) {
	// We constructed seeds for documentation purpose, but we fuzz from bytes.
	f.Fuzz(func(t *testing.T, b1, b2, b3, b4 byte) {
		v := map[int]int{int(b1): int(b2), int(b3): int(b4)}
		ser := NewMapSer[int, int](varint.Int, varint.Int)
		testutil.Test[map[int]int]([]map[int]int{v}, ser, t)
		testutil.TestSkip[map[int]int]([]map[int]int{v}, ser, t)
	})
}

func FuzzMapUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidMapSer[int, int](varint.Int, varint.Int,
		mapops.WithLenValidator[int, int](
			com.ValidatorFn[int](func(v int) error {
				if v > maxLen {
					return errors.New("too large length")
				}
				return nil
			}),
		))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// string ----------------------------------------------------------------------

func FuzzString(f *testing.F) {
	seeds := []string{"", "hello", "world", "mus-format"}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) > maxLen {
			v = v[:maxLen]
		}
		testutil.Test[string]([]string{v}, String, t)
		testutil.TestSkip[string]([]string{v}, String, t)
	})
}

func FuzzStringUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidStringSer(strops.WithLenValidator(
		com.ValidatorFn[int](func(v int) error {
			if v > maxLen {
				return errors.New("too large length")
			}
			return nil
		}),
	))
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}

// ptr -------------------------------------------------------------------------

func FuzzPtr(f *testing.F) {
	seeds := []int{1, 2, 3}
	for _, seed := range seeds {
		f.Add(true, seed)
	}
	f.Add(false, 0)
	f.Fuzz(func(t *testing.T, isNotNil bool, val int) {
		var v *int
		if isNotNil {
			v = &val
		}
		ser := NewPtrSer[int](varint.Int)
		testutil.Test[*int]([]*int{v}, ser, t)
		testutil.TestSkip[*int]([]*int{v}, ser, t)
	})
}

func FuzzPtrUnmarshal(f *testing.F) {
	ser := NewPtrSer[int](varint.Int)
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}
