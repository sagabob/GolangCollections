package hackerRank

import (
	"fmt"
	"math"
)

func getWays_Recursion(n int32, c []int64) int64 {
	// Write your code here

	var maxArray int64 = int64(n)
	var maxCoin int64 = 0

	for i := 0; i < len(c); i++ {
		maxArray = int64(math.Max(float64(c[i]), float64(maxArray)))
		maxCoin = int64(math.Max(float64(c[i]), float64(maxCoin)))
	}

	store := make([]int64, maxArray+1)

	output := calChange_Recusion(int64(n), int64(n), c, store, maxCoin)

	//fmt.Println(c)
	//fmt.Println(store)
	return output

}

//DP
func getWays(n int32, c []int64) int64 {
	// Write your code here

	var maxArray int64 = int64(n)
	var maxCoin int64 = 0

	for i := 0; i < len(c); i++ {
		maxArray = int64(math.Max(float64(c[i]), float64(maxArray)))
		maxCoin = int64(math.Max(float64(c[i]), float64(maxCoin)))
	}

	//fmt.Println("Max =", maxArray+1)
	store := make([][]int64, maxArray+1)

	for i := 0; i < int(maxArray+1); i++ {

		store[i] = make([]int64, maxArray+1)
		for j := 0; j < int(maxArray+1); j++ {

			store[i][j] = 0
		}
	}

	output := calChange_DP(int64(n), int64(n), c, store, maxCoin)
	return output

}

func calChange_DP(original int64, amount int64, c []int64, store [][]int64, prev int64) int64 {

	if amount < 0 {
		return 0
	}

	if amount == 0 {
		return 1
	}

	var total int64 = 0

	for i := 0; i < len(c); i++ {

		if c[i] <= prev {

			if amount-c[i] >= 0 {
				if store[amount-c[i]][c[i]] > 0 {
					total += store[amount-c[i]][c[i]]
				} else {
					total += calChange_DP(original, amount-c[i], c, store, c[i])
				}
			}
		}
	}
	store[amount][prev] = total
	return total
}

func calChange_Recusion(original int64, amount int64, c []int64, store []int64, prev int64) int64 {

	if amount < 0 {
		return 0
	}

	if amount == 0 {
		return 1
	}

	var total int64 = 0
	for i := 0; i < len(c); i++ {

		if c[i] <= prev {
			total += calChange_Recusion(original, amount-c[i], c, store, c[i])
		}
	}

	return total
}

func Output_initializeChange() {

	c := []int64{8, 47, 13, 24, 25, 31, 32, 35, 3, 19, 40, 48, 1, 4, 17, 38, 22, 30, 33, 15, 44, 46, 36, 9, 20, 49}
	fmt.Println(getWays(250, c))
	fmt.Println(getWays_Recursion(250, c))
}
