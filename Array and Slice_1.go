package main

import "fmt"

func main() {
  var a [5]int
  fmt.Println("empty set:", a)
  
  fruits := [3]string{"Apple", "Banana", "Melon"}
  fmt.Println("fruits:", fruits)
  fmt.Println("fruits[0]:", fruits[0])
  fmt.Println("fruits[2]:", fruits[2])
}

/*
  ✓ 배열은 특정 데이터의 집합을 정의할 때 사용한다. 한 번 집합의 크기를 정하면 변경할 수 없다.
  ✓ 배열에는 동일한 타입의 데이만 저장할 수 있다.
  ✓ 배열의 원소에 접근할 때는 원소의 인덱스를 0번부터 시작해서 접근한다.
*/
