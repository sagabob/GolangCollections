package hackerRank

import "math"

//https://www.hackerrank.com/challenges/strange-code/problem
// Complete the angryProfessor function below.
func strangeCounter(t int64) int64 {

	if t <= 3 {
		return 3 - t + 1
	}

	numberPower2 := int64(math.Log2(float64(t)))

	store := make([]int64, numberPower2)

	store[0] = 3
	var i int64 = 1
	var starter int64 = 0
	for ; i < numberPower2; i++ {

		store[i] = int64(3*math.Pow(float64(2), float64(i))) + store[i-1]

		if store[i] >= t {
			starter = store[i]
			break
		}
	}

	return (starter - t + 1)
}
