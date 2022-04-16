package a

import "fmt"

func ExampleCond() {

	and := And(IsMod(3), IsMod(5))
	or := Or(IsMod(3), IsMod(5))

	fmt.Println("15 % And(3, 5):", and(15))
	fmt.Println("3  % And(3, 5):", and(3))
	fmt.Println("5  % And(3, 5):", and(5))
	fmt.Println("2  % And(3, 5):", and(2))

	fmt.Println("15 %  Or(3, 5):", or(15))
	fmt.Println("3  %  Or(3, 5):", or(3))
	fmt.Println("5  %  Or(3, 5):", or(5))
	fmt.Println("2  %  Or(3, 5):", or(2))

	// Output:
	// 15 % And(3, 5): true
	// 3  % And(3, 5): false
	// 5  % And(3, 5): false
	// 2  % And(3, 5): false
	// 15 %  Or(3, 5): true
	// 3  %  Or(3, 5): true
	// 5  %  Or(3, 5): true
	// 2  %  Or(3, 5): false
}
