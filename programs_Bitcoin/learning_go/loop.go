package main
import "fmt"

func main() {
	/* for k := 0; k <= 10; k++ {
		fmt.Printf("Helo, stephen\n")
	}

	r := 0
	for; r <= 10; r++ {
		fmt.Printf("%d\n", r)
		if r == 5 {
			fmt.Printf("this is it, the special number!\n")
		}
	}

	// break statement was used here

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		} 
		}*/
	for w := 0; w < 15; {
		w++
		if w == 10 {
			continue
		}
		fmt.Printf("are you still there?\n")
	}

}