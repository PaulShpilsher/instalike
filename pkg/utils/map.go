package utils

// Map function maps elements of source array to another one
func Map[S any, D any](source []S, mapFunc func(S) D) []D {

	dest := make([]D, len(source))

	for i, src := range source {
		dest[i] = mapFunc(src)
	}

	return dest
}
