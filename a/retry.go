package a

// Retry retries a function in case of error.
func (f Func[T]) Retry(n int) Func[T] {
	return f.Catch(func(t T, err error) error {
		for i := 0; i < n; i++ {
			if err = f(t); err == nil {
				return nil
			}
		}
		return err
	})
}
