package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// requestHandler는 context를 활용하여 타임아웃을 처리하는 예제 핸들러
func requestHandler(w http.ResponseWriter, r *http.Request) {
	// 2초짜리 context 생성 (타임아웃)
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	resultChan := make(chan string, 1)

	// 비동기 작업 실행
	go func() {
		time.Sleep(3 * time.Second) // 의도적으로 긴 작업 (2초보다 길게)
		resultChan <- "작업 완료"
	}()

	select {
	case <-ctx.Done():
		http.Error(w, "요청이 타임아웃되었습니다.", http.StatusRequestTimeout)
	case result := <-resultChan:
		fmt.Fprintln(w, result)
	}
}

func main() {
	mux := http.NewServeMux()

	// 핸들러 등록
	mux.HandleFunc("/process", requestHandler)

	// 서버 객체 생성
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// TODO: OS 종료 시그널을 감지하고 서버를 안전하게 종료하도록 구현
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill) // Ctrl+C (SIGINT) 또는 프로세스 종료 (SIGKILL)

	go func() {
		log.Println("서버 시작: 포트 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-stop // 종료 시그널을 받을 때까지 대기

	fmt.Println("\n서버 종료 중...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: server.Shutdown(ctx)를 호출하여 서버를 정상적으로 종료하도록 구현
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("서버 종료 실패: %v", err)
	}

	fmt.Println("서버가 정상적으로 종료되었습니다.")
}
