package fp

// ReduceWithFilter applies fn to each items
// element that satisfies the predicate, starting with initial.
func ReduceWithFilter[T any, K any](
	items []T, initial K,
	fn func(acc K, item T) K,
	predicate func(item T) bool) K {
	acc := initial
	for _, item := range items {
		if predicate(item) {
			acc = fn(acc, item)
		}
	}
	return acc
}
