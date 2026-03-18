package test

import (
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/mus-format/mus-go"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func Test[T any](cases []T, ser mus.Serializer[T], t *testing.T) {
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			bs   = make([]byte, size)
			n    int
			v    T
		)
		n = ser.Marshal(cases[i], bs)
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		v, n, err := ser.Unmarshal(bs)
		if err != nil {
			t.Fatal(err)
		}
		if n != size {
			t.Errorf("case '%v', unexpected n, want '%v' actual '%v'", i, size, n)
		}
		if tm, ok := any(v).(time.Time); ok {
			tm1 := any(cases[i]).(time.Time)
			if !tm.Equal(tm1) {
				t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
			}
		} else {
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
			if !reflect.DeepEqual(v, cases[i]) {
				t.Errorf("case '%v', unexpected v, want '%v' actual '%v'", i, cases[i], v)
			}
		}
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
		if err != nil {
			t.Fatal(err)
		}
		if n != len(bs) {
			t.Fatal("skipped not enough")
		}
	}
}

func TestValidation[T any](cases []T, ser mus.Serializer[T], wantErr error,
	t *testing.T,
) {
	for i := range cases {
		var (
			size = ser.Size(cases[i])
			bs   = make([]byte, size)
		)
		ser.Marshal(cases[i], bs)
		_, _, err := ser.Unmarshal(bs)
		asserterror.EqualError(t, err, wantErr)
	}
}
