package main

import "fmt"

func main() {
  // An array is a numbered sequence of elements of a specific length. In idiomatic Go, slices are much more common and arrays are only useful in specific scenarios.
  
  // Array that holds exactly 5 ints
  var a [5] int
  fmt.Println("emp:", a)
  
  a[4] = 100
  fmt.Println("set:", a)
  fmt.Println("get", a[4])
  
  // Returns the length of an array
  fmt.Println("len", len(a))
  
  // Declare and initialise an array in one line
  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)
  
  // The compiler will count the number of elements for you with `...`
  b = [...]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)
  
  // If you specify the index with `:`, the elements in between will be zeroed
  b = [...]int{100, 3: 400, 500}
  fmt.Println("idx:", b)
  
  // You can compose types to build multi-dimensional structuresx
  var twoD [2][3]int
  for i := range 2 {
    for j := range 2 {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
  
  twoD = [2][3]int{
    {1, 2, 3},
    {1, 2, 3},
  }
  fmt.Println("2d: ", twoD)
}