## Struct implementation

This branch implements an Option[T] type using the struct approach.

Special cases for Map and OrElse from T -> T are defined as part of the Option[T] method set. For transformations from T -> U helper functions are provided taking the first argument as Option[T].