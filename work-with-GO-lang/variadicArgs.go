package main
import (
    "fmt"
)

func main() {
    my_print(6, 7, 8)
}

func my_print(num ...int) {
    for i:=0; i < len(num); i++ { 
	    fmt.Println(num[i])
    }
}