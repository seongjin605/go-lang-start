package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println(len("Hello World"))
	// 문자열에는 1이 아닌 0부터 시작하는 "인덱스"가 지정돼 있다. [1]은 첫 번째 요소가 아닌 두 번째 요소를 반환한다.
	// 프로그램을 실행하면 e이 아닌 101(아스키코드10진수)이 출력된다는 점을 눈여겨보자.
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}
