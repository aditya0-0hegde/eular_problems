package main
import "fmt"
func sum_of_numbers(n uint64) uint64{
    r := (n*(n+1))/uint64(2)
    return r*r
}
func sum_of_squares(n uint64) uint64{
    return (n*(n+1)*((2*n)+1))/uint64(6)
}
func main(){
    fmt.Println(sum_of_numbers(100)-sum_of_squares(100))
//    fmt.Println(sum_of_numbers(5),sum_of_squares(5))
}
