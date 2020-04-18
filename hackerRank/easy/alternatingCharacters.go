package hackerRank

//https://www.hackerrank.com/challenges/alternating-characters/problem
func alternatingCharacters(s string) int32 {

	length := len(s)

	indexStr := s[0]
	var totalDeletion int32 = 0
	for i := 1; i < length; i++ {

		if indexStr == s[i] {
			totalDeletion = totalDeletion + 1

		} else {
			indexStr = s[i]
		}
	}
	return totalDeletion
}
