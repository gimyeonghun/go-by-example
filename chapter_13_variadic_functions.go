package main

import "fmt"

// Function that will take arbitary number of ints as arguments
// Within the function, the type of nums is equivalent to `[]int`
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}

func main() {
  // Variadic functions can be called in the usual way with individual arguments
  sum(1, 2)
  sum(1, 2, 3)
  
  nums := []int{1, 2, 3, 4}
  sum(nums...)
}