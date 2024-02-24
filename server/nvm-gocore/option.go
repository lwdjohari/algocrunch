package nvmgocore

import "fmt"

// Option is a generic type that may or may not hold a value of type T.
type Option[T any] struct {
	value *T
}

// Some returns an Option with a value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None returns an Option without a value.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// IsSome checks if the Option holds a value.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone checks if the Option does not hold a value.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Unwrap returns the value if the Option is Some; panics if the Option is None.
func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("attempted to unwrap an Option that is None")
	}
	return *o.value
}

// ExampleOption usage of Option<T>
func ExampleOption() {
	someOption := Some(42)
	if someOption.IsSome() {
		fmt.Println("The option contains:", someOption.Unwrap())
	}

	noneOption := None[int]()
	if noneOption.IsNone() {
		fmt.Println("The option does not contain a value.")
	}
}
