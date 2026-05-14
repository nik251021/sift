package sift

type Query[T any] struct {
	items []T
}

func From[T any](items []T) Query[T] {
	return Query[T]{items}
}
