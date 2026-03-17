package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hello world hello go hello old公 hello world"

	//拆成单词切片
	words := strings.Fields(text)

	//map统计
	count := make(map[string]int)

	//遍历 words ，统计
	//。。。
	for _, words := range words {
		count[words]++
	}

	//输出结果
	fmt.Println("单词统计:")
	for word, num := range count {
		fmt.Printf("%s: %d 次\n", word, num)
	}

}
