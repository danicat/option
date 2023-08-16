package option_test

import (
	"errors"
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
	div := func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("can't divide by zero")
		}
		return a / b, nil
	}

	full := option.Lift2(div)
	fmt.Println(full(1, 0))
	fmt.Println(full(10, 2))
	// output:
	// None[int]
	// Some[int]=5
}

func ExampleLift2_withMixedTypes() {
	convertAndDivide := func(a string, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("can't divide by zero")
		}

		num, err := strconv.Atoi(a)
		if err != nil {
			return 0, err
		}

		return num / b, nil
	}

	full := option.Lift2(convertAndDivide)
	fmt.Println(full("not a number", 1))
	fmt.Println(full("100", 0))
	fmt.Println(full("100", 2))
	// output:
	// None[int]
	// None[int]
	// Some[int]=50
}
