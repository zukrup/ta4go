package Number

var _ Num = &NaN{}

type NaN struct {
}

func (NaN) EPS() float64 {
	return 0
}

func (NaN) ValueOf32f(i float32) NaN {
	return NaN{}
}

func (NaN) ValueOf64f(i float64) NaN {
	return NaN{}
}

// Apply implements Function
func (NaN) Apply(input Num) Num {
	panic("unimplemented")
}

/*
   default <V> Function<V, R> compose(Function<? super V, ? extends T> before) {
       Objects.requireNonNull(before);
       return (V v) -> apply(before.apply(v));
   }
*/
// Apply implements Function
func (d *NaN) Compose(input Function[float64, Num]) Function[float64, Num] {
	panic("unimplemented")
}

// Apply implements Function
func (d *NaN) ThenAfter(input Function[Num, float64]) Function[Num, float64] {
	panic("unimplemented")
}

// Function implements Num
func (d *NaN) Function() Function[Num, Num] {
	panic("unimplemented")
}

// abs implements Num
func (d *NaN) Abs() Num {
	return d
}

// ceil implements Num
func (d *NaN) Ceil() Num {
	return d
}

// dividedBy implements Num
func (d *NaN) DividedBy(n Num) Num {
	if n == nil {
		return nil
	}
	return d
}

// equals implements Num
func (d *NaN) Equals(n Num) bool {
	if n == nil {
		return false
	}

	return n == d
}

/*

@Override
    public boolean equals(Object obj) {
        if (!(obj instanceof NaN)) {
            return false;
        }

        NaN NaNObj = (NaN) obj;
        return Math.abs(delegate - NaNObj.delegate) < EPS;
    }
*/

// floor implements Num
func (d *NaN) Floor() Num {

	return d
}

// getDelegate implements Num
func (d *NaN) GetDelegate() float64 {
	return 0
}

// getName implements Num
func (*NaN) GetName() string {
	panic("unimplemented")
}

// hashCode implements Num
func (*NaN) HashCode() int {
	panic("unimplemented")
}

// isGreaterThan implements Num
func (d *NaN) IsGreaterThan(n Num) bool {
	return false
}

// isGreaterThanOrEqual implements Num
func (d *NaN) IsGreaterThanOrEqual(n Num) bool {
	return false
}

// isLessThan implements Num
func (d *NaN) IsLessThan(n Num) bool {
	return false
}

// isLessThanOrEqual implements Num
func (d *NaN) IsLessThanOrEqual(n Num) bool {
	return false
}

// isNan implements Num
func (*NaN) IsNaN() bool {
	return true
}

// isNegative implements Num
func (d *NaN) IsNegative() bool {
	return false
}

// isNegativeOrZero implements Num
func (d *NaN) IsNegativeOrZero() bool {
	return false
}

// isPositive implements Num
func (d *NaN) IsPositive() bool {
	return false
}

// log implements Num
func (d *NaN) Log() Num {

	return d
}

// max implements Num
func (d *NaN) Max(n Num) Num {

	return d
}

// min implements Num
func (d *NaN) Min(n Num) Num {

	return d
}

// minus implements Num
func (d *NaN) Minus(n Num) Num {

	return d
}

// multipliedBy implements Num
func (d *NaN) MultipliedBy(n Num) Num {

	return d
}

// negate implements Num
func (d *NaN) Negate() Num {

	return d
}

// numOf implements Num
func (NaN) NumOf(val float64) Num {
	return &NaN{}
}

// plus implements Num
func (d *NaN) Plus(n Num) Num {

	return d
}

// pow implements Num
func (d *NaN) Pow(i int) Num {

	return d
}

// remainder implements Num
func (d *NaN) Remainder(divisor Num) Num {
	return d
}

// sqrt implements Num
func (d *NaN) Sqrt() Num {
	return d
}

// toString implements Num
func (d *NaN) ToString() string {
	return "NaN"
}

// compareTo implements Num
func (d *NaN) CompareTo(T Num) int {
	if d == nil || T == nil {
		return 0
	}

	return 0
}
