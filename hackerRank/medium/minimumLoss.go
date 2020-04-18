package hackerRank

//https://www.hackerrank.com/challenges/minimum-loss/problem
//N*N solution -> need to be optimized to logN

func minimumLoss(price []int64) int32 {

	var buyIndex int = -1
	var sellIndex int = -1
	var minLoss int64 = -1

	for i := 0; i < len(price)-1; i++ {

		for j := i + 1; j < len(price); j++ {

			if minLoss < 0 {
				currentLoss := price[i] - price[j]

				if currentLoss > 0 {
					minLoss = currentLoss
					buyIndex = i
					sellIndex = j
				}
			} else {

				if (price[i]-price[j] < price[buyIndex]-price[sellIndex]) && (price[i]-price[j] > 0) {
					minLoss = price[i] - price[j]
					buyIndex = i
					sellIndex = j
				}
			}
		}
	}
	return int32(minLoss)
}
