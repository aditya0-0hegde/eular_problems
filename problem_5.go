package main
import "fmt"
func gcd(x int, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}
func lcm(x int, y int) int{
    return (x*y/gcd(x,y))
}
func main(){
    a :=2520
    for i:=10; i<=20;i++{
        a = lcm(a,i)
    }
    fmt.Println(a)
}

