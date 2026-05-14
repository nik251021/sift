package sift

func (q Query[T]) Where(filter func(T) bool) Query[T] {
	var result []T

	for _, v := range q.items {
		if filter(v) == true {
			result = append(result, v)
		}
	}

	return Query[T]{items: result}
}

func (q Query[T]) Take(n int) Query[T] {
	if n <= 0 {
		return Query[T]{}
	}
	if n > q.Count() {
		n = q.Count()
	}
	return Query[T]{items: q.items[:n]}
}

func (q Query[T]) Skip(n int) Query[T] {
	if n <= 0 {
		return q
	}
	if n >= q.Count() {
		return Query[T]{}
	}
	return Query[T]{items: q.items[n:]}
}
