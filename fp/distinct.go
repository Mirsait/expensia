package fp

func Distinct[T comparable](data []T) []T {
	unique := make(map[T]bool)
	var result []T
	for _, item := range data {
		if !unique[item] {
			result = append(result, item)
			unique[item] = true
		}
	}
	return result
}
