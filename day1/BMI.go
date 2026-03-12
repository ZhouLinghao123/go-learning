package main

import "fmt"

func main() {
	weight := 0.0
	height := 0.0

	fmt.Print("Enter your weight (kg):")
	fmt.Scan(&weight)

	fmt.Print("Enter your height (m):")
	fmt.Scan(&height)

	bmi := weight / (height * height)
	fmt.Printf("your BMI is: %2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Overweight")

	}

	//练习：加常量MIN_BMI=18.5，用在if中

}
