package main

import "fmt"

func main() {
  num := 2
  if num%2 == 0 {
    fmt.Println("짝수")
  } else {
    fmt.Println("홀수")
  }
  
  score := 92
  
  if score > 90 {
    fmt.Println(score, ": A+")
  } else if score >= 80 {
    fmt.Println(score, ": A0")
  } else {
    fmt.Println(score, ": F")
  }
}

/*
  ✓ if 문의 입력값은 bool 타입의 값을 받는다.
  ✓ 입력값이 true면 바로 다음 분기의 코드를 실행하며, false인 경우에는 else/else-if 구문의 분기로 넘어간다.
  ✓ else if를 이용해서 추가 조건을 입력할 수 있다.
*/
