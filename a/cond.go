package a

// Or applies OR logic to the conditions.
func Or[T any](cs ...Cond[T]) Cond[T] {
	return func(v T) bool {
		for _, c := range cs {
			if c(v) {
				return true
			}
		}
		return false
	}
}

// And applies AND logic to the conditions.
func And[T any](cs ...Cond[T]) Cond[T] {
	return func(v T) bool {
		for _, c := range cs {
			if !c(v) {
				return false
			}
		}
		return true
	}
}

// Cond is a condition function.
type Cond[T any] func(T) bool
