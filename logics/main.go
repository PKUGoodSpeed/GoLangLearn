package main

import "fmt"

import (
    L "./loops"
    C "./conditions"
    M "./maps"
    R "./ranges"
)

func main(){
    fmt.Println("Loops:")
    fmt.Println(L.IterativeFactorial(5))
    fmt.Println(L.RecursiveFactorial(5))
    L.ShowArrayUsage()
    L.ShowSliceUsage()
    fmt.Println("\n\n")
    fmt.Println("Conditions")
    C.AgeForDrive(17)
    C.AgeForDrive(31)
    C.AgeForDrive(56)
    C.WhichDayToday()
    fmt.Println("\n\n")
    fmt.Println("Maps:")
    M.ShowMapUsage()
    fmt.Println("\n\n")
    fmt.Println("Range:")
    R.ShowRangeUsage()
    fmt.Println("\n\n")
}