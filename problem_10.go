package main
import(
    "fmt";
    "math"
)

func is_prime(n int) (bool){
    if n==1{
        return false
    }
    if (n==2 || n==3){
        return true
    }
    for i:=2;i<=int(math.Sqrt(float64(n)));i++{
        if n%i==0{
            return false
        }
    }
    return true
}
func main(){
    sum := 0
    for i:=1;i<=2000000; i++{
        if is_prime(i){
            sum+=i
        }
    }
    fmt.Println(sum)
}
