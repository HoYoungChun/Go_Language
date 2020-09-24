package main

import "fmt"

func main() {
  var a int = 10
  fmt.Println("a =", a)
  
  var b = true // 타입 추론
  fmt.Println(b)
  
  var c int
  fmt.Println(c)
  
  d := "seoul" // 짧은 표현
  fmt.Println(d)
}

/*
  ✓ 문자열, 정수, 부동소수점, boolean 등의 기본 값이 존재한다.
  ✓ 기본적인 사칙연산을 제공한다.
  ✓ 존재 하지 않는 값을 표현할 때는 nil을 이용한다.
*/
