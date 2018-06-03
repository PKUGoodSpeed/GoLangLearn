package conditions

import "time"
import "fmt"

func WhichDayToday(){
    switch time.Now().Weekday(){
        case time.Monday:
            fmt.Println("Today is Monday!")
        case time.Tuesday:
            fmt.Println("Today is Tuesday!")
        case time.Wednesday:
            fmt.Println("Today is Wednesday!")
        case time.Thursday:
            fmt.Println("Today is Thursday!")
        case time.Friday:
            fmt.Println("Today is Friday!")
        default:
            fmt.Println("Today is weekend!!")
    }
}