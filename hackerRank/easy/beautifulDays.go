package hackerRank

//https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
func reverse_int(n int32) int32 {

	var new_int int32 = 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}

func beautifulDays(i int32, j int32, k int32) int32 {

	var v int32 = i

	var total int32 = 0

	for ; v <= j; v++ {

		output := reverse_int(v)

		if (output-v)%k == 0 {
			total++
		}
	}

	return total
}
