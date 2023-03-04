package main
import "fmt"
// program to calculate the match result.
func main() {
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    var a int
    var b string
    fmt.Scanln(&b)
    results = append(results, b)
    for _, y := range results {
        switch{
            case y == "w":
            a+=3
            case y == "l":
            a+=0
            case y == "d":
            a+=1
        }
    }
    fmt.Println(a)
}