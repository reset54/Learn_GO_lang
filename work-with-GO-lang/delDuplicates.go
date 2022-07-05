package main

import "fmt"

func RemoveDuplicates(input []string) []string {
    output := make([]string, len(input))
    copy(output, input)

    inputSet := make(map[string]struct{}, len(input))
    outputIdx := 0
    for _, v := range input {
        if _, ok := inputSet[v]; !ok {
            output[outputIdx] = v
            outputIdx++
        }
        inputSet[v] = struct{}{}
    }

    return output[:outputIdx]
}

func main() {
    input := []string{
        "cat",
        "dog",
        "bird",
        "dog",
        "parrot",
        "cat",
    }
    fmt.Println(input)

    fmt.Println(RemoveDuplicates(input))
} 