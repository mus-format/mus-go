package dts

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go/varint"
)

// DTMSer serializes DTM values. It implements the mus.Serializer[com.DTM]
// interface.
var DTMSer = dtmSer{}

type dtmSer struct{}

func (s dtmSer) Marshal(dtm com.DTM, bs []byte) (n int) {
	return varint.PositiveInt.Marshal(int(dtm), bs)
}

func (s dtmSer) Unmarshal(bs []byte) (dtm com.DTM, n int, err error) {
	num, n, err := varint.PositiveInt.Unmarshal(bs)
	if err != nil {
		return
	}
	dtm = com.DTM(num)
	return
}

func (s dtmSer) Size(dtm com.DTM) (size int) {
	return varint.PositiveInt.Size(int(dtm))
}

func (s dtmSer) Skip(bs []byte) (n int, err error) {
	return varint.PositiveInt.Skip(bs)
}
