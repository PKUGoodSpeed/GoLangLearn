package variadic

import "fmt"

func Sum(nums ... int) int{
    fmt.Print(nums, " ")
    var res int = 0
    for _, n := range nums{
        res += n
    }
    return res
}