package main

import "fmt"

func main() {
	//用 map 存学生姓名，分数
	scores := make(map[string]int)

	// 添加几个学生
	scores["herb"] = 92
	scores["grok"] = 100
	scores["rose"] = 88

	//查询某个学生
	name := "herb"
	if score, ok := scores[name]; ok {
		fmt.Printf("%s 的分数是 %d\n", name, score)
	} else {
		fmt.Println("没找到该学生")
	}

	//计算平均分
	total := 0
	for _, s := range scores {
		total += s
	}
	avg := float64(total) / float64(len(scores))
	fmt.Printf("平均分: %.2f\n", avg)

	//统计及格/不及格人数
	var pass, fail int
	for _, s := range scores {
		if s >= 60 {
			pass++
		} else {
			fail++
		}
	}
	fmt.Printf("及格： %d 人，不及格: %d 人\n", pass, fail)
}
