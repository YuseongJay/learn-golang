package main

const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func main() {
	println("MaxInt8", MaxInt8)
	println("MinInt8", MinInt8)
	println("MaxInt16", MaxInt16)
	println("MinInt16", MinInt16)
	println("MaxInt32", MaxInt32)
	println("MinInt32", MinInt32)
	println("MaxInt64", MaxInt64)
	println("MinInt64", MinInt64)
	println("MaxUint8", MaxUint8)
	println("MaxUint16", MaxUint16)
	println("MaxUint32", MaxUint32)
	println("MaxUint64", uint64(MaxUint64))
}
