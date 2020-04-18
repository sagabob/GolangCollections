package topics

import (
	"fmt"

	"github.com/sagabob/GolangCollections/apps/calculator"
)

func FromCalculator() {

	fmt.Println("Add(6, 3) =>", calculator.Add(6, 3))
	fmt.Println("Subtract(6, 3) =>", calculator.Subtract(6, 3))
	fmt.Println("Multiply(6, 3) =>", calculator.Multiply(6, 3))
	fmt.Println("Divide(6, 3) =>", calculator.Divide(6, 3))

}
