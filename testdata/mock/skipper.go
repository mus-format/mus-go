package mock

import "github.com/ymz-ncnk/mok"

func NewSkipper() Skipper {
	return Skipper{mok.New("Skipper")}
}

type Skipper struct {
	*mok.Mock
}

func (u Skipper) RegisterSkipMUS(
	fn func(bs []byte) (n int, err error)) Skipper {
	u.Register("SkipperMUS", fn)
	return u
}

func (u Skipper) RegisterNSkipMUS(n int,
	fn func(bs []byte) (n int, err error)) Skipper {
	u.RegisterN("SkipperMUS", n, fn)
	return u
}

func (u Skipper) SkipMUS(bs []byte) (n int, err error) {
	result, err := u.Call("SkipperMUS", bs)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}
