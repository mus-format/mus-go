package testdata

import (
	"bytes"
	"testing"

	mock "github.com/mus-format/mus-go/testdata/mock"
)

func m[T comparable](wantV T, r []byte, t *testing.T) mock.MarshalFn[T] {
	return func(v T, bs []byte) (n int) {
		if v == wantV {
			return copy(bs, r)
		}
		t.Fatalf("ser.Marshal: unexepcted value, want %v actual %v", wantV, v)
		return
	}
}

func u[T any](wantBs []byte, r T, t *testing.T) mock.UnmarshalFn[T] {
	return func(bs []byte) (v T, n int, err error) {
		if bytes.Equal(bs[:len(wantBs)], wantBs) {
			return r, len(wantBs), nil
		}
		t.Fatalf("ser.Unmarshal: unexepcted bs, want '%v' actual '%v'",
			wantBs, bs)
		return
	}
}

func s[T comparable](wantV T, r int, t *testing.T) mock.SizeFn[T] {
	return func(v T) (size int) {
		if v == wantV {
			return r
		}
		t.Fatalf("ser.Size: unexepcted value, want %v actual %v", wantV, v)
		return
	}
}

func sk(wantBs []byte, t *testing.T) mock.SkipFn {
	return func(bs []byte) (n int, err error) {
		if bytes.Equal(bs[:len(wantBs)], wantBs) {
			return len(wantBs), nil
		}
		t.Fatalf("ser.Skip: unexepcted bs, want '%v' actual '%v'",
			wantBs, bs)
		return
	}
}
