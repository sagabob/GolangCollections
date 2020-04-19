package hackerRank

import (
	"fmt"
	"math"
)

//https://www.hackerrank.com/challenges/max-array-sum/
// optimized -> recheck DP
func maxSubsetSum(arr []int32) int32 {

	arrLength := len(arr)

	trackArray := make([]int32, arrLength)
	trackArray[0] = int32(math.Max(float64(arr[0]), 0))
	trackArray[1] = int32(math.Max(float64(arr[1]), float64(trackArray[0])))

	for i := 2; i < arrLength; i++ {
		trackArray[i] = int32(math.Max(float64(trackArray[i-1]), float64(arr[i]+trackArray[i-2])))
	}

	return trackArray[arrLength-1]
}

func Output_maxSubsetSum() {
	array := []int32{2, 1, 5, 8, 4}
	fmt.Println(maxSubsetSum(array))
}
