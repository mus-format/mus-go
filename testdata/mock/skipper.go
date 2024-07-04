package mock

import "github.com/ymz-ncnk/mok"

func NewSkipper() Skipper {
	return Skipper{mok.New("Skipper")}
}

type Skipper struct {
	*mok.Mock
}

func (u Skipper) RegisterSkip(
	fn func(bs []byte) (n int, err error)) Skipper {
	u.Register("Skip", fn)
	return u
}

func (u Skipper) RegisterNSkip(n int,
	fn func(bs []byte) (n int, err error)) Skipper {
	u.RegisterN("Skip", n, fn)
	return u
}

func (u Skipper) Skip(bs []byte) (n int, err error) {
	result, err := u.Call("Skip", bs)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
