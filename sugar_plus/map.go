package sugar_plus

func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}

func FiltrateBy[K comparable, V any](in map[K]V, filtrate func(K, V) bool) map[K]V {
	result := map[K]V{}

	for k, v := range in {
		if filtrate(k, v) {
			result[k] = v
		}
	}

	return result
}
