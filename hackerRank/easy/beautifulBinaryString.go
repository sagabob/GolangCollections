package hackerRank

import "fmt"

//https://www.hackerrank.com/challenges/beautiful-binary-string/

func beautifulBinaryString(b string) int32 {

	defaultString := "010"

	var lenStr int = (len(b))

	defaultRune := []rune(defaultString)

	if lenStr < 3 {
		return 0
	}

	var totalSwitch int32 = 0

	for i := 0; i < lenStr-2; i++ {

		currentStr := b[i : i+3]

		if currentStr == defaultString {

			b = replaceAtIndex(b, defaultRune[1], i+2)
			fmt.Println("Adjust -> ", b)
			totalSwitch++
		}
	}

	return totalSwitch
}

func replaceAtIndex(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}
