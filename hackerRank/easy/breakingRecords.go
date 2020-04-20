package hackerRank

func breakingRecords(scores []int32) []int32 {

	output := make([]int32, 2)

	var good int32 = 0
	var bad int32 = 0

	var goodRecord int32 = scores[0]
	var badRecord int32 = scores[0]

	for i := 0; i < len(scores); i++ {

		if scores[i] < badRecord {
			bad++
			badRecord = scores[i]
		} else if scores[i] > goodRecord {
			good++
			goodRecord = scores[i]
		}

	}
	output[0] = good
	output[1] = bad
	return output
}
