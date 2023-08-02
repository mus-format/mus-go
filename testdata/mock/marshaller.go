package mock

import (
	"reflect"

	"github.com/ymz-ncnk/mok"
)

func NewMarshaller[T any]() Marshaller[T] {
	return Marshaller[T]{mok.New("Marshaller")}
}

type Marshaller[T any] struct {
	*mok.Mock
}

func (m Marshaller[T]) RegisterMarshalMUS(fn func(t T, bs []byte) (n int)) Marshaller[T] {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller[T]) RegisterNMarshalMUS(n int, fn func(t T, bs []byte) (n int)) Marshaller[T] {
	m.RegisterN("MarshalMUS", n, fn)
	return m
}

func (m Marshaller[T]) MarshalMUS(t T, bs []byte) (n int) {
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
