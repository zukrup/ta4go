package core

type Num interface {
	Comparable[Number]
	GetDelegate() Number
	GetName() string
	Plus(Num) Num
	Minus(Num) Num
	MultipliedBy(Num) Num
	DividedBy(Num) Num
	Remainder(Num) Num
	Floor() Num
	Ceil() Num
	Pow(int Num) Num
	Log() Num
	Sqrt() Num
	Abs() Num
	Negate() Num
	IsPositive() bool
	IsNegative() bool
	IsNegativeOrZero() bool
	IsGreaterThan(Num) bool
	IsGreaterThanOrEqual(Num) bool
	IsLessThan(Num) bool
	IsLessThanOrEqual(Num) bool
	Min(Num) Num
	Max(Num) Num
	NumOf(Number) Num
	Function(Number, Num)
	IsNan() bool
	HashCode() int
	ToString() string
	Equals(Num) bool
}
