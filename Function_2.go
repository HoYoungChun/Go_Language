package main

import "fmt"

func eval(a int, b int) (int, int, int, int) {
  return (a + b), (a - b), (a * b), (a / b)
}

func main() {
  add, sub, mul, div := eval(5, 8)

  fmt.Println("5+8", add)
  fmt.Println("5-8", sub)
  fmt.Println("5*8", mul)
  fmt.Println("5/8", div)
}

// ✓ Go언어에서 함수는 2개 이상의 값을 리턴할 수 있다.
