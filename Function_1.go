package main

import "fmt"

func add(a int, b int) int {
  return a + b
}

func main() {
  result1 := add(1, 2)
  fmt.Println("1+2:", result1)
  
  result2 := add(4, 5)
  fmt.Println("4+5:", result2)
  
  result3 := add(8, -2)
  fmt.Println("8+(-2):", result3)
}

/*
  ✓ 코드의 재사용성을 높이기 위해서, 일정 단위를 함수라는 이름으로 정의하여 사용한다.
  ✓ 함수는 func [name](...params) type { ... } 형태로 선언해서 사용한다.
  ✓ 함수의 결과값을 리턴하기 위해서는 명시적으로 return 지시자를 사용해야한다.
*/

