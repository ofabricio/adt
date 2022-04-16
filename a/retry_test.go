package a

import "fmt"

func ExampleFunc_Retry() {

	RetError := func() Func[int] {
		return func(i int) error {
			fmt.Println(i)
			return fmt.Errorf("oops")
		}
	}

	all := RetError().Retry(3)

	err := all.Run(2)

	fmt.Println("Run:", err)

	// Output:
	// 2
	// 2
	// 2
	// 2
	// Run: oops
}
