package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

/**
* 규모가 큰 프로그램은 다수의 더 작은 프로그램으로 구성된다.
* 예를 들어 웹 서버는 웹 브라우저에서 오는 요청을 처리해 HTML 웹 페이지를 응답으로 보낸다. 이때 각 요청은 자그마한 프로그램처럼 처리된다.
* 이 같은 작업을 동시에 각기 더 작은 구성요소로 실행할 수 있다면 이상적일 것이다(웹 서버의 경우 여러 요청을 처리하는 것이 여기에 해당한다).
* 하나 이상의 작업을 동시에 진행하는 것을 동시성(concurrency)라 한다. Go에서는 고루틴(goroutine)과 채널(channel)을 통해 동시성을 풍부하게 지원한다.
 */
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
