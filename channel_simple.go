package main

import (
	"fmt"
)

func sender(ch chan string) {
	ch <- "msg1"
	ch <- "msg2"
	ch <- "msg3"
	close(ch) // 채널 닫기: range가 종료되기 위한 조건
}

type StringChannel chan string

func main() {
	ch := make(StringChannel)

	go sender(ch)

	// 채널이 닫힐 때까지 반복해서 수신
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}
