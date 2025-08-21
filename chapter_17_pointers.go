package main

import "fmt"

// Pointers allow you to pass references to values and records within your program

func zeroval(ival int) {
  // `zeroval` has an `int` parameter, so arguments will be passed to it by value. 
  // `zeroval` will get a copy of `ival` distinct from the one in the function
  ival = 0
}

func zeroptr(iptr *int) {
  // `zeroptr` has an `int` paramter, meaning that it takes an `int` pointer.
  // The `iptr` code in the function body then dereferences the pointer from its memory address to the current value at that address.
  // Assigning a value to a dereferenced pointer changes the value at the referenced address
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("initial:", i)
  
  // `zeroval` doesn't change the `i` in `main`, but `zeroptr` does because it has a reference to the memory address for that variable.
  zeroval(i)
  fmt.Println("zeroval:", i)
  
  // The `&i` syntax gives the memory address of i, ie. a pointer i
  zeroptr(&i)
  fmt.Println("zeroptr:", i)
  
  fmt.Println("pointer", &i)
}