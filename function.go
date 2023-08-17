package option

type PartialFunction1[T, R any] func(T) (R, error)
type PartialFunction2[T1, T2, R any] func(T1, T2) (R, error)
type PartialFunction3[T1, T2, T3, R any] func(T1, T2, T3) (R, error)
type PartialFunction4[T1, T2, T3, T4, R any] func(T1, T2, T3, T4) (R, error)
type PartialFunction5[T1, T2, T3, T4, T5, R any] func(T1, T2, T3, T4, T5) (R, error)
type PartialFunction6[T1, T2, T3, T4, T5, T6, R any] func(T1, T2, T3, T4, T5, T6) (R, error)
type PartialFunction7[T1, T2, T3, T4, T5, T6, T7, R any] func(T1, T2, T3, T4, T5, T6, T7) (R, error)
type PartialFunction8[T1, T2, T3, T4, T5, T6, T7, T8, R any] func(T1, T2, T3, T4, T5, T6, T7, T8) (R, error)
type PartialFunction9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any] func(T1, T2, T3, T4, T5, T6, T7, T8, T9) (R, error)

type Function1[T, R any] func(T) Option[R]
type Function2[T1, T2, R any] func(T1, T2) Option[R]
type Function3[T1, T2, T3, R any] func(T1, T2, T3) Option[R]
type Function4[T1, T2, T3, T4, R any] func(T1, T2, T3, T4) Option[R]
type Function5[T1, T2, T3, T4, T5, R any] func(T1, T2, T3, T4, T5) Option[R]
type Function6[T1, T2, T3, T4, T5, T6, R any] func(T1, T2, T3, T4, T5, T6) Option[R]
type Function7[T1, T2, T3, T4, T5, T6, T7, R any] func(T1, T2, T3, T4, T5, T6, T7) Option[R]
type Function8[T1, T2, T3, T4, T5, T6, T7, T8, R any] func(T1, T2, T3, T4, T5, T6, T7, T8) Option[R]
type Function9[T1, T2, T3, T4, T5, T6, T7, T8, T9, R any] func(T1, T2, T3, T4, T5, T6, T7, T8, T9) Option[R]
