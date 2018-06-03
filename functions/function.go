package main

import "fmt"
import V "./variadic"
import R "./recu"

func Febnacci(x0, x1 int) (int, int, int, int, int){
    A := [5] int{x0, x1, 0, 0, 0}
    for i:=2; i<5 ; i++{
        A[i] = A[i-1] + A[i-2]
    }
    return A[0], A[1], A[2], A[3], A[4]
}

func main(){
    a, b, c, d, e := Febnacci(10, 20)
    fmt.Println("Febnacci Results: ", a, b, c, d, e)
    fmt.Println(V.Sum(1,2))
    fmt.Println(V.Sum(1,2,3))
    fmt.Println(V.Sum(1,2,3,4))
    fmt.Println(V.Sum([] int {1, 2, 3, 4, 5} ...))
    f := func () func() int {
        i := 0
        return func() int {
            i++
            return i
        }
    }()
    for i:=0; i<5; i++ {
        fmt.Println(f())
    }
    
    fmt.Println("\n\n")
    for i:=0; i<32; i++{
        fmt.Println(R.Factorial(i))
    }
}