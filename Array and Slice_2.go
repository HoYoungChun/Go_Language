package main

import "fmt"

func main() {
  var vector [3][3]int
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      vector[i][j] = i + j
    }
  }
  fmt.Println("vector[3][3]:", vector)
}

// ✓ 아래 예시코드와 같이 다차원 배열을 정의할 수 있다.
