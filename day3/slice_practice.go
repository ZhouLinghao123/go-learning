package main

import "fmt"

func main() {
	//创建一个切片
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("原始切片:", numbers)

	//遍历切片（带索引和值）
	fmt.Println("遍历切片：", numbers)
	for i, num := range numbers {
		fmt.Printf("索引 %d: 值 %d\n", i, num)
	}

	//忽略索引，只取值
	fmt.Println("只取值:")
	for _, num := range numbers {
		fmt.Println(num)
	}

	//切片操作（append，切片截取）
	numbers = append(numbers, 60) //在末尾加 60
	fmt.Println("append 后:", numbers)

	sub := numbers[1:4] //从索引1到3
	fmt.Println("截取[1:4]:", sub)

	//修改切片元素
	for i := range numbers {
		numbers[i] *= 2 //每个元素乘2
	}
	fmt.Println("每个元素乘 2 后:", numbers)

	//求和功能
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("所有元素的和:", sum)

	//过滤偶数
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Println("偶数切片:", evens)

	//反转切片
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	fmt.Println("反转后的切片:", numbers)

	//range
	str := "hello,世界你好"
	fmt.Println("字符串遍历：")
	for i, r := range str {
		fmt.Printf("位置 %d: 字符 %c\n", i, r)
	}
}
