package main

import (
    "fmt"
)

func main() {
    for i := 1; i < 100; i++ {
        switch {
            case i % 3 == 0 && i % 5 == 0:
                fmt.Println("FuzzBuzz")
            case i % 3 == 0 && i % 5 != 0:
                fmt.Println("Fizz")
            case i % 3 != 0 && i % 5 == 0:
                fmt.Println("Buzz")
            default:
                fmt.Println(i)
        }
    } 
}