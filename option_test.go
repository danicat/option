package option_test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danicat/option"
)

func ExampleNew() {
	opt := option.New(1337)
	print := func(n int) { fmt.Println(n) }
	opt.ForEach(print)
	// output: 1337
}

func ExampleOption_Filter() {
	opt := option.New(1337)
	isEven := func(n int) bool { return n%2 == 0 }
	print := func(n int) { fmt.Println(n) }
	opt.Filter(isEven).ForEach(print)
	// output:
}

func ExampleMap() {
	opt := option.New(1234)
	startsWith := func(prefix string) func(string) bool {
		return func(s string) bool {
			return strings.HasPrefix(s, prefix)
		}
	}
	print := func(s string) { fmt.Println(s) }
	option.Map(opt, strconv.Itoa).Map(strings.TrimSpace).Filter(startsWith("1")).ForEach(print)
	// output: 1234
}

func ExampleOption_OrElse() {
	opt := option.New("the quick black cat jumped over the lazy dog")
	startsWith := func(prefix string) func(string) bool {
		return func(s string) bool {
			return strings.HasPrefix(s, prefix)
		}
	}
	print := func(s string) { fmt.Println(s) }
	option.Map(opt, strings.ToUpper).Filter(startsWith("S")).OrElse("meow").ForEach(print)
	// output: meow
}

func ExampleOption_GetOrElse() {
	opt := option.New("the quick black cat")
	startsWith := func(prefix string) func(string) bool {
		return func(s string) bool {
			return strings.HasPrefix(s, prefix)
		}
	}

	res := opt.Map(strings.ToUpper).Filter(startsWith("s")).GetOrElse("jumped over the lazy dog")
	fmt.Println(res)
	// output: jumped over the lazy dog
}
