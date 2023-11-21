package main

import (
	"testing"
	"unsafe"
)

// Standard library conversion: string to byte slice
func StringToByteSliceStandard(str string) []byte {
	return []byte(str)
}

// Standard library conversion: byte slice to string
func ByteSliceToStringStandard(b []byte) string {
	return string(b)
}

// Unsafe conversion: string to byte slice
func StringToByteSliceUnsafe(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

// Unsafe conversion: byte slice to string
func ByteSliceToStringUnsafe(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// BenchmarkStringToByteSliceStandard-12    	103879900	        11.42 ns/op	      32 B/op	       1 allocs/op
// BenchmarkByteSliceToStringStandard-12    	100000000	        10.04 ns/op	      32 B/op	       1 allocs/op
// BenchmarkStringToByteSliceUnsafe-12      	1000000000	        0.6024 ns/op	  0 B/op	       0 allocs/op
// BenchmarkByteSliceToStringUnsafe-12      	1000000000	        0.5285 ns/op	  0 B/op	       0 allocs/op

var byteSink []byte
var stringSink string

// Benchmark for standard library string to byte slice conversion
func BenchmarkStringToByteSliceStandard(b *testing.B) {
	testString := "test string for benchmark"
	for i := 0; i < b.N; i++ {
		byteSink = StringToByteSliceStandard(testString)
	}
}

// Benchmark for standard library byte slice to string conversion
func BenchmarkByteSliceToStringStandard(b *testing.B) {
	testBytes := []byte("test string for benchmark")
	for i := 0; i < b.N; i++ {
		stringSink = ByteSliceToStringStandard(testBytes)
	}
}

// Benchmark for unsafe string to byte slice conversion
func BenchmarkStringToByteSliceUnsafe(b *testing.B) {
	testString := "test string for benchmark"
	for i := 0; i < b.N; i++ {
		byteSink = StringToByteSliceUnsafe(testString)
	}
}

// Benchmark for unsafe byte slice to string conversion
func BenchmarkByteSliceToStringUnsafe(b *testing.B) {
	testBytes := []byte("test string for benchmark")
	for i := 0; i < b.N; i++ {
		stringSink = ByteSliceToStringUnsafe(testBytes)
	}
}
