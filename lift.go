package option

// Lift extends the domain of a PartialFunction1 to the full domain by
// converting it to a Function1 returning an Option[U].
func Lift[T, U any](f PartialFunction1[T, U]) Function1[T, U] {
	return func(v T) Option[U] {
		res, err := f(v)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

// Lift2 extends the domain of a PartialFunction2 to the full domain by
// converting it to a Function2 returning an Option[U].
func Lift2[T1, T2, U any](f PartialFunction2[T1, T2, U]) Function2[T1, T2, U] {
	return func(v1 T1, v2 T2) Option[U] {
		res, err := f(v1, v2)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

// Lift3 extends the domain of a PartialFunction3 to the full domain by
// converting it to a Function3 returning an Option[U].
func Lift3[T1, T2, T3, U any](f PartialFunction3[T1, T2, T3, U]) Function3[T1, T2, T3, U] {
	return func(v1 T1, v2 T2, v3 T3) Option[U] {
		res, err := f(v1, v2, v3)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

// Lift4 extends the domain of a PartialFunction4 to the full domain by
// converting it to a Function4 returning an Option[U].
func Lift4[T1, T2, T3, T4, U any](f PartialFunction4[T1, T2, T3, T4, U]) Function4[T1, T2, T3, T4, U] {
	return func(v1 T1, v2 T2, v3 T3, v4 T4) Option[U] {
		res, err := f(v1, v2, v3, v4)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

// Lift5 extends the domain of a PartialFunction4 to the full domain by
// converting it to a Function4 returning an Option[U].
func Lift5[T1, T2, T3, T4, T5, U any](f PartialFunction5[T1, T2, T3, T4, T5, U]) Function5[T1, T2, T3, T4, T5, U] {
	return func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Option[U] {
		res, err := f(v1, v2, v3, v4, v5)
		if err != nil {
			return Option[U]{}
		}

		return New(res)
	}
}

// TODO: implement Lift6 to Lift9
