package main
import (
    "fmt"
    "encoding/csv"
    "io"
    "os"
    "bufio"
    "strconv"
)
func main(){
    var grid  [20][20]int
    f,_ := os.Open("./grid_problem_11.csv")
    r  := csv.NewReader(bufio.NewReader(f))
    i:=0
    for {
        record,err:=r.Read()
        if err==io.EOF{
            break
        }
        for j, b := range(record){
            temp,_ := strconv.Atoi(string(b))
            grid[i][j]=temp
        }
        i++
    }
    max_hor:=max_horizontal(grid)
    max_ver:=max_vertical(grid)
    max_dia:=max_digonal_forward(grid)
    fmt.Println(max_hor)
    fmt.Println(max_ver)
    fmt.Println(max_dia)
    fmt.Println(max_digonal_backward(grid))
}


func max_horizontal(grid [20][20]int) int{
    max:=0
    for i:=0;i<20;i++{
        max_temp:=0
        for j:=0;j<=15;j++{
            prod:=1
            for k:=j;k<=j+3;k++{
                prod=prod*grid[i][k]
            }
            if prod>max_temp{
                max_temp=prod
            }
        }
        if max_temp>max{
            max=max_temp
        }
    }
    return max
}


func max_vertical(grid [20][20]int) int{
    max := 0
    for j:=0;j<20;j++{
        max_temp:=0
        for i:=0;i<=15;i++{
            prod:=1
            for k:=i;k<=i+3;k++{
                prod*=grid[k][j]
            }
            if prod>max_temp{
                max_temp=prod
            }
        }
        if max_temp>max{
            max=max_temp
        }
    }
    return max
}


func max_adjfour(list []int) int{
    l:=len(list)
    max:=0
    for i:=0;i<l-3;i++{
        prod:=1
        for j:=i;j<=i+3;j++{
            prod*=list[j]
        }
        if max<prod{
            max=prod
        }
    }
    return max
}

func generate_diagonal_forward(grid [20][20]int, i int, j int) []int{
    list := make([]int, 20-(i+j))
    if j==0{
        for a:=0;a<20-i;a++{
            list[a]=grid[a+i][a]
        }
    }
    if i==0{
        for a:=0;a<20-j;a++{
            list[a]=grid[a][a+j]
        }
    }
    return list
}

func generate_diagonal_backward(grid [20][20]int, i int, j int) []int{
    list := make([]int, 1+j-i)
    if i==0{
        for a:=0;a<=j;a++{
            list[a]=grid[a][j-a]
        }
    }
    if i!=0{
        for a:=0;a<20-i;a++{
            list[a]=grid[i+a][j-a]
        }
    }
    return list
}

func max_digonal_forward(grid [20][20]int) int{
    max:=0
    for i:=0;i<17;i++{
        l:=max_adjfour(generate_diagonal_forward(grid,i,0))
        if l>max{
            max = l
        }
    }
    for j:=0; j<17;j++{
        l:=max_adjfour(generate_diagonal_forward(grid,0,j))
        if max < l{
            max = l
        }
    }
    return max
}

func max_digonal_backward(grid [20][20]int) int{
    max:=0
    for j:=19;j>=3;j--{
        prod:=max_adjfour(generate_diagonal_backward(grid,0,j))
        if max<prod{
            max=prod
        }
    }
    for i:=1; i<=16;i++{
        prod:=max_adjfour(generate_diagonal_backward(grid,i,19))
        if max < prod{
            max = prod
        }
    }
    return max
}
