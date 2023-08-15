## Interface implementation

This branch implements an Option[T] type using the interface approach.

Special cases for Map and OrElse from T -> T are defined as part of the interface. For transformations from T -> U helper functions are provided taking the first argument as Option[T].
