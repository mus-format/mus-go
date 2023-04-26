package testdata

import "math"

// Unsigned test cases ---------------------------------------------------------
var Uint64TestCases = []uint64{
	0,
	1000,
	math.MaxUint64,
}

var Uint32TestCases = []uint32{
	0,
	1000,
	math.MaxUint32,
}

var Uint16TestCases = []uint16{
	0,
	1000,
	math.MaxUint16,
}

var Uint8TestCases = []uint8{
	0,
	100,
	math.MaxUint8,
}

var UintTestCases = []uint{
	0,
	1000000,
	math.MaxUint,
}

// Signed test cases -----------------------------------------------------------
var Int64TestCases = []int64{
	math.MinInt64,
	-119283019,
	0,
	1000,
	math.MaxInt64,
}

var Int32TestCases = []int32{
	math.MinInt32,
	-1112202828,
	0,
	1000,
	math.MaxInt32,
}

var Int16TestCases = []int16{
	math.MinInt16,
	-1000,
	0,
	1000,
	math.MaxInt16,
}

var Int8TestCases = []int8{
	math.MinInt8,
	-10,
	0,
	100,
	math.MaxInt8,
}

var IntTestCases = []int{
	math.MinInt,
	-127637,
	0,
	100,
	math.MaxInt,
}

// Byte test cases -------------------------------------------------------------
var ByteTestCases = []byte{
	0,
	19,
	255,
}

// Float test cases ------------------------------------------------------------
var Float64TestCases = []float64{
	-math.MaxFloat64,
	-math.SmallestNonzeroFloat64,
	0,
	math.SmallestNonzeroFloat64,
	100,
	math.MaxFloat64,
}

var Float32TestCases = []float32{
	-math.MaxFloat32,
	-math.SmallestNonzeroFloat32,
	0,
	math.SmallestNonzeroFloat32,
	100,
	math.MaxFloat32,
}

// Bool test cases -------------------------------------------------------------
var BoolTestCases = []bool{
	true,
	false,
}

// String test cases -----------------------------------------------------------
var StringTestCases = []string{
	"",
	"alkjsdlfkj",
}

// Pointer test cases ----------------------------------------------------------
var PointerTestCases = func() []*string {
	var (
		str1 = "str1"
		str2 = "str2"
		str3 = ""
	)
	return []*string{
		&str1,
		&str2,
		&str3,
	}
}()

// Slice test cases ------------------------------------------------------------
var SliceTestCases = [][]int{
	{},
	{1, 20, 3, math.MaxInt, math.MinInt},
}

// Map test cases --------------------------------------------------------------
var MapTestCases = []map[float32]uint8{
	{},
	{1.0: 8, 20.182736: 110, math.MaxFloat32: math.MaxUint8},
	{math.SmallestNonzeroFloat32: 0, 0: 0},
}
