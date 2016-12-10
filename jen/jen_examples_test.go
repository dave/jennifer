package jen

import "fmt"

func ExampleFunc() {
	fmt.Printf("%#v", Func().Id("a").Params())
	// Output: func a()
}

func ExampleFunc1() {
	fmt.Printf("%#v", Func().Braces())
	// Output: func{}
}
