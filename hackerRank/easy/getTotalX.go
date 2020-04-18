package hackerRank

import "fmt"

// Complete the angryProfessor function below.
// https://www.hackerrank.com/challenges/angry-professor/problem
func getTotalX(a []int32, b []int32) int32 {
	// Write your code here

	gcdValue := GCDArray(b)
	lcmValue := LCMArray(a)

	if lcmValue > gcdValue {
		return 0
	}

	var counter int32 = 0
	var i int32 = 1
	for ; i < gcdValue; i++ {
		temp := lcmValue * i
		if gcdValue >= temp && gcdValue%temp == 0 {
			fmt.Println(temp)
			counter++
		}
	}

	// need to include the gcd in this case
	if lcmValue == 1 {
		counter++
	}

	return counter

}

func GCD(a, b int32) int32 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int32) int32 {

	result := a * b / GCD(a, b)

	return result
}

func LCMArray(integers []int32) int32 {

	result := integers[0]

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCDArray(integers []int32) int32 {
	result := integers[0]

	for i := 0; i < len(integers); i++ {
		result = GCD(result, integers[i])
	}

	return result
}

func Output_getTotalX() {
	a := []int32{1}
	b := []int32{72, 48}
	fmt.Println(getTotalX(a, b))
}
