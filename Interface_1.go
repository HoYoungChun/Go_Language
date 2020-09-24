package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float64
}

type rect struct{ width, height float64 }
type circle struct{ radius float64 }

func (r rect) area() float64 {
  return r.width * r.height
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

/*
  ✓ 메서드의 집합을 정의한 타입이다. 객체지향 언어에서 부모 클래스와 동일한 개념으로 이용가능하다.
  ✓ Go 언어에서는 인터페이스의 메서드만 구현된 구조체라면 모두 동일한 인터페이스로 취급할 수 있는데 이를 덕 타이핑 이라고 부른다.
*/
