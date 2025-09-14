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
	for {
		var ok bool
		some, ok = Unwrap[T](some)
		if !ok {
			return some
		}
	}
}

func UnwrapUntilCast[T any, V any](some T) (V, bool) {
	if val, ok := any(some).(V); ok {
		return val, true
	}
	var empty V
	for {
		s, ok := Unwrap[T](some)
		if val, success := any(s).(V); success {
			return val, true
		}
		if !ok {
			return empty, false
		}
		some = s
	}
}
