package conditions

import "fmt"

func AgeForDrive(age int){
    if age < 18{
        fmt.Printf("%d is too young to drive.\n", age)
    }else if age < 50{
        fmt.Printf("%d can drive.\n", age)
    }else{
        fmt.Printf("%d is too old to drive.\n", age)
    }
}