package main

import (
  "fmt"
  "maps"
)

func main() {
  // Use the built-in `make` to make an empty map
  m := make(map[string]int)
  
  // Set key/value pairs using typical name[key] = val syntax
  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)
  
  // Get a value for a key with name[key]
  v1 := m["k1"]
  fmt.Println("v1:", v1)
  
  // If the key doesn't exist, the zero value of the value type is returned
  v3 := m["k3"]
  fmt.Println("v3:", v3)
  
  fmt.Println("len:", len(m))
  
  // The built-in `delete` removes key/value pairs from a map
  delete(m, "k2")
  fmt.Println("map:", m)
  
  // Use the `clear` builtin to remove all key/value pairs
  clear(m)
  fmt.Println("map:", m)
  
  _, prs := m["k2"]
  fmt.Println("prs:", prs)
  
  // Declare and intialise a new map in the same line with this syntax
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
  
  n2 := map[string]int{"foo": 2, "bar": 2}
  if maps.Equal(n, n2) {
    fmt.Println("n == n2")
  }
}