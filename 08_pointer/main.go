package main

import "fmt"

// zero 함수는 main 함수에 있는 원본 x 변수를 변경하지 않을 것이다.
func zero(x int) {
	x = 0
}

// 포인터는 값 자체보다는 값이 저장된 메모리상의 위치를 가리킨다(포인터는 다른 뭔가를 가리킨다).
// 포인터 (*int)를 이용하면 zero 함수가 원본 변수를 수정할 수 있게 된다.
func changeZero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x는 여전히 5

	xPtr := 5
	changeZero(&xPtr)
	fmt.Println(xPtr) // xPtr값이 0 으로 변경됨.
}
