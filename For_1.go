package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }
}

/*
  ✓ Go언어에서는 while, do ... while 구문은 존재하지 않는다. 오로지 for 구문만 이용한다.
  ✓ Go언어에서는 지시자 뒤에 괄호가 붙지 않는다.
  ✓ for문은 입력으로 bool 타입의 값을 받는다. (true | false), false인 경우에는 반복 구문이 종료된다.
*/
