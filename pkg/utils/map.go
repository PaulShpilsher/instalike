package utils

// Map function - creates a new array populated with the results of calling a provided function on every element in the source array
// similar to javascripts Array.map function
func Map[S any, D any](source []S, mapFunc func(S) D) []D {

	dest := make([]D, len(source))

	for i, src := range source {
		dest[i] = mapFunc(src)
	}

	return dest
}
