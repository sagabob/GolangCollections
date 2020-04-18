package hackerRank

import (
	"fmt"
	"math"
)

//https://www.hackerrank.com/challenges/beautiful-triplets/problem
//optimized
func beautifulTriplets(d int32, arr []int32) int32 {

	var totalLeft int32 = 0

	var totalRight int32 = 0

	var total int32 = 0

	for j := len(arr) - 1; j > 0; j-- {

		totalRight = 0
		totalLeft = 0
		for i := 0; i < j; i++ {
			if arr[j]-arr[i] < d {
				break
			} else if arr[j]-arr[i] == d {
				totalLeft++

			}

		}
		for k := len(arr) - 1; k > j; k-- {
			if arr[k]-arr[j] < d {
				break
			} else if arr[k]-arr[j] == d {
				totalRight++

			}
		}
		if totalLeft > 0 && totalRight > 0 {
			total += int32(math.Max(float64(totalLeft), float64(totalRight)))
		}

	}

	return total

}

func Output_beautifulTriplets() {
	v := []int32{1, 6, 7, 7, 8, 10, 12, 13, 14, 19}
	var d int32 = 3
	fmt.Println(beautifulTriplets(d, v))
}
