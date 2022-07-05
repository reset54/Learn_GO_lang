package main
import (
    "fmt"
)

func main () {
    first()    
    second()
}

func first(){
    var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}   // Variable - array
    
    sumTemp := 0                            // sum (counter)
    for i :=  0; i < len(weekTemp); i++ {
        sumTemp += weekTemp[i]
        // fmt.Println(sumTemp)
    }
    
    average := sumTemp/len(weekTemp)         // middle value
    fmt.Println(sumTemp, len(weekTemp), average)
}

func second(){

    var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5} 

    sumTemp := 0
    for _, temp := range weekTemp {
        sumTemp += temp
        // fmt.Println(sumTemp)
    }
    
    average := sumTemp/len(weekTemp) 
    fmt.Println(sumTemp, len(weekTemp), average)
}