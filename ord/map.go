package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	mapops "github.com/mus-format/mus-go/options/map"
	"github.com/mus-format/mus-go/varint"
)

// NewMapSer returns a new map serializer with the given key and value
// serializers. To specify a length, key or value validator, use NewValidMapSer
// instead.
func NewMapSer[T comparable, V any](keySer mus.Serializer[T],
	valueSer mus.Serializer[V], ops ...mapops.SetOption[T, V]) mapSer[T, V] {
	o := mapops.Options[T, V]{}
	mapops.Apply(ops, &o)

	return newMapSer(keySer, valueSer, o)
}

// NewValidMapSer returns a new valid map serializer.
func NewValidMapSer[T comparable, V any](keySer mus.Serializer[T],
	valueSer mus.Serializer[V], ops ...mapops.SetOption[T, V]) validMapSer[T, V] {
	o := mapops.Options[T, V]{}
	mapops.Apply(ops, &o)

	var (
		lenVl   com.Validator[int]
		keyVl   com.Validator[T]
		valueVl com.Validator[V]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.KeyVl != nil {
		keyVl = o.KeyVl
	}
	if o.ValueVl != nil {
		valueVl = o.ValueVl
	}

	return validMapSer[T, V]{
		mapSer:  newMapSer(keySer, valueSer, o),
		lenVl:   lenVl,
		keyVl:   keyVl,
		valueVl: valueVl,
	}
}

func newMapSer[T comparable, V any](keySer mus.Serializer[T],
	valueSer mus.Serializer[V], o mapops.Options[T, V]) mapSer[T, V] {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return mapSer[T, V]{lenSer, keySer, valueSer}
}

type mapSer[T comparable, V any] struct {
	lenSer   mus.Serializer[int]
	keySer   mus.Serializer[T]
	valueSer mus.Serializer[V]
}

// Marshal fills bs with an encoded map value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s mapSer[T, V]) Marshal(v map[T]V, bs []byte) (n int) {
	n = s.lenSer.Marshal(len(v), bs)
	for k, v := range v {
		n += s.keySer.Marshal(k, bs[n:])
		n += s.valueSer.Marshal(v, bs[n:])
	}
	return
}

// Unmarshal parses an encoded map value from bs.
//
// In addition to the map value and the number of used bytes, it may also return
// com.ErrNegativeLength, or a length/key/value unmarshalling error.
func (s mapSer[T, V]) Unmarshal(bs []byte) (v map[T]V, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1  int
		k   T
		val V
	)
	v = make(map[T]V, length)
	for i := 0; i < length; i++ {
		k, n1, err = s.keySer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		val, n1, err = s.valueSer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		v[k] = val
	}
	return
}

// Size returns the size of an encoded map value.
func (s mapSer[T, V]) Size(v map[T]V) (size int) {
	size = s.lenSer.Size(len(v))
	for k, v := range v {
		size += s.keySer.Size(k)
		size += s.valueSer.Size(v)
	}
	return
}

// Skip skips an encoded map value.
//
// In addition to the number of skipped bytes, it may also return
// com.ErrNegativeLength, a length unmarshalling error, or a key/value skipping
// error.
func (s mapSer[T, V]) Skip(bs []byte) (n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var n1 int
	for i := 0; i < length; i++ {
		n1, err = s.keySer.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
		n1, err = s.valueSer.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// -----------------------------------------------------------------------------

type validMapSer[T comparable, V any] struct {
	mapSer[T, V]
	lenVl   com.Validator[int]
	keyVl   com.Validator[T]
	valueVl com.Validator[V]
}

// Unmarshal parses an encoded map value from bs.
//
// In addition to the map value and the number of used bytes, it may also return
// com.ErrNegativeLength, a length/key/value unmarshalling error, or a
// length/key/value validation error.
func (s validMapSer[T, V]) Unmarshal(bs []byte) (v map[T]V, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	var (
		n1  int
		k   T
		val V
	)
	v = make(map[T]V, length)
	for i := 0; i < length; i++ {
		k, n1, err = s.keySer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if s.keyVl != nil {
			if err = s.keyVl.Validate(k); err != nil {
				return
			}
		}
		val, n1, err = s.valueSer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if s.valueVl != nil {
			if err = s.valueVl.Validate(val); err != nil {
				return
			}
		}
		v[k] = val
	}
	return
}
