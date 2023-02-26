package main
import "fmt"

// Function to get the sum average of the team

func main() {
    team := map[string]float32{
        "P1": 1.98, 
        "P2": 2.05,
        "P3": 1.89,
        "P4": 2.0,
        "P5": 2.11}
        
    var sum float32 = 0.0

    for _, v := range team{
        sum += v
    }
    fmt.Println(sum/5)
}