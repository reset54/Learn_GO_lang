package main

import (
    "fmt"
)

func main () {
    var k, limit int
    fmt.Scanln(&limit, &k)
    generateArray := genArr(limit)
    fmt.Println(generateArray, k)
    fmt.Println(find(generateArray, k))
}

func genArr(limit int) []int {
    arrayNums := make([]int, 0, limit)
    for i := 0; i < limit; i++ {
        arrayNums = append(arrayNums, i+1)
    } 
    return arrayNums
}

func find(arr []int, k int) []int {
    m := make(map[int]int)       // Создадим пустую map  
    for i, a := range arr {      // будем складывать в неё индексы массива, а в качестве ключей использовать само значение 
        if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
            return []int{i, j}
        }
        m[a] = i    // если искомого значения нет, то добавляем текущий индекс и значение в map
    }
    return nil     // если пары подходящих чисел не найдены []
}      