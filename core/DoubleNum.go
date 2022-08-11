package core

import "fmt"

var _ Num = &DoubleNum{}

type DoubleNum struct {
	delegate float64
}

func New(
	val float64,
) Num {
	return &DoubleNum{
		delegate: val,
	}
}

func (DoubleNum) EPS() float64 {
	return 0.00001
}

func (DoubleNum) ValueOf32f(i float32) DoubleNum {
	return DoubleNum{float64(i)}
}

func (DoubleNum) ValueOf64f(i float64) DoubleNum {
	return DoubleNum{i}
}

// Function implements Num
func (*DoubleNum) Function(Number, Num) {
	panic("unimplemented")
}

// abs implements Num
func (*DoubleNum) Abs() Num {
	panic("unimplemented")
}

// ceil implements Num
func (*DoubleNum) Ceil() Num {
	panic("unimplemented")
}

// dividedBy implements Num
func (*DoubleNum) DividedBy(Num) Num {
	panic("unimplemented")
}

// equals implements Num
func (*DoubleNum) Equals(Num) bool {
	panic("unimplemented")
}

// floor implements Num
func (*DoubleNum) Floor() Num {
	panic("unimplemented")
}

// getDelegate implements Num
func (*DoubleNum) GetDelegate() Number {
	panic("unimplemented")
}

// getName implements Num
func (*DoubleNum) GetName() string {
	panic("unimplemented")
}

// hashCode implements Num
func (*DoubleNum) HashCode() int {
	panic("unimplemented")
}

// isGreaterThan implements Num
func (d *DoubleNum) IsGreaterThan(n Num) bool {
	return d != nil && d.CompareTo(n.GetDelegate()) < 1
}

// isGreaterThanOrEqual implements Num
func (*DoubleNum) IsGreaterThanOrEqual(Num) bool {
	panic("unimplemented")
}

// isLessThan implements Num
func (d *DoubleNum) IsLessThan(n Num) bool {
	return d != nil && d.CompareTo(n.GetDelegate()) < 1
}

// isLessThanOrEqual implements Num
func (*DoubleNum) IsLessThanOrEqual(Num) bool {
	panic("unimplemented")
}

// isNan implements Num
func (*DoubleNum) IsNan() bool {
	panic("unimplemented")
}

// isNegative implements Num
func (*DoubleNum) IsNegative() bool {
	panic("unimplemented")
}

// isNegativeOrZero implements Num
func (*DoubleNum) IsNegativeOrZero() bool {
	panic("unimplemented")
}

// isPositive implements Num
func (*DoubleNum) IsPositive() bool {
	panic("unimplemented")
}

// log implements Num
func (*DoubleNum) Log() Num {
	panic("unimplemented")
}

// max implements Num
func (*DoubleNum) Max(Num) Num {
	panic("unimplemented")
}

// min implements Num
func (*DoubleNum) Min(Num) Num {
	panic("unimplemented")
}

// minus implements Num
func (*DoubleNum) Minus(Num) Num {
	panic("unimplemented")
}

// multipliedBy implements Num
func (*DoubleNum) MultipliedBy(Num) Num {
	panic("unimplemented")
}

// negate implements Num
func (*DoubleNum) Negate() Num {
	panic("unimplemented")
}

// numOf implements Num
func (*DoubleNum) NumOf(Number) Num {
	panic("unimplemented")
}

// plus implements Num
func (*DoubleNum) Plus(Num) Num {
	panic("unimplemented")
}

// pow implements Num
func (*DoubleNum) Pow(int Num) Num {
	panic("unimplemented")
}

// remainder implements Num
func (*DoubleNum) Remainder(Num) Num {
	panic("unimplemented")
}

// sqrt implements Num
func (*DoubleNum) Sqrt() Num {
	panic("unimplemented")
}

// toString implements Num
func (d *DoubleNum) ToString() string {
	return fmt.Sprintf("%f", d.delegate)
}

// compareTo implements Num
func (d *DoubleNum) CompareTo(T Number) int {
	if d == nil || T == nil {
		return 0
	}

	return int(d.delegate) - int(T.intValue())
}
