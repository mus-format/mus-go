package dts

import (
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
	asserterror "github.com/ymz-ncnk/assert/error"
)

const FooDTM com.DTM = 1

type Foo struct {
	Num int
	Str string
}

var FooSer = fooSer{}

type fooSer struct{}

func (s fooSer) Marshal(foo Foo, bs []byte) (n int) {
	n = varint.Int.Marshal(foo.Num, bs)
	n += ord.String.Marshal(foo.Str, bs[n:])
	return
}

func (s fooSer) Unmarshal(bs []byte) (foo Foo, n int, err error) {
	foo.Num, n, err = varint.Int.Unmarshal(bs)
	if err != nil {
		return
	}
	var n1 int
	foo.Str, n1, err = ord.String.Unmarshal(bs[n:])
	n += n1
	return
}

func (s fooSer) Size(foo Foo) (size int) {
	size = varint.Int.Size(foo.Num)
	return size + ord.String.Size(foo.Str)
}

func (s fooSer) Skip(bs []byte) (n int, err error) {
	n, err = varint.Int.Skip(bs)
	if err != nil {
		return
	}
	var n1 int
	n1, err = ord.String.Skip(bs[n:])
	n += n1
	return
}

func TestDTS(t *testing.T) {
	t.Run("Marshal, Unmarshal, Size, Skip methods should succeed",
		func(t *testing.T) {
			var (
				foo    = Foo{Num: 11, Str: "hello world"}
				fooDTS = New[Foo](FooDTM, FooSer)
				bs     = make([]byte, fooDTS.Size(foo))
			)
			n := fooDTS.Marshal(foo, bs)
			asserterror.Equal(t, len(bs), n)

			afoo, n, err := fooDTS.Unmarshal(bs)
			asserterror.EqualError(t, nil, err)
			asserterror.Equal(t, len(bs), n)
			asserterror.EqualDeep(t, foo, afoo)

			n1, err := fooDTS.Skip(bs)
			asserterror.EqualError(t, nil, err)
			asserterror.Equal(t, n, n1)
		})

	t.Run("Marshal, UnmarshalDTM, UnmarshalData, Size, SkipDTM, SkipData methods should succeed",
		func(t *testing.T) {
			var (
				wantDTSize = 1
				foo        = Foo{Num: 11, Str: "hello world"}
				fooDTS     = New[Foo](FooDTM, FooSer)
				bs         = make([]byte, fooDTS.Size(foo))
			)
			n := fooDTS.Marshal(foo, bs)
			asserterror.Equal(t, len(bs), n)

			dtm, n, err := DTMSer.Unmarshal(bs)
			asserterror.EqualError(t, nil, err)
			asserterror.Equal(t, FooDTM, dtm)
			asserterror.Equal(t, wantDTSize, n)

			afoo, n1, err := fooDTS.UnmarshalData(bs[n:])
			asserterror.EqualError(t, nil, err)
			asserterror.EqualDeep(t, foo, afoo)
			asserterror.Equal(t, len(bs)-wantDTSize, n1)

			fooDTS.Marshal(foo, bs)
			n, err = DTMSer.Skip(bs)
			asserterror.EqualError(t, nil, err)

			n1, err = fooDTS.SkipData(bs[n:])
			asserterror.EqualError(t, nil, err)
			asserterror.Equal(t, len(bs)-wantDTSize, n1)
		})

	t.Run("DTM method should return correct DTM", func(t *testing.T) {
		var (
			fooDTS = New[Foo](FooDTM, nil)
			dtm    = fooDTS.DTM()
		)
		asserterror.Equal(t, FooDTM, dtm)
	})

	t.Run("Unamrshal should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantDTSize = 1
				wantErr    = com.NewWrongDTMError(FooDTM, actualDTM)

				bs     = []byte{byte(actualDTM)}
				fooDTS = New[Foo](FooDTM, nil)
			)
			foo, n, err := fooDTS.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.EqualDeep(t, Foo{}, foo)
			asserterror.Equal(t, wantDTSize, n)
		})

	t.Run("Skip should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM = FooDTM + 3

				wantDTSize = 1
				wantErr    = com.NewWrongDTMError(FooDTM, actualDTM)

				dtm    = FooDTM + 3
				bs     = []byte{byte(dtm)}
				fooDTS = New[Foo](FooDTM, nil)
			)
			n, err := fooDTS.Skip(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, wantDTSize, n)
		})

	t.Run("If UnmarshalDTM fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantFoo = Foo{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice

				bs     = []byte{}
				fooDTS = New[Foo](FooDTM, nil)
			)
			foo, n, err := fooDTS.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.EqualDeep(t, wantFoo, foo)
			asserterror.Equal(t, wantN, n)
		})

	t.Run("If UnmarshalDTM fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice

				bs     = []byte{}
				fooDTS = New[Foo](FooDTM, nil)
			)
			n, err := fooDTS.Skip(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, wantN, n)
		})

	t.Run("If varint.PositiveInt.Unmarshal fails with an error, UnmarshalDTM should return it",
		func(t *testing.T) {
			var (
				wantDTM com.DTM = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice

				bs = []byte{}
			)

			dtm, n, err := DTMSer.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, wantDTM, dtm)
			asserterror.Equal(t, wantN, n)
		})
}
