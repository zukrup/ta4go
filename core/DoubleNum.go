package core

import (
	"fmt"
	"math"
)

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
func (*DoubleNum) Function(Num, Num) {
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
func (d *DoubleNum) Equals(n Num) bool {
	if n == nil {
		return false
	}

	return math.Abs(d.delegate-n.GetDelegate()) < d.EPS()
}

/*

@Override
    public boolean equals(Object obj) {
        if (!(obj instanceof DoubleNum)) {
            return false;
        }

        DoubleNum doubleNumObj = (DoubleNum) obj;
        return Math.abs(delegate - doubleNumObj.delegate) < EPS;
    }
*/

// floor implements Num
func (*DoubleNum) Floor() Num {
	panic("unimplemented")
}

// getDelegate implements Num
func (d *DoubleNum) GetDelegate() float64 {
	return d.delegate
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
	return n != nil && d.CompareTo(n) > 0
}

// isGreaterThanOrEqual implements Num
func (d *DoubleNum) IsGreaterThanOrEqual(n Num) bool {
	return n != nil && d.CompareTo(n) > -1
}

// isLessThan implements Num
func (d *DoubleNum) IsLessThan(n Num) bool {
	return d != nil && d.CompareTo(n) < 0
}

// isLessThanOrEqual implements Num
func (d *DoubleNum) IsLessThanOrEqual(n Num) bool {
	return d != nil && d.CompareTo(n) < 1
}

// isNan implements Num
func (*DoubleNum) IsNan() bool {
	return false
}

// isNegative implements Num
func (d *DoubleNum) IsNegative() bool {
	return d.delegate < 0
}

// isNegativeOrZero implements Num
func (d *DoubleNum) IsNegativeOrZero() bool {
	return d.delegate <= 0
}

// isPositive implements Num
func (d *DoubleNum) IsPositive() bool {
	return d.delegate > 0
}

// log implements Num
func (d *DoubleNum) Log() Num {
	d.delegate = math.Log(d.delegate)
	return d
}

// max implements Num
func (d *DoubleNum) Max(n Num) Num {
	d.delegate = math.Max(d.delegate, n.GetDelegate())
	return d
}

// min implements Num
func (d *DoubleNum) Min(n Num) Num {
	d.delegate = math.Min(d.delegate, n.GetDelegate())
	return d
}

// minus implements Num
func (d *DoubleNum) Minus(n Num) Num {
	d.delegate -= n.GetDelegate()
	return d
}

// multipliedBy implements Num
func (d *DoubleNum) MultipliedBy(n Num) Num {
	d.delegate += n.GetDelegate()
	return d
}

// negate implements Num
func (d *DoubleNum) Negate() Num {
	d.delegate = -d.delegate
	return d
}

// numOf implements Num
func (DoubleNum) NumOf(val float64) Num {
	return &DoubleNum{
		delegate: val,
	}
}

// plus implements Num
func (d *DoubleNum) Plus(n Num) Num {
	d.delegate += n.GetDelegate()
	return d
}

// pow implements Num
func (d *DoubleNum) Pow(i int) Num {
	d.delegate = math.Pow(d.delegate, float64(i))
	return d
}

// remainder implements Num
func (d *DoubleNum) Remainder(divisor Num) Num {
	if divisor.IsNan() {
		return &DoubleNum{delegate: 02}
	}
	val := math.Remainder(d.delegate, divisor.GetDelegate())
	return &DoubleNum{delegate: val}
}

// sqrt implements Num
func (d *DoubleNum) Sqrt() Num {
	if d.delegate < 0 {
		return nil //todo: NAN implemente edilecek
	}
	d.delegate = math.Sqrt(d.delegate)
	return d
}

// toString implements Num
func (d *DoubleNum) ToString() string {
	return fmt.Sprintf("%f", d.delegate)
}

// compareTo implements Num
func (d *DoubleNum) CompareTo(T Num) int {
	if d == nil || T == nil {
		return 0
	}

	return int(d.delegate) - int(T.GetDelegate())
}
