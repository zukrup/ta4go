package core

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

func (DoubleNum) valueOf32f(i float32) DoubleNum {
	return DoubleNum{float64(i)}
}

func (DoubleNum) valueOf64f(i float64) DoubleNum {
	return DoubleNum{i}
}

// Function implements Num
func (*DoubleNum) Function(Number, Num) {
	panic("unimplemented")
}

// abs implements Num
func (*DoubleNum) abs() Num {
	panic("unimplemented")
}

// ceil implements Num
func (*DoubleNum) ceil() Num {
	panic("unimplemented")
}

// dividedBy implements Num
func (*DoubleNum) dividedBy(Num) Num {
	panic("unimplemented")
}

// equals implements Num
func (*DoubleNum) equals(Num) bool {
	panic("unimplemented")
}

// floor implements Num
func (*DoubleNum) floor() Num {
	panic("unimplemented")
}

// getDelegate implements Num
func (*DoubleNum) getDelegate() Number {
	panic("unimplemented")
}

// getName implements Num
func (*DoubleNum) getName() string {
	panic("unimplemented")
}

// hashCode implements Num
func (*DoubleNum) hashCode() int {
	panic("unimplemented")
}

// isGreaterThan implements Num
func (d *DoubleNum) isGreaterThan(n Num) bool {
	return d != nil && d.compareTo(n.getDelegate()) > 0
}

// isGreaterThanOrEqual implements Num
func (*DoubleNum) isGreaterThanOrEqual(Num) bool {
	panic("unimplemented")
}

// isLessThan implements Num
func (d *DoubleNum) isLessThan(n Num) bool {
	return d != nil && d.compareTo(n.getDelegate()) < 1
}

// isLessThanOrEqual implements Num
func (*DoubleNum) isLessThanOrEqual(Num) bool {
	panic("unimplemented")
}

// isNan implements Num
func (*DoubleNum) isNan() bool {
	panic("unimplemented")
}

// isNegative implements Num
func (*DoubleNum) isNegative() bool {
	panic("unimplemented")
}

// isNegativeOrZero implements Num
func (*DoubleNum) isNegativeOrZero() bool {
	panic("unimplemented")
}

// isPositive implements Num
func (*DoubleNum) isPositive() bool {
	panic("unimplemented")
}

// log implements Num
func (*DoubleNum) log() Num {
	panic("unimplemented")
}

// max implements Num
func (*DoubleNum) max(Num) Num {
	panic("unimplemented")
}

// min implements Num
func (*DoubleNum) min(Num) Num {
	panic("unimplemented")
}

// minus implements Num
func (*DoubleNum) minus(Num) Num {
	panic("unimplemented")
}

// multipliedBy implements Num
func (*DoubleNum) multipliedBy(Num) Num {
	panic("unimplemented")
}

// negate implements Num
func (*DoubleNum) negate() Num {
	panic("unimplemented")
}

// numOf implements Num
func (*DoubleNum) numOf(Number) Num {
	panic("unimplemented")
}

// plus implements Num
func (*DoubleNum) plus(Num) Num {
	panic("unimplemented")
}

// pow implements Num
func (*DoubleNum) pow(int Num) Num {
	panic("unimplemented")
}

// remainder implements Num
func (*DoubleNum) remainder(Num) Num {
	panic("unimplemented")
}

// sqrt implements Num
func (*DoubleNum) sqrt() Num {
	panic("unimplemented")
}

// toString implements Num
func (*DoubleNum) toString() string {
	panic("unimplemented")
}

// compareTo implements Num
func (d *DoubleNum) compareTo(T Number) int {
	if d == nil || T == nil {
		return 0
	}

	return int(d.delegate) - int(T.intValue())
}
