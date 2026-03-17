package main

import "fmt"

func main() {
	score := []int{85, 92, 78, 60, 45, 88}

	var excellent, pass, fail int
	var failedScores []int

	//用range遍历 score
	//用if 判断每个分数属于哪类
	//统计人数+收集不及格分数
	for _, r := range score {
		if r >= 90 {
			excellent++
		} else if r >= 60 {
			pass++
		} else {
			fail++
			failedScores = append(failedScores, r)
		}
	}

	fmt.Printf("优秀: %d人\n及格: %d人\n不及格: %d人\n", excellent, pass, fail)
	fmt.Println("不及格分数:", failedScores)
}
