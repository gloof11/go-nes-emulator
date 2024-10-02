package main

import (
  "fmt"
)

func main() {
  bus := NewBus()
  cpu := NewOlc6502(bus)
  fmt.Println(cpu)
  fmt.Println(bus)
}
