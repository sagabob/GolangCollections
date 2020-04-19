package core

import "fmt"

func helloGo() {
	fmt.Println("Hello Go from a Function")
}

func incrementCounter() func() int {

	var initializedNumber = 0

	return func() int {
		initializedNumber++
		return initializedNumber
	}
}

func Closures_Sample01() {

	helloGo()

	func() {
		fmt.Println("Hello Go from an Anonymous Function")
	}()

	var hello func() = func() {
		fmt.Println("Hello Go from an Anonymous	Function Variable")
	}
	hello()
}
