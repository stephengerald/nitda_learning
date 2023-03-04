package main
import "fmt"

//create the route() function

func route(cities...string){
    
    for _, x := range cities{
        var y string = ""
        y += (x+"->")
        fmt.Println (y)}
}


func main() {
   /* var n int
    fmt.Scanln(&n)

    var cities []string
    //take n strings as input and append them to the slice
  
  for i:=0; i<n; i++{
  var city_name string
 fmt.Scanln(&city_name)
cities = append (cities, city_name,)
    //
    }
    route(cities...)  */
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x[2:4])
}