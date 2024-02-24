package nvmgocore

// Result is a generic type that holds either a value of type T or an error.
type Result[T any] struct {
	Value T
	Err   error
}

// NewResult creates a new Result with the provided value and error.
func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Err: err}
}

// IsOk checks if the Result holds a value (i.e., no error).
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

// IsErr checks if the Result holds an error.
func (r Result[T]) IsErr() bool {
	return r.Err != nil
}
