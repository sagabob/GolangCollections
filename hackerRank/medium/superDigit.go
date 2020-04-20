package hackerRank

import (
	"math"
	"strconv"
)

func superDigit(n string, k int32) int64 {

	var total int64 = 0

	for i := 0; i < len(n); i++ {

		digit, err := strconv.ParseInt(n[i:i+1], 10, 64)

		if err != nil {
			panic("Something wrong!!!")
		}

		total += int64(digit)

	}
	total = total * int64(k)

	if total < 10 {
		return total
	}
	return calDigit(total)
}

func calDigit(sum int64) int64 {

	numberOfDigits := int32(math.Floor(math.Log10(float64(sum)) + 1))

	var totalDigits int64 = 0
	for i := 1; int32(i) <= numberOfDigits; i++ {

		digit := sum % 10
		totalDigits += digit
		sum = sum / 10
	}

	if totalDigits < 10 {
		return totalDigits
	}
	return calDigit(totalDigits)

}
