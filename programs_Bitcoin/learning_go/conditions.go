package main

import "fmt"

func main() {
	i := 20
	if i > 30 {
		fmt.Printf("number is above limit")
	} else if i < 20 {
		fmt.Printf("number is below limit")
	} else {
		fmt.Printf("this is the optimal limit")
	}
}