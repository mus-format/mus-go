// Package unsafe provides high-performance serializers for byte, bool,
// string, byte slice, time.Time, uint, int, and float types.
//
// These serializers utilize Go's unsafe package for direct memory access,
// enabling maximum performance.
//
// The unsafe package is intended for use cases where performance is critical
// and the developer is aware of the potential risks associated with unsafe
// operations.
package unsafe
