package Number

type Function[T any, R any] interface {
	Apply(T) R
	Compose(Function[float64, T]) Function[float64, R]
	ThenAfter(Function[R, float64]) Function[T, float64]
}
