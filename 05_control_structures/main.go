package main

import "fmt"

/*
다른 프로그래밍 언어에는 다양한 종류의 루프(while, do, until, foreach, ...)가 있지만,
Go에서는 다양한 방식으로 사용할 수 있는 단 하나의 루프만 제공한다.
*/
func main() {
	/*
		go lang에서는 아래의 형태처럼 for 루프를 돌릴 수 있음.

				//	var i int = 1
				i := 1
				for i <= 100 {
					fmt.Println("i=", i)
					i = i + 1
				}
	*/

	// 기존의 프로그래밍 방식으로도 가능
	for j := 1; j <= 100; j++ {
		fmt.Println("j=", j)
		// switch case 문법
		switch j {
		case 0:
			fmt.Println("영")
		case 1:
			fmt.Println("일")
		case 2:
			fmt.Println("이")
		case 3:
			fmt.Println("삼")
		case 4:
			fmt.Println("사")
		case 5:
			fmt.Println("오")
		default:
			fmt.Println("알 수 없는 숫자")
		}
	}
}
