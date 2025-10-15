package unsafe_example

import (
	"encoding/binary"
	"math/bits"
	"testing"
	"unsafe"
)

type Data struct {
	Value  uint32   // 4 bytes
	Label  [10]byte // 10 bytes
	Active bool     // 1 byte
}

const dataSize = unsafe.Sizeof(Data{})

var isLE bool

var inputData = Data{
	Value:  8675309,
	Label:  [10]byte{80, 104, 111, 110, 101},
	Active: true,
}
var input = [dataSize]byte{0x0, 0x84, 0x5f, 0xed, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}
var inputSlice = []byte{0x0, 0x84, 0x5f, 0xed, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0}

var bh [dataSize]byte
var bhs []byte
var dh Data

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

func DataFromBytes(b [dataSize]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func DataFromBytesUnsafe(b [dataSize]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

func BytesFromData(d Data) [dataSize]byte {
	out := [dataSize]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func BytesFromDataUnsafe(d Data) [dataSize]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	b := *(*[dataSize]byte)(unsafe.Pointer(&d))
	return b
}

//
// slice
//

func BytesFromDataUnsafeSlice(d Data) []byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	bs := unsafe.Slice((*byte)(unsafe.Pointer(&d)), dataSize)
	return bs
}

func DataFromBytesUnsafeSlice(b []byte) Data {
	data := *(*Data)((unsafe.Pointer)(unsafe.SliceData(b)))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

// BenchmarkBytesFromData-20
// 138902151	         8.725 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBytesFromDataUnsafe-20
// 254475913	         4.685 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBytesFromDataUnsafeSlice-20
// 101026154	        11.86 ns/op	      16 B/op	       1 allocs/op
// BenchmarkDataFromBytes-20
// 138038656	         8.863 ns/op	       0 B/op	       0 allocs/op
// BenchmarkDataFromBytesUnsafe-20
// 257361504	         4.712 ns/op	       0 B/op	       0 allocs/op
// BenchmarkDataFromBytesUnsafeSlice-20
// 260674116	         4.686 ns/op	       0 B/op	       0 allocs/op

func BenchmarkBytesFromData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bh = BytesFromData(inputData)
	}
}

func BenchmarkBytesFromDataUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bh = BytesFromDataUnsafe(inputData)
	}
}

func BenchmarkBytesFromDataUnsafeSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bhs = BytesFromDataUnsafeSlice(inputData)
	}
}

func BenchmarkDataFromBytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dh = DataFromBytes(input)
	}
}

func BenchmarkDataFromBytesUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dh = DataFromBytesUnsafe(input)
	}
}

func BenchmarkDataFromBytesUnsafeSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dh = DataFromBytesUnsafeSlice(inputSlice)
	}
}
