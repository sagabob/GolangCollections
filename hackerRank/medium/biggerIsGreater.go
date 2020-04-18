package hackerRank

//https://www.hackerrank.com/challenges/bigger-is-greater/problem

//https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
func biggerIsGreater(w string) string {

	runeW := []byte(w)

	i := len(runeW) - 1

	for ; i > 0 && (runeW[i-1] >= runeW[i]); i-- {

	}

	if i <= 0 {
		return "no answer"
	}

	// Find successor to pivot
	j := len(runeW) - 1
	for ; runeW[j] <= runeW[i-1]; j-- {
	}

	temp := runeW[i-1]
	runeW[i-1] = runeW[j]
	runeW[j] = temp

	j = len(runeW) - 1
	for ; i < j; j-- {
		temp = runeW[i]
		runeW[i] = runeW[j]
		runeW[j] = temp
		i++

	}

	return string(runeW)

}
