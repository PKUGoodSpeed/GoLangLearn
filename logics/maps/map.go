package maps

import "fmt"

func ShowMapUsage(){
    m := make(map[string]int)
    m["zebo"] = 17
    m["shaocong"] = 25
    fmt.Println("\n\nmap: ", m)

    v1, v2 := m["shaocong"]
    fmt.Println("shaocong: ", v1, v2)
    fmt.Println("len: ", len(m))
    
    delete(m, "shaocong")
    fmt.Println("shaocong: ", v1)
    fmt.Println("m[shaocong]: ", m["shaocong"])
    fmt.Println("new Len: ", len(m))

    n := map[string]int{"x": 123}
    fmt.Println("new Map: ", n)
}