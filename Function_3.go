package main

import "fmt"

func createSequence() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {
increment := createSequence()

fmt.Println(increment())
fmt.Println(increment())
fmt.Println(increment())

newInc := createSequence()
fmt.Println(newInc())
}

/*
  ✓ 아래와 같이 익명 함수를 정의하여 사용할 수 있다.
  ✓ 클로저 및 람다 함수와 비슷한 용도로 Go 언어의 많은 부분에서 이용된다.  
*/
