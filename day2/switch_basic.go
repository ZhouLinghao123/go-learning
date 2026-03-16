package main

import "fmt"

func main() {
	day := 7

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuseday")
	case 3:
		fmt.Println("Wensday")
	case 4, 5: //多值
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekend")
	}

	// 无表达式 switch
	score := 0
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D or F")
	}
}
