package main

import (
  "fmt"
  "math"
)

// Here's a basic interface for geometric shapes
type geometry interface {
  area() float64
  perim() float64
}

type rect struct {
  width, height float64
}

type circle struct {
  radius float64
}

// We'll implement the interface on `circle` and `rect` types
// To implement an interface, all the methods in the interface need to be implemented.
func (r rect) area() float64 {
  return r.width * r.height
}

func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}

// The generic `measure` function can take advantage to work on any `geometry`
func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}

// Sometimes it's useful to know the runtime type of an interface value. Use a `type assertion` as shown here.
func detectCircle(g geometry) {
  if c, ok := g.(circle); ok {
    fmt.Println("circle with radius", c.radius)
  }
}

func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}
  
  measure(r)
  measure(c)
  
  detectCircle(r)
  detectCircle(c)
}