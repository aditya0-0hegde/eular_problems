package main
import "fmt"
func palindrome(i int) bool{
  reverse := 0
  orignal := i
  for i!=0{
    reverse=(reverse*10) + i%10
    i= i /10
  }
  if orignal == reverse{
    return true
  }
  return false
}
func main()  {
  max:=0
  for i:=100; i<1000;i++{
    for j:=100;j<1000;j++{
      if (i*j>max) && palindrome(i*j){
        max = i*j
      }
    }
  }
  fmt.Println(max)
}
