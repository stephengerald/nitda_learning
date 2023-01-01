package main
 
import (
	"fmt"
	"time"
)
 
func the_day() {
	today := time.Now()
 
	switch today.Day() {
	case 1:
		fmt.Println("Today is 1st. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
	// The default statement is used if no match is found.
}

func the_days() {
	today := time.Now()
 
	switch today.Day() {
	case 1:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
/*The fallthrough keyword used to force the execution flow to fall 
through the successive case block.*/

func day_condition() {
	today := time.Now()
 
	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func main() {
	the_day()
	fmt.Printf("\n")
	the_days()
	fmt.Printf("\n")
	day_condition()
}