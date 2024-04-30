package main

import (
    "fmt"
    "strings"
    "unicode/utf8"
)

var order = 4 
var grain = "*" 

func main() {
    // Initialize the pyramid with one element containing the base unit surrounded by spaces
    t := []string{grain + strings.Repeat(" ", utf8.RuneCountInString(grain))}
    fmt.Println("Initial t:", t)
    
    for ; order > 0; order-- {
        fmt.Println("Current order:", order)
        
        // Calculate the number of spaces needed to center the base unit in the current level
        sp := strings.Repeat(" ", utf8.RuneCountInString(t[0])/2)
        fmt.Println("Current sp:", sp)
        
        top := make([]string, len(t))
        
        for i, s := range t {
            fmt.Printf("Current i:%d, s:%s\n", i, s)
            
            // Add spaces on both sides of the base unit to center it in the current level
            top[i] = sp + s + sp
            fmt.Println("Current top:", top)
            
            t[i] += s
            fmt.Println("Updated t[i]:", t[i])
        }
        
        t = append(top, t...)
        fmt.Println("Updated t:", t)
    }
    
    fmt.Println("Final t:")
    for _, r := range t {
        fmt.Println(r)
    }
}

