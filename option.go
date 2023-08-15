package option

// Option defines the interface that both Some[T] and None[T] must implement
type Option[T any] interface {
	IsDefined() bool
	Get() T
	Filter(func(T) bool) Option[T]
	ForEach(func(T))

	// Convenience functions for the special case T -> T
	Map(func(T) T) Option[T]
	OrElse(v T) Option[T]
}

// New creates a new Some[T] with the given value
func New[T any](value T) Option[T] {
	return Some[T]{value: value}
}

// Some implements an Option that has a value
type Some[T any] struct {
	value T
}

// Some is always defined
func (_ Some[T]) IsDefined() bool {
	return true
}

// Get returns the value of the option
func (s Some[T]) Get() T {
	return s.value
}

// Filter applies function f and returns some if the condition is true, none otherwise
func (s Some[T]) Filter(f func(T) bool) Option[T] {
	if f(s.value) {
		return s
	}
	return None[T]{}
}

// ForEach calls function f for each value
func (s Some[T]) ForEach(f func(T)) {
	f(s.value)
}

// Map applies a function f: T -> T
func (s Some[T]) Map(f func(T) T) Option[T] {
	return New(f(s.value))
}

// OrElse returns a default value if option is None
func (s Some[T]) OrElse(v T) Option[T] {
	return s
}

// None implements the Option[T] that doesn't contain a value
type None[T any] struct{}

// None is never defined
func (_ None[T]) IsDefined() bool {
	return false
}

// Get returns the zero value for T
func (_ None[T]) Get() T {
	return *new(T)
}

// Filter on None types always return None
func (n None[T]) Filter(_ func(T) bool) Option[T] {
	return n
}

// ForEach calls the function f for each value. Does nothing for None
func (_ None[T]) ForEach(_ func(T)) {
	// nothing to do
}

// OrElse returns a default value if option is None
func (_ None[T]) OrElse(v T) Option[T] {
	return New(v)
}

// Map applies a function f: T -> T. Do nothing for None
func (n None[T]) Map(_ func(T) T) Option[T] {
	return n
}

// General functions Option[T] -> Option[U]

// Map applies a function f: T -> U converting Option[T] to Option[U]
func Map[T any, U any](opt Option[T], f func(T) U) Option[U] {
	if opt.IsDefined() {
		return New(f(opt.Get()))
	}
	return None[U]{}
}

// OrElse from Option[T] to Option[U] is not implementable in Go!
// func OrElse[T any, U any](opt Option[T], value U) Option[U] {
// 	if !opt.IsDefined() {
// 		return New(value)
// 	}
// 	return ???
// }
