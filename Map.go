package main

import "fmt"

func main() {
  score := make(map[string]int)
  score["Alice"] = 80
  score["Bob"] = 72
  score["Carol"] = 65
  
  _, isDiana := score["Diana"]
  fmt.Println("is Diana in score?: ", isDiana)
  
  _, isAlice := score["Alice"]
  fmt.Println("is Alice in score?: ", isAlice)
}

// ✓ 데이터 접근하는 socre[key] 구문은 두 개의 값을 리턴하는데, 두 번째 결과값은 해당 값이 존재하는지 여부를 알려준다.
