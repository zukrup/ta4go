package core

type Function[T any, R any] interface {
	apply(T) R
}
