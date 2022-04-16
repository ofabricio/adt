package a

import (
	"fmt"
)

func ExampleFunc_Catch() {

	Handle := func(v int, err error) error {
		fmt.Println(v, err)
		return nil
	}

	RetErr := func(v int, err error) error {
		return fmt.Errorf("wrap(%d, %v)", v, err)
	}

	one := Print("A").Catch(Handle)
	two := RetError().Catch(Handle)
	six := RetError().Catch(RetErr, Handle)
	ten := RetError().Catch(RetErr)

	fmt.Println("-- Handler is not called --")
	err1 := one.Run(1)
	fmt.Println(err1)

	fmt.Println("-- Handler supress the error --")
	err2 := two.Run(2)
	fmt.Println(err2)

	fmt.Println("-- Handler supress the error --")
	err3 := six.Run(6)
	fmt.Println(err3)

	fmt.Println("-- Handler return an error --")
	err4 := ten.Run(10)
	fmt.Println(err4)

	// Output:
	// -- Handler is not called --
	// A <nil>
	// -- Handler supress the error --
	// 2 oops
	// <nil>
	// -- Handler supress the error --
	// 6 wrap(6, oops)
	// <nil>
	// -- Handler return an error --
	// wrap(10, oops)
}

func ExampleFunc_On() {

	Handle := func(v int, err error) error {
		fmt.Println(v, err)
		return err
	}

	one := Print("A").On(Handle)
	two := RetError().On(Handle)

	fmt.Println("-- Handle called on success --")
	err1 := one.Run(2)
	fmt.Println(err1)

	fmt.Println("-- Handle called on error --")
	err2 := two.Run(3)
	fmt.Println(err2)

	// Output:
	// -- Handle called on success --
	// A 2 <nil>
	// <nil>
	// -- Handle called on error --
	// 3 oops
	// oops
}

func RetError() Func[int] {
	return func(n int) error {
		return fmt.Errorf("oops")
	}
}
