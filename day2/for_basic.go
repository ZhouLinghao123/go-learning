package main

import "fmt"

func main() {
	//经典for
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//while 式
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// range slice
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}

	//忽略索引
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
