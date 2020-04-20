package hackerRank

import (
	"math/big"
)

// Complete the fibonacciModified function below.
func fibonacciModified(t1 int32, t2 int32, n int32) string {

	var n0 = big.NewInt(int64(t1))
	var n1 = big.NewInt(int64(t2))
	var n2 = big.NewInt(1)

	for i := 2; int32(i) < n; i++ {
		n2.Set(n1)
		n2.Mul(n2, n1)
		n2.Add(n2, n0)

		//fmt.Println(n0)
		//fmt.Println(n1)
		//fmt.Println(n2)

		//reset
		n0.Set(n1)
		n1.Set(n2)

		//fmt.Println("-----")
	}

	return n1.String()
}

func Output_fibonacciModified() {
	fibonacciModified(0, 1, 10)
}
