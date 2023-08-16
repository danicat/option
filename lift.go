package option

// Lift extends the domain of a partial function to the full domain by
// returning an Option instead of a value-error pair
func Lift[T, U any](f PartialFunction1[T, U]) Function1[T, U] {
	return func(v T) Option[U] {
		res, err := f(v)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

func Lift2[T1, T2, U any](f PartialFunction2[T1, T2, U]) Function2[T1, T2, U] {
	return func(v1 T1, v2 T2) Option[U] {
		res, err := f(v1, v2)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}
