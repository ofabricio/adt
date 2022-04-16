package a

// Catch catches an error and handles it.
func (f Func[T]) Catch(cf ...CatchFunc[T]) Func[T] {
	return func(v T) error {
		err := f(v)
		if err != nil {
			for _, c := range cf {
				err = c(v, err)
			}
		}
		return err
	}
}

// On runs c function on either success or error.
func (f Func[T]) On(c CatchFunc[T]) Func[T] {
	return func(v T) error {
		return c(v, f(v))
	}
}

// CatchFunc is a function used to handle errors.
type CatchFunc[T any] func(T, error) error
