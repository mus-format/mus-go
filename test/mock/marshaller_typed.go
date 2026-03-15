package mock

import (
	"github.com/ymz-ncnk/mok"
)

type (
	MarshalTypedMUSFn func(bs []byte) (n int)
	SizeTypedMUSFn    func() (size int)
)

type MarshallerTyped struct {
	*mok.Mock
}

func NewMarshallerTyped() MarshallerTyped {
	return MarshallerTyped{mok.New("MarshallerTyped")}
}

func (m MarshallerTyped) RegisterMarshalTypedMUS(fn MarshalTypedMUSFn) MarshallerTyped {
	m.Register("MarshalTypedMUS", fn)
	return m
}

func (m MarshallerTyped) RegisterSizeTypedMUS(fn SizeTypedMUSFn) MarshallerTyped {
	m.Register("SizeTypedMUS", fn)
	return m
}

func (m MarshallerTyped) MarshalTypedMUS(bs []byte) (n int) {
	result, err := m.Call("MarshalTypedMUS", bs)
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}

func (m MarshallerTyped) SizeTypedMUS() (size int) {
	result, err := m.Call("SizeTypedMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
