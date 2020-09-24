package main

import "fmt"

type rect struct {
  width, height int
}

func (r rect) area() int {
  return r.width * r.height
}

func (r rect) minPerfectSquare() rect {
  if r.width > r.height {
    return rect{r.height, r.height}
  } else {
    return rect{r.width, r.width}
  }
}

func main() {
  r := rect{20, 25}
  fmt.Println("area: ", r.area())
  fmt.Println("minPerfectSquare:", r.minPerfectSquare())
}

//✓ 구조체에 메서드(특정 데이터 타입에 속한 함수)를 정의할 수 있다.
