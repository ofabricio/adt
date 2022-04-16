package a

// All runs functions sequentially. It stops
// if one of them returns an error.
func All[T any](fs ...Func[T]) Func[T] {
	return func(v T) error {
		for _, f := range fs {
			if err := f(v); err != nil {
				return err
			}
		}
		return nil
	}
}

// Func is a function used to do some work.
type Func[T any] func(T) error

// Run runs a function with an initial value.
func (f Func[T]) Run(ini T) error {
	return f(ini)
}
