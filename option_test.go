package option_test

import (
	"fmt"
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
	opt := option.New("the quick black cat jumped over the lazy dog")
	startsWith := func(prefix string) func(string) bool {
		return func(s string) bool {
			return strings.HasPrefix(s, prefix)
		}
	}
	print := func(s string) { fmt.Println(s) }
	option.Map(opt, strings.ToUpper).Filter(startsWith("T")).ForEach(print)
	// output: THE QUICK BLACK CAT JUMPED OVER THE LAZY DOG
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
