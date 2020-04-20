package hackerRank

import (
	"fmt"
	"math"
)

func kaprekarNumbers(p int32, q int32) {

	hasValue := false

	for k := p; k <= q; k++ {
		if checkNumber(k, (int64(k) * int64(k))) {
			//fmt.Print(k, " ")
			hasValue = true
		}
	}

	if !hasValue {
		fmt.Println("INVALID RANGE")
	}
}

func checkNumber(target int32, output int64) bool {

	var digits int = int(math.Floor(math.Log10(float64(target)) + 1))

	tenPower := int64(math.Pow10(int(digits)))

	firstNumber := output / tenPower
	secondNumber := output - (firstNumber * tenPower)

	if (firstNumber + secondNumber) == int64(target) {
		fmt.Println(firstNumber, " + ", secondNumber, " =  ", target)
		return true
	}
	return false
}

func Output_kaprekarNumbers() {
	kaprekarNumbers(1, 99999)
}
