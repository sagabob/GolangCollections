package hackerRank

import (
	"fmt"
	"math/big"
)

//https://www.hackerrank.com/challenges/extra-long-factorials/problem
func extraLongFactorials(n int32) {

	FibonacciBig(n)
}

func FibonacciBig(n int32) {
	if n <= 1 {
		fmt.Println(n)
	} else {
		var n1 = big.NewInt(1)

		for i := int64(1); i <= int64(n); i++ {
			n1.Mul(n1, big.NewInt(i))
		}
		fmt.Println(n1)
	}

}
