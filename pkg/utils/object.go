package utils

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

func Either[T any](v1 *T, v2 *T) *T {
	if v1 != nil {
		return v1
	}

	return v2
}

func ToPointer[T any](v T) *T {
	return &v
}

func GetValue[T any](v *T) T {
	if v != nil {
		return *v
	}
	return *new(T)
}

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

func IfThenElse[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}

	return b
}

// ArrayDiff finds the difference between two string slices.
//
// Parameters:
// slice1: the first string slice
// slice2: the second string slice
// []string: the resulting string slice containing elements that are in slice1 but not in slice2
func ArrayDiff(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, item := range slice2 {
		m[item] = true
	}

	var diff []string
	for _, item := range slice1 {
		if !m[item] {
			diff = append(diff, item)
		}
	}
	return diff
}
