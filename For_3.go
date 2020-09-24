package main

import "fmt"

func main() {
  i := 1
  for {
    if i == 10 {
      break
      }
    fmt.Println(i)
    i++
  }
}

/*
  ✓ for 문에 아무런 조건을 입력하지 않으면 입력값은 true로 간주한다.
  ✓ break 명령어로 반복을 종료할 수 있다.
*/
