package main

import "fmt"

func main() {
  // Strings can be added with `+`
  fmt.Println("go" + "lang")
  
  // Integers...
  fmt.Println("1 + 1=", 1+1)
   
  // ...and floats
  fmt.Println("7.0/3.0=", 7.0/3.0)

  // Booleans
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
}