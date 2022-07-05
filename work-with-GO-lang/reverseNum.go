package main

import "fmt"

func main() {
    limit := 100
    ourSlice := make([]int, 0, limit)
    //fmt.Println(ourSlice)

    for i := 0; i < limit; i++ {
        ourSlice = append(ourSlice, i+1)
    } 
    //fmt.Println(ourSlice, cap(ourSlice))

    length := len(ourSlice)
    //fmt.Println(length)

    for i := 0; i < length/2; i++ {
        ourSlice[i], ourSlice[length-1-i] = ourSlice[length-1-i], ourSlice[i]
    }
    fmt.Println(ourSlice)
}
