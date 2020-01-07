package main

import "fmt"

/*
함수는 x에 접근할 수 없으므로, 컴파일 에러가 난다.

func f() {
	fmt.Println(x)
}
func main() {
	x := 5
	f()
}
*/

// 다음과 같이 접근해서 x값을 파라미터로 넘기면 가능.

func callInputValue(x int) {
	fmt.Println(x)
}

//func main() {
//	x := 5
//	f(x)
//}

// 다음과 같이도 선언 가능.
var five int = 5

func callFive() {
	fmt.Println(five)
}

// Multi-valued expressions(다중값 반환 표현식)
func callMultiValues() (int, int) {
	return 5, 6
}

func main() {
	callInputValue(10)
	callFive()
	x, y := callMultiValues()
	fmt.Println(x, y)
}
