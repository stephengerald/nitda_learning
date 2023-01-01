package main

import "fmt"

/* func main() {
	fmt.Printf("hello world\n")
} */

// learning functions

/* func hello(name string) string {
	message := fmt.Sprintf("hello %v, welcome!.\n", name)
	return message
}

func main() {
	fmt.Printf(hello("Stephen Gerald"))
} */

// variable scope: local
func addsum(a, b int) int {
	var sum = a + b
	return sum
}

var c, d int
// global variable
func subnum(c, d int) int {
	var sum = c + d
	return sum
}
func main() {
	fmt.Printf("%v\n", addsum(3, 6))
	fmt.Println(subnum(5, 6))
}
