package nvmgocore

// MapGetOneOf returns an arbitrary key-value pair from the provided map.
func MapGetOneOf[K comparable, V any](m map[K]V) (K, V, bool) {
	for key, value := range m {
		return key, value, true // Return the first encountered key-value pair
	}
	var zeroK K
	var zeroV V
	return zeroK, zeroV, false // Return zero values if the map is empty
}
