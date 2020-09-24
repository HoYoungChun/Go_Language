package main

import "fmt"

func incval(val int) {
  val++
}

func incptr(val *int) {
  *val++
}

func main() {
  i := 1
  incval(i)
  fmt.Println("incval:", i)
  
  incptr(&i)
  fmt.Println("incptr:", i)
}

/*
  ✓ Go 언어에서 함수의 파라미터는 기본적으로 Call-by-Value를 따른다.
  ✓ Call-by-Reference를 이용하기 위해서는 포인터를 정의한 후 사용해야 한다.
*/
