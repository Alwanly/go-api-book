package utils

func AnyInSlice(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}

func AnySliceInSlice(s1 []string, s2 []string) bool {
	for _, i := range s1 {
		if AnyInSlice(s2, i) {
			return true
		}
	}
	return false
}

func DeleteAtIndexSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func DeleteValueFromSlice[T comparable](slice []T, val T) []T {
	for i, v := range slice {
		if v == val {
			slice = DeleteAtIndexSlice(slice, i)
		}
	}
	return slice
}
