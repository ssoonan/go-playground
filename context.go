package main

import (
	"context"
	"fmt"
	"time"
)

// 작업을 수행하는 함수
func slowOperation(ctx context.Context, resultChan chan<- string) {
	select {
	case <-time.After(3 * time.Second): // 실제 작업 (3초 걸림)
		resultChan <- "작업 완료!"
	case <-ctx.Done(): // context가 취소되면
		resultChan <- "작업이 취소됨: " + ctx.Err().Error()
	}
}

func main2() {
	// 2초 후 타임아웃이 발생하는 컨텍스트 생성
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 메모리 누수를 방지하기 위해 취소

	resultChan := make(chan string)

	// 별도 고루틴에서 작업 수행
	go slowOperation(ctx, resultChan)

	// 결과 출력
	fmt.Println("작업 시작...")
	fmt.Println(<-resultChan) // 결과를 출력
}
