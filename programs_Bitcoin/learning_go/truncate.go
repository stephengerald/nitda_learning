package main
import "fmt"

func main() {
	var decNum float32

	fmt.Printf("Enter a decimal number:")
	fmt.Scanf("%f\n", &decNum)
	fmt.Printf("the number %f to integer is: %d\n", decNum, int(decNum))
}