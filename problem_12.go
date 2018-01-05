package main
import ("fmt"
)
func no_divisors(n int) int{
    divisors:=1
    for i:=2;i*i<=n;i++{
        count:=0
        for n%i==0{
            count++
            n=n/i
        }
        divisors*=(count+1)
    }
    if n>2{
        divisors*=2
    }
    return divisors
}

func triange_numbers(n int) int{
    triangle :=0
    for i:=1; i <=n; i++{
        triangle+=i
    }
    return triangle
}

func main(){
    i:=1
    for true{
        triangle := triange_numbers(i)
        n := no_divisors(triangle)
        if n>=500{
            fmt.Println(triangle,n)
            break
        }
        i++
    }
}
