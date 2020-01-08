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

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func third() {
	fmt.Println("3rd")
}
func main() {
	callInputValue(10)
	callFive()
	x, y := callMultiValues()
	fmt.Println(x, y)

	// 함수안에 함수를 만듬(클로저)
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	// defer는 어떤 식으로든 자원을 해제해야 할 때 자주 사용된다.
	// 이것은 해당 함수가 실행을 완료했을 때 실행을 위해 호출 스케줄을 잡는다.
	// 예를 들어, 파일을 열 때 나중에 해당 파일을 반드시 닫아야 한다.
	// second가 항상 맨마지막에 실행이되므로, 자원해제 함수를 defer로 선언하면 됨.
	defer second()
	first()
	third()

}
