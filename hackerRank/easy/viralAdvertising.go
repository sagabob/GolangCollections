package hackerRank

// Complete the viralAdvertising function below.
//https://www.hackerrank.com/challenges/strange-advertising/problem?h_r=next-challenge&h_v=zen
func viralAdvertising(n int32) int32 {

	output := make([]int32, n)

	output[0] = (5 / 2) * 3

	var i int32 = 1

	var total int32 = 2

	for ; i < n; i++ {
		currentTotal := output[i-1] / 2
		output[i] = currentTotal * 3
		total += currentTotal
	}

	return total
}
