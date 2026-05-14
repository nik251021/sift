package sift

func (q Query[T]) ToSlice() []T {
	return q.items
}
func (q Query[T]) Count() int {
	return len(q.items)
}
func (q Query[T]) Any(filter func(T) bool) bool {
	for _, v := range q.ToSlice() {
		if filter(v) == true {
			return true
		}
	}
	return false
}
func (q Query[T]) All(filter func(T) bool) bool {
	for _, v := range q.ToSlice() {
		if filter(v) != true {
			return false
		}
	}
	return true
}
func (q Query[T]) First() (T, bool) {
	if q.Count() == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}
func (q Query[T]) Last() (T, bool) {
	if q.Count() == 0 {
		var zero T
		return zero, false
	}
	return q.items[q.Count()-1], true
}
func (q Query[T]) Find(filter func(T) bool) (T, bool) {
	for _, v := range q.items {
		if filter(v) == true {
			return v, true
		}
	}
	var zero T
	return zero, false
}
func (q Query[T]) At(index int) (T, bool) {
	if index < 0 || index >= q.Count() {
		var zero T
		return zero, false
	}
	return q.items[index], true
}
