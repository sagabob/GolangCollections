package core

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct {
	S string
}

func Interface_Sample01() {

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	var t T = T{"hello"}
	describe(t)

}
