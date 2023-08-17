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

func ExampleLift4() {
	p4 := func(v1, v2, v3, v4 int) (int, error) {
		if v1 < 0 {
			return 0, errors.New("v1 should not be negative")
		}
		return v1 + v2 + v3 + v4, nil
	}

	f4 := option.Lift4(p4)
	fmt.Println(f4(-1, 2, 3, 4))
	// output:
	// None[int]
}

func ExampleLift5() {
	p5 := func(v1, v2, v3, v4, v5 int) (int, error) {
		if v1 < 0 {
			return 0, errors.New("v1 should not be negative")
		}
		return v1 + v2 + v3 + v4 + v5, nil
	}

	f5 := option.Lift5(p5)
	fmt.Println(f5(1, 2, 3, 4, 5))
	// output:
	// Some[int]=15
}
