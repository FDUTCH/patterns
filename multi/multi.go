package multi

func Join[T any](opts ...JoinFunc[T]) JoinFunc[T] {
	return func(t T) T {
		for _, f := range opts {
			t = f(t)
		}
		return t
	}
}

func Apply[T any](val T, opts ...func(T)) {
	for _, opt := range opts {
		opt(val)
	}
}

type JoinFunc[T any] = func(T) T
