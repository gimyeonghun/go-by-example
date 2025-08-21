package main

import "fmt"

type person struct {
  name string
  age int
}

// `newPerson` constructs a new person strcut with the given name
// Because Go is a garbage collected language; you can safely return a pointer to a local variable
func newPerson(name string) *person {
  p := person{name: name}
  p.age = 42
  return &p
}

func main() {
  // This syntax creates a new struct
  fmt.Println(person{"Bob", 20})
  
  // You can name the fields when initialising a struct
  fmt.Println(person{name: "Alice", age: 30})
  
  // Omitted fields will be zero-valued
  fmt.Println(person{name: "Fred"})
  
  // An & prefix yields a pointer to the struct
  fmt.Println(&person{name: "Ann", age: 40})
  
  // It's idiomatic to encapsulate new struct creation in the constructor functions
  fmt.Println(newPerson("Jon"))
  
  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)
  
  sp := &s
  fmt.Println(s.age)
  
  // Structs are mutable
  sp.age = 51
  fmt.Println(sp.age)
  
  // If a struct type is only used for a single value, we don’t have to give it a name. 
  dog := struct {
    name string
    isGood bool
  }{
    "Rex", 
    true,
  }
  fmt.Println(dog)
}