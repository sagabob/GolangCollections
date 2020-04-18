package hackerRank

//https://www.hackerrank.com/challenges/ctci-array-left-rotation
func rotLeft(a []int32, d int32) []int32 {

	inputLength := len(a)

	output := make([]int32, inputLength)

	copy(output, a)

	for i := 0; i < inputLength; i++ {

		currentIndex := (i + inputLength - int(d)) % inputLength
		output[currentIndex] = a[i]

	}
	return output
}
