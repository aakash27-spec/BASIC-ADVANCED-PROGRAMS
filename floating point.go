package main 
  

import "fmt"
  

func main() { 
    val := 1.1 
  
   
    switch { 
    case val == 1.2: 
        fmt.Println("One Point Two") 
    case val == 1.3: 
        fmt.Println("One Point Three") 
    case val == 1.1: 
        fmt.Println("One Point One") 
    } 
} 