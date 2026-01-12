package fp

// Reduce applies the reduceFn function to each item,
// starting with initial, and returns the resulting value
func Reduce[T any, K any](items []T, initial K, fn func(acc K, item T) K) K {
	acc := initial
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}
