package main

import "fmt"

func main() {
  // `var` declares 1 or more variables
  var a = "initial"
  fmt.Println(a)
  
  // Go will infer the type of initialised variables
  var b, c int = 1, 2
  fmt.Println(b, c)
  
  var d = true
  fmt.Println(d)
  
  var e int
  fmt.Println(e)
  
  // The `:=` syntax will declare and initialised a variable
  f := "apple"
  fmt.Println(f)
}