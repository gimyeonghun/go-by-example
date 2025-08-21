package main

import "fmt"

func vals() (int, int) {
  // The function signature shows that the function returns 2 ints
  return 3, 7
}

 func main() {
   a, b := vals()
   fmt.Println(a)
   fmt.Println(b)
   
   _, c := vals()
   fmt.Println(c)
 }