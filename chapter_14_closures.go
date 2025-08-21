package main

import "fmt"

func intSeq() func() int {
  // Function `intSeq` returns another function, which we define anonymously in the body of `intSeq`
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
  nextInt := intSeq()
  
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  
  newInts := intSeq()
  fmt.Println(newInts())
}