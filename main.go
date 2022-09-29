package main

import "fmt"

func sum(a int, b int) int {
	return a * b
}

func main() {
	var texts = deck{"mom", "drain", "gang"}

	total := sum(2, 3)
	fmt.Println("Try To Loop")

	print(texts)

	fmt.Println(total)
}
