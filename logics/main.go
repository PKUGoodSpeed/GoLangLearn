package main

import "fmt"

import (
    L "./loops"
)

func main(){
    fmt.Println(L.IterativeFactorial(5))
    fmt.Println(L.RecursiveFactorial(5))
}