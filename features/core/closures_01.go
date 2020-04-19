package core

import (
	"fmt"
	"sort"
)

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

func Closures_01_Sample01() {

	helloGo()

	func() {
		fmt.Println("Hello Go from an Anonymous Function")
	}()

	var hello func() = func() {
		fmt.Println("Hello Go from an Anonymous -> Function Variable")
	}
	hello()

	n1 := incrementCounter()
	fmt.Println("n1 increment counter #1: ", n1())
	fmt.Println("n1 increment counter #2: ", n1()) // Notice the second invocation; n1 is called twice, so n1 == 2

	n2 := incrementCounter()                       // New instance of initializedNumber
	fmt.Println("n2 increment counter #1: ", n2()) // n2 is only called once, so n2 == 1
	fmt.Println("n1 increment counter #3: ", n1()) // state of n1 is not changed with the n2 calls
}

func Closures_01_Sample02() {
	input := []string{"foo", "bar", "baz"}
	var result []string
	// closure callback
	func() {
		result = append(input, "abc")       // Append to the array
		result = append(result, "def")      // Append to the array again
		sort.Sort(sort.StringSlice(result)) // Sort the larger array
	}()
	fmt.Print(result)
}
