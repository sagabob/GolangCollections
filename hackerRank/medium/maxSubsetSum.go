package hackerRank

import (
	"fmt"
	"math"
)

//https://www.hackerrank.com/challenges/max-array-sum/
// need a relook, too slow
func maxSubsetSum(arr []int32) int32 {

	arrLength := len(arr)

	maxNumber := arrLength / 2

	if arrLength%2 == 1 {
		maxNumber = arrLength/2 + 1
	}

	indexArray := make([]int32, 3)

	indexArray[1] = math.MinInt32

	fmt.Println("Size -> ", maxNumber)

	for i := 2; i <= maxNumber; i++ {
		fmt.Println("Starting...", i)
		indexArray[2] = 0
		subArray(arr, indexArray, int32(i))
	}
	return indexArray[1]
}

func subArray(arr []int32, indexArray []int32, num int32) {

	if num == 1 {

		maxValue := getMax(arr, indexArray)
		fmt.Println("Final Pick -> ", maxValue, " num =", num)
		indexArray[0] = indexArray[0] + maxValue

		if indexArray[0] > indexArray[1] {
			indexArray[1] = indexArray[0]
			fmt.Println("Output -> ", indexArray[1])
		}

		indexArray[0] = 0 //reset

		return
	}

	starter := indexArray[2]

	for i := starter; i < int32(len(arr)); i++ {

		fmt.Println("Pick -> ", arr[i], " num =", num)
		indexArray[0] = indexArray[0] + arr[i]
		indexArray[2] = i + 2
		subArray(arr, indexArray, num-1)
	}

}

func getMax(arr []int32, indexArray []int32) int32 {

	starter := indexArray[2]

	var max int32 = math.MinInt32
	var hasValue bool = false
	for i := starter; i < int32(len(arr)); i++ {
		if arr[i] >= max {

			max = arr[i]
			hasValue = true
		}
	}
	if hasValue {
		return max
	}
	return 0
}

func Output_maxSubsetSum() {
	array := []int32{2, 1, 5, 8, 4}
	fmt.Println(maxSubsetSum(array))
}
