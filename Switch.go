package main

import (
  "fmt"
  "time"
)

func main() {
  switch time.Now().Weekday() {
  case time.Sunday:
    fmt.Println("Today is Sunday")
  case time.Monday:
    fmt.Println("Today is Monday")
  case time.Tuesday:
    fmt.Println("Today is Tuesday")
  case time.Wednesday:
    fmt.Println("Today is Wednesday")
  case time.Thursday:
    fmt.Println("Today is Thursday")
  case time.Friday:
    fmt.Println("Today is Friday")
  case time.Saturday:
    fmt.Println("Today is Saturday")
  default:
    fmt.Println("-")
  }
}

//✓ if 문의 추가 조건이 많아질 때 switch 구문을 쓰면 간결하게 코드를 작성할 수 있다.
