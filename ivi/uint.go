package ivarint

import "math/bits"

// 1 					127											1 << 7
// 01					16383										1 << 14
// 001				2097151									1 << 21
// 0001				268435455								1 << 28
// 00001			34359738367							1 << 35
// 000001			4398046511103						1 << 42
// 0000001		562949953421311					1 << 49
// 00000001		72057594037927935				1 << 56

const (
	B1_threshold uint64 = 1 << 7
	B2_threshold uint64 = 1 << 14
	B3_threshold uint64 = 1 << 21
	B4_threshold uint64 = 1 << 28
	B5_threshold uint64 = 1 << 35
	B6_threshold uint64 = 1 << 42
	B7_threshold uint64 = 1 << 49
	B8_threshold uint64 = 1 << 56
)

const (
	B1_marshal_mask byte = 0x80
	B2_marshal_mask byte = 0x40
	B3_marshal_mask byte = 0x20
	B4_marshal_mask byte = 0x10
	B5_marshal_mask byte = 0x8
	B6_marshal_mask byte = 0x4
	B7_marshal_mask byte = 0x2
	B8_marshal_mask byte = 0x1
)

const (
	B1_unmarshal_mask byte = 0x7f
	B2_unmarshal_mask byte = 0x3f
	B3_unmarshal_mask byte = 0x1f
	B4_unmarshal_mask byte = 0xf
	B5_unmarshal_mask byte = 0x7
	B6_unmarshal_mask byte = 0x3
	B7_unmarshal_mask byte = 0x1
	B8_unmarshal_mask byte = 0x0
)

func marshalUint64(t uint64, bs []byte) (n int) {
	switch {
	case t < B1_threshold:
		bs[0] = byte(t) | B1_marshal_mask
		return 1
	case t < B2_threshold:
		bs[0] = byte(t>>8) | B2_marshal_mask
		bs[1] = byte(t)
		return 2
	case t < B3_threshold:
		bs[0] = byte(t>>16) | B3_marshal_mask
		bs[1] = byte(t >> 8)
		bs[2] = byte(t)
		return 3
	case t < B4_threshold:
		bs[0] = byte(t>>32) | B4_marshal_mask
		bs[1] = byte(t >> 16)
		bs[2] = byte(t >> 8)
		bs[3] = byte(t)
		return 4
	case t < B5_threshold:
		bs[0] = byte(t>>32) | B5_marshal_mask
		bs[1] = byte(t >> 24)
		bs[2] = byte(t >> 16)
		bs[3] = byte(t >> 8)
		bs[4] = byte(t)
		return 5
	case t < B6_threshold:
		bs[0] = byte(t>>40) | B6_marshal_mask
		bs[1] = byte(t >> 32)
		bs[2] = byte(t >> 24)
		bs[3] = byte(t >> 16)
		bs[4] = byte(t >> 8)
		bs[5] = byte(t)
		return 6
	case t < B7_threshold:
		bs[0] = byte(t>>48) | B7_marshal_mask
		bs[1] = byte(t >> 40)
		bs[2] = byte(t >> 32)
		bs[3] = byte(t >> 24)
		bs[4] = byte(t >> 16)
		bs[5] = byte(t >> 8)
		bs[6] = byte(t)
		return 7
	default:
		bs[0] = B8_marshal_mask
		bs[1] = byte(t >> 56)
		bs[2] = byte(t >> 48)
		bs[3] = byte(t >> 40)
		bs[4] = byte(t >> 32)
		bs[5] = byte(t >> 24)
		bs[6] = byte(t >> 16)
		bs[7] = byte(t >> 8)
		bs[8] = byte(t)
		return 9
	}
}

func unmarshalUint64(bs []byte) (t uint64, n int, err error) {
	switch bits.LeadingZeros8(uint8(bs[0])) {
	case 0:
		t = uint64(bs[0] & B1_unmarshal_mask)
		n = 1
	case 1:
		t = uint64(bs[1])
		t |= uint64(bs[0]&B2_unmarshal_mask) << 8
		n = 2
	case 2:
		t = uint64(bs[2])
		t |= uint64(bs[1]) << 8
		t |= uint64(bs[0]&B3_unmarshal_mask) << 16
		n = 3
	case 3:
		t = uint64(bs[3])
		t |= uint64(bs[2]) << 8
		t |= uint64(bs[1]) << 16
		t |= uint64(bs[0]&B4_unmarshal_mask) << 24
		n = 4
	case 4:
		t = uint64(bs[4])
		t |= uint64(bs[3]) << 8
		t |= uint64(bs[2]) << 16
		t |= uint64(bs[1]) << 24
		t |= uint64(bs[0]&B5_unmarshal_mask) << 32
		n = 5
	case 5:
		t = uint64(bs[5])
		t |= uint64(bs[4]) << 8
		t |= uint64(bs[3]) << 16
		t |= uint64(bs[2]) << 24
		t |= uint64(bs[1]) << 32
		t |= uint64(bs[0]&B6_unmarshal_mask) << 40
		n = 6
	case 6:
		t = uint64(bs[6])
		t |= uint64(bs[5]) << 8
		t |= uint64(bs[4]) << 16
		t |= uint64(bs[3]) << 24
		t |= uint64(bs[2]) << 32
		t |= uint64(bs[1]) << 40
		t |= uint64(bs[0]&B7_unmarshal_mask) << 48
		n = 7
	case 7:
		t = uint64(bs[8])
		t |= uint64(bs[7]) << 8
		t |= uint64(bs[6]) << 16
		t |= uint64(bs[5]) << 24
		t |= uint64(bs[4]) << 32
		t |= uint64(bs[3]) << 40
		t |= uint64(bs[2]) << 48
		t |= uint64(bs[1]) << 56
		n = 9
	}
	return
}
