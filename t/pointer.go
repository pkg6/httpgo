package t

func Pointer[T any](any T) *T {
	return &any
}
