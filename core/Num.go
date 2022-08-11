package core

type Num interface {
	Comparable[Number]
	getDelegate() Number
	getName() string
	plus(Num) Num
	minus(Num) Num
	multipliedBy(Num) Num
	dividedBy(Num) Num
	remainder(Num) Num
	floor() Num
	ceil() Num
	pow(int Num) Num
	log() Num
	sqrt() Num
	abs() Num
	negate() Num
	isPositive() bool
	isNegative() bool
	isNegativeOrZero() bool
	isGreaterThan(Num) bool
	isGreaterThanOrEqual(Num) bool
	isLessThan(Num) bool
	isLessThanOrEqual(Num) bool
	min(Num) Num
	max(Num) Num
	numOf(Number) Num
	Function(Number, Num)
	isNan() bool
	hashCode() int
	toString() string
	equals(Num) bool
}
