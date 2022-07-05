package main

import (
    "fmt"
)

func main() { 
    dict := make(map[string]string)
    dict["foo"] = "bar"
    dict["bazz"] = "yup"
    for k, v := range dict {
        // k будет перебирать ключи, v — соответствующие этим ключам значения
        fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
    }
}