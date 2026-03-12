package main

import "fmt"

func main() {
	const PI float64 = 3.14159
	//PI = 3.14 //故意加着一行

	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	fmt.Println("PI:", PI)
	fmt.Println("Monday:", Monday)
	fmt.Println("Saturday:", Saturday)

	//练习：用iota做月份枚举（January=1到December=12），打印April
	const (
		January = iota + 1
		February
		March
		April
		May
		June
		July
		August
		September
		October
		November
		December
	)
	fmt.Println("April:", April)
}
