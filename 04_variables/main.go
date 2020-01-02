package main

import "fmt"

func main() {
	var str string = "Hello World"
	fmt.Println(str)

	str = "Hello World2"
	fmt.Println(str)

	// 여기서 = 앞에 :가 있고 타입을 지정하지 않았다는 점을 눈여겨보자.
	//  Go 컴파일러가 변수에 할당하는 리터럴 값을 토대로 타입을 추론할 수 있기 때문에 타입은 필요하지 않다.
	// (문자열 리터럴을 할당하고 있으므로 x에는 string 타입이 지정된다)
	// 컴파일러는 var 문장에서도 타입을 추론할 수 있다.

	valueOfNumber := 5
	fmt.Println(valueOfNumber)

	const welcomeGo = "WELCOME  GO"
	fmt.Println(welcomeGo)

}
