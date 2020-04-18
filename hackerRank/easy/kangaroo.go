package hackerRank

//https://www.hackerrank.com/challenges/kangaroo/problem
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if v1 == v2 && x1 == x2 {
		return "YES"
	}

	if x1 < x2 {
		if v1 < v2 {
			return "NO"
		} else if v1 > v2 {
			startDist := x2 - x1

			gap := startDist % (v1 - v2)

			if gap == 0 {
				return "YES"
			}
		}
		return "NO"
	} else {
		if v1 > v2 {
			return "NO"
		} else if v1 < v2 {
			startDist := x1 - x2

			gap := startDist % (v2 - v1)

			if gap == 0 {
				return "YES"
			}
		}
		return "NO"
	}
}
