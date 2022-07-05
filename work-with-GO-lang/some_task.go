
//https://stepik.org/lesson/560080/step/6?unit=554126
package main

import "fmt"

func main() { 
    var input string
    fmt.Scanln(&input)
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    fmt.Println(scoring(results, input))
}

func scoring(s []string, value string) int {
    var cnt int = 0
    for i, value := range s {
        if s[i] == value {
            fmt.Println(s[i], value)
            cnt++
        }
	}
    return cnt
}