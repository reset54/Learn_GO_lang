package main

import (
    "fmt"
    "time"
)

func UpdatePersonWithLastVisisted (p *Person) {
    p.lastVisited = time.Now() 
} 

func main() {
    a := 1
    p := &a
    b := &p
    fmt.Printf("%d %d",*p, **b)
    fmt.Println()
    
    
    *p = 3
    fmt.Println("%d %d %d",*p, **b, a)
    **b = 5
    fmt.Println("%d %d %d",*p, **b, a)
    
    a += 4 + *p + **b
    
    fmt.Printf("%d",*p)
} 


p := P {
    Name: "Alex",
    Age: 25,
    lastVisited: time.Time{} 
}
UpdatePersonWithLastVisisted(&p) 