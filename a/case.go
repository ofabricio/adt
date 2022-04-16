package a

// Branch runs each case and stops at
// the first truthy case found. It is
// like a switch statement.
func Branch[T any](cs ...CaseFunc[T]) Func[T] {
	return func(v T) error {
		for _, c := range cs {
			if cnd, f := c(); cnd(v) {
				return f(v)
			}
		}
		return nil
	}
}

// Case runs f when c is true.
func Case[T any](c Cond[T], f Func[T]) CaseFunc[T] {
	return func() (Cond[T], Func[T]) {
		return c, f
	}
}

// Else is the else condition of a Branch.
// It must be the last case of a Branch.
func Else[T any](f Func[T]) CaseFunc[T] {
	True := func(T) bool { return true }
	return Case(True, f)
}

// CaseFunc is a function used
// to create a case condition.
type CaseFunc[T any] func() (Cond[T], Func[T])
