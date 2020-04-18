package hackerRank

// Complete the angryProfessor function below.
// https://www.hackerrank.com/challenges/angry-professor/problem
func angryProfessor(k int32, a []int32) string {

	var count int32 = 0
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			count++
		}
	}

	if count >= k {
		return "NO"
	}

	return "YES"

}
