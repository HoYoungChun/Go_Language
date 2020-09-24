package main

import "fmt"

func main() {
  fruits := make([]string, 3)
  fruits[0] = "Apple"
  fruits[1] = "Banana"
  fruits[2] = "Melon"
  
  fmt.Println("fruits: ", fruits)
  fruits = append(fruits, "Kiwi")
  fmt.Println("fruits: ", fruits)
  fmt.Println("fruits[1:3]: ", fruits[1:3])
  fmt.Println("fruits[2:]: ", fruits[2:])
  fmt.Println("fruits[:4]: ", fruits[:4])
}

/*
  ✓ Slice는 배열에서 조금 더 확장된 기능을 제공한다.
  ✓ 집합의 크기를 자유롭게 늘릴 수 있으며 파이썬 문법과 같이 구간 접근을 할 수 있다.
  ✓ Slice는 정의할 때 make 함수를 이용한다.
  ✓ Slice의 구간 접근은 slice[start:end] 문법으로 접근이 가능한데, end에 위치한 원소는 포함되지 않는다.
*/
