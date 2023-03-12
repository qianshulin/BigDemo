package main

import "fmt"

func main() {
	strings := make(chan string, 10)
	for i := 0; i < 100; i++ {
		strings <- "你好世界"
		fmt.Println(<-strings)
	}
}
