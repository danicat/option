package option_test

import (
	"fmt"
	"strconv"

	"github.com/danicat/option"
)

func ExampleLift() {
	full := option.Lift(strconv.Atoi)
	fmt.Println(full("foobar"))
	fmt.Println(full("123"))
	// output:
	// None[int]
	// Some[int]=123
}

func ExampleLift2() {
	parse := option.Lift2(strconv.ParseFloat)
	fmt.Println(parse("", 0))
	fmt.Println(parse("3.1415", 32))
	// output:
	// None[float64]
	// Some[float64]=3.1414999961853027
}

func ExampleLift3() {
	parse := option.Lift3(strconv.ParseInt)
	fmt.Println(parse("", 0, 16))
	fmt.Println(parse("3.1415", 0, 16))
	fmt.Println(parse("3", 0, 16))
	// output:
	// None[int64]
	// None[int64]
	// Some[int64]=3
}
