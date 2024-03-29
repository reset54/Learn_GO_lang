package main

import "fmt"

// Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных. 
// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
// Лучше создать слайс рун из строки и развернуть его
// Например, так:

func main() {
    input := "The quick brown 狐 jumped over the lazy 犬j"
    // Get Unicode code points. 
    n := 0
    runes := make([]rune, len(input))        // создаём слайс рун 
    for _, r := range input {                // добавляем руны в слайс
        runes[n] = r
        n++
    }
    fmt.Println(runes, n)                       // слайс рун
    
    runes = runes[0:n]                       // убираем лишние нулевые руны
    for i := 0; i < n/2; i++ {               // разворачиваем
        runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
    }
    output := string(runes)                  // преобразуем в строку UTF-8. 
    fmt.Println(output)
}