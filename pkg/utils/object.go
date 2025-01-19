package utils

func ToPointer[T any](v T) *T {
	return &v
}

func GetValue[T any](v *T) T {
	if v != nil {
		return *v
	}
	return *new(T)
}

func Either[T any](v1 *T, v2 *T) *T {
	if v1 != nil {
		return v1
	}

	return v2
}

func IfThenElse[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}

	return b
}
