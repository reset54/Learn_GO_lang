package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    myBill := reader.ReadString()
}

func reader_1(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
}