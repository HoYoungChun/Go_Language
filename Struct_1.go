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

/*
  ✓ 서로 다른 타입을 가진 필드를 하나의 데이터 타입으로 정의할 때 구조체를 이용한다.
  ✓ Go 언어에서는 구조체를 이용해 객체지향 기법을 적용 할 수 있다.
  ✓ type [name] struct { fields } 의 형식으로 구조체를 선언한다.
*/
