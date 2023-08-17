package option

type PartialFunction1[T, R any] func(T) (R, error)
type PartialFunction2[T1, T2, R any] func(T1, T2) (R, error)
type PartialFunction3[T1, T2, T3, R any] func(T1, T2, T3) (R, error)

type Function1[T, R any] func(T) Option[R]
type Function2[T1, T2, R any] func(T1, T2) Option[R]
type Function3[T1, T2, T3, R any] func(T1, T2, T3) Option[R]
