package main

import "fmt"

func main() {
	//显示声明
	var age int = 25
	var name string = "herb" //改成你的名字
	var height float64       //未初始化，零值

	//短声明
	score := 95.5
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Height:", height)
	fmt.Println("Score:", score)

	//类型转换
	ageFloat := float64(age)
	fmt.Printf("Age as float: %.2f\n", ageFloat)

	//练习：加一个bool变量isAdult := true,打印它
	isAdult := true
	fmt.Println("Is Adult:", isAdult)

}
