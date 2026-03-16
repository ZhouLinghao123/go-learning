package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("成人")
	} else {
		fmt.Println("未成年")
	}

	// 第一部分：硬编码分数演示 else if 链
	score := 85
	fmt.Println("硬编码分数演示：")
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// 第二部分：从终端输入分数（复用 score 变量）
	fmt.Print("请输入分数（0-100）：")
	fmt.Scan(&score) // 直接用前面的 score，不需要再 var

	fmt.Println("你输入的分数判断：")
	if score < 0 || score > 100 {
		fmt.Println("输入错误！分数必须在 0-100 之间")
	} else if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
