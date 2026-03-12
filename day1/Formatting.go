package main

import "fmt"

func main() {
	var name string
	var age int

	//fmt.Print("Enter your name:")
	//fmt.Scan(&name) //输入你的名字

	//fmt.Print("Enter your age:")
	//fmt.Scan(&age) //输入你的年龄

	fmt.Print("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	//练习：用fmt.Scanf("%s %d", &name, &age) 一次输入年龄和姓名

}
