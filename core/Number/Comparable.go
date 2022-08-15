package Number

type Comparable[T Num] interface {
	CompareTo(T Num) int
}
