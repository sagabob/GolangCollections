package hackerRank

//https://www.hackerrank.com/challenges/grading/problem
func gradingStudents(grades []int32) []int32 {
	// Write your code here

	b := make([]int32, len(grades))

	for i := 0; i < len(grades); i++ {

		if grades[i] < 38 {
			b[i] = grades[i]
		} else {

			if grades[i]%5 >= 3 {
				b[i] = grades[i] + 5 - grades[i]%5
			} else {
				b[i] = grades[i]
			}
		}
	}
	return b
}
