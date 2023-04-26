package mock

import (
	"reflect"

	"github.com/ymz-ncnk/mok"
)

func NewMarshaler[T any]() Marshaler[T] {
	return Marshaler[T]{mok.New("Marshaler")}
}

type Marshaler[T any] struct {
	*mok.Mock
}

func (m Marshaler[T]) RegisterMarshalMUS(fn func(t T, bs []byte) (n int)) Marshaler[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaler[T]) RegisterNMarshalMUS(n int, fn func(t T, bs []byte) (n int)) Marshaler[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaler[T]) MarshalMUS(t T, bs []byte) (n int) {
	var tVal reflect.Value
	if v := reflect.ValueOf(t); (v.Kind() == reflect.Ptr) && v.IsNil() {
		tVal = reflect.Zero(reflect.TypeOf((*T)(nil)).Elem())
	} else {
		tVal = reflect.ValueOf(t)
	}
	result, err := m.Call("MarshalMUS", tVal, bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
