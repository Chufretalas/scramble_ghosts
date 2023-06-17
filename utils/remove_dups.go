package utils

func RemoveDups[T comparable](slice []T) []T {
	new_slice := make([]T, 0, len(slice))
	for _, v := range slice {
		if !InSlice(new_slice, v) {
			new_slice = append(new_slice, v)
		}
	}
	return new_slice
}
