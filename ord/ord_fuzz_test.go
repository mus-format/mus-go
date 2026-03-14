package ord

import (
	"errors"
	"testing"

	com "github.com/mus-format/common-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	mapops "github.com/mus-format/mus-go/options/map"
	slops "github.com/mus-format/mus-go/options/slice"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/test"
	"github.com/mus-format/mus-go/varint"
)

const maxLen = 1000

// byte_slice ------------------------------------------------------------------

func FuzzOrd_ByteSlice(f *testing.F) {
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
		test.Test([][]byte{v}, ByteSlice, t)
		test.TestSkip([][]byte{v}, ByteSlice, t)
	})
}

func FuzzOrd_ByteSliceUnmarshal(f *testing.F) {
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

func FuzzOrd_Slice(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		if len(bs) > maxLen {
			bs = bs[:maxLen]
		}
		v := make([]int, len(bs))
		for i, b := range bs {
			v[i] = int(b)
		}
		ser := NewSliceSer(varint.Int)
		test.Test([][]int{v}, ser, t)
		test.TestSkip([][]int{v}, ser, t)
	})
}

func FuzzOrd_SliceUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidSliceSer(varint.Int, slops.WithLenValidator[int](
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

func FuzzOrd_Bool(f *testing.F) {
	seeds := []bool{true, false}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v bool) {
		test.Test([]bool{v}, Bool, t)
		test.TestSkip([]bool{v}, Bool, t)
	})
}

func FuzzOrd_BoolUnmarshal(f *testing.F) {
	f.Fuzz(func(t *testing.T, bs []byte) {
		Bool.Unmarshal(bs)
		Bool.Skip(bs)
	})
}

// map -------------------------------------------------------------------------

func FuzzOrd_Map(f *testing.F) {
	// We constructed seeds for documentation purpose, but we fuzz from bytes.
	f.Fuzz(func(t *testing.T, b1, b2, b3, b4 byte) {
		v := map[int]int{int(b1): int(b2), int(b3): int(b4)}
		ser := NewMapSer(varint.Int, varint.Int)
		test.Test([]map[int]int{v}, ser, t)
		test.TestSkip([]map[int]int{v}, ser, t)
	})
}

func FuzzOrd_MapUnmarshal(f *testing.F) {
	// We use Valid serializer to avoid OOM during fuzzing.
	ser := NewValidMapSer(varint.Int, varint.Int,
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

func FuzzOrd_String(f *testing.F) {
	seeds := []string{"", "hello", "world", "mus-format"}
	for _, seed := range seeds {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) > maxLen {
			v = v[:maxLen]
		}
		test.Test([]string{v}, String, t)
		test.TestSkip([]string{v}, String, t)
	})
}

func FuzzOrd_StringUnmarshal(f *testing.F) {
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

func FuzzOrd_Ptr(f *testing.F) {
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
		ser := NewPtrSer(varint.Int)
		test.Test([]*int{v}, ser, t)
		test.TestSkip([]*int{v}, ser, t)
	})
}

func FuzzOrd_PtrUnmarshal(f *testing.F) {
	ser := NewPtrSer(varint.Int)
	f.Fuzz(func(t *testing.T, bs []byte) {
		ser.Unmarshal(bs)
		ser.Skip(bs)
	})
}
