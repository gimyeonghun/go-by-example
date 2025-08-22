package main

import "fmt"

// Enumerated types are a special type of sum types. They have a fixed number of possible values, each with a distinct name
// Our enum type `ServerState` has an underlying `int` type
type ServerState int

// The possible values for `ServerState` are defined as constants. The special keyword `iota` generates successive constant values automatically.
const (
  StateIdle ServerState = iota
  StateConnected
  StateError
  StateRetrying
)

// By implementing the `fmt.Stringer` interface, values of `ServerState` can be printed out or converted to string
var stateName = map[ServerState]string {
  StateIdle: "idle",
  StateConnected: "connected",
  StateError: "error",
  StateRetrying: "retrying",
}

func (ss ServerState) String() string {
  return stateName[ss]
}

func main() {
  ns := transition(StateIdle)
  fmt.Println(ns)
  
  ns2 := transition(ns)
  fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
  switch s {
    case StateIdle:
      return StateConnected
    case StateConnected, StateRetrying:
      return StateIdle
    case StateError:
      return StateError
    default:
      panic(fmt.Errorf("unknown state: %s", s))
  }
}