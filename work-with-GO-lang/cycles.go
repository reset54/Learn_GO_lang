package main

import (
    "fmt"
)
func main() {
    /* Трёхкомпонентный цикл, создаём переменную */
    v := 0
    for i := 1; i < 10; i++ {
        v++
        fmt.Println(v)
    }
    

    /* Заполнять каждую компоненту необязательно — можно опускать.
    Есть ещё два варианта бесконечного цикла, но в форме трёх компонент:
    for ;; {}
    for ; true; {} 
    Компоненты цикла могут принимать более комплексный вид:*/
    for a, b := 5, 10; a < 10 && b < 20; a, b = a + 1, b + 2 { 
        fmt.Println(a, b)
    } 

    // Цикл while 
    i := 0
    for i < 5 {
        i++
    }
    fmt.Println(i) 
    // Цикл while похож на трёхкомпонентный, но здесь оставлено только основное условие.

    // Цикл range
    array := [3]int{1, 2, 3}
    for arrayIndex, arrayValue := range array {
        fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
    } 
}