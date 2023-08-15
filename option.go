package option

// Option defines an option type that can be defined (Some[T]) or undefined (None)
type Option[T any] struct {
	value   T
	defined bool
}

// New creates a new Some[T] with the given value
func New[T any](value T) Option[T] {
	return Option[T]{
		value:   value,
		defined: true,
	}
}

// IsDefined is true if option holds a value (Some) or false if empty (None)
func (o Option[T]) IsDefined() bool {
	return o.defined
}

// Get returns the value of the option
func (o Option[T]) Get() T {
	return o.value
}

// Filter applies function f and returns some if the condition is true, none otherwise
func (o Option[T]) Filter(f func(T) bool) Option[T] {
	if o.defined && f(o.value) {
		return o
	}
	return Option[T]{}
}

// ForEach calls function f for each value
func (o Option[T]) ForEach(f func(T)) {
	if o.defined {
		f(o.value)
	}
}

// Map applies a function f: T -> T (special case)
func (o Option[T]) Map(f func(T) T) Option[T] {
	if o.defined {
		return New(f(o.value))
	}
	return Option[T]{}
}

// OrElse returns a default value if option is None (special case)
func (o Option[T]) OrElse(v T) Option[T] {
	if o.defined {
		return o
	}
	return New(v)
}

// General functions Option[T] -> Option[U]

// Map applies a function f: T -> U converting Option[T] to Option[U]
func Map[T any, U any](opt Option[T], f func(T) U) Option[U] {
	if opt.IsDefined() {
		return New(f(opt.Get()))
	}
	return Option[U]{}
}
