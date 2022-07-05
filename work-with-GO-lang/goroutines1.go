//!go run goroutines1.go
package main
import (
    "fmt"
    "time"
)

// out  с аргументами типа integer!
func out(from, to int) {
    for i := from; i <= to; i++ {
        time.Sleep( 0.00000001 * time.Millisecond )
        fmt.Println(i)
    }
}

func main() {
    go out(0, 6)                        // run first goroutine
    go out(6, 10)                       // run first goroutine
    time.Sleep(500 * time.Millisecond)  // без этой строки вывода не будет
}