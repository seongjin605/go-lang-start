package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var grade [5]float64
	grade[0] = 98
	grade[1] = 93
	grade[2] = 77
	grade[3] = 82
	grade[4] = 83

	// var total float64 = 0
	// for i := 0; i < 5; i++ {
	// 	total += grade[i]
	// }
	// fmt.Println(total / 5)

	var total float64 = 0
	for i := 0; i < len(grade); i++ {
		total += grade[i]
	}

	//  invalid operation: total / 5 ERROR 발생.
	//  total은 float64인 반면 len(x)는 int다
	// fmt.Println(total / len(x))
	fmt.Println(total / float64(len(grade)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5) // 1,2,3,4,5
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	/**
	panic: runtime error: assignment to entry in nil map
	goroutine 1 [running]:
	main.main()
	main.go:7 +0x4d

	지금까지는 컴파일 시점 오류(compile-time error)만 봤다. 이것은 런타임 오류의 한 예다.
	컴파일 시점 오류가 프로그램을 컴파일할 때 발생하는 데 반해 이름에서 알 수 있듯이 런타임 오류는 프로그램을 실행할 때 발생한다.
	이 프로그램에서 발생한 문제는 맵을 사용하기 전에 초기화해야 한다는 것이다. 즉, 다음과 같이 작성해야 했다.
	*/
	// var key map[string]int
	// key["key1"] = 10
	// fmt.Println(x)

	key := make(map[string]int)
	key["key"] = 10
	fmt.Println(key["key"])

}
