package sift

func Select[T any, E any](q Query[T], editor func(T) E) Query[E] {
	result := make([]E, 0, len(q.items))
	for _, v := range q.items {
		result = append(result, editor(v))
	}
	return From(result)
}
