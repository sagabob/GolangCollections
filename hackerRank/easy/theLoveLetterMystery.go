package hackerRank

import "fmt"

//https://www.hackerrank.com/challenges/the-love-letter-mystery/

func theLoveLetterMystery(s string) int32 {

	len := len(s)

	counter := len / 2

	b := []byte(s)
	var total int32 = 0
	for i := 0; i < counter; i++ {
		fmt.Println(b[i], " vs ", b[len-i-1])
		if s[i:i+1] != s[(len-1-i):(len-i)] {

			if b[i] > b[len-i-1] {
				total += int32(b[i] - b[len-i-1])
			} else {
				total += int32(-b[i] + b[len-i-1])
			}
		}
	}
	return total

}
