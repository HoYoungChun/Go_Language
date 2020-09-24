package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Println(i)
  }
}

// ✓ continue 명령어로 현재 분기의 코드를 무시하고 다음 반복 주기로 넘어갈 수 있다.
