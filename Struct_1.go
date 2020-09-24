package main

import "fmt"

type rect struct {
  width int
  height int
}
func main() {
  rct1 := rect{80, 20}
  rct2 := rect{40, 45}
  fmt.Println(rct1)
  fmt.Println(rct2)
}
