package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Enter expected rate of return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateValues(investmentAmount, expectedReturnRate, inflationRate, years)

	fmt.Println("futureValue: ", futureValue)
	fmt.Printf("future Value adjusting inflation: %.2f\n", futureRealValue)
}

func calculateValues(investmentAmount float64, expectedReturnRate float64, inflationRate float64, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate)/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue

}
