package a

import (
	"fmt"
)

func ExampleAll() {

	all := All(Print("A"), Print("B"))

	all.Run(2)

	// Output:
	// A B
}

func ExampleAll_error() {

	all := All(Print("A"), RetError(), Print("B"))

	err := all.Run(2)

	fmt.Println(err)

	// Output:
	// A oops
}
