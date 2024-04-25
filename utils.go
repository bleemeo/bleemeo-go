package bleemeo

// cloneMap returns a shallow clone of the given map.
func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := make(map[K]V, len(m))

	for k, v := range m {
		m2[k] = v
	}

	return m2
}
