

package main

import (
    "fmt"
)

func main() {
    weekTempArr := [7]int{1,2,3,4,5,6,7}
    workDaysSlice := weekTempArr[:5]
    weekendSlice := weekTempArr[5:]
    fromTuesdayToThursDaySlice := weekTempArr[1:4] 
    weekTempSlice := weekTempArr[:]
    
    fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2 
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice),cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6 
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7


    fmt.Println("\n")
    a := []int{1, 2, 3}
    fmt.Println(a, len(a), cap(a)) // [1 2 3] 3 3
    b := append(a, 4)
    fmt.Println(a, len(a), cap(a)) // [1 2 3] 3 3
    fmt.Println(b, len(b), cap(b)) // [1 2 3 4] 4 6 



    fmt.Println("\n")
    s := make([]int, 4, 7) 
    fmt.Println(s, len(s), cap(s)) // [0 0 0 0], len = 4 cap = 7
    // 1. Создаём слайс s с базовым массивом на 7 элементов. 
    // Четыре первых элемента будут доступны в слайсе.
    
    slice1 := append(s[:2], 2, 3, 4)  
    fmt.Println(s, slice1, len(slice1), cap(slice1)) // [0 0 2 3] [0 0 2 3 4]
    // 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
    // Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7), 
    // то базовый массив остаётся прежним.
    // Слайс s тоже изменился, но его длина осталась прежней
    
    slice2 := append(s[1:2], 7) 
    fmt.Println(s, slice1, slice2, len(slice2), cap(slice2)) // [0 0 7 3] [0 0 7 3 4] [0 7]
    // 3. Здесь также базовый массив остаётся прежним, изменились все три слайса
    
    slice3 := append(s, slice1[1:]...)
    fmt.Println(s, slice3, len(slice3), cap(slice3))  // 8 14
    // 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,  
    // что больше ёмкости базового массива.
    // Будет создан новый базовый массив ёмкостью 14,
    // ёмкость нового базового массива подбирается автоматически 
    // и зависит от текущего размера и количества добавленных элементов
    
    // 5. Легко проверить, что slice3 ссылается на новый базовый массив
    
    
    fmt.Println("\n")
    s[1] = 99
    fmt.Println(s, slice1, slice2, slice3) 
    // [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4] 
}