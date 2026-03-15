package mock

import (
	"github.com/ymz-ncnk/mok"
)

type (
	MarshalMUSFn func(bs []byte) (n int)
	SizeMUSFn    func() (size int)
)

type Marshaller struct {
	*mok.Mock
}

func NewMarshaller() Marshaller {
	return Marshaller{mok.New("Marshaller")}
}

func (m Marshaller) RegisterMarshalMUS(fn MarshalMUSFn) Marshaller {
	m.Register("MarshalMUS", fn)
	return m
}

func (m Marshaller) RegisterSizeMUS(fn SizeMUSFn) Marshaller {
	m.Register("SizeMUS", fn)
	return m
}

func (m Marshaller) MarshalMUS(bs []byte) (n int) {
	result, err := m.Call("MarshalMUS", bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}

func (m Marshaller) SizeMUS() (size int) {
	result, err := m.Call("SizeMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
