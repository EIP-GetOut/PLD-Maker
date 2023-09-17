package tools

type Pair[T, G any] struct {
	First  T
	Second G
}

type Triple[T, G, V any] struct {
	First  T
	Second G
	Third  V
}

type Tuple[T any] []T
