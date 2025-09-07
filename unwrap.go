package patterns

type Unwrapper[T any] interface {
	Unwrap() T
}

func Unwrap[T any](some T) (T, bool) {
	w, ok := any(some).(Unwrapper[T])
	if ok {
		return w.Unwrap(), true
	}
	return some, false
}

func FullyUnwrap[T any](some T) T {
	var ok bool
	for some, ok = Unwrap[T](some); ok; {
	}
	return some
}
