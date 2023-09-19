package tools

func Map[T any](f func(T) T, list []T) []T {
	for i := range list {
		list[i] = f(list[i])
	}
	return list
}
