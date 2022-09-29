package main

import "fmt"

func sum(a int, b int) int {
	return a * b
}

func main() {
	total := sum(2, 3)
	fmt.Println(total)

}
