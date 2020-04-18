package hackerRank

//https://www.hackerrank.com/challenges/ctci-making-anagrams

func makeAnagram(a string, b string) int32 {

	m := make(map[string]int32)

	for i := 0; i < len(a); i++ {

		var currentInt = m[a[i:i+1]]
		currentInt++
		m[a[i:i+1]] = currentInt

	}

	var left int32 = 0
	for i := 0; i < len(b); i++ {

		var currentInt = m[b[i:i+1]]

		if currentInt > 0 {
			currentInt--
			m[b[i:i+1]] = currentInt
		} else {

			left++
		}
	}

	var total int32 = 0
	for _, value := range m {
		total += value

	}

	return (total + left)

}
