package main

import (
	"fmt"
	"math"
)

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// 이렇게 하면 모든 필드에 대한 메모리가 할당되고, 각 필드는 0 값으로 설정된 후 포인터(*Circle)가 반환된다.
	// 하지만 각 필드에 값을 할당하고 싶을 때가 더 많다. 이렇게 하는 데는 두 가지 방법이 있다. 다음 코드를 보자.
	//c := new(Circle)
	c := Circle{x: 0, y: 0, r: 5}
	// 또는 필드가 정의된 순서를 알고 있을 경우 필드명을 생략할 수 있다.
	// c := Circle{0, 0, 5}
	fmt.Print("c:", c)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

type Circle struct {
	x, y, r float64
}
