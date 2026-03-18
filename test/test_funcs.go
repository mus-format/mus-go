package test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/mus-format/mus-go"
	asserterror "github.com/ymz-ncnk/assert/error"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
	"github.com/ymz-ncnk/mok"
)

type UnmarshalResult[T any] struct {
	V   T
	N   int
	Err error
}

type SkipResult struct {
	N   int
	Err error
}

func Test[T any](cases []T, ser mus.Serializer[T], t *testing.T) {
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			bs   = make([]byte, size)
			n    int
			v    T
		)
		n = ser.Marshal(cases[i], bs)
		asserterror.Equal(t, n, size,
			fmt.Sprintf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n))

		v, n, err := ser.Unmarshal(bs)
		assertfatal.EqualError(t, err, nil)
		asserterror.Equal(t, n, size,
			fmt.Sprintf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n))

		if tm, ok := any(v).(time.Time); ok {
			asserterror.Equal(t, tm.Equal(any(cases[i]).(time.Time)), true,
				fmt.Sprintf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v))
			return
		}

		if f64, ok := any(v).(float64); ok {
			if math.Float64bits(f64) == math.Float64bits(any(cases[i]).(float64)) {
				continue
			}
		}
		if f32, ok := any(v).(float32); ok {
			if math.Float32bits(f32) == math.Float32bits(any(cases[i]).(float32)) {
				continue
			}
		}
		asserterror.EqualDeep(t, v, cases[i],
			fmt.Sprintf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v))
	}
}

func TestSkip[T any](cases []T, ser mus.Serializer[T], t *testing.T) {
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			bs   = make([]byte, size)
		)
		ser.Marshal(cases[i], bs)
		n, err := ser.Skip(bs)
		assertfatal.EqualError(t, err, nil,
			fmt.Sprintf("case '%v', unexpected error", i))
		asserterror.Equal(t, n, len(bs),
			fmt.Sprintf("case '%v', skipped not enough", i))
	}
}

func TestValidation[T any](testCase T, ser mus.Serializer[T], wantErr error,
	t *testing.T,
) {
	var (
		size = ser.Size(testCase)
		bs   = make([]byte, size)
	)
	ser.Marshal(testCase, bs)
	_, _, err := ser.Unmarshal(bs)
	asserterror.EqualError(t, err, wantErr, "unexpected error")
}

func TestUnmarshalOnly[T any](bs []byte, ser mus.Serializer[T],
	want UnmarshalResult[T], mocks []*mok.Mock, t *testing.T,
) {
	v, n, err := ser.Unmarshal(bs)
	asserterror.EqualDeep(t, v, want.V, "unexpected v")
	asserterror.Equal(t, n, want.N, "unexpected n")
	asserterror.EqualError(t, err, want.Err, "unexpected err")
	asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap, "unexpected mocks")
}

func TestSkipOnly[T any](bs []byte, ser mus.Serializer[T],
	want SkipResult, mocks []*mok.Mock, t *testing.T,
) {
	n, err := ser.Skip(bs)
	asserterror.Equal(t, n, want.N, "unexpected n")
	asserterror.EqualError(t, err, want.Err, "unexpected err")
	asserterror.EqualDeep(t, mok.CheckCalls(mocks), mok.EmptyInfomap, "unexpected mocks")
}
