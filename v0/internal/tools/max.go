package tools

func Max(values ...int) int {
	var max int
	for i, v := range values {
		if i == 0 || max < v {
			max = v
		}
	}
	return max
}
