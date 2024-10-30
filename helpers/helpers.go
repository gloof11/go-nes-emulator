package helpers

import (
  "reflect"
  "fmt"
)

func FindFunc(function any) uintptr {
  return reflect.ValueOf(function).Pointer()
}

func Hex(number any, width int) string {
  format := fmt.Sprintf("%%0%dX", width)

  return fmt.Sprintf(format, number)
}
