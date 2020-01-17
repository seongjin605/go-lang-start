package main

import (
	"fmt"
	"time"
)

func f() {
	// 시작 시간
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		time.Sleep(1000)
	}
	// 경과 시간
	elapsedTime := time.Since(startTime)
	fmt.Printf("실행시간: %s\n", elapsedTime)
}

// goroutine async 동작 방식 파악
// https://golang.org/ref/spec#Go_statements
func main() {
	go f()
	time.Sleep(10 * time.Second)
}
