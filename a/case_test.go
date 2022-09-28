package a

import "fmt"

func ExampleCase() {

	fizzbuzz := Branch(
		Case(And(IsMod(3), IsMod(5)), Print("FizzBuzz")),
		Case(IsMod(3), Print("Fizz")),
		Case(IsMod(5), Print("Buzz")),
		Else(PrintIndex),
	)

	for i := 1; i <= 15; i++ {
		fizzbuzz.Run(i)
	}

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func IsMod(v int) Cond[int] {
	return func(i int) bool {
		return i%v == 0
	}
}

func Print(s string) Func[int] {
	return func(int) error {
		fmt.Print(s, " ")
		return nil
	}
}

func PrintIndex(i int) error {
	fmt.Print(i, " ")
	return nil
}

func ExampleBranch_no_else() {

	adt := Branch(
		Case(IsMod(3), Print("Three")),
	)

	err := adt.Run(1)

	fmt.Println(err)

	// Output:
	// <nil>
}
