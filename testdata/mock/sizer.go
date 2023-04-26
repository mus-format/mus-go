package mock

import (
	"reflect"

	"github.com/ymz-ncnk/mok"
)

func NewSizer[T any]() Sizer[T] {
	return Sizer[T]{mok.New("Sizer")}
}

type Sizer[T any] struct {
	*mok.Mock
}

func (m Sizer[T]) RegisterSizeMUS(fn func(t T) (size int)) Sizer[T] {
	m.Register("SizeMUS", fn)
	return m
}

func (m Sizer[T]) RegisterNSizeMUS(n int, fn func(t T) (size int)) Sizer[T] {
	m.RegisterN("SizeMUS", n, fn)
	return m
}

func (m Sizer[T]) SizeMUS(t T) (size int) {
	var tVal reflect.Value
	if v := reflect.ValueOf(t); (v.Kind() == reflect.Ptr) && v.IsNil() {
		tVal = reflect.Zero(reflect.TypeOf((*T)(nil)).Elem())
	} else {
		tVal = reflect.ValueOf(t)
	}
	result, err := m.Call("SizeMUS", tVal)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
