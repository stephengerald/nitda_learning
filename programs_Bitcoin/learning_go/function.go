package main

import "fmt"

/* func main() {
	fmt.Printf("hello world\n")
} */

func hello(name string) string {
	message := fmt.Sprintf("hello %v, welcome!.\n", name)
	return message
}

func main() {
	fmt.Printf(hello("Stephen Gerald"))
}
