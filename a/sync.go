package a

// Parallel runs fs functions concurrently.
// It awaits until all functions are complete.
// It returns the first error found.
func Parallel[T any](fs ...Func[T]) Func[T] {
	return Semaphore(len(fs), fs...)
}

// Semaphore is like Parallel, but you control
// the number of go routines it uses. If n == 1
// it behaves like the All operator.
func Semaphore[T any](n int, fs ...Func[T]) Func[T] {
	return func(v T) error {
		sema := make(chan int, n)
		done := make(chan error, len(fs))
		defer close(sema)
		defer close(done)

		for _, fn := range fs {
			sema <- 1
			go func(f Func[T]) {
				done <- f(v)
				<-sema
			}(fn)
		}

		var err error
		for range fs {
			if e := <-done; e != nil && err == nil {
				err = e
			}
		}
		return err
	}
}
