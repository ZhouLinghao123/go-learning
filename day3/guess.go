package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置随机种子，让每次运行都不同
	rand.Seed(time.Now().UnixNano())

	//生成1-100的随机数
	secret := rand.Intn(100) + 1

	for attempt := 1; attempt <= 7; attempt++ {
		var guess int
		fmt.Printf("第%d次猜测:", attempt)
		fmt.Scan(&guess)

		if guess == secret {
			fmt.Println("猜对了！")
		} else if guess < secret {
			fmt.Println("太小了！")
		} else if guess > secret {
			fmt.Println("太大了！")
		}
	}
	fmt.Printf("正确答案是 %d\n", secret)

}
