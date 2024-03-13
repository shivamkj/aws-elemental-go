package utils

func DeepCopyMap[K comparable, V any](orginal map[K]V) map[K]V {
	copy := make(map[K]V, len(orginal))
	for k := range orginal {
		copy[k] = orginal[k]
	}

	return copy
}
