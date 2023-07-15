package tools

func Ternary[T any](b bool, v_one T, v_two T) T {
	if b {
		return v_one
	}
	return v_two
}

func TernaryRef[T any](b bool, v_one *T, v_two T) T {
	if b {
		return *v_one
	}
	return v_two
}
