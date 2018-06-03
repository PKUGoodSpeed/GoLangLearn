package ranges

import "fmt"

func ShowRangeUsage(){
    nums := []int{2, 3, 4}
    nums = append(nums, 5)
    for i, n := range nums{
        fmt.Println(i, n)
    }
    
    m := map[string]string{
        "a": "apple",
        "b": "banana",
        "c": "cast",
        "d": "death",
    }
    for k, v := range m{
        fmt.Println(k, v)
    }
}