package core

type Comparable[T Num] interface {
	CompareTo(T Num) int
}
