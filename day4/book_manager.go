package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	books := []string{}                 //书名列表
	borrowCount := make(map[string]int) //借阅次数统计

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入命令 （add/list/borrow/exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) > 1 {
				book := parts[1]
				books = append(books, book)
				fmt.Println("添加成功: ", book)
			} else {
				fmt.Println("请提供书名，列如 add 三国演义")
			}

		case "list":
			if len(books) == 0 {
				fmt.Println("当前图书列表为空~")
			} else {
				fmt.Println("当前图书列表:")
				for _, b := range books {
					fmt.Printf("- %s (借阅 %d 次)\n", b, borrowCount[b])
				}
			}
		case "borrow":
			if len(parts) > 1 {
				book := parts[1]
				found := false
				for _, b := range books {
					if b == book {
						found = true
						borrowCount[book]++
						fmt.Println("借阅成功: ", book)
						break
					}
				}
				if !found {
					fmt.Println("没找到这本书")
				}

			}
		case "exit":
			fmt.Println("再见~")
			return

		default:
			fmt.Println("无效命令")
		}
	}
}
