package main
import "fmt"
const s string = "seoul"
func main() {
fmt.Println(s)
s = "busan" //여기서 오류발생한다!
}

/*
  ✓ 상수를 선언할 때는 const 키워드를 이용한다.
  ✓ 상수로 선언된 변수는 초기값을 변경할 수 없다.
*/

