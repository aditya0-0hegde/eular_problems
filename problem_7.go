package main
import (
    "fmt";
    "math"
)
func is_prime(n int) bool {
    if n==1{
        return false
    }
    if (n==2 || n==3){
        return true
    }
    for i:=2; i<= int(math.Sqrt(float64(n)));i++{
        if n%i==0{
            return false
        }
    }
    return true
}

func main(){
    count :=6
    ini := 13
    for true{
        if is_prime(ini){
            count++
        }
        if count==10001{
            fmt.Println(ini)
            break
        }
        ini+=2
    }
}
