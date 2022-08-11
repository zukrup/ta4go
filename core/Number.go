package core

type Number interface {
	intValue() int32
	longValue() int64
	doubleValue() float64
	byteValue() byte
	shortValue() int16
}
