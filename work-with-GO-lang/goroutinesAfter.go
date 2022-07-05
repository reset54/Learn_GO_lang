package main
import (
    "fmt"
    "time"
)

func main() {
    doneChan := make(chan string)
    go func() {
     // Сделать что-нибудь полезное
        doneChan <- “I’m all done!”
    }()

    <-doneChan // чтение блокируется до сигнала о завершении работы
}