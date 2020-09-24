package main

import "fmt"

func main() {
  nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  sum := 0
  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)
  
  score := make(map[string]int)
  score["Alice"] = 80
  score["Bob"] = 72
  score["Carol"] = 60
  
  for key, value := range score {
    fmt.Printf("%s -> %s\n", key, value)
  }
}

// ✓ 배열(Array) 혹은 맵(Map) 데이터 자료구조의 내용을 반복하는 용도로 많이 이용한다.
