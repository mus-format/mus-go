package typed

import (
	"testing"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/test/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

const FooDTM com.DTM = 0

type Foo struct {
	Num int
	Str string
}

func TestTyped(t *testing.T) {
	t.Run("Marshal, Unmarshal, Size, Skip methods should succeed",
		func(t *testing.T) {
			var (
				foo  = Foo{Num: 11, Str: "hello world"}
				size = 5
				n    = 5
				ser  = mock.NewSerializer[Foo]().RegisterSize(
					func(foo Foo) int { return size },
				).RegisterMarshal(
					func(foo Foo, bs []byte) int { return n },
				).RegisterUnmarshal(
					func(bs []byte) (Foo, int, error) { return foo, n, nil },
				).RegisterSkip(
					func(bs []byte) (int, error) { return n, nil },
				)
				fooTypedMUS = NewTypedSer(FooDTM, ser)
				bs          = make([]byte, fooTypedMUS.Size(foo))
			)
			// Marshal
			n1 := fooTypedMUS.Marshal(foo, bs)
			asserterror.Equal(t, n+DTMSer.Size(FooDTM), n1)

			// Unmarshal
			afoo, n1, err := fooTypedMUS.Unmarshal(bs)
			asserterror.Equal(t, nil, err)
			asserterror.Equal(t, n+DTMSer.Size(FooDTM), n1)
			asserterror.EqualDeep(t, foo, afoo)

			// Skip
			n1, err = fooTypedMUS.Skip(bs)
			asserterror.Equal(t, nil, err)
			asserterror.Equal(t, n+DTMSer.Size(FooDTM), n1)
		})

	t.Run("Marshal, UnmarshalDTM, UnmarshalData, Size, SkipDTM, SkipData methods should succeed",
		func(t *testing.T) {
			var (
				foo  = Foo{Num: 11, Str: "hello world"}
				size = 5
				n    = 5
				mfn  = func(foo Foo, bs []byte) int { return n }
				ser  = mock.NewSerializer[Foo]().RegisterSize(
					func(foo Foo) int { return size },
				).RegisterMarshalN(2, mfn).RegisterUnmarshal(
					func(bs []byte) (Foo, int, error) { return foo, n, nil },
				).RegisterSkip(
					func(bs []byte) (int, error) { return n, nil },
				)
				fooTypedMUS = NewTypedSer(FooDTM, ser)
				bs          = make([]byte, fooTypedMUS.Size(foo))
			)
			n1 := fooTypedMUS.Marshal(foo, bs)
			asserterror.Equal(t, len(bs), n1)

			dtm, n1, err := DTMSer.Unmarshal(bs)
			asserterror.Equal(t, nil, err)
			asserterror.Equal(t, FooDTM, dtm)
			asserterror.Equal(t, DTMSer.Size(FooDTM), n1)

			afoo, n2, err := fooTypedMUS.UnmarshalData(bs[n1:])
			asserterror.Equal(t, nil, err)
			asserterror.EqualDeep(t, foo, afoo)
			asserterror.Equal(t, n, n2)

			fooTypedMUS.Marshal(foo, bs)
			n1, err = DTMSer.Skip(bs)
			asserterror.Equal(t, nil, err)

			n2, err = fooTypedMUS.SkipData(bs[n1:])
			asserterror.Equal(t, nil, err)
			asserterror.Equal(t, n, n2)
		})

	t.Run("DTM method should return correct DTM", func(t *testing.T) {
		var (
			fooTypedMUS = NewTypedSer[Foo](FooDTM, nil)
			dtm         = fooTypedMUS.DTM()
		)
		asserterror.Equal(t, FooDTM, dtm)
	})

	t.Run("Unmarshal should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM   = FooDTM + 3
				wantErr     = com.NewWrongDTMError(FooDTM, actualDTM)
				bs          = make([]byte, DTMSer.Size(actualDTM))
				fooTypedMUS = NewTypedSer[Foo](FooDTM, nil)
			)
			DTMSer.Marshal(actualDTM, bs)
			foo, n, err := fooTypedMUS.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.EqualDeep(t, Foo{}, foo)
			asserterror.Equal(t, len(bs), n)
		})

	t.Run("Skip should fail with ErrWrongDTM, if meets another DTM",
		func(t *testing.T) {
			var (
				actualDTM   = FooDTM + 3
				wantErr     = com.NewWrongDTMError(FooDTM, actualDTM)
				bs          = make([]byte, DTMSer.Size(actualDTM))
				fooTypedMUS = NewTypedSer[Foo](FooDTM, nil)
			)
			DTMSer.Marshal(actualDTM, bs)
			n, err := fooTypedMUS.Skip(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, len(bs), n)
		})

	t.Run("If UnmarshalDTM fails with an error, Unmarshal should return it",
		func(t *testing.T) {
			var (
				wantFoo = Foo{}
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice

				bs          = []byte{}
				fooTypedMUS = NewTypedSer[Foo](FooDTM, nil)
			)
			foo, n, err := fooTypedMUS.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.EqualDeep(t, wantFoo, foo)
			asserterror.Equal(t, wantN, n)
		})

	t.Run("If UnmarshalDTM fails with an error, Skip should return it",
		func(t *testing.T) {
			var (
				wantN   = 0
				wantErr = mus.ErrTooSmallByteSlice

				bs          = []byte{}
				fooTypedMUS = NewTypedSer[Foo](FooDTM, nil)
			)
			n, err := fooTypedMUS.Skip(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, wantN, n)
		})

	t.Run("If varint.PositiveInt.Unmarshal fails with an error, UnmarshalDTM should return it",
		func(t *testing.T) {
			var (
				wantDTM com.DTM = 0
				wantN           = 0
				wantErr         = mus.ErrTooSmallByteSlice
				bs              = []byte{}
			)
			dtm, n, err := DTMSer.Unmarshal(bs)
			asserterror.EqualError(t, wantErr, err)
			asserterror.Equal(t, wantDTM, dtm)
			asserterror.Equal(t, wantN, n)
		})
}
