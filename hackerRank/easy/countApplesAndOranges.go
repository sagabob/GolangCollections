package hackerRank

import "fmt"

//https://www.hackerrank.com/challenges/apple-and-orange/problem
// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	fmt.Println(CalDist(apples, a, s, t))
	fmt.Println(CalDist(oranges, b, s, t))
}

func CalDist(apples []int32, a int32, s int32, t int32) int {
	totalApples := 0
	for i := 0; i < len(apples); i++ {

		dist := apples[i] + a

		if dist >= s && dist <= t {
			totalApples++
		}
	}
	return totalApples
}
