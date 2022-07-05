package main

import (
    "fmt"
)

func main () {
    
    productsList := map[string]int{"хлеб": 50, 
                                   "молоко": 100, 
                                   "масло": 200, 
                                   "колбаса": 500, 
                                   "соль": 20, 
                                   "огурцы": 200, 
                                   "сыр": 600, 
                                   "ветчина": 700, 
                                   "буженина": 900, 
                                   "помидоры": 250, 
                                   "рыба": 300, 
                                   "хамон": 1500,
    }
    // перечень продуктов-деликатесов, стоящих больше 500 
    for k, v := range productsList {
        if v > 500 {
            fmt.Printf("%v %d\n", string(k), v)
        }
    }

    // рассчёт стоимости заказа
    order := []string{"хлеб", "буженина", "сыр", "огурцы"}
    summOrder := 0
    for _, v := range order{
        summOrder += productsList[v]
    }
    fmt.Println("Стоимость заказа ", summOrder) 
}
