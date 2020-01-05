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

	var total1 float64 = 0
	for i, value := range x {
		total1 += value
	}
	fmt.Println(total1 / float64(len(x)))

}
