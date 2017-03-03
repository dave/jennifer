package jen

import "fmt"

func ExampleCaseStateful() {
	c := Switch(Id("person")).Block(
		Case(Id("John"), Id("Peter")).Block(
			Return(Lit("male")),
		),
		Case(Id("Gill")).Block(
			Return(Lit("female")),
		),
	)
	fmt.Printf("%#v", c)
	// Output:
	// switch person {
	// case John, Peter:
	// 	return "male"
	// case Gill:
	// 	return "female"
	// }
}
